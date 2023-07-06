package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type RoleRepo interface {
	ListRoleAll(ctx context.Context) ([]*Role, error)
	CreateRole(ctx context.Context, role *Role) (*Role, error)
	GetRole(ctx context.Context, params map[string]interface{}) (*Role, error)
	UpdateRole(ctx context.Context, role *Role) error
	DeleteRole(ctx context.Context, role *Role) error
}

type Role struct {
	Id        int64
	Name      string
	Domain    string
	CreatedAt string
	UpdatedAt string
}


type RoleUseCase struct {
	repo   RoleRepo
	logger *log.Helper
}

func NewRoleUseCase(repo RoleRepo, logger log.Logger) *RoleUseCase {
	return &RoleUseCase{repo: repo, logger: log.NewHelper(log.With(logger, "module", "UseCase/role"))}
}

func (c *RoleUseCase) ListRoleAll(ctx context.Context) ([]*Role, error) {
	return c.repo.ListRoleAll(ctx)
}

func (c *RoleUseCase) CreateRole(ctx context.Context, role *Role) (*Role, error) {
	return c.repo.CreateRole(ctx, role)
}

func (c *RoleUseCase) UpdateRole(ctx context.Context, role *Role) error {
	return c.repo.UpdateRole(ctx, role)
}

func (c *RoleUseCase) DeleteRole(ctx context.Context, role *Role) error {
	return c.repo.DeleteRole(ctx, role)
}
