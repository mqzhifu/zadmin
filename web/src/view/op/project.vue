<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="类型">
          <el-select v-model="searchInfo.type" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in PROJECT_TYPEOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>





        <el-form-item label="状态1正常2关闭">
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
        <el-table-column align="left" label="类型" prop="type" width="">
            <template #default="scope">
            {{ filterDict(scope.row.type,PROJECT_TYPEOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="baseAuth 认证KEY" prop="access" width="" />
        <el-table-column align="left" label="密钥" prop="secretKey" width="" />
        <el-table-column align="left" label="状态1正常2关闭" prop="status" width="">
            <template #default="scope">
            {{ filterDict(scope.row.status,PROJECT_STATUSOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="描述信息" prop="desc" width="" />
        <el-table-column align="left" label="git仓地址" prop="git" width="" />
          <el-table-column align="left" label="日期" width="">
            <template #default="scope">{{ formatUnixtDate(scope.row.CreatedAt) }}</template>
          </el-table-column>

        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateProjectFunc(scope.row)">变更</el-button>
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
      <el-form ref="projectFormVV" :rules="rules" :model="formData" label-position="right" label-width="80px">

        <el-form-item label="ID:" v-if="isEdit">
          <el-input v-model="formData.ID" :disabled="isEdit" />
        </el-form-item>


        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="类型:" prop="type">
          <el-select v-model="formData.type" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in PROJECT_TYPEOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="baseAuth 认证KEY:" prop="access">
          <el-input v-model="formData.access" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密钥:">
          <el-input v-model="formData.secretKey" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态1正常2关闭:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in PROJECT_STATUSOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="描述信息:"  prop="desc">
          <el-input v-model="formData.desc" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="git仓地址:" prop="git">
          <el-input v-model="formData.git" clearable placeholder="请输入" />
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
  name: 'Project'
}
</script>

<script setup>
import {
  createProject,
  deleteProject,
  deleteProjectByIds,
  updateProject,
  findProject,
  getProjectList
} from '@/api/project'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate,formatUnixtDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const PROJECT_TYPEOptions = ref([])
const PROJECT_STATUSOptions = ref([])

const projectFormVV = ref (null)
const rules = ref({
  name: [
    { required: true, message: '请输入 名称', trigger: 'blur' },
    // {pattern: /^[0-9]*$/, message: `请输入数字`, trigger: 'blur'}
    {pattern: /^[A-Z]{1}[A-Za-z]+$/, message: `只能是字母，且首字母大写`, trigger: 'blur'}
  ],
  type: [
    { required: true, message: '不允许为空', trigger: 'blur' }
  ],
  status: [
    { required: true, message: '不允许为空', trigger: 'blur' }
  ],
  access: [
    { required: true, message: '请输入 access', trigger: 'blur' }
  ],

  desc: [
    { required: true, message: '请输入 描述信息', trigger: 'blur' }
  ],
  git: [
    { required: true, message: '请输入git clone 地址,用于cicd', trigger: 'blur' }
  ]

})



const formData = ref({
        name: '',
        type: undefined,
        access: '',
        secretKey: '',
        status: undefined,
        desc: '',
        git: '',
})






// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const isEdit = ref(false)

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
  const table = await getProjectList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    PROJECT_TYPEOptions.value = await getDictFunc('PROJECT_TYPE')
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
            deleteProjectFunc(row)
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
      const res = await deleteProjectByIds({ ids })
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
const updateProjectFunc = async(row) => {

    const res = await findProject({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      isEdit.value = true
        formData.value = res.data.reproject
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteProjectFunc = async (row) => {
    const res = await deleteProject({ ID: row.ID })
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
  projectFormVV.value.resetFields()
    isEdit.value = false
    dialogFormVisible.value = false
    formData.value = {
        name: '',
        type: undefined,
        access: '',
        secretKey: '',
        status: undefined,
        desc: '',
        git: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
  // projectFormVV
  console.log("enterDialog:",formData.value)
  projectFormVV.value.validate(async valid => {
    if (valid) {
      let res
      switch (type.value) {
        case 'create':
          res = await createProject(formData.value)
          break
        case 'update':
          res = await updateProject(formData.value)
          break
        default:
          res = await createProject(formData.value)
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
  })
}
</script>

<style>
</style>
