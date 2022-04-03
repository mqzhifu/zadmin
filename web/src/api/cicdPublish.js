import service from '@/utils/request'

// @Tags CicdPublish
// @Summary 创建CicdPublish
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CicdPublish true "创建CicdPublish"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cicdPublish/createCicdPublish [post]
export const createCicdPublish = (data) => {
  return service({
    url: '/cicdPublish/createCicdPublish',
    method: 'post',
    data
  })
}

// @Tags CicdPublish
// @Summary 删除CicdPublish
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CicdPublish true "删除CicdPublish"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cicdPublish/deleteCicdPublish [delete]
export const deleteCicdPublish = (data) => {
  return service({
    url: '/cicdPublish/deleteCicdPublish',
    method: 'delete',
    data
  })
}

// @Tags CicdPublish
// @Summary 删除CicdPublish
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CicdPublish"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cicdPublish/deleteCicdPublish [delete]
export const deleteCicdPublishByIds = (data) => {
  return service({
    url: '/cicdPublish/deleteCicdPublishByIds',
    method: 'delete',
    data
  })
}

// @Tags CicdPublish
// @Summary 更新CicdPublish
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CicdPublish true "更新CicdPublish"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cicdPublish/updateCicdPublish [put]
export const updateCicdPublish = (data) => {
  return service({
    url: '/cicdPublish/updateCicdPublish',
    method: 'put',
    data
  })
}

// @Tags CicdPublish
// @Summary 用id查询CicdPublish
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CicdPublish true "用id查询CicdPublish"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cicdPublish/findCicdPublish [get]
export const findCicdPublish = (params) => {
  return service({
    url: '/cicdPublish/findCicdPublish',
    method: 'get',
    params
  })
}

// @Tags CicdPublish
// @Summary 分页获取CicdPublish列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CicdPublish列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cicdPublish/getCicdPublishList [get]
export const getCicdPublishList = (params) => {
  return service({
    url: '/cicdPublish/getCicdPublishList',
    method: 'get',
    params
  })
}
