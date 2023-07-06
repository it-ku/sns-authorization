package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "sns-authorization/api/authorization/v1"
	"sns-authorization/internal/biz"
)

type AuthorizationService struct {
	v1.UnimplementedAuthorizationServiceServer
	apiUseCase  *biz.ApiUseCase
	roleUseCase *biz.RoleUseCase
	log         *log.Helper
}

func NewAuthorizationService(
	apiUseCase *biz.ApiUseCase,
	roleUseCase *biz.RoleUseCase,
	logger log.Logger,
) *AuthorizationService {
	return &AuthorizationService{
		log:         log.NewHelper(log.With(logger, "module", "service/authorization")),
		apiUseCase:  apiUseCase,
		roleUseCase: roleUseCase,
	}
}

func (s *AuthorizationService) CheckAuthorization(ctx context.Context, req *v1.CheckAuthorizationReq) (*v1.CheckAuthorizationReply, error) {
	return &v1.CheckAuthorizationReply{}, nil
}
