package main

import (
	"context"
	"fmt"
	"github.com/tangrc99/gohelloblog/global"
	"github.com/tangrc99/gohelloblog/internal/router"
	"github.com/tangrc99/gohelloblog/pkg/db"
	"github.com/tangrc99/gohelloblog/pkg/setting"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"
)

func main() {

	configPath := "conf/conf.yaml"

	if len(os.Args) > 1 && os.Args[1] != "" {
		configPath = os.Args[1]
	}

	// 解析路径、文件名、拓展名
	p, f := path.Split(configPath)
	ext := path.Ext(f)
	f = strings.TrimSuffix(f, ext)
	ext = strings.TrimLeft(ext, ".")

	// 读取配置文件，并且分配参数
	s := setting.ReadFile(p, f, ext)
	global.MySQLSetting = s.SetupMySQL()
	global.MongoSetting = s.SetupMongo()
	global.ServerSetting = s.SetupServer()
	global.JWTSetting = s.SetupJWT()

	// 根据参数初始化全局变量
	global.Mongo = db.NewMongo(global.MongoSetting.Url, global.MongoSetting.Db)
	global.MongoLog = global.Mongo.GetCollection(global.MongoSetting.AccessLog)
	global.MongoArticle = global.Mongo.GetCollection(global.MongoSetting.Article)
	global.MySQL = db.NewMySQLFrom(global.MySQLSetting)

	var engine = router.New()

	server := http.Server{
		Addr:         global.ServerSetting.Url,
		Handler:      engine,
		ReadTimeout:  global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
		//TLSConfig: &tls.Config{
		//	MinVersion:               tls.VersionTLS13,
		//	PreferServerCipherSuites: true,
		//},
	}

	//pprof.Register(engine)

	fmt.Printf("Server listen on %s\n", server.Addr)

	go func() {
		err := server.ListenAndServe()
		//err := server.ListenAndServeTLS("cert.pem", "key.pem")
		if err != nil {
			return
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 接受软中断信号并且传递到 channel

	<-quit

	// http 的官方退出方式，留下五秒的善后时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {

	}

}
