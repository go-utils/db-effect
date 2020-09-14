package repositories

import (
	"errors"
	. "github.com/go-utils/db-effect"
)

// ------------ int / Insert ------------//

type InsertInt struct {
	TableName string
	NewValue  int
}

func (this InsertInt) Apply(ip Interpreter) RuntimeContext {
	if ip.Insert == nil {
		panic(errors.New("Interpreter for Insert has not been implemented"))
	}
	return ip.Insert(this)
}

func (this InsertInt) InsertionParam() (string, Any) {
	return this.TableName, this.NewValue
}

// ---------- string / Select ---------- //

type SelectString struct {
	TableName string
	Query     string
}

func (this SelectString) Apply(ip Interpreter) RuntimeContext {
	if ip.Select == nil {
		panic(errors.New("Interpreter for Select has not been implemented"))
	}
	return ip.Select(this)
}

func (this SelectString) SelectionParam() (string, Any) {
	return this.TableName, this.Query
}

// ---------- string / Update ---------- //

type UpdateString struct {
	TableName string
	Query     string
	Upsert    bool
}

func (this UpdateString) Apply(ip Interpreter) RuntimeContext {
	if ip.Update == nil {
		panic(errors.New("Interpreter for Update has not been implemented"))
	}
	return ip.Update(this)
}

func (this UpdateString) UpdateParam() (string, Any, bool) {
	return this.TableName, this.Query, this.Upsert
}

// ------------ int / Delete ------------//

type DeleteInt struct {
	TableName string
	OldValue  int
}

func (this DeleteInt) Apply(ip Interpreter) RuntimeContext {
	if ip.Delete == nil {
		panic(errors.New("Interpreter for Delete has not been implemented"))
	}
	return ip.Delete(this)
}

func (this DeleteInt) DeletionParam() (string, Any) {
	return this.TableName, this.OldValue
}
