package op

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type BusinessDb struct {
	Gorm *gorm.DB
}

func NewBusinessDb() *BusinessDb {
	businessDb := new(BusinessDb)
	businessDb.Gorm = nil
	return businessDb
}

func (businessDb *BusinessDb) GetDbInc() (gormDb *gorm.DB) {
	if businessDb.Gorm != nil {
		return businessDb.Gorm
	}

	businessDb.Init()
	return businessDb.Gorm
}

func (businessDb *BusinessDb) Init() (gormDb *gorm.DB, err error) {
	m := global.GVA_CONFIG.Mysql
	dns := "root" + ":" + "mqzhifu" + "@tcp(" + "8.142.177.235" + ":" + m.Port + ")/" + "test" + "?" + m.Config
	fmt.Println("mysql link info: " + dns)
	//dns := m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dns,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	diyConfig := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true, NamingStrategy: schema.NamingStrategy{SingularTable: true}}
	db, err := gorm.Open(mysql.New(mysqlConfig), diyConfig)
	fmt.Println("gorm.Open er  info:", err)
	if err != nil {
		return nil, err
	}
	businessDb.Gorm = db
	//db = db.Debug()
	//sqlDB, _ := db.DB()
	//sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	//sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	return db, nil
}

var myBusinessDb = NewBusinessDb()
