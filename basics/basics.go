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
	
	fmt.Println("=========================================================")
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

	// Structs better to not use them in the functions declare them else where
	fmt.Println("=========================================================")

	myCar := Car{Engine{volume: 1.2, cylinders: 6} }
	myCarEngineVolume := myCar.engine.cylinders
	fmt.Println(myCar.calculateSomething())
	fmt.Printf("%v - It's the engine volume of my car", myCarEngineVolume)

	//  ANONYMOUS struct
	// myCar := struct {
	//   Make string
	//   Model string
	// } {
	//   Make: "tesla",
	//   Model: "model 3",
	// }
	fmt.Println(	getUserExpenses(User{
		userName: "Chad",
		email: "chad@gmail.com",
		costPerSMS: 1.2,
		totalMessages: 22,
	}))
	fmt.Println(getUserExpenses(User{
		userName: "Chad Player",
		email: "chadplayer@gmail.com",
		costPerSMS: 1,
		totalMessages: 432,
	}))
	fmt.Println("=========================================================")

	// if 12 < 11 {
	// 	return errors.New("STOP IT, IT'S NOT POSSIBLE")
	// }
	for i:= 0; i < 10; i++ {
		fmt.Printf("Counter: %v", i)
	}

	// totalCost := 74 
 
	//While loop in go
	// for totalCost > 1 {
	// 	// Do something
	// } 
	fmt.Println("=================================================================")
	fmt.Println("Lists:")
	itemList := [3]string{"Javascript 5/10", "I love go lang","Rust is veeery interesting"}
	fmt.Println(itemList[0])
	fmt.Println("Slices:")
	goodLangSlices := itemList[1:]
	allLanguage := itemList[:]
	fmt.Println(goodLangSlices)
	fmt.Println(allLanguage)

	newSlice := make([]int, 10)
	fmt.Println(newSlice)
	fmt.Println("Slice len", len(newSlice))

	println(sumSlice(1,2,3,4,6))


	newLanguages := append(allLanguage, "Zig")
	fmt.Println(newLanguages)


	// Maps
	fmt.Println("===========================================================")
	fmt.Println("MAPS")

	map1 := make(map[string]int)
	map1["Key1"] = 1
	fmt.Println(map1)

	map2 := map[string]int{
		"KEY1": 12,
		"KEY2": 13,
	}
	fmt.Println(map2)
	delete(map2, "KEY1")

	elem, ok := map2["KEY2"]

	fmt.Println(elem, ok)

}

func sumSlice(nums ...int)int {
	total := 0	
	for i :=0; i < len(nums); i++ {
	total = total + nums[i]
	}
	return total 
}

func getUserExpenses(u userData) string {
	return fmt.Sprintf("%s, %.1f - Here is customer expenses", u.getUserData(),u.getTotalMessageCost())
}

type userData interface {
	getTotalMessageCost() float64
	getUserData() string
} 
type User struct {
	userName string
	email string
	costPerSMS float64
	totalMessages int
}

func (u User) getTotalMessageCost() float64 {
	return u.costPerSMS * float64(u.totalMessages)
} 
func (u User) getUserData() string {
	return fmt.Sprintf("User: %s, Email: %s", u.userName, u.email) 
} 



type Player struct {
	playerName string
	email string
	costPerSMS float64
	totalMessages int 
}


func (p Player) getTotalMessageCost() float64 {
	return p.costPerSMS * float64(p.totalMessages)
} 

func (p Player) getUserData() string {
	return fmt.Sprintf("User: %s, Email: %s", p.playerName, p.email)
} 
type Engine struct {
	volume float32; 
	cylinders int	
} 
type Car struct {
	engine Engine	
} 
func (c Car) calculateSomething() int{	
	return c.engine.cylinders + int(c.engine.volume)
}
// Struct test
// type messageToSend struct {
// 	message   string
// 	sender    user
// 	recipient user
// }

// type user struct {
// 	name   string
// 	number int
// }

// func canSendMessage(mToSend messageToSend) bool {
// 	// ?
// 	sender := mToSend.sender
// 	recipient := mToSend.recipient 

// 	if sender.name == "" || recipient.name == "" {
// 		return false
// 	}

// 	if sender.number == 0 || recipient.number == 0 {
// 		return false
// 	} 
// 	return true
// }

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