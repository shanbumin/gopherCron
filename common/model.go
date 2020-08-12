package common



/*
CREATE TABLE `gc_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL COMMENT '用户名称',
  `permission` varchar(100) NOT NULL COMMENT '用户权限',
  `account` varchar(100) NOT NULL COMMENT '用户账号',
  `password` varchar(255) NOT NULL COMMENT '用户密码',
  `salt` varchar(6) NOT NULL COMMENT '密码盐',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `name` (`name`),
  KEY `account` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

 */
type User struct {
	ID         int64  `json:"id" gorm:"column:id;pirmary_key;auto_increment"`
	Name       string `json:"name" gorm:"column:name;index:name;type:varchar(100);not null;comment:'用户名称'"`
	Permission string `json:"permission" gorm:"column:permission;type:varchar(100);not null;comment:'用户权限'"`
	Account    string `json:"account" gorm:"column:account;index:account;type:varchar(100);not null;comment:'用户账号'"`
	Password   string `json:"-" gorm:"password;type:varchar(255);not null;comment:'用户密码'"`
	Salt       string `json:"-" gorm:"salt;type:varchar(6);not null;comment:'密码盐'"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time;type:bigint(20);not null;comment:'创建时间'"`
}


/*
CREATE TABLE `gc_project` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL COMMENT '关联用户id',
  `title` varchar(100) NOT NULL COMMENT '项目名称',
  `remark` varchar(255) NOT NULL COMMENT '项目备注',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `title` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/
type Project struct {
	ID     int64  `json:"id" gorm:"column:id;pirmary_key;auto_increment"`
	UID    int64  `json:"uid" gorm:"column:uid;index:uid;type:bigint(20);not null;comment:'关联用户id'"`
	Title  string `json:"title" gorm:"column:title;index:title;type:varchar(100);not null;comment:'项目名称'"`
	Remark string `json:"remark" gorm:"column:remark;type:varchar(255);not null;comment:'项目备注'"`
}

/*

CREATE TABLE `gc_project_relevance` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL COMMENT '关联用户id',
  `project_id` bigint(20) NOT NULL COMMENT '关联项目id',
  `create_time` bigint(20) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `project_id` (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

 */
type ProjectRelevance struct {
	ID         int64 `json:"id" gorm:"column:id;pirmary_key;auto_increment"`
	UID        int64 `json:"uid" gorm:"column:uid;index:uid;type:bigint(20);not null;comment:'关联用户id'"`
	ProjectID  int64 `json:"project_id" gorm:"column:project_id;index:project_id;type:bigint(20);not null;comment:'关联项目id'"`
	CreateTime int64 `json:"create_time" gorm:"column:create_time;type:bigint(20);not null;comment:'创建时间'"`
}

/*

CREATE TABLE `gc_task_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `project_id` bigint(20) NOT NULL COMMENT '关联项目id',
  `task_id` bigint(20) NOT NULL COMMENT '关联任务id',
  `project` varchar(100) NOT NULL COMMENT '项目名称',
  `name` varchar(100) NOT NULL COMMENT '任务名称',
  `result` varchar(20) NOT NULL COMMENT '任务执行结果',
  `start_time` bigint(20) NOT NULL COMMENT '任务开始时间',
  `end_time` bigint(20) NOT NULL COMMENT '任务结束时间',
  `command` varchar(255) NOT NULL COMMENT '任务指令',
  `with_error` int(11) NOT NULL COMMENT '是否发生错误',
  `client_ip` varchar(20) NOT NULL COMMENT '节点ip',
  PRIMARY KEY (`id`),
  KEY `client_ip` (`client_ip`),
  KEY `project_id` (`project_id`),
  KEY `task_id` (`task_id`),
  KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/
type TaskLog struct {
	ID        int64  `json:"id" gorm:"column:id;pirmary_key;auto_increment"`
	ProjectID int64  `json:"project_id" gorm:"column:project_id;index:project_id;type:bigint(20);not null;comment:'关联项目id'"`
	TaskID    string `json:"task_id" gorm:"column:task_id;index:task_id;type:bigint(20);not null;comment:'关联任务id'"`
	Project   string `json:"project" gorm:"column:project;type:varchar(100);not null;comment:'项目名称'"`

	Name      string `json:"name" gorm:"column:name;index:name;type:varchar(100);not null;comment:'任务名称'"`
	Result    string `json:"result" gorm:"column:result;type:varchar(20);not null;comment:'任务执行结果'"`
	StartTime int64  `json:"start_time" gorm:"column:start_time;type:bigint(20);not null;comment:'任务开始时间'"`
	EndTime   int64  `json:"end_time" gorm:"column:end_time;type:bigint(20);not null;comment:'任务结束时间'"`
	Command   string `json:"command" gorm:"column:command;type:varchar(255);not null;comment:'任务指令'"`
	WithError int    `json:"with_error" gorm:"column:with_error;type:int(11);not null;comment:'是否发生错误'"`
	ClientIP  string `json:"client_ip" gorm:"client_ip;index:client_ip;type:varchar(20);not null;comment:'节点ip'"`
}

// MonitorInfo 监控信息
type MonitorInfo struct {
	IP            string `json:"ip"`
	CpuPercent    string `json:"cpu_percent"`
	MemoryPercent string `json:"memory_percent"`
	MemoryTotal   uint64 `json:"memory_total"`
	MemoryFree    uint64 `json:"memory_free"`
}
