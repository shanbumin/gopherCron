package sqlStore

import (
	"fmt"
	"reflect"

	"github.com/holdno/gopherCron/common"
	"github.com/holdno/gopherCron/config"
	"github.com/holdno/gopherCron/pkg/store"
	"github.com/holdno/gopherCron/utils"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//-------------------------- SqlStore接口 -------------------------------------------------------------------------------------------
//@todo 就是为安装数据库初始化而准备的接口
//@author sam@2020-08-12 13:48:46
type SqlStore interface {
	User() store.UserStore
	Project() store.ProjectStore
	ProjectRelevance() store.ProjectRelevanceStore
	TaskLog() store.TaskLogStore
	BeginTx() *gorm.DB
	Install()
	Shutdown()
}
//------------------------------------ SqlProvider结构体 ---------------------------------------------------------------
//@todo 这个结构体真伟大，不仅与db打交道，而且还实现了安装接口，什么找它都能拿到额
type SqlProvider struct {
	master   *gorm.DB
	replicas []*gorm.DB
	stores   SqlProviderStores  //默认需要初始化的所有表结构体作为属性组成的结构体
	logger   *logrus.Logger
}

func (p *SqlProvider) Logger() *logrus.Logger {
	return p.logger
}

func (s *SqlProvider) GetMaster() *gorm.DB {
	return s.master
}

func (s *SqlProvider) GetReplica() *gorm.DB {
	return s.replicas[utils.Random(0, len(s.replicas)-1)]
}

//该方法就是批量轮流执行stores们的某个方法吧了
//@reviser  sam@2020-07-22 16:11:47
func (s *SqlProvider) batchExecStoreFuncs(fname string) {
	val := reflect.ValueOf(s.stores)
	//获取所有的stores
	num := val.NumField()
	for i := 0; i < num; i++ {
		//调用每个store的CheckSelf()，不需要传递任何参数，所以用[]reflect.Value{}
		val.Field(i).MethodByName(fname).Call([]reflect.Value{})
	}
}

//统一执行AutoMigrate方法
//@reviser  sam@2020-07-22 16:12:53
func (s *SqlProvider) Install() {
	s.batchExecStoreFuncs("AutoMigrate")
	s.Logger().Info("All stores are installed!")
}

//检测所有的store是否都准备好了
//@reviser sam@2020-07-22 16:07:26
func (s *SqlProvider) CheckStores() {
	s.batchExecStoreFuncs("CheckSelf")
	s.Logger().Info("-----------All stores are ready!------------------")
}


func (s *SqlProvider) User() store.UserStore {
	return s.stores.User
}

func (s *SqlProvider) Project() store.ProjectStore {
	return s.stores.Project
}

func (s *SqlProvider) ProjectRelevance() store.ProjectRelevanceStore {
	return s.stores.ProjectRelevance
}

func (s *SqlProvider) TaskLog() store.TaskLogStore {
	return s.stores.TaskLog
}

func (s *SqlProvider) BeginTx() *gorm.DB {
	return s.GetMaster().Begin()
}

// Shutdown close all connect
func (s *SqlProvider) Shutdown() {
	s.master.Close()
	for _, v := range s.replicas {
		v.Close()
	}
}

//mysql连接建立
//@reviser sam@2020-07-22 11:31:20
func (s *SqlProvider) initConnect(conf *config.MysqlConf) error {
	var (
		err    error
		mc     mysql.Config
		engine *gorm.DB
	)
	mc = mysql.Config{
		User:                 conf.Username,
		Passwd:               conf.Password,
		Net:                  "tcp",
		Addr:                 conf.Service,
		DBName:               conf.Database,
		AllowNativePasswords: true,
	}
	if engine, err = gorm.Open("mysql", mc.FormatDSN()); err != nil {
		return fmt.Errorf("failed to create seller, %w", err)
	}
	if err = engine.DB().Ping(); err != nil {
		return fmt.Errorf("connect to database, but ping was failed, %w", err)
	}

	s.master = engine
	s.replicas = append(s.replicas, engine)

	return nil
}
//---------------------------------------------------------------------------------------------------------------------


type SqlProviderStores struct {
	User             store.UserStore
	Project          store.ProjectStore
	ProjectRelevance store.ProjectRelevanceStore
	TaskLog          store.TaskLogStore
}

//启动mysql
//@reviser sam@2020-07-22 11:13:49
func MustSetup(conf *config.MysqlConf, logger *logrus.Logger, install bool) SqlStore {
	//创建provider
	provider := new(SqlProvider)
	//设置logger
	provider.logger = logger
	//初始化连接,创建引擎  provider.master ...
	if err := provider.initConnect(conf); err != nil {
		logger.Error("init Connect error")
		panic(err)
	}
	//初始化相关基础数据表以及对应的默认数据
	provider.stores.User = NewUserStore(provider)
	provider.stores.Project = NewProjectStore(provider)
	provider.stores.TaskLog = NewTaskLogStore(provider)
	provider.stores.ProjectRelevance = NewProjectRelevanceStore(provider)
	provider.CheckStores() //仅仅只全部检测所有的stores是否准备好了,本质是执行每个表结构提对应的CheckSelf方法而已
	//install为真
	if install {
		provider.logger.Info("start install database ...")
		provider.Install() //真正的安装在这里,本质是执行每个表结构体对应的AutoMigrate方法而已
		provider.logger.Info("------------------------finish--------------------------------")
	}
	//检测是否需要创建管理员
	admin, err := provider.stores.User.GetAdminUser()
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}
	if admin == nil {
		provider.logger.Info("start create admin user")
		if err = provider.stores.User.CreateAdminUser(); err != nil {
			panic(err)
		}
		provider.logger.WithFields(logrus.Fields{
			"account":  common.ADMIN_USER_ACCOUNT,
			"password": common.ADMIN_USER_PASSWORD,
		}).Info("admin user created")
	}
	return provider
}





