package setting

import (
	"time"
)

type Server struct {
	HttpPort          int
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	WsPort            int
	WsPongWait        time.Duration
	WsReadBufferSize  int
	WsWriteBufferSize int
	WsWriteWait       time.Duration
	WsMaxMessageSize  int64
	IntranetIp        string
	MainRedisMode     string
	TeamRedisMode     string
}

var ServerSetting = &Server{}

type MySql struct {
	User        string
	Password    string
	Host        string
	Port        int
	DbName      string
	Timeout     string
	MaxConn     int
	MaxIdleConn int
}

var DatabaseSetting = &MySql{}

type Mongo struct {
	ApplyURI        string
	Hosts           string
	MaxConnIdleTime uint
	MaxPoolSize     uint64
	MinPoolSize     uint64
}
