// // package main

// // import (
// // 	"fmt"
// // 	"math/rand"
// // )

// // func main() {

// // 	randomNumber := rand.Intn(50)
// // 	for i := 0; i < 5; i++ {
// // 		fmt.Println("guess the number")
// // 		inputNumber, _ := fmt.Scanln(int)
// // 		if inputNumber > randomNumber {
// // 			fmt.Println("Your number is greater than the random number")
// // 		} else if inputNumber < randomNumber {
// // 			fmt.Println("Your number is smaller than the random number")
// // 		} else {
// // 			fmt.Println("great you have found the random number")
// // 			break
// // 		}
// // 	}
// // }
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// 	// to genereate random number each time we need to change seed value
// 	rand.Seed(time.Now().UnixNano())
// 	//random number generation
// 	randomSecretNumber := rand.Intn(50) + 1
// 	fmt.Println(randomSecretNumber)

// 	fmt.Println("Guess a number between 1 and 50")
// 	fmt.Println("Please type your guessed number")

// 	for i := 0; i < 5; i++ {

// 		// declaring variable reader which is initialised in taking input
// 		//returns a method NewReader
// 		reader := bufio.NewReader(os.Stdin)

// 		// takes input (string) till it scans the first occuring \n
// 		input, _ := reader.ReadString('\n')

// 		//err check
// 		// if err != nil {
// 		// 	fmt.Println("An error occured while reading input. Try again", err)
// 		// 	continue
// 		// }

// 		//removing the \n from the string using TrimSuffix
// 		input = strings.TrimSuffix(input, "\n")

// 		// converting string into int  using Atoi()method from strconv
// 		inputInt, _ := strconv.Atoi(input)

// 		//checking error if any
// 		// if err != nil {
// 		// 	fmt.Println("Invalid input. Please enter an integer value")
// 		// 	continue
// 		// }
// 		// printing the string after converting to int
// 		fmt.Println("Your guess is", inputInt)

// 		//now applying the conditions to the input integer
// 		if inputInt > randomSecretNumber {
// 			fmt.Println("Your guess is bigger than the secret random number. Try again")
// 		} else if inputInt < randomSecretNumber {
// 			fmt.Println("Your guess is smaller than the secret random number. Try again")
// 		} else {
// 			fmt.Println("You guessed right after", i, "attempts")
// 			break
// 		}
// 	}

// }
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// to genereate random number each time,	 we need to change seed value
	rand.Seed(time.Now().UnixNano())
	//random number generation
	randomSecretNumber := rand.Intn(50) + 1
// 	fmt.Println(randomSecretNumber)

	fmt.Println("Guess a number between 1 and 50")
	fmt.Println("Please type your guessed number")

	for i := 0; i < 5; i++ {

		var guessedNumber int
		fmt.Scan(&guessedNumber)
		// fmt.Println("read number", i, "from stdin")

		// declaring variable reader which is initialised in taking input
		//returns a method NewReader
		// reader := bufio.NewReader(os.Stdin)

		// takes input (string) till it scans the first occuring \n
		// input, _ := reader.ReadString('\n')

		//err check
		// if err != nil {
		// 	fmt.Println("An error occured while reading input. Try again", err)
		// 	continue
		// }

		//removing the \n from the string using TrimSuffix
		// input = strings.TrimSuffix(input, "\n")

		// converting string into int  using Atoi()method from strconv
		// inputInt, _ := strconv.Atoi(input)

		//checking error if any
		// if err != nil {
		// 	fmt.Println("Invalid input. Please enter an integer value")
		// 	continue
		// }
		// printing the string after converting to int
		fmt.Println("Your guess is", guessedNumber)
		if guessedNumber >= 1 && guessedNumber <= 50 {

			//now applying the conditions to the input integer
			if guessedNumber > randomSecretNumber {
				fmt.Println("Your guess is bigger than the secret random number. Try again")
			} else if guessedNumber < randomSecretNumber {
				fmt.Println("Your guess is smaller than the secret random number. Try again")
			} else {
				fmt.Println("You guessed right after", i, "attempts")
				break
			}
		} else {
			fmt.Println("enter a number between 1 and 50 only")
		}
	}

}
