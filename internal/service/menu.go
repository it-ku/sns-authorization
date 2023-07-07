package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "sns-authorization/api/authorization/v1"
)

func (s *AuthorizationService) GetMenuAll(ctx context.Context, req *emptypb.Empty) (*v1.GetMenuTreeReply, error) {
	return &v1.GetMenuTreeReply{}, nil
}
func (s *AuthorizationService) GetMenuTree(ctx context.Context, req *emptypb.Empty) (*v1.GetMenuTreeReply, error) {
	return &v1.GetMenuTreeReply{}, nil
}
func (s *AuthorizationService) CreateMenu(ctx context.Context, req *v1.CreateMenuReq) (*v1.Menu, error) {
	return &v1.Menu{}, nil
}
func (s *AuthorizationService) UpdateMenu(ctx context.Context, req *v1.UpdateMenuReq) (*v1.CheckReply, error) {
	return &v1.CheckReply{}, nil
}
func (s *AuthorizationService) DeleteMenu(ctx context.Context, req *v1.IdReq) (*v1.CheckReply, error) {
	return &v1.CheckReply{}, nil
}
