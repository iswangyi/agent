package settings

import (
	"agent/logger"
	"agent/utils"
	"encoding/json"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type GlobalConfig struct {
	Debug     bool             `json:"debug"`
	HostName  string           `json:"host_name"`
	Ip        string           `json:"ip"`
	Pid       string           `json:"pid"`
	Logfile   string           `json:"logfile"`
	Transfer  *TransferConfig  `json:"transfer"`
	Http      *HttpConfig      `json:"http"`
	Collector *CollectorConfig `json:"collector"`
	Mysql     *MysqlConfig     `json:"mysql"`
	HeartBeat *HeartBeatConfig `json:"heartbeat"`
	Registrar *RegistrarConfig `json:"registrar"`
}
type RegistrarConfig struct {
	Enable bool     `json:"enable"`
	Addrs  []string `json:"addrs"`
}

type HeartBeatConfig struct {
	Enable   bool `json:"enable"`
	Interval int  `json:"interval"`
}
type MysqlConfig struct {
	MysqlHost   string `json:"mysqlhost"`
	MysqlPort   string `json:"mysqlport"`
	MysqlUser   string `json:"mysqluser"`
	MysqlPasswd string `json:"mysqlpasswd"`
}

type TransferConfig struct {
	Enable   bool     `json:"enable"`
	Addr     []string `json:"addr"`
	Interval uint8    `json:"interval"`
	Timeout  int      `json:"timeout"`
}
type HttpConfig struct {
	Listen string `json:"listen"`
}

type CollectorConfig struct {
	Collector []string `json:"collector"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

// LoadConfiguration 加载配置文件
func LoadConfiguration() {
	var builder strings.Builder
	str, _ := os.Getwd()
	builder.Write([]byte(str))
	//builder.WriteString("/cfg.json")
	filePath := "/cfg.json"
	builder.WriteString(filePath)
	cfg := builder.String()
	if !utils.IsExist(cfg) {
		logger.Fatal("config file isn't exists:", cfg)
	}
	ConfigFile = cfg
	ConfigContent, err := utils.ToTrimString(cfg)
	if err != nil {
		logger.Fatal("read config file", cfg, "fail", err)
	}
	var c GlobalConfig
	err = json.Unmarshal([]byte(ConfigContent), &c)
	if err != nil {
		logger.Fatal("parse config file", cfg, "fail", err)
	}
	lock.Lock()
	defer lock.Unlock()
	config = &c
	logger.Info("read config file", cfg, "successfully")
}

// Hostname 获取主机名称
func Hostname() (string, error) {
	hostname := Config().HostName
	if hostname != "" {
		return hostname, nil
	}
	hostname, err := os.Hostname()
	if err != nil {
		logger.Info("get hostname err")
	}
	return hostname, nil
}

var LocalIp string

//初始化IP
func InitLocalIp() {
	Ip := ""
	addr := []string{"180.101.49.12:80"}
	for _, ip := range addr {
		conn, err := net.DialTimeout("tcp", ip, time.Second*5)
		if err != nil {
			logger.Info("get local addr failed")
		} else {
			Ip = strings.Split(conn.LocalAddr().String(), ":")[0]
			logger.StartupInfo(Ip)
			conn.Close()
			break
		}
	}
	if Ip != "" {
		LocalIp = Ip
	} else {
		logger.Fatal("get local addr failed")
	}
}

func IP() string {
	//如果配置了IP，直接使用
	ip := Config().Ip
	if ip != "" {
		return ip
	}
	//使用获取到的IP
	if len(LocalIp) > 0 {
		ip = LocalIp
	}
	return ip
}
