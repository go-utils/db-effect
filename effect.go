package go_repo_gen

// Insert
type Insert interface {
	Apply(Interpreter) RuntimeContext
	InsertionParam() (string, Any)
}

// Select
type Select interface {
	Apply(Interpreter) RuntimeContext
	SelectionParam() (string, Any, int)
}

// TODO: add Update, Delete

type DBEffect interface {
	Apply(ip Interpreter) RuntimeContext
}
