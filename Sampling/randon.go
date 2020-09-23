package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println("Hello world")
	add(5,5)
	fmt.Println(sqrt(16))
	grade(41)
	fmt.Println(square(3))
	fmt.Println(BMI(45,80))

}

func add(a int ,b int  ){
	var sum int
	sum =a +b
	fmt.Println(sum)


}
func sqrt(c float64)float64{
	fmt.Println("The square root of",c,"is")
	return math.Sqrt(c)
}
func square(a int) int{
	fmt.Println("The square of ", a,"is")
	return a*a
}
func grade(mark int){
	
	var grade string

	if mark >=70{
		grade="A"
		fmt.Println("Got",grade,"Passed")
	}else if mark >=60{
		grade="B"
		fmt.Println("Got",grade,"Passed")
	}else if mark >=50{
		grade="C"
		fmt.Println("Got",grade,"Passed")
	}else if mark >=40{
		grade="D"
		fmt.Println("Got",grade,"Passed")
	}else{
		grade="E"
		fmt.Println("Got",grade,"Failed")
	}


}
func BMI(weight float64,height float64 )float64{
	return weight/height
}





