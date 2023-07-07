package data

import (
	"context"

	"sns-authorization/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type ApiEntity struct {
	Id     int64  `gorm:"primarykey;type:int;comment:主键id"`
	Domain string `gorm:"type:varchar(255);not null;comment:所在域"`
	Name   string `gorm:"type:varchar(255);not null;comment:名称"`
	Group  string `gorm:"type:varchar(255);not null;comment:分组"`
	Method string `gorm:"type:varchar(255);not null;comment:请求方式"`
	Path   string `gorm:"type:varchar(255);not null;comment:请求路径"`
}

func (ApiEntity) TableName() string {
	return "apis"
}

type ApiRepo struct {
	data *Data
	log  *log.Helper
}

func NewApiRepo(data *Data, logger log.Logger) biz.ApiRepo {
	return &ApiRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/api")),
	}
}

func (repo ApiRepo) ListApiAll(ctx context.Context) ([]*biz.Api, error) {
	return nil, nil
}

func (repo ApiRepo) ListApi(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*biz.Api, int64, error) {
	return nil, 0, nil
}

func (repo ApiRepo) CreateApi(ctx context.Context, Api *biz.Api) (*biz.Api, error) {
	return nil, nil
}

func (repo ApiRepo) UpdateApi(ctx context.Context, Api *biz.Api) error {
	return nil
}

func (repo ApiRepo) DeleteApi(ctx context.Context, Api *biz.Api) error {
	return nil
}
