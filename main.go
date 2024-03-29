package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"systemMoniter-Server/dao/mysql"
	"systemMoniter-Server/logger"
	"systemMoniter-Server/models"
	"systemMoniter-Server/routes"
	"systemMoniter-Server/settings"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	//设置东八中文时区
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	//1.加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed err:%v\n", err)
		return
	}
	//2、初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	//3.初始化mysql连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed err:%v\n", err)
		return
	}
	mysql.Migrate()
	err := mysql.SetOldNodeStatus()
	if err != nil {
		zap.L().Error("Update node status failed")
	}
	// //4.初始化Redis连接
	// if err := redis.Init(); err != nil {
	// 	fmt.Printf("init redis failed err:%v\n", err)
	// 	return
	// }
	// defer redis.Close()
	//5.注册路由
	models.SetConfig()
	// basic := local.NewBasic()
	// basic.Start()
	// defer basic.Stop()
	// netSpeed := local.NewNetSpeed()
	// netSpeed.Start()
	// defer netSpeed.Stop()
	// if models.IsOpen == true {
	// 	local.PingValue.IpStatus = false
	// 	p10086 := local.NewPing()
	// 	defer p10086.Stop()
	// 	p10086.RunCM()
	// 	p10010 := local.NewPing()
	// 	defer p10010.Stop()
	// 	p10010.RunCU()
	// 	p189 := local.NewPing()
	// 	defer p189.Stop()
	// 	p189.RunCT()
	// }

	r := routes.SetUp()

	//6.启动服务（优雅关机）
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", viper.GetInt("port")),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
