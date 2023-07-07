package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sns-authorization/internal/biz"
)

type RoleEntity struct {
	Id     int64  `gorm:"primarykey;type:int;comment:主键id"`
	Domain string `gorm:"type:varchar(255);not null;comment:所在域"`
	Name   string `gorm:"type:varchar(255);not null;comment:名称"`
}

func (RoleEntity) TableName() string {
	return "roles"
}

type RoleRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &RoleRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/role")),
	}
}

func (repo RoleRepo) ListRoleAll(ctx context.Context) ([]*biz.Role, error) {
	return nil, nil
}

func (repo RoleRepo) CreateRole(ctx context.Context, role *biz.Role) (*biz.Role, error) {
	return nil, nil
}

func (repo RoleRepo) GetRole(ctx context.Context, params map[string]interface{}) (*biz.Role, error) {
	return nil, nil
}

func (repo RoleRepo) UpdateRole(ctx context.Context, role *biz.Role) error {
	return nil
}

func (repo RoleRepo) DeleteRole(ctx context.Context, role *biz.Role) error {
	return nil
}
