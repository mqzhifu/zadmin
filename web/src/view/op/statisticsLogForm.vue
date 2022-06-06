<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="项目ID:">
          <el-input v-model.number="formData.projectId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户ID:">
          <el-input v-model.number="formData.uid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分类，暂未使用:">
          <el-input v-model.number="formData.category" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="动作标识:">
          <el-input v-model="formData.action" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="http请求头客户端基础信息:">
          <el-input v-model="formData.headerBase" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="http公共请求头信息:">
          <el-input v-model="formData.headerCommon" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="自定义消息体:">
          <el-input v-model="formData.msg" clearable placeholder="请输入" />
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
  name: 'StatisticsLog'
}
</script>

<script setup>
import {
  createStatisticsLog,
  updateStatisticsLog,
  findStatisticsLog
} from '@/api/statisticsLog'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        projectId: 0,
        uid: 0,
        category: 0,
        action: '',
        headerBase: '',
        headerCommon: '',
        msg: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findStatisticsLog({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.restatisticsLog
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
          res = await createStatisticsLog(formData.value)
          break
        case 'update':
          res = await updateStatisticsLog(formData.value)
          break
        default:
          res = await createStatisticsLog(formData.value)
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
