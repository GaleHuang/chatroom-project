package worker

import (
	"fmt"
	"github.com/galehuang/chatroom-project/model"
	"sync"
)

type MessageReceiveWorkerPool struct {
	messageDisplayModel *model.MessageDisplayModel
}

func (p *MessageReceiveWorkerPool) DispatchWorker(id int, receiveChan <-chan string, routineWait *sync.WaitGroup)  {
	defer routineWait.Done()

	for{
		select {
		case data, ok := <- receiveChan:
			if !ok{
				fmt.Println("error when receiving message")
			}

		}
	}
}