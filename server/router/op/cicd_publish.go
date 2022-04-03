package op

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CicdPublishRouter struct {
}

// InitCicdPublishRouter 初始化 CicdPublish 路由信息
func (s *CicdPublishRouter) InitCicdPublishRouter(Router *gin.RouterGroup) {
	cicdPublishRouter := Router.Group("cicdPublish").Use(middleware.OperationRecord())
	cicdPublishRouterWithoutRecord := Router.Group("cicdPublish")
	var cicdPublishApi = v1.ApiGroupApp.OpApiGroup.CicdPublishApi
	{
		cicdPublishRouter.POST("createCicdPublish", cicdPublishApi.CreateCicdPublish)   // 新建CicdPublish
		cicdPublishRouter.DELETE("deleteCicdPublish", cicdPublishApi.DeleteCicdPublish) // 删除CicdPublish
		cicdPublishRouter.DELETE("deleteCicdPublishByIds", cicdPublishApi.DeleteCicdPublishByIds) // 批量删除CicdPublish
		cicdPublishRouter.PUT("updateCicdPublish", cicdPublishApi.UpdateCicdPublish)    // 更新CicdPublish
	}
	{
		cicdPublishRouterWithoutRecord.GET("findCicdPublish", cicdPublishApi.FindCicdPublish)        // 根据ID获取CicdPublish
		cicdPublishRouterWithoutRecord.GET("getCicdPublishList", cicdPublishApi.GetCicdPublishList)  // 获取CicdPublish列表
	}
}
