package worker

import (
	"fmt"
	"github.com/galehuang/chatroom-project/model"
	"sync"
)

type MessageReceiveWorkerPool struct {
}

func (p *MessageReceiveWorkerPool) DispatchWorker(id int, receiveChan <-chan string, routineWait *sync.WaitGroup)  {
	defer routineWait.Done()
	messageDisplayModel := model.MessageDisplayModel{}
	for{
		data, ok := <- receiveChan
		if !ok{
			fmt.Printf("message receive worker[%d] error when receiving message\n", id)
		}
		err := messageDisplayModel.ProcessData(data)
		if err != nil{
			fmt.Printf("message receive worker[%d] error when processing data err=[%v]", id, err)
		}
	}
}