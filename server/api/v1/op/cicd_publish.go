package op

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/op"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    opReq "github.com/flipped-aurora/gin-vue-admin/server/model/op/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CicdPublishApi struct {
}

var cicdPublishService = service.ServiceGroupApp.OpServiceGroup.CicdPublishService


// CreateCicdPublish 创建CicdPublish
// @Tags CicdPublish
// @Summary 创建CicdPublish
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body op.CicdPublish true "创建CicdPublish"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cicdPublish/createCicdPublish [post]
func (cicdPublishApi *CicdPublishApi) CreateCicdPublish(c *gin.Context) {
	var cicdPublish op.CicdPublish
	_ = c.ShouldBindJSON(&cicdPublish)
	if err := cicdPublishService.CreateCicdPublish(cicdPublish); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCicdPublish 删除CicdPublish
// @Tags CicdPublish
// @Summary 删除CicdPublish
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body op.CicdPublish true "删除CicdPublish"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cicdPublish/deleteCicdPublish [delete]
func (cicdPublishApi *CicdPublishApi) DeleteCicdPublish(c *gin.Context) {
	var cicdPublish op.CicdPublish
	_ = c.ShouldBindJSON(&cicdPublish)
	if err := cicdPublishService.DeleteCicdPublish(cicdPublish); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCicdPublishByIds 批量删除CicdPublish
// @Tags CicdPublish
// @Summary 批量删除CicdPublish
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CicdPublish"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cicdPublish/deleteCicdPublishByIds [delete]
func (cicdPublishApi *CicdPublishApi) DeleteCicdPublishByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := cicdPublishService.DeleteCicdPublishByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCicdPublish 更新CicdPublish
// @Tags CicdPublish
// @Summary 更新CicdPublish
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body op.CicdPublish true "更新CicdPublish"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cicdPublish/updateCicdPublish [put]
func (cicdPublishApi *CicdPublishApi) UpdateCicdPublish(c *gin.Context) {
	var cicdPublish op.CicdPublish
	_ = c.ShouldBindJSON(&cicdPublish)
	if err := cicdPublishService.UpdateCicdPublish(cicdPublish); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCicdPublish 用id查询CicdPublish
// @Tags CicdPublish
// @Summary 用id查询CicdPublish
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query op.CicdPublish true "用id查询CicdPublish"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cicdPublish/findCicdPublish [get]
func (cicdPublishApi *CicdPublishApi) FindCicdPublish(c *gin.Context) {
	var cicdPublish op.CicdPublish
	_ = c.ShouldBindQuery(&cicdPublish)
	if err, recicdPublish := cicdPublishService.GetCicdPublish(cicdPublish.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recicdPublish": recicdPublish}, c)
	}
}

// GetCicdPublishList 分页获取CicdPublish列表
// @Tags CicdPublish
// @Summary 分页获取CicdPublish列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query opReq.CicdPublishSearch true "分页获取CicdPublish列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cicdPublish/getCicdPublishList [get]
func (cicdPublishApi *CicdPublishApi) GetCicdPublishList(c *gin.Context) {
	var pageInfo opReq.CicdPublishSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := cicdPublishService.GetCicdPublishInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
