package main

import (
	. "github.com/go-utils/db-effect"
	repos "github.com/go-utils/db-effect/sample-project/repositories"
	"github.com/go-utils/db-effect/sample-project/structs"
)

func someOperation() []DBEffect {
	return []DBEffect{
		repos.InsertInt{"table1", 10},
		repos.SelectString{"table2", "someValue", 3},
		repos.InsertInt{"table5", 27},
		ValidateInsert(repos.InsertNade{"tableNade", structs.Nade{108}}),
		ValidateInsert(repos.InsertNade{"tableNade", structs.Nade{57}}),
		ValidateInsert(repos.InsertNade{"tableNade", structs.Nade{86}}),
	}
}
