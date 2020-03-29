package main


import "fmt"

var a int

type hotdog int 

var s string = "shiva"

// string is Shiva

var b hotdog 

var t string

func main() {

	a = 42 

 	b = 43

	fmt.Println(a)

	fmt.Printf("%T\n",a)

	fmt.Println(b)



	fmt.Printf("%T\n",b)

	fmt.Println(s)

	t = fmt.Sprintf("%v\n",s)

	fmt.Println(t)

// ./OwnType2.go:27:3: cannot use b (type hotdog) as type int in assignment

      //	a = b 
      //	fmt.Println(a)
       // fmt.Printf("%T\n", b)


}
