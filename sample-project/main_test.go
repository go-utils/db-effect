package main

import (
	"fmt"
	. "github.com/go-utils/db-effect"
	"reflect"
	"testing"
)

// モック用のDB (テストが非自明になるので，なるべく使わずに済むようにする)
func makeMockDB_1() Interpreter {
	return Interpreter{
		Insert: func(effect Insert) RuntimeContext {
			return RuntimeContext{
				OK: true,
			}
		},

		Select: func(effect Select) RuntimeContext {
			// tableName, existingValue, limit
			_, existingValue, limit := effect.SelectionParam()

			// vv for Testing vv //
			// 検索に使われた文字列をひたすら返すようなDB (理由は特にない)
			foundVal := existingValue.(string)
			var res []string
			for i := 0; i < limit; i++ {
				res = append(res, foundVal)
			}
			// ^^ for Testing ^^ //

			return RuntimeContext{
				OK:  true,
				Ctx: res,
			}
		},
	}
}

// テスト本体
func TestMain(t *testing.T) {
	effects := someOperation()

	// ここで直接 `effects` を検査する

	// 一応，以下のようなテストもできる (複雑になるので，バリデーションなどが動いているかを検査したい時に使う程度にしたい。)
	DB := makeMockDB_1()
	resList, err := GetNotNilList(effects, DB) // モックDBで実行
	if err == nil {
		t.Fatal("failed to validate wrong value: ")
	}

	expected := []Any{
		[]string{"someValue", "someValue", "someValue"},
	}
	if !reflect.DeepEqual(expected, resList) {
		t.Fatal(fmt.Sprintf("response: %#v <> %#v", expected, resList))
	}
}
