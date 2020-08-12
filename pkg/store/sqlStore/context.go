package sqlStore

import (
	"github.com/holdno/gocommons/selection"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//-----------------------  SqlProviderInterface 接口 -------------------------------
//provider.go中存在实现该接口的结构体
//@reviser sam@2020-07-22 13:38:25
type SqlProviderInterface interface {
	GetMaster() *gorm.DB  //主引擎
	GetReplica() *gorm.DB  //副本引擎们
	Logger() *logrus.Logger //日志
}
//----------------------  commonFields  结构体 --------------------------------------------------------------
//每张表对应的结构体公共继承的结构体
type commonFields struct {
	provider SqlProviderInterface
	table    string
}

func (c *commonFields) SetProvider(p SqlProviderInterface) {
	c.provider = p
}

func (c *commonFields) GetTable() string {
	return c.table
}

func (c *commonFields) SetTable(table string) {
	c.table = table
}
func (c *commonFields) GetMap(selector selection.Selector) ([]map[string]interface{}, error) {
	db := parseSelector(c.GetReplica(), selector, true)
	rows, err := db.Table(c.GetTable()).Rows()
	if err != nil {
		return nil, err
	}
	res, err := scanToMap(rows)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return res, nil
}

func (c *commonFields) GetTotal(selector selection.Selector) (int, error) {
	var (
		err   error
		total int
	)

	db := parseSelector(c.GetReplica(), selector, false)

	if err = db.Table(c.GetTable()).Count(&total).Error; err != nil {
		return total, err
	}

	return total, nil
}

func (c *commonFields) GetMaster() *gorm.DB {
	return c.provider.GetMaster()
}

func (c *commonFields) GetReplica() *gorm.DB {
	return c.provider.GetReplica()
}

//检测每个表对应结构体是否准备ok了
//@todo 只要provider以及table不为空就算过关了
//@reviser sam@2020-07-22 16:00:01
func (c *commonFields) CheckSelf() {
	if c.provider == nil {
		panic("can not found provider")
	}

	if c.table == "" {
		panic("can not set table")
	}

	c.provider.Logger().Infof("store %s is ok!", c.GetTable())
}

