package main

import "fmt"

func main() {
package main

import "fmt"

var firstname string = "MOhan" // Declaration with initilization with type
var lastname string            //Declaration

var age = 29 // Declaration with initilization without type

// variable literals
var (
	emplcode, empdid int
	empadd           string = "chennai"
	empref                  = "EMIS"
)

func main() {
	var _firstname string = "MOhan" // Declaration with initilization with type
	var _lastname string            //Declaration

	var _age = 29 // Declaration with initilization without type

	//variable literals
	var (
		_emplcode, _empdid int
		_empadd            string = "chennai"
		_empref                   = "EMIS"
	)
	fathername := "Nataraj"
	fmt.Println(_firstname, _lastname, _age, _emplcode, _empdid, _empadd, _empref, fathername)
}

}
