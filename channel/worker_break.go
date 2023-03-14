package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func WorkerBreak(wg *sync.WaitGroup, jobs chan Job, results chan Result, i int, ctx context.Context) {
	// for job := range jobs {
	// 	fmt.Printf("worker %d job: %d start\n", i, job.Id)
	// 	execResult := job.Run()
	// 	fmt.Printf("worker %d job: %d end\n", i, job.Id)
	// 	execResultStr := "failed"
	// 	if execResult {
	// 		execResultStr = "successful"
	// 	}
	// 	//fmt.Printf("job id %d, result is %d\n", job.Id, job.RandNum)
	// 	result := Result{
	// 		Job:        job,
	// 		ExecResult: execResultStr,
	// 	}
	// 	results <- result
	// }
	// wg.Done()
	// fmt.Println("worker done: ", i)
	for {
		fmt.Println("---------------------worker ", i)
		select {
		case <-ctx.Done():
			wg.Done()
			fmt.Println("worker done: ", i)
			return
		case job, ok := <-jobs:
			if !ok {
				wg.Done()
				fmt.Println("worker done: ", i)
				return
			}
			fmt.Printf("worker %d job: %d start\n", i, job.Id)
			execResult := job.Run()
			fmt.Printf("worker %d job: %d end\n", i, job.Id)
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
			// default:
			// 	wg.Done()
			// 	fmt.Println("worker done: ", i)
			// 	return
		}
	}
}

func CreateWorkerPoolBreak(noOfWorker int, jobs chan Job, results chan Result, ctx context.Context) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorker; i++ {
		fmt.Println("create worker ", i)
		wg.Add(1)
		go WorkerBreak(&wg, jobs, results, i, ctx)
	}
	wg.Wait()
	fmt.Println("worker pool wait")
	close(results)
}

func AllocateBreak(noOfJob int, jobs chan Job, ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			close(jobs)
			fmt.Println("stop allocate job: ", i)
			return
		default:
		}
		if i > noOfJob {
			fmt.Println("allocate all job")
			close(jobs)
			return
		}
		fmt.Println("allocate job ", i)
		randomNo := rand.Intn(999)
		job := Job{
			Id:      i,
			RandNum: randomNo,
		}
		jobs <- job
		i++
	}
}

func execBreak() {
	fmt.Println("exec starting")
	var jobs = make(chan Job, 4)
	var results = make(chan Result, 10)
	ctx, cancel := context.WithCancel(context.Background())
	go AllocateBreak(50, jobs, ctx)
	time.Sleep(1 * time.Second)
	go CreateWorkerPoolBreak(4, jobs, results, ctx)

	for result := range results {
		// fmt.Printf("job %d result\n", result.Job.Id)
		if result.Job.Id == 6 {
			fmt.Println("job failed: ", 6)
			continue
		}
	}
	fmt.Println("all job finished")
	cancel()
}
