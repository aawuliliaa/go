package main

import (
	"err/detail"
	"fmt"
)

func main() {
	//使用wrap的方式
	err := detail.AddApp("demo")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("add app success ")
    //没有使用wrap的方式
	//err = detail.AddAppNoWrap("demo")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("add app success ")
}
