package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
	"net/http"
	"short-url/internal/conf"
	"short-url/internal/data"
	"short-url/internal/router"
	"short-url/pkg/log"

	"github.com/spf13/viper"
)

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "config", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	logger, err := log.NewJSONLogger(
		log.WithDisableConsole(),
		log.WithField("service", "short-url"),
		log.WithTimeLayout("2006-01-02 15:04:05"),
		log.WithFileP(""),
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		logger.Sync()
	}()
	var bc conf.ServerConfig

	v := viper.New()
	v.SetConfigFile(flagconf)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&bc); err != nil {
			fmt.Println(err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&bc); err != nil {
		panic(fmt.Errorf("unmarshal config failed: %s \n", err))
	}
	if err := data.Init(&bc); err != nil {
		panic(fmt.Errorf("init database failed: %s \n", err))
	}
	g := router.LoadEngine()
	server := &http.Server{
		Addr:    "",
		Handler: g,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("http server startup err", zap.Error(err))
		}
	}()
}
