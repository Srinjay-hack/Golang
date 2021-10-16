package main

import (
	"fmt"
)
var(
	name string="Srinjay"
 	company string="NEw Tech Ltd."
	doctNumber int = 3
	season int = 11
	i int = 11
)

var(
	count int=0
)



func main(){
	a:=17
	b:=5
	fmt.Println(a/b)
	fmt.Println(a%b)
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)


	var e int = 10
	var f int = 7
	//fmt.Println(e+f) it won't work as it is typecasted to integer to integer 8 bytes
	//so we have to type cast this integer common data types
	fmt.Println(e+f)


	fmt.Println(a&b)
	fmt.Println(a|b)
	fmt.Println(a^b)
	fmt.Println(a&^b)

	l := 8 // 2^3
	fmt.Println(l<<3) // 2 ^ 3 * 2 ^ 3 = 2 ^ 6
	fmt.Println(l>>3) // 2 ^ 3 / 2 ^ 3 = 2 ^ 0
	fmt.Println("")



	g := 3.14
	//g = 13.7e72
	//g =2.1E14
	fmt.Printf("%v, %T",g,g)
	 
}
