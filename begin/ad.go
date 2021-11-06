package main
import ("fmt")

func main(){


	var i int 
	i=0
	fmt.Println(i)
	var n complex64=1+2i
	var x complex64=2+3i
	//fmt.Printf("%v, %T",n*x,n*x);
	fmt.Println(n+x)
	fmt.Println(n-x)
	fmt.Println(n/x)
	fmt.Println(n*x)
	fmt.Println(real(n*x),imag(n*x))

	var b complex128=complex(6,7)
	fmt.Println(b)

	s:="hello this is th fam!!"
	fmt.Printf("%v,%T\n",s,s)
	// Ascii value hardcoded
	bt := []byte(s)
	fmt.Printf("%v,%T\n",bt,bt)
	

}
