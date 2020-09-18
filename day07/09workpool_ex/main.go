package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job 待计算的64位数
type Job struct {
	value int64
}

// Result 计算结果
type Result struct {
	Job *Job
	sum int64
}

var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)
var wg sync.WaitGroup

func producer(ch1 chan<- *Job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &Job{
			value: x,
		}
		ch1 <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func consumer(ch1 <-chan *Job, ch2 chan<- *Result) {
	defer wg.Done()
	for {
		job := <-ch1
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &Result{
			Job: job,
			sum: sum,
		}
		ch2 <- newResult
	}
}

func main() {
	// n := 123
	// sum := 0
	// for n > 0 {
	// 	sum += n % 10
	// 	n = n / 10
	// }
	// fmt.Println(sum)

	wg.Add(1)
	go producer(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go consumer(jobChan, resultChan)
	}

	for result := range resultChan {
		fmt.Printf("value: %d sum: %d\n", result.Job.value, result.sum)
	}
	wg.Wait()
}
