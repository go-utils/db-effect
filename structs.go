package db_effect

type Any interface {
}

type RuntimeContext struct {
	OK  bool
	Err error
	Ctx Any
}
