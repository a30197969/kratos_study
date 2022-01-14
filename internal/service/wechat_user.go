package service

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"kratos_study/internal/biz"

	pb "kratos_study/api/sigma/v1"
)

type WechatUserService struct {
	pb.UnimplementedWechatUserServer
	// uc 用户操作的封装，在“/biz/wechat_user.go”中
	uc  *biz.WechatUserUsecase
	log *log.Helper
}

func NewWechatUserService(uc *biz.WechatUserUsecase, logger log.Logger) *WechatUserService {
	return &WechatUserService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *WechatUserService) CreateWechatUser(ctx context.Context, req *pb.CreateWechatUserRequest) (*pb.CreateWechatUserReply, error) {

	// 获取前端发送的Body数据
	body := req.GetCreatBody()
	// 创建用户
	id, err := s.uc.CreateUser(ctx, body.GetOpenid(), body.GetNickname(), body.GetUnionid(), body.GetAppid())
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(id)
	if err != nil {
		return nil, err
	}
	// 给前端返回数据
	return &pb.CreateWechatUserReply{
		Code:    0,
		Message: "成功",
		Data:    string(data),
	}, nil
}
func (s *WechatUserService) GetWechatUser(ctx context.Context, req *pb.GetWechatUserRequest) (*pb.GetWechatUserReply, error) {
	return &pb.GetWechatUserReply{}, nil
}
func (s *WechatUserService) ListWechatUser(ctx context.Context, req *pb.ListWechatUserRequest) (*pb.ListWechatUserReply, error) {
	return &pb.ListWechatUserReply{}, nil
}
