package main

import (
	"awesomeProject/analyse file/Router"
	"gopkg.in/ini.v1"
)

var (
	//配置信息
	iniFile *ini.File
)

func main() {


	// 请求路由
	// 参数目前是我自己赋值的,参数我已经写的固定了。
	// 输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵 在浏览器输入：localhost:8000/getsolider
	// 输入士兵id获取稀有度在浏览器输入： localhost:8000/getrarity
	// 输入士兵id获取战力 在浏览器输入：localhost:8000/getcombatpoints
	// 输入cvc获取所有合法的士兵 在浏览器输入： localhost:8000/getlegalsolider
	// 获取每个阶段解锁相应士兵的json数据 在浏览器输入：localhost:8000/getsoliderjson

	Router.Getroute()


}


