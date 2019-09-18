package main

import "fmt"

func main() {
	sum := 0
	for i :=0; i<10; i++ { // the only looping construct. No parens
		sum += i
	}

	fmt.Println(sum)


	sum = 1
	for ; sum < 1000; {  // init and post statements are optional
		sum+= sum
	}

	fmt.Println(sum)



	sum = 1
	for sum < 1000 { //semicolons can be dropped
		sum+= sum
	}

	fmt.Println(sum)


	for { // equivalent of while(true){
		sum += sum
		if sum > 10000{ // braces are required
			break;
		}
	}

	fmt.Println(sum)

}



