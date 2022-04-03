package op

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/op"
	opReq "github.com/flipped-aurora/gin-vue-admin/server/model/op/request"
)

type ServerService struct {
}

// CreateServer 创建Server记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverService *ServerService) CreateServer(server op.Server) (err error) {
	err = myBusinessDb.GetDbInc().Create(&server).Error
	return err
}

// DeleteServer 删除Server记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverService *ServerService) DeleteServer(server op.Server) (err error) {
	err = myBusinessDb.GetDbInc().Delete(&server).Error
	return err
}

// DeleteServerByIds 批量删除Server记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverService *ServerService) DeleteServerByIds(ids request.IdsReq) (err error) {
	err = myBusinessDb.GetDbInc().Delete(&[]op.Server{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateServer 更新Server记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverService *ServerService) UpdateServer(server op.Server) (err error) {
	err = myBusinessDb.GetDbInc().Save(&server).Error
	return err
}

// GetServer 根据id获取Server记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverService *ServerService) GetServer(id uint) (err error, server op.Server) {
	err = myBusinessDb.GetDbInc().Where("id = ?", id).First(&server).Error
	return
}

// GetServerInfoList 分页获取Server记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverService *ServerService) GetServerInfoList(info opReq.ServerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := myBusinessDb.GetDbInc().Model(&op.Server{})
	var servers []op.Server
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&servers).Error
	return err, servers, total
}