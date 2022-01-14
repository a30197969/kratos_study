package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos_study/internal/biz"
)

// 实现 biz.WechatUserRepo 接口，完成对 用户数据模型 的数据库基础操作
type wechatUserRepo struct {
	data *Data // 数据库客户端的集合
	log  *log.Helper
}

// NewWechatUserRepo 创建结构体wechatUserRepo
func NewWechatUserRepo(data *Data, logger log.Logger) biz.WechatUserRepo {
	return &wechatUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateUser 创建用户的数据库操作
func (r *wechatUserRepo) CreateUser(ctx context.Context, model *biz.WechatUser) (uint64, error) {
	err := r.data.Database.Create(model).Error
	if err != nil {
		return 0, err
	}
	return model.ID, nil
}
