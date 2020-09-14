package main

import (
	"github.com/go-playground/validator"
	. "github.com/go-utils/db-effect"
	repos "github.com/go-utils/db-effect/sample-project/repositories"
	"github.com/go-utils/db-effect/sample-project/structs"
)

func someOperation() []DBEffect {
	v := validator.New()
	return []DBEffect{
		repos.InsertInt{"table1", 10},
		ValidateSelect(repos.SelectString{"table2", "someQuery"}, v),
		repos.InsertInt{"table5", 27},
		repos.UpdateString{"table3", "newValue", true},
		ValidateInsert(repos.InsertNade{"tableNade", structs.Nade{108}}, v),
		ValidateInsert(repos.InsertNade{"tableNade", structs.Nade{57}}, v),
		ValidateInsert(repos.InsertNade{"tableNade", structs.Nade{86}}, v),
	}
}
