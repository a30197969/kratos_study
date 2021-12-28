package service

import (
	"context"

	pb "kratos_study/api/sigma/v1"
)

type AlbumService struct {
	pb.UnimplementedAlbumServer
}

func NewAlbumService() *AlbumService {
	return &AlbumService{}
}

func (s *AlbumService) CreateAlbum(ctx context.Context, req *pb.CreateAlbumRequest) (*pb.CreateAlbumReply, error) {
	return &pb.CreateAlbumReply{}, nil
}
func (s *AlbumService) UpdateAlbum(ctx context.Context, req *pb.UpdateAlbumRequest) (*pb.UpdateAlbumReply, error) {
	return &pb.UpdateAlbumReply{}, nil
}
func (s *AlbumService) DeleteAlbum(ctx context.Context, req *pb.DeleteAlbumRequest) (*pb.DeleteAlbumReply, error) {
	return &pb.DeleteAlbumReply{}, nil
}
func (s *AlbumService) GetAlbum(ctx context.Context, req *pb.GetAlbumRequest) (*pb.GetAlbumReply, error) {
	return &pb.GetAlbumReply{}, nil
}
func (s *AlbumService) ListAlbum(ctx context.Context, req *pb.ListAlbumRequest) (*pb.ListAlbumReply, error) {
	return &pb.ListAlbumReply{}, nil
}
