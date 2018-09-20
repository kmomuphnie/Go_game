package main

import (
	"fmt"
	"testing"
)

func Test_CheckValidMove_1(t *testing.T) {
	var board [8][8]string
	for i:=0; i< 8; i++{
		for j:=0; j < 8; j++{
			board[i][j] = "_"
		}
	}

	board[0][3] = "B"
	board[2][3] = "B"
	board[4][3] = "B"
	board[3][4] = "B"
	board[4][5] = "B"

	board[1][3] = "W"
	board[3][3] = "W"
	board[2][4] = "W"
	board[4][4] = "W"

	printBoard(board[:][:], 8)

	moveList, validMoveCount := getValidMoves(board[:][:], 8, "B")

	fmt.Println(moveList)
	fmt.Println("validMoveCount:",validMoveCount)


	fmt.Println("finish Test_CheckValidMove_1")
}


func Test_executeMove_1(t *testing.T){
	var board [8][8]string
	for i:=0; i< 8; i++{
		for j:=0; j < 8; j++{
			board[i][j] = "_"
		}
	}

	board[3][3] = "W"
	board[3][4] = "B"
	board[4][4] = "W"
	board[4][3] = "B"

	printBoard(board[:][:], 8)

	executeMove(board[:][:], 8, "B", 2, 3)

	printBoard(board[:][:], 8)

	fmt.Println("finish Test_executeMove_1")
}

func Test_getValidMoves_1(t *testing.T){
	var board [8][8]string
	for i:=0; i< 8; i++{
		for j:=0; j < 8; j++{
			board[i][j] = "_"
		}
	}

	board[3][3] = "W"
	board[3][4] = "B"
	board[4][4] = "W"
	board[4][3] = "B"

	moveList, numValidMoves := getValidMoves(board[0:][0:], 8, "B")
	fmt.Println(moveList)
	fmt.Println(numValidMoves)

	fmt.Println("finish Test_executeMove_1")
}

func Test_getValidMoves_2(t *testing.T){
	var board [8][8]string
	for i:=0; i< 8; i++{
		for j:=0; j < 8; j++{
			board[i][j] = "_"
		}
	}

	board[0][0] = "W"
	board[0][1] = "W"
	board[0][2] = "W"
	board[0][3] = "W"
	board[0][4] = "W"
	board[0][5] = "W"
	board[0][6] = "W"
	board[0][7] = "W"

	board[1][0] = "W"
	board[1][1] = "B"
	board[1][2] = "B"
	board[1][3] = "W"
	board[1][4] = "W"
	board[1][5] = "W"
	board[1][6] = "W"
	board[1][7] = "W"


	board[2][0] = "W"
	board[2][1] = "W"
	board[2][2] = "B"
	board[2][3] = "B"
	board[2][4] = "B"
	board[2][5] = "W"
	board[2][6] = "W"
	board[2][7] = "W"

	board[3][0] = "W"
	board[3][1] = "W"
	board[3][2] = "W"
	board[3][3] = "B"
	board[3][4] = "W"
	board[3][5] = "W"
	board[3][6] = "B"
	board[3][7] = "W"

	board[4][0] = "W"
	board[4][1] = "W"
	board[4][2] = "W"
	board[4][3] = "W"
	board[4][4] = "B"
	board[4][5] = "B"
	board[4][6] = "B"
	board[4][7] = "W"

	board[5][0] = "W"
	board[5][1] = "W"
	board[5][2] = "W"
	board[5][3] = "B"
	board[5][4] = "B"
	board[5][5] = "B"
	board[5][6] = "B"
	board[5][7] = "B"

	board[6][0] = "W"
	board[6][1] = "W"
	board[6][2] = "B"
	board[6][3] = "B"
	board[6][4] = "B"
	board[6][5] = "B"
	board[6][6] = "B"
	board[6][7] = "B"

	//board[6][0] = "W"
	//board[6][1] = "W"
	//board[6][2] = "B"
	board[7][3] = "W"
	board[7][4] = "W"
	board[7][5] = "W"
	board[7][6] = "W"
	//board[6][7] = "B"

	printBoard(board[0:][0:], 8)

	moveList, numValidMoves := getValidMoves(board[0:][0:], 8, "B")
	fmt.Println(moveList)
	fmt.Println(numValidMoves)

	fmt.Println("finish Test_executeMove_2")



}