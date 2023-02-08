package bootstrap

import (
	"encoding/json"
	"github.com/869413421/wechatbot/handlers"
	"github.com/eatmoreapple/openwechat"
	"log"
)

func Run() {
	log.Printf("============= server start =============\n")

	// bot := openwechat.DefaultBot() // 网页版微信
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式

	// 注册消息处理函数
	bot.MessageHandler = handlers.Handler
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()

	// 执行热登录
	err := bot.HotLogin(reloadStorage)
	if err != nil {
		if err = bot.Login(); err != nil {
			log.Printf("login error: %v \n", err)
			return
		}
	}

	// get self information
	self, err := bot.GetCurrentUser()
	if err != nil {
		log.Printf("get login user failed: %v \n", err)
	} else {
		selfJson, _ := json.Marshal(self)
		log.Printf("login user: %s \n", string(selfJson))
	}

	//"UserName": "@df6be4faaba2df2fc96d38cccf88ea78",
	//"NickName": "yang",

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
