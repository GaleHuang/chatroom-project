package subscriber

import (
	"fmt"
	"github.com/galehuang/chatroom-project/common"
	"github.com/galehuang/chatroom-project/services"
	"github.com/galehuang/chatroom-project/worker"
	"sync"
	"time"
)


type MessageReceiveSubscriber struct {
	workerPool *worker.MessageReceiveWorkerPool

	messageReceiveChan chan string

	routineWait *sync.WaitGroup

	switchChannelChan chan string
}

func NewMessageReceiveSubscriber() *MessageReceiveSubscriber  {
	return &MessageReceiveSubscriber{
		workerPool:         &worker.MessageReceiveWorkerPool{},
		messageReceiveChan: make(chan string),
		routineWait:        &sync.WaitGroup{},
	}
}

func (s *MessageReceiveSubscriber) Start()  {
	// 开启工作线程池
	for i := 0; i < common.WorkerNum; i++{
		interval := 1000 / common.WorkerNum
		s.routineWait.Add(1)
		go s.workerPool.DispatchWorker(i, s.messageReceiveChan, s.routineWait)
		time.Sleep(time.Millisecond * time.Duration(interval))
	}

	// 开启Redis subscriber主监听线程
	// TODO: redis subscriber main listening routine
	s.routineWait.Add(1)
	go s.StartSubscribe()

	// 开启channel切换线程
	s.routineWait.Add(1)
	go s.ChannelSwitch()

}

func (s *MessageReceiveSubscriber) Stop()  {
	s.routineWait.Wait()
}

func (s *MessageReceiveSubscriber) StartSubscribe()  {
	defer s.routineWait.Done()
	defer close(s.messageReceiveChan)
	// TODO: redis subscribe code
	err := services.GetPubSubHelper().StartSubscribe(common.CommonBroadCast)
	if err != nil{
		panic("start subscribe err=" + err.Error())
	}
	for{
		msg, ok := <- services.GetPubSubHelper().GetSubscribeChan()
		if !ok{
			panic("subscribe chan closed")
		}

		err := s.doProcessMsg(msg.Channel, msg.Payload)
		if err != nil{
			fmt.Printf("process msg error=[%v] msg=[%v]", err, msg)
		}
	}
}

func (s *MessageReceiveSubscriber) doProcessMsg(channelName string, payLoad string) error  {
	// TODO: process msg function
	return nil
}

func (s *MessageReceiveSubscriber) ChannelSwitch()  {
	// TODO: channel switch code
	defer s.routineWait.Done()
}

