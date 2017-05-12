package store

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

var ts *Store

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	// tearDown()
	os.Exit(retCode)
}

func setUp() {
	ts = NewStore(NewVirtualRespository())
	// TODO why is this necessary? How does isolation in tests work?
	virtualFile = Data{}
}

func Test_GetEmpty(t *testing.T) {
	if ts.Get("key1") != "" {
		t.Fail()
	}
}

func Test_SetGet(t *testing.T) {
	ex := "value1"
	ts.Set("key1", ex)
	if rs := ts.Get("key1"); rs != ex {
		printMsg(ex, rs)
		t.Fail()
	}
}

func Test_SetFlushGet(t *testing.T) {
	ex := "value1"
	ts.Set("key1", ex)
	ts.Flush()
	ts2 := NewStore(NewVirtualRespository())
	if rs := ts2.Get("key1"); rs != ex {
		printMsg(ex, rs)
		t.Fail()
	}
}

func Test_SetGetAll(t *testing.T) {
	ts.Set("key1", "value1")
	ts.Set("key2", "value2")
	ex := Data{"key1": "value1", "key2": "value2"}
	if rs := ts.GetAll(); !reflect.DeepEqual(rs, ex) {
		printMsg(ex, rs)
		t.Fail()
	}
}

func printMsg(ex interface{}, rs interface{}) {
	fmt.Printf("Result was '%v' (%T) instead of '%v' (%T)\n", rs, rs, ex, ex)
}
