package Introduction

import "fmt"


func Variable(){
	var invesmtAmount = 30;
	var expectedReturnRate =  5.5
	var years =  10
	var futureValue =  invesmtAmount*int(expectedReturnRate)/years
	fmt.Println("The Investment Amount is :",futureValue)
}