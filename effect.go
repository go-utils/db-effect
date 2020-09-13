package db_effect

// Insert
type Insert interface {
	Apply(Interpreter) RuntimeContext
	InsertionParam() (DBName string, newValue Any)
}

// Select
type Select interface {
	Apply(Interpreter) RuntimeContext
	SelectionParam() (DBName string, query Any)
}

// Update
type Update interface {
	Apply(Interpreter) RuntimeContext
	UpdateParam() (DBName string, newValue Any, upsert bool)
}

// Delete
type Delete interface {
	Apply(Interpreter) RuntimeContext
	DeletionParam() (DBName string, oldValue Any)
}

type DBEffect interface {
	Apply(ip Interpreter) RuntimeContext
}
