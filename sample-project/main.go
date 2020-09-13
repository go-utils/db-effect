package main

import (
	"fmt"
	"github.com/go-utils/go-repo-gen/sample-project/interpreter"
)

func main() {
	effects := someOperation()
	resList, err := interpreter.PseudoDB_Transaction(effects)

	if err != nil {
		fmt.Println("[E] main |", err)
	} else {
		fmt.Println("[I] main |", resList)
	}
}
