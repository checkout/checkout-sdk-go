package client

import "reflect"

func Lookup(v interface{}) reflect.Value {
	return reflect.ValueOf(v).MethodByName("Get")
}
