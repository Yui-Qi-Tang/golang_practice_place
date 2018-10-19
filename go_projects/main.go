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
	// "yuki.pkg.org/serverpush"
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
	oldInfo := basictype.NewPerson(p1.GetPersonName(), p1.GetPersonAges())
	p1.SetPersonName("Yui Qi Tang")
	p1.SetPersonAges(30)
	
	fmt.Printf(
		"update people name from: %s to: %s\n",
		oldInfo.GetPersonName(),
		p1.GetPersonName(),
	)
	
	fmt.Printf(
		"update people ages from: %d to: %d yaers old\n",
		oldInfo.GetPersonAges(),
		p1.GetPersonAges(),
	)
    // :display pointer value of p1
	p1ptr := &p1
	fmt.Printf("%p\n", p1ptr)

	// :group people
	group1 := make(map[string]interface{
		GetPersonName() string // for person struct GetPersonName()
		GetPersonAges() int    // for person struct GetPersonAges()
	})
	group1["1st"] = p1
	group1["2nd"] = basictype.NewPerson("Member 2", 15)
	for k, v := range group1 {
		fmt.Printf("Key: %s, Value: {%s, %d}\n", k, v.GetPersonName(), v.GetPersonAges())
	}

	/*
	    Just run gogin web framework
	*/
	// apiRunner.RunGoGin()

	/*
		Http2 sever push demo
	*/
	// serverpush.Demonstration()
}
