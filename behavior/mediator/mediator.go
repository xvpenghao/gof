package mediator

import (
	"fmt"
	"time"
)

// 使用聊天室

type User struct {
	Name string
}

func (this *User) sendMsg(message string) {
	ChatRoomSendMsg(this, message)
}

// ChatRoomSendMsg 聊天室就是 充当一个 中介者，将 用户之间的关系 解耦
func ChatRoomSendMsg(user *User, message string) {
	fmt.Printf("%s [%s]:%s \n", time.Now(), user.Name, message)
}
