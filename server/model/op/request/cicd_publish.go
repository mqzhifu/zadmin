package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/op"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CicdPublishSearch struct{
    op.CicdPublish
    request.PageInfo
}
