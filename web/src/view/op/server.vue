<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="环境">
          <el-input v-model="searchInfo.env" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="内网IP">
          <el-input v-model="searchInfo.innerIp" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="外网IP">
          <el-input v-model="searchInfo.outIp" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="平台类型">
          <el-input v-model="searchInfo.platform" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="状态">
          <el-input v-model="searchInfo.status" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
          <el-table-column align="left" label="ID" prop="ID" width="40" />
        <el-table-column align="left" label="名称" prop="name" width="" />
        <el-table-column align="left" label="环境" prop="env" width="">
            <template #default="scope">
            {{ filterDict(scope.row.env,ENVOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="内网IP" prop="innerIp" width="" />
        <el-table-column align="left" label="外网IP" prop="outIp" width="" />
        <el-table-column align="left" label="平台类型" prop="platform" width="">
            <template #default="scope">
            {{ filterDict(scope.row.platform,SERVER_PLATFORMOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="负责人" prop="chargeUserName" width="" />
        <el-table-column align="left" label="自定义配置信息" prop="ext" width="" />
        <el-table-column align="left" label="价格" prop="price" width="" />
        <el-table-column align="left" label="结束时间" prop="endTime" width="" />
        <el-table-column align="left" label="开始时间" prop="startTime" width="" />
        <el-table-column align="left" label="状态" prop="status" width="">
            <template #default="scope">
            {{ filterDict(scope.row.status,PROJECT_STATUSOptions) }}
            </template>
        </el-table-column>

          <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatUnixtDate(scope.row.CreatedAt) }}</template>
          </el-table-column>

        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateServerFunc(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="环境:">
          <el-select v-model="formData.env" placeholder="请选择" style="width:100%" clearable>
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
          <el-select v-model="formData.platform" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in SERVER_PLATFORMOptions" :key="key" :label="item.label" :value="item.value" />
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
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in PROJECT_STATUSOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
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
  deleteServer,
  deleteServerByIds,
  updateServer,
  findServer,
  getServerList
} from '@/api/server'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatUnixtDate,formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const ENVOptions = ref([])
const SERVER_PLATFORMOptions = ref([])
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

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getServerList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    ENVOptions.value = await getDictFunc('ENV')
    SERVER_PLATFORMOptions.value = await getDictFunc('SERVER_PLATFORM')
    PROJECT_STATUSOptions.value = await getDictFunc('PROJECT_STATUS')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteServerFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteServerByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateServerFunc = async(row) => {
    const res = await findServer({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reserver
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteServerFunc = async (row) => {
    const res = await deleteServer({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
        closeDialog()
        getTableData()
      }
}
</script>

<style>
</style>
