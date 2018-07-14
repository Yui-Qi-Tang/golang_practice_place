/*
	@Description: This is a fun project for golang.
	              It's my practice.
	              Just for fun.
	@version: 0.1
	@Author: Yui-Qi Tang
	@Created: 2018/07/08
	@TODO: 3. remove -1 from result of tools.RemoveDupElement()
*/
package main // this is package name

import (
	"fmt"
	tester "yuki.pkg.org/sayhi"
	tools "yuki.pkg.org/tools"
	// apiRunner "yuki.pkg.org/webapi"
	"yuki.pkg.org/basictypes"
	"yuki.pkg.org/serverpush"
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
	    Basic type
	*/

	// :Create people instance and display
	p1 := basictype.NewPerson("Yuki Tang", 29)
	fmt.Printf("new people: %s, %d yaers old\n", p1.GetPersonName(), p1.GetPersonAges())
	// p2 := basictype.NewPerson("Ann Ke", 25)
	// fmt.Printf("new people: %s, %d yaers old\n", p2.GetPersonName(), p2.GetPersonAges())

    // :display pointer value of p1 and p2
	//p1ptr := &p1
	//p2ptr := &p2
	//fmt.Printf("%p and %p\n", p1ptr, p2ptr)

    // :update name of p2
	// p2.SetPersonName("Ann Tang")
	// fmt.Println(p2.GetPersonName())

	/*
	    Just run gogin web framework
	*/
	// apiRunner.RunGoGin()

	/*
		Http2 sever push demo
	*/
	serverpush.Demonstration()
}
