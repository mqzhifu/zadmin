// 自动生成模板Instance
package op

// Instance 结构体
// 如果含有time.Time 请自行import time包
type Instance struct {
	DIY_MODEL
	ChargeUserName string `json:"chargeUserName" form:"chargeUserName" gorm:"column:charge_user_name;comment:负责人姓名;size:50;"`
	EndTime        *int   `json:"endTime" form:"endTime" gorm:"column:end_time;comment:结束时间;size:10;"`
	Env            string `json:"env" form:"env" gorm:"column:env;comment:环境变量;size:100;"`
	Ext            string `json:"ext" form:"ext" gorm:"column:ext;comment:自定义配置信息;size:255;"`
	Host           string `json:"host" form:"host" gorm:"column:host;comment:主机地址;size:255;"`
	Name           string `json:"name" form:"name" gorm:"column:name;comment:名称;size:50;"`
	Platform       string `json:"platform" form:"platform" gorm:"column:platform;comment:平台类型1自有2阿里3腾讯4华为;size:50;"`
	Port           string `json:"port" form:"port" gorm:"column:port;comment:主机端口号;size:50;"`
	Price          *int   `json:"price" form:"price" gorm:"column:price;comment:价格;size:10;"`
	Ps             string `json:"ps" form:"ps" gorm:"column:ps;comment:密码;size:100;"`
	StartTime      *int   `json:"startTime" form:"startTime" gorm:"column:start_time;comment:开始时间;size:10;"`
	User           string `json:"user" form:"user" gorm:"column:user;comment:用户名;size:100;"`
}

// TableName Instance 表名
func (Instance) TableName() string {
	return "instance"
}
