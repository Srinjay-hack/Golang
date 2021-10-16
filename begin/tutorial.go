package main
import "fmt"
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
	var i float32
	i=42
	i+='a'
	fmt.Println(i)
	i=4.7
	var j int
	j=int(i)
	fmt.Println(j)
	//fmt.Println(i)
	//fmt.Printf("%v, %T",i,i)
	fmt.Println(i)
	var n bool=true
	fmt.Printf("%v %T\n",n,n)
	x := 1 == 1
	y := 2 == 4
	fmt.Printf("%v %T\n",x,x)
	fmt.Printf("%v %T\n",y,y)
	var z bool
	fmt.Printf("%v %T\n",z,z)

	

}
