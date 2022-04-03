<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="负责人姓名:">
          <el-input v-model="formData.chargeUserName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结束时间:">
          <el-input v-model.number="formData.endTime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="环境变量:">
          <el-input v-model="formData.env" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="自定义配置信息:">
          <el-input v-model="formData.ext" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="主机地址:">
          <el-input v-model="formData.host" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="平台类型1自有2阿里3腾讯4华为:">
          <el-input v-model="formData.platform" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="主机端口号:">
          <el-input v-model="formData.port" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="价格:">
          <el-input v-model.number="formData.price" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密码:">
          <el-input v-model="formData.ps" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开始时间:">
          <el-input v-model.number="formData.startTime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户名:">
          <el-input v-model="formData.user" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Instance'
}
</script>

<script setup>
import {
  createInstance,
  updateInstance,
  findInstance
} from '@/api/instance'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        chargeUserName: '',
        endTime: 0,
        env: '',
        ext: '',
        host: '',
        name: '',
        platform: '',
        port: '',
        price: 0,
        ps: '',
        startTime: 0,
        user: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findInstance({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reinstance
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createInstance(formData.value)
          break
        case 'update':
          res = await updateInstance(formData.value)
          break
        default:
          res = await createInstance(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
