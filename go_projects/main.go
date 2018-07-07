package main // this is package name

import (
	"fmt"
	// tester "yuki.pkg.org/sayhi"
	tools "yuki.pkg.org/tools"
	apiRunner "yuki.pkg.org/webapi"
)


func main() {
	//fmt.Println("Hello world")
	//val := tester.SayHi()
	// fmt.Println(val)
	
	// apiRunner.RunGoGin()	
	
	// tools.SetArraySize(50)
	target := []int{1,2,3,3,4}
	tools.SetArraySize(len(target))
	// tools.ShowArrayInfo()
	newArray := tools.RemoveDupElement(target)
    fmt.Println(newArray)
}
