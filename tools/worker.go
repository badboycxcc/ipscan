package tools

import (
	"sync"
)

type Workdist struct {
	Host	string
}

const (
	taskload		    = 100
	tasknum			= 100
)
var wg sync.WaitGroup


func Task(iplist []string){
	wg.Add(tasknum)
	tasks := make(chan Workdist,taskload)
	for gr:=1;gr<=tasknum;gr++ {
		go worker(tasks)
	}

	for _, host := range iplist {
		task := Workdist{
			Host:host,
		}
		tasks <- task
	}
	close(tasks)
	wg.Wait()
}

func worker(tasks chan Workdist){
	defer wg.Done()
	for {
		task,ok := <- tasks
		if !ok {
			return
		}
		host := task.Host
		//fmt.Println("开始探测:",host)
		ping(host)
	}

}

