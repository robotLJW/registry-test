package main

import (
	"fmt"
	"log"
	"net/http"
	"registry-test/pkg/calculate"
	"registry-test/pkg/config"
	"strconv"
	"sync/atomic"
	"time"
)

var endTime int64
var count int64
var data = make(chan int64, 10)

func main() {
	config.ReadCalculateConfig()
	err := calculate.Execute()
	if err != nil {
		log.Fatal(err)
	}
	if config.CalculateKVCfg.Channel > 0 {
		data = make(chan int64, config.CalculateKVCfg.Channel)
	}
	go calculate.Execute()
	http.HandleFunc("/calculate", calc)
	err = http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func calc(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	receiveTime, _ := strconv.ParseInt(values["receiveTime"][0], 10, 64)
	fmt.Println(receiveTime)
	compute(receiveTime)
	w.Write([]byte("calculate finish!"))
}

func compute(receiveTime int64) {
	if count > 0 {
		return
	}
	select {
	case data <- receiveTime:
		if len(data) == config.CalculateKVCfg.Channel {
			endTime = time.Now().UnixNano()
			atomic.AddInt64(&count, 1)
			close(data)
			computeTime()
		}
	default:
		endTime = time.Now().UnixNano()
		atomic.AddInt64(&count, 1)
		close(data)
		computeTime()
	}
}

func computeTime() {
	var maxEndTime int64
	for {
		if d, ok := <-data; ok {
			if maxEndTime < d {
				maxEndTime = d
			}
		} else {
			fmt.Println(ok)
			break
		}
	}
	fmt.Println("--------------------------")
	useTime := (endTime - calculate.StartTime) / 1000000
	fmt.Printf("startTime %v \n", calculate.StartTime)
	fmt.Printf("endTime %v\n", endTime)
	fmt.Printf("useTime %v\n", useTime)
	fmt.Printf("thoughout %v\n", int64(config.CalculateKVCfg.Channel)*1000/useTime)
	fmt.Println("--------------------------")
	dataUseTime := (maxEndTime - calculate.StartTime) / 1000000
	fmt.Printf("startTime %v\n", calculate.StartTime)
	fmt.Printf("maxEndTime %v\n", maxEndTime)
	fmt.Printf("dataUseTime %v\n", dataUseTime)
	fmt.Printf("thoughout %v\n", int64(config.CalculateKVCfg.Channel)*1000/dataUseTime)
}
