package handlers

import (
	"github.com/869413421/wechatbot/config"
	"github.com/eatmoreapple/openwechat"
	"log"
)

// MessageHandlerInterface 消息处理接口
type MessageHandlerInterface interface {
	handle(*openwechat.Message) error
	ReplyText(*openwechat.Message) error
}

type HandlerType string

const (
	GroupHandler = "group"
	UserHandler  = "user"
)

// handlers 所有消息类型类型的处理器
var handlers map[HandlerType]MessageHandlerInterface

func init() {
	handlers = make(map[HandlerType]MessageHandlerInterface)
	handlers[GroupHandler] = NewGroupMessageHandler()
	handlers[UserHandler] = NewUserMessageHandler()
}

// Handler 全局处理入口
func Handler(msg *openwechat.Message) {
	log.Printf("hadler Received msg : type:%d content: %v", msg.MsgType, msg.Content)
	if msg.IsSendByGroup() { // 处理群消息
		handlers[GroupHandler].handle(msg)
	} else if msg.IsFriendAdd() { // 好友申请
		if config.LoadConfig().AutoPass {
			_, err := msg.Agree("你好我是基于chatGPT引擎开发的微信机器人，你可以向我提问任何问题。")
			if err != nil {
				log.Fatalf("add friend agree error : %v", err)
			}
		}
	} else if msg.IsSystem() || msg.StatusNotify() { // 系统消息
		log.Printf("unsupport system msg %d %s", msg.MsgType, msg.Content)
	} else {
		// 其余消息 包含私聊
		handlers[UserHandler].handle(msg)
	}

	err := msg.AsRead()
	if err != nil {
		log.Printf("mark message as read failed %v \n", err)
	}
}
