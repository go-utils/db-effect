package go_repo_gen

// Interpreter
type Interpreter struct {
	Insert func(ctx Insert) RuntimeContext
	Select func(ctx Select) RuntimeContext
}
