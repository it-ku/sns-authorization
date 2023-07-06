package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "sns-authorization/api/authorization/v1"
)

func (s *AuthorizationService) GetApiListAll(ctx context.Context, req *emptypb.Empty) (*v1.GetApiListAllReply, error) {
	return &v1.GetApiListAllReply{}, nil
}
func (s *AuthorizationService) GetApiList(ctx context.Context, req *v1.GetApiListReq) (*v1.GetApiListReply, error) {
	return &v1.GetApiListReply{}, nil
}
func (s *AuthorizationService) CreateApi(ctx context.Context, req *v1.CreateApiReq) (*v1.Api, error) {
	return &v1.Api{}, nil
}
func (s *AuthorizationService) UpdateApi(ctx context.Context, req *v1.UpdateApiReq) (*v1.CheckReply, error) {
	return &v1.CheckReply{}, nil
}
func (s *AuthorizationService) DeleteApi(ctx context.Context, req *v1.DeleteApiReq) (*v1.CheckReply, error) {
	return &v1.CheckReply{}, nil
}
