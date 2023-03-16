package job

import (
	"fmt"
	"github.com/et-zone/ppcli/consts"
	"net/http"
	"runtime/debug"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func Run() {
	time.Sleep(time.Second * 5)
	go func() {
		for {
			wg.Add(5)
			go request(consts.JobHeap)
			go request(consts.JobBlock)
			go request(consts.JobGoroutine)
			go request(consts.JobAllocs)
			go request(consts.JobProfileCPU)
			wg.Wait()
			fmt.Println("succ exec Job ", time.Now().Format("2006-01-02 15:04:05"))
			time.Sleep(time.Second)
		}
	}()

}

func request(path string) {
	defer wg.Done()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Catch exception", err, string(debug.Stack()))
		}
	}()
	url := consts.BaseReqDataURL + consts.ExportAddr + path + consts.JobStepSeconds
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Run Job err ,path = ", path, err.Error())
	}
}
