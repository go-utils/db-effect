package repositories

import (
	. "github.com/go-utils/go-repo-gen"
	"github.com/go-utils/go-repo-gen/sample-project/structs"
)

// Insert for specific type
type InsertNade struct {
	TableName string
	NewValue  structs.Nade
}

func (this InsertNade) Apply(ip Interpreter) RuntimeContext {
	return ip.Insert(this)
}

func (this InsertNade) InsertionParam() (string, Any) {
	return this.TableName, this.NewValue
}
