package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	var name string = "Kris"
	var age int = 13
    fmt.Println("My name is", name, "I'm", age)

	
	// initialize variables here
	var smsSendingLimit int = 32  	
	var costPerSMS float32 = .2
	var hasPermission bool = true
	var username string = "Someusername"

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)


	// Short version of variable declaration  
	congrats := "Happy Birthday!"
	fmt.Println(congrats)

	averageOpenRate, displayMessage := .23, "is the average open rate of your messages" 
	fmt.Println(averageOpenRate, displayMessage)	

	
	// Reassigning type from float to int. 
	temperatureFloat := 13.44 
	temperatureInt := int16(temperatureFloat)

    fmt.Println("Temperature float is", temperatureFloat, "Temperature int is", temperatureInt)

	// Constants defined like so
	const premiumPlan = "Premium Plan."  
	fmt.Println(premiumPlan)

	const basicPlan = "Basic Plan."
	const allPlans = "We have" + premiumPlan + "and" + basicPlan 
	fmt.Println(allPlans)

	// How to parse edit strings
	fullName := "John Loosie";
	familyMemberCount := 12 
	
	msg := fmt.Sprintf("Hey %v, It's great that you have %v family members", fullName, familyMemberCount)

	fmt.Println(msg)
	// Work with conditions 

	maxMessageLength := 32	
	messageLength := 12
	if maxMessageLength > messageLength {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Max message length reached")
	}
	

	// also you can define initial statement  
	if basicLength:= 33; basicLength > maxMessageLength {
		fmt.Println("basic length is greater then max")
	}
	

	// Functions 
	// fmt.Fprintln(concat("Hello", "World"))
	fmt.Println(sub(12, 3))
	foo1(2, 3, "John")
		// SO you need to reassign x because GO operating not by using referenced but by using values 
	x := 44
	x = increment(x)
	fmt.Println(x)
    // Here the second value is ignored and cleaned from the memory. It's not just ignored by linter.
	firstName, _:= getNames()	
	// Also a great thing is that go does not allow us to have unused variable it will throw an array during the compilation 
	fmt.Println(firstName);

	yearsUntilAdult, yearUntilDrinking := yearsUntilEvents(15)
	fmt.Println(yearUntilDrinking, yearsUntilAdult)

}

func getNames()(string, string) {
	return "John", "Smith"
}

func yearsUntilEvents(age int)(yearsUntilAdult, yearsUntilDrinking int) {
	// Here we will return default value for an int which is 0
	return yearsUntilAdult, yearsUntilDrinking
}



func increment(x int) int {
	x++
	return x
}

func sub(num1, num2 int) int{
	return num1 - num2
} 

func foo1(num1, num2 int, name string) {
	sum := num1 + num2
	fmt.Println("Sum:", sum, "Name:" ,name)
}