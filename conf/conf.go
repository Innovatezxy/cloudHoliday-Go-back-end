package conf

import (
	"github.com/go-ini/ini"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	// Cfg .
	Cfg       *ini.File
	AppID     string
	AppSecret string
	Database  string
	HttpPort  string
)

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

// 加载配置文件
func init() {
	var err error
	Cfg, err = ini.Load(GetAppPath() + "/conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadApp()
	LoadServer()
	LoadBase()
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	AppID = sec.Key("APPID").String()
	AppSecret = sec.Key("APPSECRET").String()
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	HttpPort = sec.Key("HTTP_PORT").String()
}

func LoadBase() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	Database = sec.Key("URL").String()
}
