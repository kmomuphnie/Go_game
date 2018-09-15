package main

import "fmt"

func main(){
	fmt.Printf("hello, world\n")
	var boardsize int = 8
	var colour string
	//var AIFirst, AIFinish bool = false, false
	//var playerFinish bool = false
	//var stop bool = false
	//var boardIndex string = "abcdefgh"


	//fmt.Print("Enter the board dimension: \n")
	//fmt.Scanf("%d", &n)

	fmt.Printf("Computer plays (B_1/W_2) (B is the first to play) : \n")
	fmt.Scanf("%s", &colour)
	//fmt.Println(string(colour[0]))
	//fmt.Println(string(colour[0]))
	if string(colour[0]) == string('B') {//is B
		//AIFirst = true


	}

	var board [8]string
	for i:=0; i<boardsize; i++{
		board[i] = "________"//8x8 board
	}

	board[0][3] = byte('B')

	for i := 0; i < boardsize; i++ {
		fmt.Print(string('a'+i))
		for j := 0; j < boardsize; j++ {
			fmt.Print(" ",string( (board[i])[j] )," ")
		}
		fmt.Println(i)
	}



}

//func main()  {
//	test := "a b"
//	fmt.Scanf("%s", &test)
//
//	fmt.Println(test)
//	fmt.Println(string(test[0]))
//
//	fmt.Println(string(test[1]))
//}