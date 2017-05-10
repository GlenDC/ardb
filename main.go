package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var (
		flagServerCount int
		flagStartPort   int
		flagVerbose     bool
	)

	flag.BoolVar(&flagVerbose, "v", false, "print ardb output")
	flag.IntVar(&flagServerCount, "n", 1, "amount of ardb servers to run")
	flag.IntVar(&flagStartPort, "port", 16380, "first of multiple consecutive ardb-used ports")
	flag.Parse()

	if flagServerCount < 1 {
		log.Fatal("at least 1 ardb server is required")
	}

	var wg sync.WaitGroup
	wg.Add(flagServerCount)

	var servers []string
	for i := 0; i < flagServerCount; i++ {
		port := strconv.Itoa(flagStartPort + i)
		cmd := exec.Command("redis-server", "--port", port)
		cmd.Stderr = os.Stderr
		if flagVerbose {
			cmd.Stdout = os.Stdout
		}

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}

		go func() {
			defer wg.Done()
			cmd.Wait()
		}()

		servers = append(servers, fmt.Sprintf("localhost:%s", port))
	}

	fmt.Println(strings.Join(servers, ","))
	wg.Wait()
}
