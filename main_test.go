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