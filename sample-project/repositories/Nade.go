package repositories

import (
	"errors"
	. "github.com/go-utils/db-effect"
	"github.com/go-utils/db-effect/sample-project/structs"
)

// Insert for specific type
type InsertNade struct {
	TableName string
	NewValue  structs.Nade
}

func (this InsertNade) Apply(ip Interpreter) RuntimeContext {
	if ip.Insert == nil {
		panic(errors.New("Interpreter for Insert has not been implemented"))
	}
	return ip.Insert(this)
}

func (this InsertNade) InsertionParam() (string, Any) {
	return this.TableName, this.NewValue
}
