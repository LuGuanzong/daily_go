package main

import (
	"daily_plan_go/common/logger"
	"daily_plan_go/config"
	"daily_plan_go/dao"
	"daily_plan_go/router"
	"flag"
	"fmt"
	"log"
)

func main() {
	configFile := flag.String("c", "./conf/app.yaml", "-c filname (default: ./conf/app.yaml)")
	flag.Parse()
	fmt.Println("use conf", *configFile)
	config.LoadConfFromYaml(*configFile)
	fmt.Println(" - 完成配置文件读取")

	_ = logger.InitLogger()

	_ = dao.InitDb(config.Conf.Basic.Env)
	fmt.Println(" - 完成配置数据库")

	r := router.SetRouter()
	fmt.Println(" - 完成配置路由")

	// 启动服务
	log.Fatal(r.Run(":9999"))
}
