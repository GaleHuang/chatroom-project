package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
)
var gPubSubHelper *PubSubHelper

type PubSubHelper struct {
	redisClient *redis.Client
	pubSub *redis.PubSub
	routineWait *sync.WaitGroup
}

func NewPubSubHelper(client *redis.Client) *PubSubHelper {
	return &PubSubHelper{redisClient: client}
}


func (h *PubSubHelper) StartSubscribe(channelNames ...string) error {
	fmt.Printf("start subscribe channels %v\n", channelNames)
	ctx := context.Background()
	pubSub := h.redisClient.Subscribe(ctx, channelNames...)

	// Wait for confirmation that subscription is created before publishing anything

	_, err := pubSub.Receive(ctx)
	if err != nil{
		return errors.New(fmt.Sprintf("subscribe error=[%v]", err))
	}

	h.pubSub = pubSub
	return nil
}

func (h *PubSubHelper) StopSubscribe(channelNames ...string) error {
	err := h.pubSub.Close()
	if err != nil{
		fmt.Printf("stop subsribe error=[%v]", err)
	}
	return nil
}

func (h *PubSubHelper) GetSubscribeChan() <-chan *redis.Message {
	return h.pubSub.Channel()
}

func GetPubSubHelper() *PubSubHelper{
	return gPubSubHelper
}

func SetPubSubHelper(helper *PubSubHelper)  {
	gPubSubHelper = helper
}