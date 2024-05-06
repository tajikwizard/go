package main

import "fmt"


func main(){

	//conse.log()
	fmt.Println("Hello World")


	//let name = "Abdulaziz";
	var name string = "Abdulaziz"
	fmt.Println(name)

	//let age = 24
	var age int = 24
	fmt.Println(age)

	//let isMarried = false
	var isMarried bool = false
	if(isMarried) {
		fmt.Println("Married")
	}else{
		fmt.Println("Not Married")
	}
	var friends = [3]string{"Abdulaziz","Abdurasul","Manon"}
	fmt.Println(friends)

	for i := 0; i < len(friends); i++ {
		fmt.Printf(friends[i])
	}



	// check if it has Manon in the array 

	hasManon :=false
	for i:=0; i < len(friends); i++ {
		if(friends[i] == "Manon"){
			hasManon = true
			break
		}
	}

	if hasManon{
		fmt.Println("Has Manon")
	}else{
		fmt.Println("Does not have Manon")
	}
}