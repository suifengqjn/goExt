package main

import (
	"goExt/goExt"
	"fmt"
	"reflect"
)

func main()  {
	//
	//a := 1
	//var s = []*int{&a}
	//
	//
	//fmt.Println(goExt.FirstObject(s))

	//var s2 = []string{}
	//
	//var s3 = []interface{}{1, "asd", true}

	//
	//fmt.Println("1",goExt.FirstObject(s))
	//fmt.Println("1761461024
	//QJNqjn19922",goExt.FirstObject(s2))
	//fmt.Println("3",goExt.FirstObject(s3))
	//
	//s4, _ := goExt.RemoveObject(s3, 1)
	//
	//fmt.Println(3, s4)
	//
	//
	//s3 = goExt.Insert2(s3, 0,"a")
	//fmt.Println(s3)

	//test2()


	var v interface{}
	v = []interface{}{12}
	//if _, ok := v.(int); ok {
	//	fmt.Printf("int--%d\n", v)
	//}

	if _, ok := v.([]interface{}); ok {
		fmt.Printf("slice -- %d\n", v)
	}
	testFirstObject()
	testInsetObject()
	testRemoveObject()
}

func testFirstObject()  {
	var slice  = []interface{}{"a", 1, true}

	var t = goExt.FirstObject(slice)

	fmt.Printf("%T----%v\n", t, t)
}


func testInsetObject() {

	var slice = []int{1}

	newSlice, _ := goExt.InsertObject(slice, 1, 6)
	fmt.Printf("insert:%T-%v\n", newSlice, newSlice)

}

func testRemoveObject()  {

	var slice = []int{1, 2,3 ,4}

	 s, _ := goExt.RemoveObject(slice, 3)
	 fmt.Printf("remove :   %T---%v", s, s)


}

func test(a interface{}){
	value := reflect.ValueOf(a)
	fmt.Println(value.Type())
	if value.Kind() == reflect.Int {
		var b = value.Interface().(int)
		fmt.Println("int",b)
	}

	fmt.Println(value)
}

func test2()  {
	var m = make(map[string]interface{})

	m["s"] = "as"
	m["a"] = "kk"
	m["h"] = "adff"
	fmt.Println(goExt.AllKeysAscending(m))
}