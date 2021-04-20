package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/lsytj0413/myinctl/cmd/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(runtime.GOMAXPROCS(0))
	command := app.NewMyInCtlCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
