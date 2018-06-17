package goExt

import (
	"sort"
	"fmt"
)

//升序
func AllKeysAscending(m map[string]interface{}) []string  {
	var slice = make([]string,0,len(m))
	for i,v := range m {
		fmt.Println(i,"---", v)
		slice = append(slice, i)
	}
	fmt.Println(slice)
	sort.Strings(slice)
	return slice

}

//降序
func AllKeysDescending(m map[string]interface{}) []string  {
	var slice = make([]string,0,len(m))
	for i,v := range m {
		fmt.Println(i,"---", v)
		slice = append(slice, i)
	}
	fmt.Println(slice)
	sort.Sort(sort.Reverse(sort.StringSlice(slice)))
	return slice
}

func AllValues(m map[string]interface{}) []interface{}  {

	var slice = make([]interface{},0,len(m))
	for _,v := range m {
		slice = append(slice, v)
	}
	return slice
}
