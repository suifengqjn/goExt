package goExt

import "runtime"
import (
	"fmt"

)

func Log(mes interface{})  {
	pc, _, line, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Println("error:", mes, " line:",line," func:", f.Name())
}



//func test2() {
//	pc, file, line, ok := runtime.Caller(2)
//	fmt.Println(pc)
//	fmt.Println(file)
//	fmt.Println(line)
//	fmt.Println(ok)
//	f := runtime.FuncForPC(pc)
//	fmt.Println(f.Name())
//
//	fmt.Println("=====================")
//	pc, file, line, ok = runtime.Caller(0)
//	fmt.Println(pc)
//	fmt.Println(file)
//	fmt.Println(line)
//	fmt.Println(ok)
//	f = runtime.FuncForPC(pc)
//	fmt.Println(f.Name())
//	fmt.Println("=====================")
//	pc, file, line, ok = runtime.Caller(1)
//	fmt.Println(pc)
//	fmt.Println(file)
//	fmt.Println(line)
//	fmt.Println(ok)
//	f = runtime.FuncForPC(pc)
//	fmt.Println(f.Name())
//}
