package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func readCPUUsage(interval time.Duration) <-chan float64 {
	ch := make(chan float64)
	go func() {
		for {
			ps, err := cpu.Percent(interval, false)
			if err != nil {
				log.Print(err)
			}
			ch <- ps[0]
		}
	}()
	return ch
}

func drawGraph(name string, value int) {
	v := value / 4

	graph := "["
	for i := 0; i < 25; i++ {
		if i < v {
			graph += "|"
		} else {
			graph += "."
		}
	}
	graph += "]"
	fmt.Printf("%s %3d%% %s", name, value, graph)
}

func eraseGraph() {
	len := 11 + 100/4
	for i := 0; i < len; i++ {
		fmt.Print("\r \r")
	}
}

func drawHostname() {
	n, err := os.Hostname()
	if err != nil {
		log.Print(err)
	}
	fmt.Print(n)
}

func lf() {
	fmt.Print("\n")
}

func flush() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	cpuch := readCPUUsage(time.Millisecond * 1000)
	flush()
	drawHostname()
	lf()
	drawGraph("CPU", 0)
	for v := range cpuch {
		eraseGraph()
		drawGraph("CPU", int(v))
	}
}
