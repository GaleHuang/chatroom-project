package model

import (
	"encoding/json"
	"fmt"
	"github.com/galehuang/chatroom-project/bean"
)

type MessageDisplayModel struct {
	BucketSize int
	MessageBucket []*bean.Message
}

func NewMessageDisplayModel(bucketSize int) *MessageDisplayModel  {
	return &MessageDisplayModel{
		BucketSize: bucketSize,
		MessageBucket:     []*bean.Message{},
	}
}

func (m *MessageDisplayModel) FlushData() error {

}

func (m *MessageDisplayModel) ProcessData(data string) error  {
	Message := bean.Message{}
	err := json.Unmarshal([]byte(data), &Message)
	if err != nil{
		fmt.Println("error when parsing message")
	}
	m.MessageBucket = append(m.MessageBucket, &Message)
	if len(m.MessageBucket) < m.BucketSize{
		return nil
	}
	return m.FlushData()
}
