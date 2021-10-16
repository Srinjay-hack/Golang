package main
import "fmt"
var x int =7
func main(){
	var i int
	i=42
	i+='a'
	//fmt.Println(i)
	fmt.Printf("%v, %T",i,i)
	fmt.Println("")
	fmt.Printf("%v, %T",x,x)

}
