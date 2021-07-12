package Router

import (
	"awesomeProject/analyse file/Controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"os"
)

func Getroute()  {
    // 1，读取ini端口号
	cfg, err := ini.Load("/Users/fuyuanfu/go/src/awesomeProject/First/config/app.ini")
	if err != nil {
		fmt.Println("文件读取错误",err)
		os.Exit(1)
	}
	fmt.Println(cfg.Section("server").Key("HttpPort"))

 	//2，将端口号传输给router
	router := gin.Default()
 	port := cfg.Section("server").Key("HttpPort").String()
	// 匹配的url格式:
	router.GET("/getsolider", Controller.Getsolider)
 	router.GET("/getrarity", Controller.Getrarity)
	router.GET("/getcombatpoints", Controller.Getcombatpoints)
	router.GET("/getlegalsolider", Controller.Getlegalsolider)
	router.GET("/getsoliderjson", Controller.Getsoliderjson)



	router.Run(":"+port)

}

