// 自动生成模板Server
package op

// Server 结构体
// 如果含有time.Time 请自行import time包
type Server struct {
	DIY_MODEL
	ChargeUserName string `json:"chargeUserName" form:"chargeUserName" gorm:"column:charge_user_name;comment:负责人姓名;size:50;"`
	EndTime        *int   `json:"endTime" form:"endTime" gorm:"column:end_time;comment:结束时间;size:10;"`
	Env            string `json:"env" form:"env" gorm:"column:env;comment:环境变量,local dev test pre online;size:50;"`
	Ext            string `json:"ext" form:"ext" gorm:"column:ext;comment:自定义配置信息;size:255;"`
	InnerIp        string `json:"innerIp" form:"innerIp" gorm:"column:inner_ip;comment:内容IP;size:50;"`
	Name           string `json:"name" form:"name" gorm:"column:name;comment:名称;size:50;"`
	OutIp          string `json:"outIp" form:"outIp" gorm:"column:out_ip;comment:外部IP;size:255;"`
	Platform       string `json:"platform" form:"platform" gorm:"column:platform;comment:平台类型1自有2阿里3腾讯4华为;size:50;"`
	Price          *int   `json:"price" form:"price" gorm:"column:price;comment:价格;size:10;"`
	StartTime      *int   `json:"startTime" form:"startTime" gorm:"column:start_time;comment:开始时间;size:10;"`
	Status         *bool  `json:"status" form:"status" gorm:"column:status;comment:状态1正常2关闭3异常;"`
}

// TableName Server 表名
func (Server) TableName() string {
	return "server"
}
