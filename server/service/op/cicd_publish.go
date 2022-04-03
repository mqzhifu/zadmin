package op

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/op"
	opReq "github.com/flipped-aurora/gin-vue-admin/server/model/op/request"
)

type CicdPublishService struct {
}

// CreateCicdPublish 创建CicdPublish记录
// Author [piexlmax](https://github.com/piexlmax)
func (cicdPublishService *CicdPublishService) CreateCicdPublish(cicdPublish op.CicdPublish) (err error) {
	err = myBusinessDb.GetDbInc().Create(&cicdPublish).Error
	return err
}

// DeleteCicdPublish 删除CicdPublish记录
// Author [piexlmax](https://github.com/piexlmax)
func (cicdPublishService *CicdPublishService) DeleteCicdPublish(cicdPublish op.CicdPublish) (err error) {
	err = myBusinessDb.GetDbInc().Delete(&cicdPublish).Error
	return err
}

// DeleteCicdPublishByIds 批量删除CicdPublish记录
// Author [piexlmax](https://github.com/piexlmax)
func (cicdPublishService *CicdPublishService) DeleteCicdPublishByIds(ids request.IdsReq) (err error) {
	err = myBusinessDb.GetDbInc().Delete(&[]op.CicdPublish{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCicdPublish 更新CicdPublish记录
// Author [piexlmax](https://github.com/piexlmax)
func (cicdPublishService *CicdPublishService) UpdateCicdPublish(cicdPublish op.CicdPublish) (err error) {
	err = myBusinessDb.GetDbInc().Save(&cicdPublish).Error
	return err
}

// GetCicdPublish 根据id获取CicdPublish记录
// Author [piexlmax](https://github.com/piexlmax)
func (cicdPublishService *CicdPublishService) GetCicdPublish(id uint) (err error, cicdPublish op.CicdPublish) {
	err = myBusinessDb.GetDbInc().Where("id = ?", id).First(&cicdPublish).Error
	return
}

// GetCicdPublishInfoList 分页获取CicdPublish记录
// Author [piexlmax](https://github.com/piexlmax)
func (cicdPublishService *CicdPublishService) GetCicdPublishInfoList(info opReq.CicdPublishSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := myBusinessDb.GetDbInc().Model(&op.CicdPublish{})
	var cicdPublishs []op.CicdPublish
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&cicdPublishs).Error
	return err, cicdPublishs, total
}
