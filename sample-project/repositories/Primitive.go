package repositories

import (
	. "github.com/go-utils/db-effect"
)

// ------------ int / Insert ------------//

type InsertInt struct {
	TableName string
	NewValue  int
}

func (this InsertInt) Apply(ip Interpreter) RuntimeContext {
	return ip.Insert(this)
}

func (this InsertInt) InsertionParam() (string, Any) {
	return this.TableName, this.NewValue
}

// ---------- string / Select ---------- //

type SelectString struct {
	TableName     string
	ExistingValue string
	Limit         int
}

func (this SelectString) Apply(ip Interpreter) RuntimeContext {
	return ip.Select(this)
}

func (this SelectString) SelectionParam() (string, Any, int) {
	return this.TableName, this.ExistingValue, this.Limit
}
