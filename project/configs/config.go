package configs

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"os"
)

type Config struct {
	MysqlUrl string `ini:"mysqlUrl"`
	GrpcPort string `ini:"grpcPort"`
	GrpcIp string `ini:"grpcIp"`
}

var ProjectConfig *Config

func LoadConfig(configFile string) error {
	var (
		iniFile *ini.File
		section *ini.Section
	)
	ProjectConfig = &Config{}

	// default file path  ./conf/app.ini
	if configFile == "" {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		configFile = dir + "/configs/app.ini"
	}

	iniFile, err := ini.Load(configFile)
	if err != nil {
		return err
	}
	mode, exist := os.LookupEnv("MODE")
	if !exist {
		mode = iniFile.Section("").Key("runMode").String()
	}
	switch mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
		section = iniFile.Section("debug")
		break
	case "test":
		gin.SetMode(gin.TestMode)
		section = iniFile.Section("test")
		break
	case "release":
		gin.SetMode(gin.ReleaseMode)
		section = iniFile.Section("release")
		break
	default:
		err = errors.New("unknown run mode")
		return err
	}
	err = section.MapTo(ProjectConfig)
	if err != nil {
		return err
	}
	return nil
}
