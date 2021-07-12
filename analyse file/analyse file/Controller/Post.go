package Controller

import (
	"awesomeProject/analyse file/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)


// 获取json数据里面有用的map
func Solidermap() service.Vmap {
	var filename string = "/Users/fuyuanfu/go/src/awesomeProject/First/config/config.army.model.json"
	InputFile,err := os.Open(filename)
	if err != nil{
		fmt.Println("open file err =",err)

	}
	defer InputFile.Close()

	fileData, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println("file content =", string(fileData))
	}






	//反序列化
	var solider1 service.Vmap
	err = json.Unmarshal(fileData,&solider1)
	if err != nil{
		fmt.Printf("序列化错误，err=%v\n",err)
	}

	return solider1
}





// 输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
func Getsolider(c *gin.Context)  {

	c.String(http.StatusOK,"%s",service.GetSolider("1","0","1000", Solidermap()))

}




// 输入士兵id获取稀有度
func Getrarity(c *gin.Context)  {


	c.String(http.StatusOK,"%s",service.GetRarity("10102",Solidermap()))

}

// 输入士兵id获取战力
func Getcombatpoints(c *gin.Context)  {

	c.String(http.StatusOK,"%s",service.GetCombatPoints("10102",Solidermap()))

}

// 输入cvc获取所有合法的士兵
func Getlegalsolider(c *gin.Context)  {

	c.String(http.StatusOK,"%s",service.GetLegalsoldier("1000",Solidermap()))

}

// 获取每个阶段解锁相应士兵的json数据
func Getsoliderjson(c *gin.Context)  {

	c.String(http.StatusOK,"%s",service.GetSoldierjson("0",Solidermap()))

}

