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

func readHostname() string {
	n, err := os.Hostname()
	if err != nil {
		log.Print(err)
	}
	return n
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
	beginOfLine()
	fmt.Printf("%s %3d%% %s", name, value, graph)
}

func drawText(s string) {
	fmt.Print(s)
}

func newLine() {
	fmt.Print("\n")
}

func beginOfLine() {
	fmt.Print("\r")
}

func flush() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	cpuch := readCPUUsage(time.Millisecond * 1000)
	flush()
	drawText(readHostname())
	newLine()
	drawGraph("CPU", 0)
	for v := range cpuch {
		drawGraph("CPU", int(v))
	}
}
