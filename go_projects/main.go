package main // this is package name

import (
	"fmt"
	// tester "yuki.pkg.org/sayhi"
	tools "yuki.pkg.org/tools"
	// apiRunner "yuki.pkg.org/webapi"
)


func main() {
	//fmt.Println("Hello world")
	//val := tester.SayHi()
	// fmt.Println(val)
	
	// apiRunner.RunGoGin()	
	
	// tools.SetArraySize(50)
	// TODO: get max value of target array and set array size
	target := []int{1, 2, 3, 3, 19, 6, 6} // each element in target that is positive integer!!
	tools.SetArraySize(target)
	// tools.ShowArrayInfo()
	newArray := tools.RemoveDupElement(target)
    fmt.Println(newArray)
}
