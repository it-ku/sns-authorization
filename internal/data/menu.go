package data

type MenuEntity struct {
	Id        int64        `gorm:"primarykey;type:int;comment:主键id"`
	Domain    string       `gorm:"type:varchar(255);not null;comment:所在域"`
	ParentId  int64        `gorm:"type:int;comment:父级id"`
	ParentIds string       `gorm:"type:int;comment:父级id字符串 英文逗号分割"`
	Name      string       `gorm:"type:varchar(255);not null;comment:菜单名"`
	Path      string       `gorm:"type:varchar(255);not null;comment:前端路径"`
	Hidden    bool         `gorm:"not null;comment:是否隐藏 0否1是"`
	Component string       `gorm:"type:varchar(255);not null;comment:前端文件路径"`
	Sort      int64        `gorm:"type:int;comment:排序"`
	Title     string       `gorm:"type:varchar(255);not null;comment:页面名称"`
	Icon      string       `gorm:"type:varchar(255);not null;comment:菜单图标"`
	Children  []MenuEntity `gorm:"-"`
}

func (MenuEntity) TableName() string {
	return "menus"
}
