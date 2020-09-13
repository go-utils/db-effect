package interpreter

import (
	. "github.com/go-utils/db-effect"
)

func MockDB() Interpreter {
	return Interpreter{
		Insert: func(effect Insert) RuntimeContext {
			return RuntimeContext{
				OK:  true,
				Err: nil,
				Ctx: nil,
			}
		},

		Select: func(effect Select) RuntimeContext {
			// tableName, query
			_, query := effect.SelectionParam()
			return RuntimeContext{
				OK:  true,
				Err: nil,
				Ctx: query,  // should return something
			}
		},

		Update: func(effect Update) RuntimeContext {
			return RuntimeContext{
				OK:  true,
				Err: nil,
				Ctx: nil,
			}
		},

		Delete: func(effect Delete) RuntimeContext {
			return RuntimeContext{
				OK:  true,
				Err: nil,
				Ctx: nil,
			}
		},
	}
}
