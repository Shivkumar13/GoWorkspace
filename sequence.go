package main

import "fmt"

func main() {

	fmt.Println("Hello World in Go!")

    foo()

	fmt.Println("executed foo")

        for i := 1; i<100 ; i++ {
            
              if i %2 == 0{
           fmt.Println(i)
     }
     
   }
 
    bar()

}

 func foo() {

   fmt.Println("I'm in foo")

}


 func bar(){

   fmt.Println("and then we exited")

 }
//control flow

// sequence

//loop; iterative 

//conditional


