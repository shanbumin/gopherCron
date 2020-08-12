package sqlStore

import (
	"fmt"
	"github.com/holdno/gopherCron/pkg/store"
	"time"
	"github.com/holdno/gopherCron/common"
	"github.com/holdno/gopherCron/utils"
	"github.com/holdno/gocommons/selection"
)


//----------------------  userStore结构体 -------------------------------------------------------------------------------

//gc_user表
type userStore struct {
	commonFields
}

//自动创建表
//todo 表已经存在则不会进行创建表的额
//@reviser  sam@2020-07-22 16:18:40
func (s *userStore) AutoMigrate() {
	if err := s.GetMaster().Table(s.GetTable()).AutoMigrate(&common.User{}).Error; err != nil {
		panic(fmt.Errorf("unable to auto migrate %s, %w", s.GetTable(), err))
	}
	s.provider.Logger().Infof("%s, complete initialization", s.GetTable())
}

func (s *userStore) DeleteUser(id int64) error {
	if err := s.GetMaster().Table(s.GetTable()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}

//创建超级管理员
//@reviser sam@2020-07-22 17:15:23
func (s *userStore) CreateAdminUser() error {
	var (
		salt string
		err  error
	)
	salt = utils.RandomStr(6)

	user := &common.User{
		Account:    common.ADMIN_USER_ACCOUNT, //admin
		Password:   utils.BuildPassword(common.ADMIN_USER_PASSWORD, salt),
		Salt:       salt,
		Name:       common.ADMIN_USER_NAME, //administrator
		Permission: common.ADMIN_USER_PERMISSION, //"admin,user"
		CreateTime: time.Now().Unix(),
	}

	if err = s.GetMaster().Table(s.GetTable()).Create(user).Error; err != nil {
		return err
	}

	return nil
}

//获取超管admin信息
//@reviser sam@2020-07-22 17:11:13
func (s *userStore) GetAdminUser() (*common.User, error) {
	var (
		err error
		res common.User
	)
	if err = s.GetReplica().Table(s.GetTable()).Where("account = ?", common.ADMIN_USER_ACCOUNT).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

//创建新用户
func (s *userStore) CreateUser(user common.User) error {
	if err := s.GetMaster().Table(s.GetTable()).Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (s *userStore) ChangePassword(uid int64, password, salt string) error {
	if err := s.GetMaster().Table(s.GetTable()).Where("id = ?", uid).Updates(map[string]string{
		"password": password,
		"salt":     salt,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (s *userStore) GetUsers(selector selection.Selector) ([]*common.User, error) {
	var (
		err error
		res []*common.User
	)

	db := parseSelector(s.GetReplica(), selector, true)

	if err = db.Table(s.GetTable()).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

//--------------------------------------------------------------------------------------------------
//NewProjectStore
func NewUserStore(provider SqlProviderInterface) store.UserStore {
	repo := &userStore{}
	repo.SetProvider(provider)
	repo.SetTable("gc_user")
	return repo
}

