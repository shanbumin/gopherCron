package store

import (
	"github.com/holdno/gopherCron/common"

	"github.com/holdno/gocommons/selection"
	"github.com/jinzhu/gorm"
)

//todo 这里就是针对初始化的几个表对应结构体的接口声明
//因为初始化的时候可能使用mysql，也可以使用mongodb的额

// Commons defined func which can be used by other stores
type Commons interface {
	GetMap(selector selection.Selector) ([]map[string]interface{}, error)
	GetTable() string
	GetTotal(selector selection.Selector) (int, error)
	CheckSelf()
	AutoMigrate()
}


//下面是几张必须初始化的表


type ProjectStore interface {
	Commons
	CreateProject(tx *gorm.DB, obj common.Project) (int64, error)
	UpdateProject(id int64, title, remark string) error
	GetProject(selector selection.Selector) ([]*common.Project, error)
	DeleteProject(tx *gorm.DB, selector selection.Selector) error
	UpdateRelation(projectID int64, relation string) error
}

type ProjectRelevanceStore interface {
	Commons
	Create(tx *gorm.DB, r common.ProjectRelevance) error
	Delete(tx *gorm.DB, pid, uid int64) error
	GetList(selector selection.Selector) ([]*common.ProjectRelevance, error)
}

type UserStore interface {
	Commons
	CreateAdminUser() error
	DeleteUser(id int64) error
	GetAdminUser() (*common.User, error)
	CreateUser(user common.User) error
	ChangePassword(uid int64, password, salt string) error
	GetUsers(selector selection.Selector) ([]*common.User, error)
}

type TaskLogStore interface {
	Commons
	CreateTaskLog(data common.TaskLog) error
	GetList(selector selection.Selector) ([]*common.TaskLog, error)
	Clean(tx *gorm.DB, selector selection.Selector) error
}
