// 自动生成模板CicdPublish
package op

// CicdPublish 结构体
// 如果含有time.Time 请自行import time包
type CicdPublish struct {
	DIY_MODEL
	ErrInfo     string `json:"errInfo" form:"errInfo" gorm:"column:err_info;comment:错误日志;size:255;"`
	ExecTime    *int   `json:"execTime" form:"execTime" gorm:"column:exec_time;comment:执行时间;size:10;"`
	Log         string `json:"log" form:"log" gorm:"column:log;comment:日志;"`
	ServerId    *int   `json:"serverId" form:"serverId" gorm:"column:server_id;comment:服务器ID;size:10;"`
	ServerInfo  string `json:"serverInfo" form:"serverInfo" gorm:"column:server_info;comment:服务器信息-备份;size:50;"`
	ServiceId   *int   `json:"serviceId" form:"serviceId" gorm:"column:service_id;comment:服务ID;size:10;"`
	ServiceInfo string `json:"serviceInfo" form:"serviceInfo" gorm:"column:service_info;comment:服务信息-备份;size:50;"`
	Status      *bool  `json:"status" form:"status" gorm:"column:status;comment:1发布中2发布失败3发布成功;"`
}

// TableName CicdPublish 表名
func (CicdPublish) TableName() string {
	return "cicd_publish"
}
