package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"log"
)

type recievePack struct {
	Move string
	Map [8][8]string
}

type returnPack struct {
	Board     interface{}
	UserPoint int
	AIPoint int
	Winner string
}

//global for boarder
var board [8][8]string
var boardsize int
var color string
var AIFirst, AIFinish, playerFinish, stop bool


func main() {

	//initialize the board and all parameters
	GameInit()



	http.HandleFunc("/Reversi",ReversiReciever)
	http.HandleFunc("/ReversiInit",ReversiInit)

	err := http.ListenAndServe(":8007", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}

	//----------------game part-----------------------

	//for !(boardIsFull(board[0:][0:], boardsize)) && !(AIFinish && playerFinish) && !stop {
	//	if AIFirst {
	//		if !stop {
	//			AIFinish = computerMove(board[0:][0:], boardsize, color)
	//		}
	//		playerFinish = playerMove(board[0:][0:], boardsize, inverseColor(color))
	//	} else {
	//		playerFinish = playerMove(board[0:][0:], boardsize, inverseColor(color))
	//		if !stop && !(boardIsFull(board[0:][0:], boardsize)) {
	//			AIFinish = computerMove(board[0:][0:], boardsize, color)
	//		}
	//	}
	//	if boardIsFull(board[0:][0:], boardsize) {
	//		AIFinish = true
	//		playerFinish = true
	//	}
	//}
	//
	////win conditions
	//if stop {
	//	fmt.Println(color, "player wins.")
	//} else if getWinner(board[0:][0:], boardsize, color) > getWinner(board[0:][0:], boardsize, inverseColor(color)) {
	//	fmt.Println(color, "player wins.")
	//} else if getWinner(board[0:][0:], boardsize, color) < getWinner(board[0:][0:], boardsize, inverseColor(color)) {
	//	fmt.Println(inverseColor(color), "player wins.")
	//} else {
	//	fmt.Println("Draw!")
	//}

}

//---------------------------------------------------Http server-----------------------------------------------------------
func recieveData(w http.ResponseWriter, r *http.Request) recievePack {
	//decode package string->byte->struct
	strToByte := []byte(r.FormValue("first"))

	//convert to struct
	var pkg recievePack;
	err := json.Unmarshal(strToByte, &pkg);
	if err != nil {
		fmt.Println("ERROR", err);
	}

	return pkg;
}

func BoardReturn(w http.ResponseWriter,r *http.Request, input interface{}, winner string){
	fmt.Println("------Start Print Response------")

	//encode to byte[]

	userPoint := getPlayerPoint(board[0:][0:], 8, "B")
	AIPoint := getPlayerPoint(board[0:][0:], 8, "W")
	stringInfoInByte, err := json.Marshal(returnPack{input, userPoint, AIPoint, winner})
	//convert byte[] to string
	strConverted := string(stringInfoInByte)

	json.NewEncoder(w).Encode(string(strConverted))
	//check error
	if err != nil{
		fmt.Println("ERROR")
	}
	fmt.Println(stringInfoInByte)
	fmt.Println(strConverted)
}

//-----------------------------------------------------Router-----------------------------------------------------------

func ReversiReciever(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//get pkg from web
	recievePkg := recieveData(w, r)

	fmt.Println(recievePkg)
	board = recievePkg.Map

	var invalidMove bool
	playerFinish, invalidMove = playerMove(board[0:][0:], boardsize, inverseColor(color), recievePkg.Move)
	//if user play fire, dont do anything
	if invalidMove == true{
		return
	}

	AIFinish := computerMove(board[0:][0:], boardsize, color)
	full := boardIsFull(board[:][:], 8)
	fmt.Println(playerFinish, AIFinish, full)

	//check if game is finished
	if (playerFinish && AIFinish) || full{
		fmt.Println("Check Winner")
		//then check winner
		AIPoint := getPlayerPoint(board[0:][0:], boardsize, "W")
		playerPoint := getPlayerPoint(board[0:][0:], boardsize, "B")

		var winner string
		if AIPoint > playerPoint {
			winner = "AI"
		} else if AIPoint < playerPoint {
			winner = "Player"
		} else {
			winner = "Draw"
		}

		BoardReturn(w, r, board, winner)

	}else{
		//return the board as whole
		BoardReturn(w, r, board, "")
	}

	fmt.Println(recievePkg)
}

//send initial play board to web make sure sync for everthing
func ReversiInit(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//initialize the board and all parameters
	GameInit()

	//return the board as whole
	BoardReturn(w, r, board, "")
}

//---------------------------------------------------------------------------------
//initialize the board
func initialBoard(board [][8] string) {
	board[3][3] = "W"
	board[4][3] = "B"
	board[3][4] = "B"
	board[4][4] = "W"
}

//print the board
func printBoard(board [][8] string, boardsize int) {
	fmt.Print(" ")
	for i := 0; i < boardsize; i++ {
		fmt.Print(" ", i, " ")
	}
	fmt.Println()

	for i := 0; i < boardsize; i++ {
		fmt.Print(i)
		for j := 0; j < boardsize; j++ {
			fmt.Print(" ", board[i][j], " ")
		}
		fmt.Println(string('a' + i))
	}

	fmt.Print(" ")
	for i := 0; i < boardsize; i++ {
		fmt.Print(" ", string('a'+i), " ")
	}
	fmt.Println()
}

func GameInit() {
	boardsize = 8
	AIFirst, AIFinish = false, false
	playerFinish = false
	stop = false
	color = "W"

	fmt.Print("Enter the board dimension: \n")
	//fmt.Scanf("%d", &n)
	//fmt.Printf("Computer plays (B_1/W_2) (B is the first to play) : \n")
	//fmt.Scanf("%s", &color)
	//if string(color[0]) == string('B') {//is B
	//	AIFirst = true
	//}

	//var board [8]string
	//for i:=0; i<boardsize; i++{
	//	board[i] = "________"//8x8 board
	//}

	for i := 0; i < boardsize; i++ {
		for j := 0; j < boardsize; j++ {
			board[i][j] = "_" //8x8 board
		}

	}

	//board[0][0] = "_"
	//board[0][2] = "B"

	//board[0][3] = "B"
	initialBoard(board[0:][0:])          //by reference use slice
	printBoard(board[0:][0:], boardsize) //by ref
}

//inverse the color on tile
func inverseColor(color string) (string) {
	if color == "W" {
		return "B"
	}
	return "W"
}

//check if tile is within bound
func withinBound(row int, col int) bool {
	if row >= 0 && row < 8 && col >= 0 && col < 8 {
		return true
	}
	return false
}

//check whether move is valid in certain direction
func checkDirection(board [][8] string, boardsize int, row int, col int,
	color string, deltaRow int, deltaCol int) bool {
	i := deltaRow
	j := deltaCol

	//check up the direction moving out of range
	if !withinBound(row+i, col+j) {
		return false
	}

	if board[row][col] == "_" && board[row+i][col+j] == inverseColor(color) {
		for withinBound(row+i, col+j) {
			if board[row][col] == inverseColor(color) &&
				board[row+deltaRow][col+deltaCol] == color {
				return true
			}
			if board[row+deltaRow][col+deltaCol] == "_" {
				return false
			}
			row += deltaRow
			col += deltaCol
		}
	}
	return false

}

//check whether is a valid move
func isValidMove(board [][8] string, boardsize int, color string, row int, col int, ) bool {
	if board[row][col] == "_" && (
		checkDirection(board[0:][0:], boardsize, row, col, color, 1, 0) ||
		checkDirection(board[0:][0:], boardsize, row, col, color, -1, 0) ||
		checkDirection(board[0:][0:], boardsize, row, col, color, 0, 1) ||
		checkDirection(board[0:][0:], boardsize, row, col, color, 0, -1) ||
		checkDirection(board[0:][0:], boardsize, row, col, color, 1, 1) ||
		checkDirection(board[0:][0:], boardsize, row, col, color, -1, 1) ||
		checkDirection(board[0:][0:], boardsize, row, col, color, 1, -1) ||
		checkDirection(board[0:][0:], boardsize, row, col, color, -1, -1)) {
		return true
	}
	return false
}

//flip the tiles where move is done
func switchTiles(board [][8] string, boardsize int, color string, row int, col int, deltaRow int, deltaCol int) {
	var i int = row + deltaRow
	var j int = col + deltaCol
	for withinBound(i, j) {
		board[i][j] = color
		i += deltaRow
		j += deltaCol
		if board[i][j] == color {
			break;
		}
	}
}

//excecute move at point row,col
func executeMove(board [][8] string, boardsize int, color string, row int, col int) {
	//if row == 2 && col == 3{
	//	fmt.Println("------IN ExecuteMove 1------")
	//	printBoard(board, 8)
	//}

	var deltaRow, deltaCol int
	for deltaRow = -1; deltaRow <= 1; deltaRow++ {
		for deltaCol = -1; deltaCol <= 1; deltaCol ++ {
			if !(deltaCol == 0 && deltaRow == 0) {
				if checkDirection(board[0:][0:], boardsize, row, col, color, deltaRow, deltaCol) {
					switchTiles(board[0:][0:], boardsize, color, row, col, deltaRow, deltaCol)
					//if row == 2 && col == 3{
					//	fmt.Println("------Flip!------")
					//}
				}
			}
		}
	}

	//if row == 2 && col == 3{
	//	fmt.Println("------IN ExecuteMove 2------")
	//	printBoard(board, 8)
	//}

	board[row][col] = color
}

//check whether the board is full
func boardIsFull(board [][8] string, boardsize int) bool {
	var i, j int
	for i = 0; i < boardsize; i++ {
		for j = 0; j < boardsize; j++ {
			if board[i][j] == "_" {
				return false
			}
		}
	}
	return true
}

//generate the score of tiles across the board
func getScore(board [][8] string, boardsize int, color string, gameOver bool) int {
	var score int = 0
	var i, j int
	for i = 0; i < boardsize; i++ {
		for j = 0; j < boardsize; j++ {
			if board[i][j] == color {
				if (i == 0 && j == 0) ||
					(i == 0 && j == boardsize-1) ||
					(i == boardsize-1 && j == 0) ||
					(i == boardsize-1 && j == boardsize-1) &&
						(!gameOver) {
					score += boardsize * boardsize * boardsize * boardsize
				} else if (i == 1 && j == 1) ||
					(i == 1 && j == boardsize-2) ||
					(i == boardsize-2 && j == 1) ||
					(i == boardsize-2 && j == boardsize-2) {
					score -= boardsize * boardsize * boardsize
				} else if (i == 0 || j == 0 || i == boardsize-1 || j == boardsize-1) &&
					boardsize >= 16 {
					score += boardsize
				} else if i == 0 || j == 0 || i == boardsize-1 || j == boardsize-1 {
					score += boardsize
				} else if (i == 1 || j == 1 || i == boardsize-2 || j == boardsize-2) &&
					boardsize >= 8 {
					score += boardsize / 2
				} else {
					score += 1
				}
			}
		}
	}
	return score
}

//counts real score value of tiles on  the board
func getPlayerPoint(board [][8] string, boardsize int, color string) int {
	var i, j int
	var score int = 0
	for i = 0; i < boardsize; i++ {
		for j = 0; j < boardsize; j++ {
			if board[i][j] == color {
				score += 1
			}
		}
	}
	return score
}

//get max of 2 values
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//get min of 2 values
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//copy to a new board state
func generateState(board [][8] string, newBoard [][8] string, boardsize int, row int, col int, color string) {

	for i := 0; i < boardsize; i++ {
		for j := 0; j < boardsize; j++ {
			//if !(i == row && j == col) {
				newBoard[i][j] = board[i][j]
			//}
		}
	}

	//if row == 2 && col == 3{
	//	fmt.Println("------before ExecuteMove------")
	//	printBoard(newBoard, 8)
	//}

	executeMove(newBoard, boardsize, color, row, col)

	//if row == 2 && col == 3{
	//	fmt.Println("------after ExecuteMove------")
	//	printBoard(newBoard, 8)
	//}
}

//function to get a list of valid moves and return both the movelist and the total validmoves
func getValidMoves(board [][8] string, boardsize int, color string) ([64][2]int, int) {
	var numValidMoves int = 0
	var i, j int = 0, 0
	var moveList [64][2] int

	for i = 0; i < boardsize; i++ {
		for j = 0; j < boardsize; j++ {
			if isValidMove(board[0:][0:], boardsize, color, i, j) {
				numValidMoves += 1
				moveList[numValidMoves-1][0] = i
				moveList[numValidMoves-1][1] = j
			}
		}
	}


	return moveList, numValidMoves
	//var numValidMoves int = 0
	//var i,j int = 0,0
	//var moveList [64][2] string
	//
	//for i = 0; i < boardsize; i++ 	{
	//	for j = 0; j < boardsize; j++ {
	//		if isValidMove(board[0:][0:], boardsize, color, i, j) {
	//			numValidMoves += 1
	//			moveList[numValidMoves - 1][0] = strconv.Itoa(i)
	//			moveList[numValidMoves - 1][1] = strconv.Itoa(j)
	//		}
	//	}
	//}
	//return moveList, numValidMoves
}

func miniMax(boardState [][8] string, boardsize int, depth int, startDepth int,
	alpha int, beta int, color string, isMaxing bool, moveRowCol []int) int {

	var row, col int

	moveList, numValidMoves := getValidMoves(boardState[0:][0:], boardsize, color)

	//fmt.Println(boardsize, depth, startDepth, alpha, beta, color, isMaxing, moveRowCol, numValidMoves)

	if depth == 0 && (startDepth%2 != 0) {
		return getScore(boardState[0:][0:], boardsize, inverseColor(color), false) -
			getScore(boardState[0:][0:], boardsize, color, false)
	} else if isMaxing {
		maxScore := -10000000
		var preScore int

		//fmt.Println("------start max Looping------")

		for i := 0; i < numValidMoves; i++ {
			row = moveList[i][0]
			col = moveList[i][1]
			var newBoard [8][8]string
			generateState(boardState[0:][0:], newBoard[0:][0:], boardsize, row, col, color)
			preScore = maxScore

			//printBoard(newBoard[0:][0:], 8)

			//here recursively find max for further move
			maxScore = max(maxScore, miniMax(newBoard[0:][0:], boardsize, depth-1, startDepth, alpha, beta,
					inverseColor(color), false, moveRowCol))

			alpha = max(alpha, maxScore)
			if depth == startDepth && maxScore != preScore {
				moveRowCol[0] = row
				moveRowCol[1] = col
			}
			if beta <= alpha {
				return beta
			}
		}

		//fmt.Println("------end max Looping------")
		return maxScore
	} else {
		var minScore int = 10000000

		//fmt.Println("------start min Looping------")

		for i := 0; i < numValidMoves; i++ {
			row = moveList[i][0]
			col = moveList[i][1]
			var newBoard [8][8]string
			generateState(boardState[0:][0:], newBoard[0:][0:], boardsize, row, col, color)

			//printBoard(newBoard[0:][0:], 8)

			minScore = min(minScore,
				miniMax(newBoard[0:][0:], boardsize, depth-1, startDepth, alpha, beta,
					inverseColor(color), true, moveRowCol))
			beta = min(beta, minScore)
			if beta <= alpha {
				return alpha
			}
		}

		//fmt.Println("------end min Looping------")
		return minScore
	}

}

//Process computer takes to make a move
func computerMove(board [][8] string, boardsize int, color string) bool {
	var numValidMovesComp int
	compMoveRowCol := []int{0, 0} //initialize the slice
	depth := 5             //depth changes for minimax out of the time consumed
	var moveList [64][2] int
	var a int

	moveList, numValidMovesComp = getValidMoves(board[0:][0:], boardsize, color)

	if numValidMovesComp > 0 {
		a = moveList[0][0]
		//fmt.Println("------start minMax------")
		miniMax(board[0:][0:], boardsize, depth, depth, -100000, 100000,
				color, true, compMoveRowCol)
		//fmt.Println("------finish minMax------")

		fmt.Println("Computer place", color, "at", compMoveRowCol[0], compMoveRowCol[1])
		executeMove(board[0:][0:], boardsize, color, compMoveRowCol[0], compMoveRowCol[1])
		printBoard(board[0:][0:], boardsize)
		return false
	} else {
		fmt.Println(color, "player has no valid move.")
		return true
	}
	if a == 0 {
		fmt.Print("something wrong in computerMove")
	}
	return false
}

//let player make a move
func playerMove(board [][8] string, boardsize int, color string, playerMoveindex string) (bool, bool) { //return playerfinish
	numValidMovePlayer := 0
	playerMoveRowCol := []int{0, 0}
	var moveList [64][2] int
	var a int
	var invalidMove bool

	//fmt.Println("playerMoveindex:",playerMoveindex, color)
	moveList, numValidMovePlayer = getValidMoves(board[0:][0:], boardsize, color)
	if numValidMovePlayer > 0 {
		a = moveList[0][0]

		//invalidMove := true
		//for invalidMove {

			//checking invaild move
			//if invalidMove{
			//	break
			//}


			//fmt.Print("Enter move for color", color, "(RowCol): ")
			inputStr := playerMoveindex //get the input"12" to 1 2 playerMoveindex
			//fmt.Scanf("%s", &inputStr)
			var inputRowCol [2]int
			inputRowCol[0], _ = strconv.Atoi(string(inputStr[0]))
			inputRowCol[1], _ = strconv.Atoi(string(inputStr[1]))
			playerMoveRowCol[0] = inputRowCol[0]
			playerMoveRowCol[1] = inputRowCol[1]

			if !(isValidMove(board[0:][0:], boardsize, color, playerMoveRowCol[0], playerMoveRowCol[1])) {
				fmt.Println("Invalid move. Plz reinput a valid move")
				invalidMove = true
			} else {
				executeMove(board[0:][0:], boardsize, color, playerMoveRowCol[0], playerMoveRowCol[1])
				printBoard(board[0:][0:], boardsize)
				invalidMove = false
			}
		//}
	} else {
		fmt.Println(color, "player has no valid move.")
		return true, true // playerfinish = true
	}

	if a == 0 {
		fmt.Print("something wrong in playerMove")
	}
	return false, invalidMove
}
