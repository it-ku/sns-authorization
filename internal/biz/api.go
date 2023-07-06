package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)


type Api struct {
	Id        int64
	Domain    string
	Name      string
	Group     string
	Method    string
	Path      string
}

type ApiRepo interface {
	ListApiAll(ctx context.Context) ([]*Api, error)
	ListApi(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*Api, int64, error)
	CreateApi(ctx context.Context, Api *Api) (*Api, error)
	UpdateApi(ctx context.Context, Api *Api) error
	DeleteApi(ctx context.Context, Api *Api) error
}

type ApiUseCase struct {
	repo ApiRepo
	log  *log.Helper
}

func NewApiUseCase(repo ApiRepo, logger log.Logger) *ApiUseCase {
	return &ApiUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c *ApiUseCase) ListApiAll(ctx context.Context) ([]*Api, error) {
	return c.repo.ListApiAll(ctx)
}

func (c *ApiUseCase) ListApi(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*Api, int64, error) {
	return c.repo.ListApi(ctx, page, pageSize, params)
}

func (c *ApiUseCase) CreateApi(ctx context.Context, api *Api) (*Api, error) {
	return c.repo.CreateApi(ctx, api)
}

func (c *ApiUseCase) UpdateApi(ctx context.Context, api *Api) error {
	return c.repo.UpdateApi(ctx, api)
}

func (c *ApiUseCase) DeleteApi(ctx context.Context, api *Api) error {
	return c.repo.DeleteApi(ctx, api)
}