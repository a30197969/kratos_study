package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// WechatUser 用户数据模型
type WechatUser struct {
	// GORM的一个基础数据模型
	// 其中有ID、CreatedAt(创建时间)、UpdatedAt(更新时间)、DeletedAt(删除时间)
	// gorm.Model

	ID uint64 `gorm:"primaryKey"`
	// Openid 用户的唯一标识
	Openid string `gorm:"not null;type:varchar(50);index"`
	// WxID 微信号
	WxID          string `gorm:"type:varchar(50)"`
	Level         uint8  `gorm:"not null;default:1"`
	Phone         string `gorm:"type:varchar(20);index"`
	PurePhone     string `gorm:"type:varchar(20)"`
	CountryCode   string `gorm:"type:varchar(10)"`
	Nickname      string `gorm:"not null;type:varchar(50)"`
	Sex           uint8  `gorm:"not null;default:0"`
	Language      string `gorm:"not null;type:varchar(50)"`
	ProvinceCode  string `gorm:"type:varchar(10)"`
	Province      string `gorm:"type:varchar(50)"`
	CityCode      string `gorm:"type:varchar(10)"`
	City          string `gorm:"type:varchar(50)"`
	DistrictCode  string `gorm:"type:varchar(10)"`
	District      string `gorm:"type:varchar(50)"`
	Country       string `gorm:"type:varchar(50)"`
	Headimgurl    string `gorm:"type:varchar(200)"`
	Remark        string `gorm:"type:varchar(200)"`
	Groupid       uint   `gorm:"not null;default:0"`
	Privilege     string `gorm:"type:varchar(250)"`
	Unionid       string `gorm:"type:varchar(50);index"`
	Subscribe     uint   `gorm:"not null;default:0"`
	SubscribeTime uint   `gorm:"not null;default:0"`
	Dateline      string
	IsDel         uint8  `gorm:"not null;default:0"`
	Appid         string `gorm:"not null;type:varchar(50)"`
	PicID         uint   `gorm:"not null;default:0"`
	TagidList     string `gorm:"type:varchar(250)"`
	UpdateDate    string
}

// WechatUserRepo 用户数据模型 在数据库上的基础操作封装为接口
type WechatUserRepo interface {
	CreateUser(ctx context.Context, model *WechatUser) (uint64, error)
	//UpdateUser(ctx context.Context, model *WechatUser) (bool, error)
	//DeleteUser(ctx context.Context, model *WechatUser) (bool, error)
	//SelectUser(ctx context.Context, model *WechatUser) (interface{}, error)

}

// WechatUserUsecase 将对用户的操作封装
type WechatUserUsecase struct {
	repo WechatUserRepo // 我们需要使用Repo来对数据库进行基础操作
	log  *log.Helper
}

func NewWechatUserUsecase(repo WechatUserRepo, logger log.Logger) *WechatUserUsecase {
	return &WechatUserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
func (uc *WechatUserUsecase) CreateUser(ctx context.Context, openid string, nickname string, unionid string, appid string) (uint64, error) {
	t := time.Now()
	dateline := t.Format("2006-01-02 15:04:05")
	wechatUserModel := &WechatUser{
		Openid:   openid,
		Nickname: nickname,
		Unionid:  unionid,
		Appid:    appid,
		Dateline: dateline,
	}
	return uc.repo.CreateUser(ctx, wechatUserModel)
}
