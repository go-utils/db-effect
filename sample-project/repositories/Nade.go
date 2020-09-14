package repositories

import (
	"errors"
	"fmt"
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

func Any2NadeArr(values []Any, err error) ([]structs.Nade, error) {
	if err != nil {
		return nil, err
	}
	var dstArr []structs.Nade
	for _, src := range values {
		dst, ok := src.(structs.Nade)
		if ok == false {
			return dstArr, errors.New(fmt.Sprintf("couldn't convert values into Nade: %#v", src))
		}
		dstArr = append(dstArr, dst)
	}
	return dstArr, nil
}
