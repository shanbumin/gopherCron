package config

var serviceConf *ServiceConfig

//--------------------------------------------- 配置文件对应结构体 -------------------------------------------------------
// APIConfig 配置文件Root
type ServiceConfig struct {
	Env        string       `toml:"env"`
	LogLevel   string       `toml:"log_level"`
	ReportAddr string       `toml:"report_addr"`
	//######
	Deploy     *DeployConf  `toml:"deploy"` // host配置
	Etcd       *EtcdConf    `toml:"etcd"`
	MongoDB    *MongoDBConf `toml:"mongodb"`
	JWT        *JWTConf     `toml:"jwt"`
	Mysql      *MysqlConf   `toml:"mysql"`
}

// DeployConf 部署配置
type DeployConf struct {
	Environment string   `toml:"environment"` //当前的环境:dev、release
	Timeout     int      `toml:"timeout"`    //秒为单位
	Host        []string `toml:"host"`       //对外提供的端口 ["0.0.0.0:6306"]
	ViewPath    string   `toml:"view_path"`  //前端文件路径 ./view
}

// EtcdConf etcd配置
type EtcdConf struct {
	Service     []string `toml:"service"`  //["0.0.0.0:2379"]
	Username    string   `toml:"username"`
	Password    string   `toml:"password"`
	DialTimeout int      `toml:"dialtimeout"`
	Prefix      string   `toml:"prefix"` //etcd kv存储的key前缀 用来与其他业务做区分,如 /gopher_cron
	Projects    []int64  `toml:"projects,omitempty"`
	Shell       string   `toml:"shell,omitempty"`
}

// MongoDBConf mongodb连接配置
type MongoDBConf struct {
	Service       []string `toml:"service"`
	Username      string   `toml:"username"`
	Password      string   `toml:"password"`
	Table         string   `toml:"table"`
	AuthMechanism string   `toml:"auth_mechanism"`
}

type MysqlConf struct {
	Service  string `toml:"service"`  //"0.0.0.0:3306"
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

// JWTConf 签名方法配置
type JWTConf struct {
	Secret string `toml:"secret"`
	Exp    int    `toml:"exp"` //token 有效期(小时)
}
//----------------------------------------------------------------------------------------------------------------------



// InitServiceConfig 获取api相关配置
//@todo 本质就是将配置文件中的信息加载到serviceConf变量中
// @reviser sam@2020-07-21 10:32:32
func InitServiceConfig(path string) *ServiceConfig {
	if path == "" {
		return nil
	}
	var c ServiceConfig
	LoadFrom(path, &c)
	serviceConf = &c
	return &c
}

// GetServiceConfig 获取服务配置
// @todo 返回serviceConf变量
func GetServiceConfig() *ServiceConfig {
	if serviceConf != nil {
		return serviceConf
	}
	return nil
}
