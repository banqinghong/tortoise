package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	Job        Job
	ExecResult string
}

var Jobs = make(chan Job, 10)
var Results = make(chan Result, 10)

func (j *Job) Run() bool {
	if j.RandNum%2 == 0 {
		return true
	}
	return false
}

func Worker(wg *sync.WaitGroup) {
	for job := range Jobs {
		execResult := job.Run()
		execResultStr := "failed"
		if execResult {
			execResultStr = "successful"
		}
		//fmt.Printf("job id %d, result is %d\n", job.Id, job.RandNum)
		result := Result{
			Job:        job,
			ExecResult: execResultStr,
		}
		Results <- result
	}
	wg.Done()
}

func CreateWorkerPool(noOfWorker int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorker; i++ {
		wg.Add(1)
		go Worker(&wg)
	}
	wg.Wait()
	close(Results)
}

func Allocate(noOfJob int) {
	for i := 0; i < noOfJob; i++ {
		randomNo := rand.Intn(999)
		job := Job{
			Id:      i,
			RandNum: randomNo,
		}
		Jobs <- job
	}
	close(Jobs)
}

func CheckResult(done chan bool) {
	for result := range Results {
		fmt.Printf("job detail: id %d, randNum: %d,  result: %s\n", result.Job.Id, result.Job.RandNum, result.ExecResult)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	go Allocate(30)
	done := make(chan bool)
	go CheckResult(done)
	go CreateWorkerPool(3)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
