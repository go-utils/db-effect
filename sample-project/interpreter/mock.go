package interpreter

import (
	. "github.com/go-utils/db-effect"
)

func MockDB() Interpreter {
	return Interpreter{
		Insert: func(effect Insert) RuntimeContext {
			// tableName, newValue
			_, newValue := effect.InsertionParam()
			return RuntimeContext{
				OK:  true,
				Err: nil,
				Ctx: newValue,
			}
		},

		Select: func(effect Select) RuntimeContext {
			// tableName, existingValue, limit
			_, existingValue, _ := effect.SelectionParam()
			return RuntimeContext{
				OK:  true,
				Err: nil,
				Ctx: existingValue,
			}
		},
	}
}
