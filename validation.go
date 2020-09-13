package db_effect

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

// Select

type validatingSelect struct {
	effect Select
}

func (this validatingSelect) Apply(ip Interpreter) RuntimeContext {
	var validate *validator.Validate = validator.New() // FixMe: must be singleton
	// validate param
	_, query := this.effect.SelectionParam()
	err := validate.Struct(query)
	if err != nil {
		return RuntimeContext{OK: false, Err: err, Ctx: nil}
	}
	// affect
	var ctx RuntimeContext = this.effect.Apply(ip)
	// validate result
	if !ctx.OK {
		return ctx
	}
	if ctx.Ctx != nil {
		err = validate.Struct(ctx.Ctx)
		if err != nil {
			return RuntimeContext{OK: false, Err: err, Ctx: ctx.Ctx}
		}
	}
	return ctx
}

func (this validatingSelect) SelectionParam() (string, Any) {
	return this.effect.SelectionParam()
}

func ValidateSelect(effect Select) DBEffect {
	return validatingSelect{effect}
}

// Update

type validatingUpdate struct {
	effect Update
}

func (this validatingUpdate) Apply(ip Interpreter) RuntimeContext {
	var validate *validator.Validate = validator.New() // may be singleton
	_, newValue, _ := this.effect.UpdateParam()
	err := validate.Struct(newValue)
	if err != nil {
		return RuntimeContext{OK: false, Err: err, Ctx: nil}
	}
	return this.effect.Apply(ip)
}

func (this validatingUpdate) UpdateParam() (string, Any, bool) {
	return this.effect.UpdateParam()
}

func ValidateUpdate(effect Update) DBEffect {
	return validatingUpdate{effect}
}

// Delete: nope.
