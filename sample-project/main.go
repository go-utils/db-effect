package main

import (
	"fmt"
	"github.com/go-utils/db-effect/sample-project/interpreter"
	"github.com/go-utils/db-effect/sample-project/repositories"
)

func main() {
	effects := someOperation()
	resList, err := repositories.Any2NadeArr(interpreter.PseudoDB_Transaction(effects))

	if err != nil {
		fmt.Println("[E] main |", err)
	} else {
		fmt.Println("[I] main |", resList)
	}
}
