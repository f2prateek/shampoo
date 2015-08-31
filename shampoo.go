package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"sync"

	"github.com/tj/docopt"
)

const (
	Usage = `Shampoo.

Shampoo away your flaky tests.

Usage:
  shampoo [--iterations=<i>] [--parallel] <cmd> [<args>...]
  shampoo -h | --help
  shampoo --version

Options:
  -h --help                  Show this screen.
  --version                  Show version.
  --iterations=<i>           Number of iterations to run [default: 10].
  --parallel                 Run in parallel.`

	Version = `1.0.0`
)

func main() {
	arguments, err := docopt.Parse(Usage, nil, true, Version, false)
	check(err)

	cmd := arguments["<cmd>"].(string)
	args := arguments["<args>"].([]string)
	iterations := parseInt(arguments["--iterations"].(string))
	parallel := arguments["--parallel"].(bool)

	fmt.Println("Running", cmd, args, iterations, "times.")

	if parallel {
		executeParallel(iterations, cmd, args...)
	} else {
		executeSerial(iterations, cmd, args...)
	}

	fmt.Println()
	fmt.Println("Completed", iterations, "iterations without any errors.")
}

var executor = execCommand

func execCommand(cmd string, args ...string) ([]byte, error) {
	return exec.Command(cmd, args...).Output()
}

func execute(cmd string, args ...string) {
	fmt.Printf(".")
	output, err := executor(cmd, args...)
	if err != nil {
		// Dump the output.
		fmt.Println(output)
		// Print the error.
		log.Fatal(err)
	}
}

func executeSerial(iterations int, cmd string, args ...string) {
	for i := 0; i < iterations; i++ {
		execute(cmd, args...)
	}
}

func executeParallel(iterations int, cmd string, args ...string) {
	var wg sync.WaitGroup
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			execute(cmd, args...)
		}()
	}
	wg.Wait()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInt(arg string) int {
	num, err := strconv.ParseInt(arg, 10, 32)
	check(err)
	return int(num)
}
