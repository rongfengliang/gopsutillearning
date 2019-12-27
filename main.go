package main

import (
	"fmt"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func printmeminfo() {
	for {
		// fetch virtual mem info
		mem, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("some wrong", err)
		} else {
			fmt.Println((mem.String()))
		}
	}
}

func printloadinfo() {
	for {
		// fetch virtual mem info
		load, err := load.Avg()
		if err != nil {
			fmt.Println("some wrong", err)
		} else {
			fmt.Println(load.String())
		}
	}
}
func main() {
	go printmeminfo()
	go printloadinfo()
	log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
}
