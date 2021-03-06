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

type ServerApi struct {
}

var serverService = service.ServiceGroupApp.OpServiceGroup.ServerService


// CreateServer 创建Server
// @Tags Server
// @Summary 创建Server
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body op.Server true "创建Server"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /server/createServer [post]
func (serverApi *ServerApi) CreateServer(c *gin.Context) {
	var server op.Server
	_ = c.ShouldBindJSON(&server)
	if err := serverService.CreateServer(server); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteServer 删除Server
// @Tags Server
// @Summary 删除Server
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body op.Server true "删除Server"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /server/deleteServer [delete]
func (serverApi *ServerApi) DeleteServer(c *gin.Context) {
	var server op.Server
	_ = c.ShouldBindJSON(&server)
	if err := serverService.DeleteServer(server); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteServerByIds 批量删除Server
// @Tags Server
// @Summary 批量删除Server
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Server"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /server/deleteServerByIds [delete]
func (serverApi *ServerApi) DeleteServerByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := serverService.DeleteServerByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateServer 更新Server
// @Tags Server
// @Summary 更新Server
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body op.Server true "更新Server"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /server/updateServer [put]
func (serverApi *ServerApi) UpdateServer(c *gin.Context) {
	var server op.Server
	_ = c.ShouldBindJSON(&server)
	if err := serverService.UpdateServer(server); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindServer 用id查询Server
// @Tags Server
// @Summary 用id查询Server
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query op.Server true "用id查询Server"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /server/findServer [get]
func (serverApi *ServerApi) FindServer(c *gin.Context) {
	var server op.Server
	_ = c.ShouldBindQuery(&server)
	if err, reserver := serverService.GetServer(server.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reserver": reserver}, c)
	}
}

// GetServerList 分页获取Server列表
// @Tags Server
// @Summary 分页获取Server列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query opReq.ServerSearch true "分页获取Server列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /server/getServerList [get]
func (serverApi *ServerApi) GetServerList(c *gin.Context) {
	var pageInfo opReq.ServerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := serverService.GetServerInfoList(pageInfo); err != nil {
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
