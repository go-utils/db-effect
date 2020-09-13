package db_effect

// Interpreter
type Interpreter struct {
	Insert func(ctx Insert) RuntimeContext
	Select func(ctx Select) RuntimeContext
	Update func(ctx Update) RuntimeContext
	Delete func(ctx Delete) RuntimeContext
}
