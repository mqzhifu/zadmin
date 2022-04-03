<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="错误日志:">
          <el-input v-model="formData.errInfo" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="执行时间:">
          <el-input v-model.number="formData.execTime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="日志:">
          <el-input v-model="formData.log" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="服务器ID:">
          <el-input v-model.number="formData.serverId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="服务器信息-备份:">
          <el-input v-model="formData.serverInfo" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="服务ID:">
          <el-input v-model.number="formData.serviceId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="服务信息-备份:">
          <el-input v-model="formData.serviceInfo" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="1发布中2发布失败3发布成功:">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  name: 'CicdPublish'
}
</script>

<script setup>
import {
  createCicdPublish,
  updateCicdPublish,
  findCicdPublish
} from '@/api/cicdPublish'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        errInfo: '',
        execTime: 0,
        log: '',
        serverId: 0,
        serverInfo: '',
        serviceId: 0,
        serviceInfo: '',
        status: false,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCicdPublish({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recicdPublish
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
          res = await createCicdPublish(formData.value)
          break
        case 'update':
          res = await updateCicdPublish(formData.value)
          break
        default:
          res = await createCicdPublish(formData.value)
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
