package main

import (
	//"encoding/json"
	"fmt"
	//"test/morestrings"
	//"test/logger"
	//"github.com/PolarPanda611/trinity"
	//"test/exception"
	"time"
)

type PartialPaymentStruct struct {
	Number int32     `json:"num"`
	Rates  []float64 `json:"rates"`
}

func main() {
	fmt.Println("Hello world!")
	//fmt.Println(morestrings.ReverseRunes("Hello world!"))
	//trinity.SetRunMode("local")

	/* 	pp := PartialPaymentStruct{2, []float64{20.2, 79.8}}
	   	fmt.Println(pp.Number)
	   	fmt.Println(pp.Rates)
	   	str, err := json.Marshal(&pp)
	   	if err == nil {
	   		fmt.Println(string(str))
	   	}

	   	err = json.Unmarshal([]byte(string(str)), &pp)
	   	if err == nil {
	   		fmt.Println(pp)
	   	}
	*/
	//logger.TestLog()
	//exception.TestException()

	go func() {
		go func() {
			for {
				i := 1
				fmt.Println(i)
				time.Sleep(time.Second)
				i = i++
			}
		}()
		time.Sleep(5 * time.Second)
	}()

	time.Sleep(500 * time.Second)
	abc
}
