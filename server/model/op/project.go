// 自动生成模板Project
package op

// Project 结构体
// 如果含有time.Time 请自行import time包
type Project struct {
	DIY_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:名称;size:50;"`
	Type      *int   `json:"type" form:"type" gorm:"column:type;comment:类型,1service 2frontend 3backend 4app;"`
	Access    string `json:"access" form:"access" gorm:"column:access;comment:baseAuth 认证KEY;size:255;"`
	SecretKey string `json:"secretKey" form:"secretKey" gorm:"column:secret_key;comment:密钥;size:100;"`
	Status    *int   `json:"status" form:"status" gorm:"column:status;comment:状态1正常2关闭;"`
	Desc      string `json:"desc" form:"desc" gorm:"column:desc;comment:描述信息;size:255;"`
	Git       string `json:"git" form:"git" gorm:"column:git;comment:git仓地址;size:255;"`
}

// TableName Project 表名
func (Project) TableName() string {
	return "project"
}
