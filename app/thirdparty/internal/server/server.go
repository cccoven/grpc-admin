package server

import (
	"context"
	"grpc-admin/app/thirdparty/internal/model/migrate"
	"grpc-admin/app/thirdparty/internal/pkg/db"
	"grpc-admin/app/thirdparty/internal/repo"
	"grpc-admin/app/thirdparty/thirdparty"
)

type server struct {
	thirdparty.UnimplementedThirdPartyServer
	repo repo.IThirdPartyRepo
}

func NewThirdPartyServer() *server {
	migrate.Do(db.NewGormDB())
	
	return &server{
		repo: repo.NewThirdPartyRepo(),
	}
}

func (s *server) SendSMS(ctx context.Context, in *thirdparty.SendSMSRequest) (*thirdparty.Empty, error) {
	// TODO 发送短信
	
	return nil, nil
}
