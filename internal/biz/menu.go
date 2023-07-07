package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Menu struct {
	Id        int64
	Name      string
	Path      string
	ParentId  int64
	ParentIds string
	Hidden    bool
	Sort      int64
	Component string
	Title     string
	Icon      string
	Children  []Menu
}

type MenuRepo interface {
	GetMenuTree(ctx context.Context) ([]*Menu, error)
	GetMenuAll(ctx context.Context) ([]*Menu, error)
	CreateMenu(ctx context.Context, menu *Menu) (*Menu, error)
	UpdateMenu(ctx context.Context, menu *Menu) error
	DeleteMenu(ctx context.Context, id int64) error
	SaveRoleMenu(ctx context.Context, roleId int64, menuIds []int64) error
	GetRoleMenu(ctx context.Context, role string) ([]*Menu, error)
	GetRoleMenuTree(ctx context.Context, role string) ([]*Menu, error)
}

type MenuUseCase struct {
	repo   MenuRepo
	logger *log.Helper
}

func NewMenuUseCase(repo MenuRepo, logger log.Logger) *MenuUseCase {
	return &MenuUseCase{repo: repo, logger: log.NewHelper(log.With(logger, "module", "UseCase/menu"))}
}

func (c *MenuUseCase) CreateMenu(ctx context.Context, data *Menu) (*Menu, error) {
	result, err := c.repo.CreateMenu(ctx, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *MenuUseCase) DeleteMenu(ctx context.Context, id int64) error {
	return c.repo.DeleteMenu(ctx, id)
}

func (c *MenuUseCase) UpdateMenu(ctx context.Context, data *Menu) error {
	return c.repo.UpdateMenu(ctx, data)
}

func (c *MenuUseCase) GetMenuTree(ctx context.Context) ([]*Menu, error) {
	result, err := c.repo.GetMenuTree(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *MenuUseCase) GetMenuAll(ctx context.Context) ([]*Menu, error) {
	result, err := c.repo.GetMenuAll(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *MenuUseCase) SaveRoleMenu(ctx context.Context, roleId int64, menuIds []int64) error {
	return c.repo.SaveRoleMenu(ctx, roleId, menuIds)
}

func (c *MenuUseCase) GetRoleMenu(ctx context.Context, role string) ([]*Menu, error) {
	result, err := c.repo.GetRoleMenu(ctx, role)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *MenuUseCase) GetRoleMenuTree(ctx context.Context, role string) ([]*Menu, error) {
	result, err := c.repo.GetRoleMenuTree(ctx, role)
	if err != nil {
		return nil, err
	}
	return result, nil
}
