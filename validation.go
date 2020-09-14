package db_effect

import (
	"github.com/go-playground/validator"
)

// Insert

type validatingInsert struct {
	effect   Insert
	validate *validator.Validate
}

func (this validatingInsert) Apply(ip Interpreter) RuntimeContext {
	_, newValue := this.effect.InsertionParam()
	err := this.validate.Struct(newValue)
	if err != nil {
		_, isVE := err.(validator.ValidationErrors)
		if isVE {
			return RuntimeContext{OK: false, Err: err, Ctx: nil}
		}
	}
	return this.effect.Apply(ip)
}

func (this validatingInsert) InsertionParam() (string, Any) {
	return this.effect.InsertionParam()
}

func ValidateInsert(effect Insert, validate *validator.Validate) DBEffect {
	return validatingInsert{effect, validate}
}

// Select

type validatingSelect struct {
	effect   Select
	validate *validator.Validate
}

func (this validatingSelect) Apply(ip Interpreter) RuntimeContext {
	// validate param
	_, query := this.effect.SelectionParam()
	err := this.validate.Struct(query)
	// FixMe: Validation が空の時にも InvalidValidationError が出るので無視しているが，そのうち見分けられるようにしたい
	if err != nil {
		_, isVE := err.(validator.ValidationErrors)
		if isVE {
			return RuntimeContext{OK: false, Err: err, Ctx: nil}
		}
	}
	// affect
	var ctx RuntimeContext = this.effect.Apply(ip)
	// validate result
	if !ctx.OK {
		return ctx
	}
	if ctx.Ctx != nil {
		err = this.validate.Struct(ctx.Ctx)
		if err != nil {
			_, isVE := err.(validator.ValidationErrors)
			if isVE {
				return RuntimeContext{OK: false, Err: err, Ctx: ctx.Ctx}
			}
		}
	}
	return ctx
}

func (this validatingSelect) SelectionParam() (string, Any) {
	return this.effect.SelectionParam()
}

func ValidateSelect(effect Select, validate *validator.Validate) DBEffect {
	return validatingSelect{effect, validate}
}

// Update

type validatingUpdate struct {
	effect   Update
	validate *validator.Validate
}

func (this validatingUpdate) Apply(ip Interpreter) RuntimeContext {
	_, newValue, _ := this.effect.UpdateParam()
	err := this.validate.Struct(newValue)
	if err != nil {
		_, isVE := err.(validator.ValidationErrors)
		if isVE {
			return RuntimeContext{OK: false, Err: err, Ctx: nil}
		}
	}
	return this.effect.Apply(ip)
}

func (this validatingUpdate) UpdateParam() (string, Any, bool) {
	return this.effect.UpdateParam()
}

func ValidateUpdate(effect Update, validate *validator.Validate) DBEffect {
	return validatingUpdate{effect, validate}
}

// Delete: nope.
