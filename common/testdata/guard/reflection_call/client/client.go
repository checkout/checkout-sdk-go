package client

import "reflect"

func Call(v interface{}) {
	reflect.ValueOf(v).Method(0).Call(nil)
}
