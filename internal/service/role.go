package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "sns-authorization/api/authorization/v1"
)

func (s *AuthorizationService) GetRoleAll(ctx context.Context, req *emptypb.Empty) (*v1.GetRoleAllReply, error) {
	return &v1.GetRoleAllReply{}, nil
}
func (s *AuthorizationService) GetRoleList(ctx context.Context, req *v1.GetRoleListReq) (*v1.GetRoleListReply, error) {
	return &v1.GetRoleListReply{}, nil
}
func (s *AuthorizationService) CreateRole(ctx context.Context, req *v1.CreateRoleReq) (*v1.Role, error) {
	return &v1.Role{}, nil
}
func (s *AuthorizationService) UpdateRole(ctx context.Context, req *v1.UpdateRoleReq) (*v1.CheckReply, error) {
	return &v1.CheckReply{}, nil
}
func (s *AuthorizationService) DeleteRole(ctx context.Context, req *v1.DeleteRoleReq) (*v1.CheckReply, error) {
	return &v1.CheckReply{}, nil
}
