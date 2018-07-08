/*
	@Description: This is a fun project for golang.
	              It's my practice.
	              Just for fun.
	@version: 0.1
	@Author: Yui-Qi Tang
	@Created: 2018/07/08
	@TODO: 2. add golang base type practice
		   3. remove -1 from result of tools.RemoveDupElement()
*/
package main // this is package name

import (
	"fmt"
	tester "yuki.pkg.org/sayhi"
	tools "yuki.pkg.org/tools"
	// apiRunner "yuki.pkg.org/webapi"
)


func main() {
	/*
		Print the init string of sayhi package, and use SayHi function exported from package
	*/
	tester.SayHi() // return "Hi"
	
	/*
		A practice for golang that remove duplicated element of a integer array in O(n)
	*/	
	target := []int{1, 2, 3, 3, 19, 6, 6}
	tools.SetArraySize(target)
	newArray := tools.RemoveDupElement(target)
	fmt.Println("The duplicated element answer: ")
	fmt.Println(newArray)
	
	/*
	    Just run gogin web framework
	*/
	// apiRunner.RunGoGin()	
}
