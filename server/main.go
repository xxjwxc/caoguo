package main

import (
	"fmt"
	"os"
	"time"

	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/server"

	"github.com/gmsec/goplugins/plugin"

	"caoguo/internal/config"
	"caoguo/internal/routers"

	_ "net/http/pprof"

	// _ "caoguo/internal/routers" // debug模式需要添加[mod]/routers 注册注解路由

	"caoguo/internal/service/order"

	"github.com/gin-gonic/gin"
	"github.com/xxjwxc/public/mydoc/myswagger"
	"github.com/xxjwxc/public/timerDeal"
	// "caoguo/internal/model"
)

// CallBack service call backe
func CallBack() {
	// 支付回调
	var timeDeal timerDeal.TimerDeal
	timeDeal.SetCallBackTimer(30 * time.Minute) // 10分钟一次
	timeDeal.AddOneCall(order.WxQueryCallBack)
	timeDeal.OnSart()
	// -----end

	// order.CheckOrderTrack()
	// notify.NotifyEmail("20200822150028671700799101744786")
	// order.CheckOrderTrack()
	// 运单回调
	timerDeal.OnPeDay(20, 0, 0, order.CheckOrderTrack) // 每天事件
	// -----end

	// swagger

	myswagger.SetHost("http://localhost:9001/")
	myswagger.SetBasePath("caoguo")
	myswagger.SetSchemes(true, false)
	// -----end --

	// reg := registry.NewDNSNamingRegistry()
	// grpc 相关 初始化服务
	// service := micro.NewService(
	// 	micro.WithName("lp.caoguo.eg1"),
	// 	// micro.WithRegisterTTL(time.Second*30),      //指定服务注册时间
	// 	micro.WithRegisterInterval(time.Second*15), //让服务在指定时间内重新注册
	// 	// micro.WithRegistryNaming(reg),
	// )

	router := gin.Default()
	router.Use(routers.Cors())
	v1 := router.Group("/commcn/api/v1")
	routers.OnInitRouter(v1) // 自定义初始化

	// plg, b := plugin.Run(plugin.WithMicro(service),
	// 	plugin.WithGin(router),
	// 	plugin.WithAddr(":"+config.GetPort()))

	plg, b := plugin.RunHTTP(plugin.WithGin(router),
		plugin.WithAddr(":"+config.GetPort()))
	if b == nil {
		plg.Wait()
	} else {
		mylog.Error(b)
	}
	fmt.Println("done")
}

func main() {
	if config.GetIsDev() || len(os.Args) == 0 {
		CallBack()
	} else {
		mylog.SetLog(mylog.GetDefaultZap())
		server.On(config.GetServiceConfig()).Start(CallBack)
	}
}
