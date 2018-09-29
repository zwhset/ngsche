package types

// nginx.conf include /.../conf.d/upstreams.conf -> Nodes info
// include vhosts/*.conf vhosts == Server

type nginxConf struct {
	User            string `json:"user"`
	WorkerProcesses string `json:"worker_processes"`
	ErrorLog        string `json:"error_log"`
	Pid             string `json:"pid"`
	Include         string `json:"include"`
	Events          Events `json:"events"`
	Http            HTTP   `json:"http"`
}

type Events struct {
	WorkerConnections int `json:"worker_connections"`
}

type HTTP struct {
	LogFormat        string `json:"log_format"`
	AccessLog        string `json:"access_log"`
	DefaultType      string `json:"default_type"`
	KeepaliveTimeout int    `json:"keepalive_timeout"`

	// include Upstream
	Include string `json:"include"`
}

type Server struct {
	Listen     int    `json:"listen"`
	Root       string `json:"root"`
	ServerName string `json:"server_name"`
	Index      string `json:"index"`

	Location Location `json:"location"`
}

type Location struct {
	// proxy_pass http://{{Upstream.Name}}
	ProxyPass     string `json:"proxy_pass"`
	IsHealthCheck string `json:"is_health_check"`
}

type Upstream []Node

type Node struct {
	// 服务器信息
	Server string `json:"server"`
	// 端口
	Port int `json:"port"`
	// 服务器权重
	Weight int `json:"weight"`
	// 失败重试次数
	FailTimeout int `json:"fail_timeout"`
	// 当server恢复健康状态时多久后启用
	SLowStart int `json:"s_low_start"`
	// 是否是记录类型 DNS
	IsResolve bool `json:"resolve"`
	// 是否是备份节点
	IsBackup bool `json:"backup"`
	// 是否标为不可用节点
	IsDown bool `json:"down"`
}
