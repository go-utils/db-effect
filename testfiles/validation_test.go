package db_effect

import (
	"errors"
	"github.com/go-playground/validator"
	. "github.com/go-utils/db-effect"
	// "reflect"
	"testing"
)

// define repository for `test_validation_struct`

type test_validation_struct struct {
	Value_70_120 int32 `validate:"gte=70,lte=120"`
}

type InsertTestValidationStruct struct {
	TableName string
	NewValue  test_validation_struct
}

func (this InsertTestValidationStruct) Apply(ip Interpreter) RuntimeContext {
	if ip.Insert == nil {
		panic(errors.New("Interpreter for Insert has not been implemented"))
	}
	return ip.Insert(this)
}

func (this InsertTestValidationStruct) InsertionParam() (string, Any) {
	return this.TableName, this.NewValue
}

// create mock

func Mock() Interpreter {
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
				Ctx: query, // should return something
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

// test main

func TestValidation(t *testing.T) {
	mock := Mock()
	v := validator.New()

	validEntity := test_validation_struct{90}
	invalidEntity := test_validation_struct{50}

	// invalid entity should be errored
	effectSeq_invalid :=
		[]DBEffect{
			ValidateInsert(InsertTestValidationStruct{"aaa", validEntity}, v),
			ValidateInsert(InsertTestValidationStruct{"bbb", invalidEntity}, v),
			ValidateInsert(InsertTestValidationStruct{"bbb", validEntity}, v),
		}

	_, err := GetNotNilList(effectSeq_invalid, mock)
	if err == nil {
		t.Fatal("failed to validate wrong value")
	}

	// valid entity should not be errored
	effectSeq_valid :=
		[]DBEffect{
			ValidateInsert(InsertTestValidationStruct{"aaa", validEntity}, v),
			ValidateInsert(InsertTestValidationStruct{"bbb", validEntity}, v),
		}
	_, err = GetNotNilList(effectSeq_valid, mock)
	if err != nil {
		t.Fatal("too strict validation for valid value")
	}
}
