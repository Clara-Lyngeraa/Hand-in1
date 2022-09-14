// package main

// 	leftForkAnswer := make(chan string)
// 	ph0AswerLeft := make(chan string)
// 	rightForkAnswer := make(chan string)
// 	Ph0AnswerRight := make(chan string)

// func main(){
// 	go philosopher();
// 	go fork();

// }

// func philosopher (){
// 	for i := 0; i<3; i++{
// 		select {
// 		case answerLeft  = <- leftForkAnswer:
// 			select {
// 			 case answerRight  = <- leftForkAnswer:
// 				fmt.PrintLn("ph0 eating");
// 				ph0AswerLeft <- "i am done eating"
// 				ph0AswerRight <- "i am done eating"

// 			}
// 		}
// 	}
// }

// func fork () {
// 	var available bool = true;
// 	leftForkAnswer <- "I am available"
// 	rightForkAnswer <- "I am available"

// 	select {
// 	case request =  <- ph0AswerLeft:
// 			available = false;
// 			if(available){
// 				leftForkAnswer <- "I am available"
// 			}
// 		}
// }


