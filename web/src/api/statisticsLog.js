import service from '@/utils/request'

// @Tags StatisticsLog
// @Summary 创建StatisticsLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StatisticsLog true "创建StatisticsLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /statisticsLog/createStatisticsLog [post]
export const createStatisticsLog = (data) => {
  return service({
    url: '/statisticsLog/createStatisticsLog',
    method: 'post',
    data
  })
}

// @Tags StatisticsLog
// @Summary 删除StatisticsLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StatisticsLog true "删除StatisticsLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /statisticsLog/deleteStatisticsLog [delete]
export const deleteStatisticsLog = (data) => {
  return service({
    url: '/statisticsLog/deleteStatisticsLog',
    method: 'delete',
    data
  })
}

// @Tags StatisticsLog
// @Summary 删除StatisticsLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除StatisticsLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /statisticsLog/deleteStatisticsLog [delete]
export const deleteStatisticsLogByIds = (data) => {
  return service({
    url: '/statisticsLog/deleteStatisticsLogByIds',
    method: 'delete',
    data
  })
}

// @Tags StatisticsLog
// @Summary 更新StatisticsLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StatisticsLog true "更新StatisticsLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /statisticsLog/updateStatisticsLog [put]
export const updateStatisticsLog = (data) => {
  return service({
    url: '/statisticsLog/updateStatisticsLog',
    method: 'put',
    data
  })
}

// @Tags StatisticsLog
// @Summary 用id查询StatisticsLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.StatisticsLog true "用id查询StatisticsLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /statisticsLog/findStatisticsLog [get]
export const findStatisticsLog = (params) => {
  return service({
    url: '/statisticsLog/findStatisticsLog',
    method: 'get',
    params
  })
}

// @Tags StatisticsLog
// @Summary 分页获取StatisticsLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取StatisticsLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /statisticsLog/getStatisticsLogList [get]
export const getStatisticsLogList = (params) => {
  return service({
    url: '/statisticsLog/getStatisticsLogList',
    method: 'get',
    params
  })
}
