package goExt

import (
	"reflect"
	"errors"
	"fmt"
)

func FirstObject(slice interface{}) interface{} {

	rValue := reflect.ValueOf(slice)
	if  rValue.Kind() != reflect.Slice || rValue.IsNil(){
		fmt.Println("NULL")
		return nil
	} else if rValue.Kind() == reflect.Slice {
		v := reflect.Append(rValue, rValue)
		return v.Index(0).Interface()
	}
	return nil

}

func RemoveObject(slice interface{}, index int) (interface{}, error)  {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	if slice == nil || length == 0 || (length-1) < index {
		return nil, errors.New("slice is nil")
	}
	if length-1 == index {
		return sliceValue.Slice(0, index).Interface(), nil
	} else if (length - 1) >= index {
		return reflect.AppendSlice(sliceValue.Slice(0, index), sliceValue.Slice(index+1, length)).Interface(), nil
	}
	return nil, errors.New("error")
}


func InsertObject(slice interface{}, pos int, value interface{}) (interface{}, error) {

	v := reflect.ValueOf(slice)
	length := v.Len()
	if slice == nil || length == 0 || length < pos {
		return nil, errors.New("slice is nil")
	}
	v = reflect.Append(v, reflect.ValueOf(value))
	reflect.Copy(v.Slice(pos+1, v.Len()), v.Slice(pos, v.Len()))
	v.Index(pos).Set(reflect.ValueOf(value))
	return v.Interface(), nil

}