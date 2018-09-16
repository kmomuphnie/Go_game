package main

import (
	"fmt"
	"strconv"
)

func main(){
	var boardsize int= 8
	var colour string
	//var AIFirst, AIFinish bool = false, false
	//var playerFinish bool = false
	//var stop bool = false
	//var boardIndex string = "abcdefgh"


	//fmt.Print("Enter the board dimension: \n")
	//fmt.Scanf("%d", &n)
	fmt.Printf("Computer plays (B_1/W_2) (B is the first to play) : \n")
	fmt.Scanf("%s", &colour)
	if string(colour[0]) == string('B') {//is B
		//AIFirst = true
	}

	//var board [8]string
	//for i:=0; i<boardsize; i++{
	//	board[i] = "________"//8x8 board
	//}
	var board [8][8]string
	for i:=0; i<boardsize; i++{
		for j := 0; j < boardsize; j++ {
			board[i][j] = "_"//8x8 board
		}

	}

	//board[0][3] = "B"
	initialBoard(board[0:][0:])//by reference use slice
	printBoard(board, boardsize)//by copy


}
//---------------------------------------------------------------------------------
//initialize the board
func initialBoard (board[][8] string){
	board[3][3] = "B"
	board[4][3] = "W"
	board[3][4] = "W"
	board[4][4] = "B"
}

//print the board
func printBoard(board[8][8] string, boardsize int){
	fmt.Print(" ")
	for i := 0; i < boardsize; i++{
		fmt.Print(" ",i, " ")
	}
	fmt.Println()

	for i := 0; i < boardsize; i++ {
		fmt.Print(i)
		for j := 0; j < boardsize; j++ {
			fmt.Print(" ",board[i][j]," ")
		}
		fmt.Println(string('a'+i))
	}

	fmt.Print(" ")
	for i := 0; i < boardsize; i++{
		fmt.Print(" ",string('a'+i), " ")
	}
	fmt.Println()
}

//inverse the color on tile
func inverseColor(color string)(string){
	if color == "W" {
		return "B"
	}
	return "W"
}

//check if tile is within bound
func withinBound(row int, col int) bool{
	if row >= 0 && row < 8 && col >=0 && col < 8 {
		return true
	}
	return false
}

//check whether move is valid in certain direction
func checkDirection(board[8][8] string, boardsize int, row int, col int,
					color string, deltaRow int, deltaCol int) bool{
	var i int = deltaRow
	var j int = deltaCol
	if board[row][col] == "_" && board[row +i][col +j] == inverseColor(color){
	   	for withinBound(row+i, col +j) {
			if board[row][col] == inverseColor(color) &&
			   board[row +deltaRow][col+deltaCol] == color	{
				return true
			}
			if board[row +deltaRow][col +deltaCol] == "_" {
				return false
			}
			row += deltaRow
			col += deltaCol
		}
	}
	return false

}


//check whether is a valid move
func isValidMove(board[8][8] string, boardsize int,color string, row int, col int, ) bool{
	if board[row][col] == "_" && (checkDirection(board, boardsize,row, col, color,1, 0) ||
							      checkDirection(board, boardsize,row, col, color,-1,0) ||
		                          checkDirection(board, boardsize,row, col, color, 0,1) ||
		                          checkDirection(board, boardsize,row, col, color,0,-1) ||
							      checkDirection(board, boardsize,row, col, color,1,1)  ||
							      checkDirection(board, boardsize,row, col, color,-1,1) ||
							      checkDirection(board, boardsize,row, col, color,1,-1) ||
					    	      checkDirection(board, boardsize,row, col, color,-1,-1))  {
		return true
	}
	return false
}

//function to get a list of valid moves and return both the movelist and the total validmoves
func getValidMoves(board[8][8] string, boardsize int,color string)([64][2]string, int){
	var validMoveNum int = 0
	var i,j int = 0,0
	var moveList [64][2] string

	for i = 0; i < boardsize; i++ 	{
		for j = 0; j < boardsize; j++ {
			if isValidMove(board, boardsize, color, i, j) {
				validMoveNum += 1
				moveList[validMoveNum - 1][0] = strconv.Itoa(i)
				moveList[validMoveNum - 1][1] = strconv.Itoa(j)
			}
		}
	}



}








//func test(board[][8] string){
//	if board[3][3] == "B" {
//		board[3][3] = "1"
//	}
//}







//func main()  {
//	test := "a b"
//	fmt.Scanf("%s", &test)
//
//	fmt.Println(test)
//	fmt.Println(string(test[0]))
//
//	fmt.Println(string(test[1]))
//}