package goExt

import (
	"reflect"
	"fmt"
)

func FirstObject(slice interface{}) int{
	if reflect.ValueOf(slice).IsNil() {
		return 1
	}

	fmt.Println(reflect.ValueOf(slice))
	return 1
}
