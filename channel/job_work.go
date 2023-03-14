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

func (j *Job) Run() bool {
	time.Sleep(2 * time.Second)
	if j.RandNum%2 == 0 {
		return true
	}
	return false
}

func Worker(wg *sync.WaitGroup, jobs chan Job, results chan Result) {
	for job := range jobs {
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
		results <- result
	}
	wg.Done()
}

func CreateWorkerPool(noOfWorker int, jobs chan Job, results chan Result) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorker; i++ {
		wg.Add(1)
		go Worker(&wg, jobs, results)
	}
	wg.Wait()
	close(results)
}

func Allocate(noOfJob int, jobs chan Job) {
	for i := 0; i < noOfJob; i++ {
		randomNo := rand.Intn(999)
		job := Job{
			Id:      i,
			RandNum: randomNo,
		}
		jobs <- job
	}
	close(jobs)
}

func CheckResult(done chan bool, results chan Result) {
	for result := range results {
		fmt.Printf("job detail: id %d, randNum: %d,  result: %s\n", result.Job.Id, result.Job.RandNum, result.ExecResult)
	}
	done <- true
}

func exec() {
	fmt.Println("exec starting")
	var jobs = make(chan Job, 10)
	var results = make(chan Result, 10)
	go Allocate(30, jobs)
	done := make(chan bool)
	go CheckResult(done, results)
	go CreateWorkerPool(3, jobs, results)
	<-done
	fmt.Println("exec ending")
}

func main() {
	startTime := time.Now()
	go exec()
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
