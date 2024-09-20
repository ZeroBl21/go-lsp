package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ZeroBl21/go-lsp/rpc"
)

func main() {
	fmt.Println("Hello there")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}
}

func handleMessage(msg string) {
	panic("unimplemented")
}
