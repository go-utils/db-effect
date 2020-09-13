package go_repo_gen

type Any interface {
}

type RuntimeContext struct {
	OK  bool
	Err error
	Ctx Any
}
