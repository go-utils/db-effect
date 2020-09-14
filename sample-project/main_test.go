package main

import (
	"fmt"
	. "github.com/go-utils/db-effect"
	"github.com/go-utils/db-effect/sample-project/interpreter"
	"reflect"
	"testing"
)

// モック用のDBが必要なら以下のように自作する (テストが非自明になるので，なるべく使わずに済むようにする)
func makeMockDB_1() Interpreter {
	return Interpreter{
		Insert: func(effect Insert) RuntimeContext {
			return RuntimeContext{
				OK: true,
			}
		},

		Select: func(effect Select) RuntimeContext {
			// tableName, query
			_, query := effect.SelectionParam()

			// vv for Testing vv //
			// 検索に使われた文字列を4つ返すようなDB (理由は特にない)
			foundVal := query.(string)
			var res []string
			for i := 0; i < 4; i++ {
				res = append(res, foundVal)
			}
			// ^^ for Testing ^^ //

			return RuntimeContext{
				OK:  true,
				Ctx: res,
			}
		},

		Update: func(effect Update) RuntimeContext {
			return RuntimeContext{
				OK: true,
			}
		},

		// `Delete` implemention is unneeded
	}
}

// テスト本体
func TestMain(t *testing.T) {
	effects := someOperation()

	// ここで直接 `effects` を検査する

	talbesToAffect := map[string]bool{
		"table1":    true,
		"table2":    true,
		"table3":    true,
		"table4":    false, // talbe4 を弄ってはいけない
		"table5":    true,
		"tableNade": true,
	}
	for _, effect := range effects {
		switch effect := effect.(type) {
		case Insert:
			tableName, _ := effect.InsertionParam()
			chk1, chk2 := talbesToAffect[tableName]
			if !chk1 || !chk2 {
				t.Fatal(fmt.Sprintf("accessed to %s, which is not allowed", tableName))
			}
			// fmt.Printf("%#v\n", effect)
		case Select:
			// nothing to do
		case Update:
			tableName, _, upsert := effect.UpdateParam()
			chk1, chk2 := talbesToAffect[tableName]
			if !chk1 || !chk2 {
				t.Fatal(fmt.Sprintf("accessed to %s, which is not allowed", tableName))
			}
			if !upsert {
				t.Fatal("updates in this effect must be \"upsert\"")
			}
			// fmt.Printf("%#v\n", effect)
		case Delete:
			t.Fatal("this effect must not delete any value")
		}
	}

	// バリデーション等は Mock を放り込んでテストする
	mockDB := interpreter.MockDB()
	_, err := GetNotNilList(effects, mockDB) // 普通のモックDBで実行
	if err == nil {
		t.Fatal("failed to validate wrong value")
	}

	// 一応，以下のように疑似的なDBを放り込んでのテストもできる (複雑になるので，複雑なアルゴリズムを検証する時に使う程度にしたい。)
	myMockDB := makeMockDB_1()
	resList, err := GetNotNilList(effects, myMockDB) // 自作のモックDBで実行
	if err == nil {
		t.Fatal("failed to validate wrong value")
	}

	expected := []Any{
		[]string{"someQuery", "someQuery", "someQuery", "someQuery"},
	}
	if !reflect.DeepEqual(expected, resList) {
		t.Fatal(fmt.Sprintf("response: %#v <> %#v", expected, resList))
	}
}
