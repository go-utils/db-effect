package interpreter

import (
	"fmt"
	. "github.com/go-utils/db-effect"
)

func PseudoDB() Interpreter {
	return Interpreter{
		Insert: func(effect Insert) RuntimeContext {
			tableName, newValue := effect.InsertionParam()
			fmt.Printf("[D] PseudoDB | Insert | tableName: %s, inserting: %v\n", tableName, newValue)
			return RuntimeContext{OK: true}
		},
		Select: func(effect Select) RuntimeContext {
			tableName, existingValue, limit := effect.SelectionParam()
			fmt.Printf("[D] PseudoDB | Update | tableName: %s, getting: %s, limit: %d\n", tableName, existingValue, limit)
			return RuntimeContext{OK: true}
		},
	}
}

// 以下は必須ではないが，DB固有機能を使うような便利ツールの類

func PseudoDB_Transaction(effects []DBEffect) ([]Any, error){
	DB := PseudoDB()
	var resList []Any
	// TODO: Begin Transaction Here !!
	resList, err := GetNotNilList(effects, DB)
	if(err != nil){
		// Rollback Here !!
		return resList, err
	}
	// TODO: Commit Transaction Here !!
	return resList, nil
}
