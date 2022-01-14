package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql" // 引入GORM的mysql驱动
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"kratos_study/internal/biz"
	"kratos_study/internal/conf"
	"time"
)

// ProviderSet is data providers.
//var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)
var ProviderSet = wire.NewSet(NewData, NewDatabase, NewWechatUserRepo)

// Data 封装数据库客户端
type Data struct {
	// TODO wrapped database client
	Database *gorm.DB
}

// NewData 在NewData函数中添加*gorm.DB参数，使得我们创建出来的Data结构体中带有数据库客户端
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		Database: db,
	}, cleanup, nil
}

// NewDatabase 初始化数据库
func NewDatabase(c *conf.Data) (*gorm.DB, error) {
	// dsn 数据库连接
	// "用户名":"密码"@tcp("IP":"端口")/"数据库名称"?charset=utf8mb4&parseTime=True&loc=Local
	//dsn := "fengniao:fengniao123@tcp(172.16.151.61:3306)/fn_sigma?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "fengniao:fengniao123@tcp(172.16.151.61:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			// 更改创建时间使用的函数
			NowFunc: func() time.Time {
				return time.Now().Local()
			},
			NamingStrategy: schema.NamingStrategy{
				//TablePrefix: "t_",   // 表名前缀，`User`表为`t_users`
				SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
			},
			// 自动创建数据库外键约束
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	// 设置连接池
	// 空闲
	sqlDb.SetMaxIdleConns(50)
	// 打开
	sqlDb.SetMaxOpenConns(100)
	// 超时
	sqlDb.SetConnMaxLifetime(time.Second * 30)

	err = DBAutoMigrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// DBAutoMigrate 数据库模型自动迁移
func DBAutoMigrate(db *gorm.DB) error {
	// 在这里让GORM知道那些结构体是我们的数据模型，GORM将完成自动建表
	err := db.AutoMigrate(
		&biz.WechatUser{}, // 加入用户数据模型
	)
	if err != nil {
		return err
	}
	return nil
}
