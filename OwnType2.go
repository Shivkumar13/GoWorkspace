package main


import "fmt"

var a int

type hotdog int 

var b hotdog 


func main() {

	a = 42 

 	b = 43

	fmt.Println(a)

	fmt.Printf("%T\n",a)

	fmt.Println(b)

	fmt.Printf("%T\n",b)

// ./OwnType2.go:27:3: cannot use b (type hotdog) as type int in assignment

      //	a = b 
      //	fmt.Println(a)
       // fmt.Printf("%T\n", b)


}
