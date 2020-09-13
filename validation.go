package go_repo_gen

import (
	"github.com/go-playground/validator"
)

// Insert

type validatingInsert struct {
	effect Insert
}

func (this validatingInsert) Apply(ip Interpreter) RuntimeContext {
	var validate *validator.Validate = validator.New() // may be singleton
	_, newValue := this.effect.InsertionParam()
	err := validate.Struct(newValue)
	if err != nil {
		return RuntimeContext{OK: false, Err: err, Ctx: nil}
	}
	return this.effect.Apply(ip)
}

func (this validatingInsert) InsertionParam() (string, Any) {
	return this.effect.InsertionParam()
}

func ValidateInsert(effect Insert) DBEffect {
	return validatingInsert{effect}
}
