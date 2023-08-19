package main

import "fmt"

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
}
