<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="环境:">
          <el-select v-model="formData.env" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in ENVOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="内网IP:">
          <el-input v-model="formData.innerIp" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="外网IP:">
          <el-input v-model="formData.outIp" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="平台类型:">
          <el-select v-model="formData.platform" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in PLATFORMOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="负责人:">
          <el-input v-model="formData.chargeUserName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="自定义配置信息:">
          <el-input v-model="formData.ext" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="价格:">
          <el-input v-model.number="formData.price" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结束时间:">
          <el-input v-model.number="formData.endTime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="开始时间:">
          <el-input v-model.number="formData.startTime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:">
          <el-select v-model="formData.status" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in PROJECT_STATUSOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  name: 'Server'
}
</script>

<script setup>
import {
  createServer,
  updateServer,
  findServer
} from '@/api/server'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const ENVOptions = ref([])
const PLATFORMOptions = ref([])
const PROJECT_STATUSOptions = ref([])
const formData = ref({
        name: '',
        env: undefined,
        innerIp: '',
        outIp: '',
        platform: undefined,
        chargeUserName: '',
        ext: '',
        price: 0,
        endTime: 0,
        startTime: 0,
        status: undefined,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findServer({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reserver
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    ENVOptions.value = await getDictFunc('ENV')
    PLATFORMOptions.value = await getDictFunc('PLATFORM')
    PROJECT_STATUSOptions.value = await getDictFunc('PROJECT_STATUS')
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createServer(formData.value)
          break
        case 'update':
          res = await updateServer(formData.value)
          break
        default:
          res = await createServer(formData.value)
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
