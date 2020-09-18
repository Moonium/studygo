package conf

// AppConf 全局配置文件
type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}

// KafkaConf 配置文件结构体
type KafkaConf struct {
	Address     string `ini:"address"`
	Topic       string `ini:"topic"`
	ChanMaxSize int    `int:"chan_max_size"`
}

// EtcdConf 配置文件结构体
type EtcdConf struct {
	Address string `ini:"address"`
	Key     string `ini:"collect_log_key"`
	Timeout int    `ini:"timeout"`
}

// TaillogConf 配置文件结构体
type TaillogConf struct {
	FileName string `ini:"filename"`
}
