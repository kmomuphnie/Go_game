#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>




//creates a starting board
void generateStartBoard(int n, char newBoard[][26]){
    int row,col;
    for(row = 0; row < n; ++row){
        for( col = 0; col < n; ++col){
            //first special row
            if((row == n/2 -1 && col == n/2 )|| (row == n /2 && col == n/2-1)){
                newBoard[row][col] = 'B';
            }
            else if ((row == n/2 && col == n/2) || (row == n/2-1 && col == n/2 -1))
                newBoard[row][col] = 'W';
            else
                newBoard[row][col] = '_';
        }
    }
}


//prints board
void printBoard(char board[][26], int n){
    int row, col,i;
    printf("  ");
    for(i = 0; i < n; ++i){
        printf("%c ", 'a'+ i);
    }
    printf("\n");
    for(row = 0; row < n; ++row){
        printf("%c ", 'a' + row);
        for(col = 0; col < n; ++col){
            printf("%c ", board[row][col]);
        }
        printf("\n");//new line so board looks like a board
    }
}

//returns opposite color W->B B->W
char oppositecolor(char color){
    if(color == 'W')
        return 'B';
    return 'W';
}

//checks to see if (row,col) is in bounds
bool positionInBounds(char board[][26], int n, char row, char col){
    if((row-'a')>=0 && (col-'a') >=0 && (row-'a') < n && (col - 'a') < n){
        return true;
    }
    return false;
}


//checks to see if move is legal in specified direction
bool checkLegalInDirection(char board[][26], int n, char row, char col,
                           char color, int deltaRow, int deltaCol){
    
    int i = deltaRow;
    int j = deltaCol;

    if(board[row-'a'][col-'a'] == '_' && board[row-'a' +i][col-'a'+j] == oppositecolor(color)){
	//printf("%d %d %c %c\n", i, j, row, col);
        while(positionInBounds(board,n,row+i,col+j)){
            if(board[row-'a'][col-'a'] == oppositecolor(color) &&
               board[row-'a'+deltaRow][col-'a'+deltaCol] == color){
                return true;
            }
            if(board[row-'a'+deltaRow][col-'a'+deltaCol]=='_'){
                return false;
            }
            row+=deltaRow;
            col+=deltaCol;
        }
    }
    return false;
    
}



//checks to see if a move is valid
bool isValidMove(char board[][26], int n, char color, char row, char col){
    if(board[row-'a'][col-'a'] == '_' &&(
                                         checkLegalInDirection(board, n, row, col, color,1,0)
                                         || checkLegalInDirection(board, n, row, col, color,-1,0)
                                         || checkLegalInDirection(board, n, row, col, color,0,1)
                                         || checkLegalInDirection(board, n, row, col, color,0,-1)
                                         || checkLegalInDirection(board, n, row, col, color,1,1)
                                         || checkLegalInDirection(board, n, row, col, color,-1,1)
                                         || checkLegalInDirection(board, n, row, col, color,1,-1)
                                         || checkLegalInDirection(board, n, row, col, color,-1,-1))){
        return true;
    }
    return false;
}

//gets a list of valid moves for a board and for player of specified color
char** getValidMoves(char board[][26], int n, char color, int * validMoves){
    int numValidMoves = 0;
    char i,j;
    char** moveList = malloc(sizeof(char) * numValidMoves);
    for(i = 'a'; i < n+'a'; ++i){
        for(j = 'a'; j < n+'a'; ++j){
            if(isValidMove(board,n,color, i,j)){
                
                numValidMoves++;
                moveList = realloc(moveList, sizeof(char*) * numValidMoves);
                if(moveList == NULL){
                    printf("moveList has failed to reallocate\n");
                }
                else{
                    //printf("hi\n");
                    moveList[numValidMoves - 1] = malloc(sizeof(char)*2);
                    moveList[numValidMoves-1][0] = i;
                    moveList[numValidMoves-1][1] = j;
                }
            }
        }
    }
    *validMoves = numValidMoves;
	


    return moveList;
}

//switches tiles where move is excecuted
void switchTiles(char board[][26], int n, char color, char row, char col,
                 int deltaRow, int deltaCol){
    int i = row - 'a'+deltaRow;
    int j = col - 'a'+deltaCol;
    while(positionInBounds(board,n, i+'a',j+'a')){
        board[i][j] = color;
        i += deltaRow;
        j += deltaCol;
        if(board[i][j] ==color){
            break;
        }
    }
}

//excecutes move at (row,col)
void executeMove(char board[][26], int n, char color, char row, char col){
    int deltaRow, deltaCol;
    for(deltaRow = -1; deltaRow <= 1; ++deltaRow){
        for(deltaCol = -1; deltaCol <= 1; ++deltaCol){
            if(!(deltaCol ==0 && deltaRow ==0)){
                if(checkLegalInDirection(board, n, row, col, color,deltaRow,deltaCol)){
                    switchTiles(board, n, color, row, col,deltaRow, deltaCol);
                }
            }
        }
    }
    board[row-'a'][col-'a'] = color;
}

//checks to see if board is full
bool boardIsFull(char board[][26], int n){
    int i,j;
    for(i = 0; i < n; ++i){
        for(j = 0; j < n; ++j){
            if(board[i][j] == '_'){
                return false;
            }
        }
    }
    return true;
}

//calculates a weighted score of tiles across board
int getScore(char board[][26], int n, char color, bool gameOver){
    int score = 0;
    int i,j;
    for( i = 0; i < n; ++i){
        for(j = 0; j < n; ++j){
            
            if(board[i][j] == color){
                if((i == 0 && j ==0) || (i ==0 && j == n-1)
                   ||(i == n-1 && j == 0)||(i ==n-1 && j == n-1) && ! gameOver){//corners
                    score += n*n*n*n;
                }
                else if((i == 1 && j ==1) || (i ==1 && j == n-2)
                        ||(i == n-2 && j == 1)||(i ==n-2 && j == n-2)){
                    score -= n*n*n;
                }
                
                else if((i == 0 || j == 0 || i == n-1 || j == n -1) && n>=16){
                    score += n;
                }
                else if((i == 0 || j == 0 || i == n-1 || j == n -1)){
                    score += n;
                }
                
                
                else if((i == 1 || j == 1 || i == n-2 || j == n -2) && n >= 8 ){//&& n > 14){ //&& n <= 8){
                    score += n/2;
                }
                
                else{//normal piece
                    score += 1;
                }
            }
        }
        
    }
    return score;
}

//counts real score value of tiles across a board
int getWinner(char board[][26], int n, char color){
    int i ,j;
    int score = 0;
    for( i = 0; i < n; ++i){
        for(j = 0; j < n; ++j){
            if(board[i][j] == color) score +=1;
        }
    }
    return score;
}

//returns max of two values
int max(int a, int b){
    if(a>b){
        return a;
    }
    else
        return b;
}

//returns min of two values
int min(int a, int b){
    if(a < b){
        return a;
    }
    return b;
}

//generates a new board state
void generateState(char board[][26], char newBoard[][26],int n , char row, char col,char color){
    
    
    char i,j;
    for(i = 'a'; i < n +'a'; ++i){
        for(j = 'a'; j < n +'a'; ++j){
            if(!(i == row && j == col)){
                newBoard[i-'a'][j-'a'] = board[i-'a'][j-'a'];
            }
        }
    }
    executeMove(newBoard, n, color,row,col);
    
    
}

//finds the optimal move at specified depth, is actually alphaBeta. Name is kinda a misnomer (sorry!)
int miniMax(char boardState[][26], int n, int depth, int startDepth, int alpha, int beta,char color,bool isMaxing
            , char* moveRow, char* moveCol){
    
    char row, col;
    int i ,numValidMoves;
    char** moveList = getValidMoves(boardState, n, color, &numValidMoves);
    //moveList = getValidMoves(boardState, n, color, &numValidMovesOpp);


printf("%d %d %d %d %d %c", n, depth, startDepth,alpha, beta,color);
if(isMaxing == true){
	printf(" true ");
}else{
	printf(" false ");
}

printf("%d %d %d\n", *moveRow, *moveCol, numValidMoves);


    if(numValidMoves > 27 && n > 16){//cut off to ensure time is okay
        numValidMoves = 27;
    }
    else if(numValidMoves > 24 && n > 20){
        numValidMoves = 24;
    }

    
    if(depth == 0 && startDepth % 2 != 0){
        
        return
        getScore(boardState,n,oppositecolor(color),false)
        -getScore(boardState,n,color, false);//fix gameIsOver
    }
    
    else if (isMaxing){
        //int numValidMoves, i;
        int maxScore = -10000000;
        int prevScore;
        
        printf("------start max Looping------\n");
        for(i = 0 ; i < numValidMoves; ++i){
            
            row = moveList[i][0];
            col = moveList[i][1];
            char newBoard[n][26];
            generateState(boardState, newBoard,n, row, col, color);//generate new board

printBoard(newBoard, n);


            prevScore = maxScore;
            maxScore = max(maxScore,
                           miniMax(newBoard, n, depth - 1, startDepth, alpha,beta,oppositecolor(color), false, moveRow, moveCol));
            
            alpha = max(alpha, maxScore);
            if(depth == startDepth && maxScore != prevScore){//possible highest move//makes next move maximal move
                *moveRow = row;
                *moveCol = col;
            }
            if(beta <= alpha){//checks to see if number is out of range
                return beta;
            }
        }

	printf("------end max Looping------\n");
        free(moveList);//free memory
        return maxScore;
    }
    
    
    else{
        int minScore = 10000000;
        
        
        printf("------start min Looping------\n");
        for(i = 0 ; i < numValidMoves; ++i){
            row = moveList[i][0];
            col = moveList[i][1];
            char newBoard[n][26];
            generateState(boardState, newBoard,n, row, col, color);
            minScore = min(minScore, miniMax(newBoard, n, depth - 1, startDepth,alpha, beta,oppositecolor(color), true, moveRow, moveCol));
            beta = min(beta , minScore);
            if(beta <= alpha){
                return alpha;
            }
        }
	printf("------end min Looping------\n");
        free(moveList);
        return minScore;
    }
    
    
    
}


//Process computer takes to make a move
void computerMove(char board[][26], int n, char color, bool* AIFinish){
    
    int numValidMovesComp;
    char compMoveRow, compMoveCol;
    int depth = 3;//depth changes for minimax based on amount of time I can afford
    if(n < 8 ){
        depth = 3;
    }
    else if(n < 10){
        depth = 5;
    }
    

    getValidMoves(board,n,color,&numValidMovesComp);//get number of valid moves
	printf("numValidMovesComp:%d\n", numValidMovesComp);
    if(numValidMovesComp > 0){
        
        
        //if you have a move exceute minimax and make a move
        miniMax(board,n,depth,depth,-100000,100000,color,true, &compMoveRow, &compMoveCol);
        printf("Computer places %c at %c%c.\n", color, compMoveRow,compMoveCol);
        executeMove(board, n, color, compMoveRow, compMoveCol);
        printBoard(board,n);
    }
    else{
        printf("%c player has no valid move.\n", color);
        *AIFinish = true;
        
    }
}


//allows the player to make a move
void playerMove(char board[][26], int n, char color, bool* playerFinish, bool* stop){
    int numValidMovesPlayer = 0;
    int playerMoveRow, playerMoveCol;
    
    getValidMoves(board,n,color,&numValidMovesPlayer);
    if(numValidMovesPlayer > 0){
        
        bool invalidMove = true;
        while (invalidMove) {
            printf("Enter move for color %c (RowCol): ", color);
            char inputStr[2];
            scanf("%s", inputStr);
            playerMoveRow = inputStr[0];
            playerMoveCol = inputStr[1];
            
            
            if(!(isValidMove(board,n,color,playerMoveRow,playerMoveCol))){
                printf("Invalid move. Plz reinput a valid move\n");
                //*stop = true;
                invalidMove = true;
            }
            else{
                executeMove(board,n,color, playerMoveRow,playerMoveCol);
                printBoard(board,n);
                invalidMove = false;
            }
        }
        

    }
    else{
        printf("%c player has no valid move.\n", color);
        *playerFinish = true;
    }
    
}

int main(void){
    int n;
    char color;
    bool AIFirst = false;
    bool AIFinish = false;
    bool playerFinish = false;
    bool stop = false;
    
    
    
    printf("Enter the board dimension: ");
    scanf("%d", &n);
    printf("Computer plays (B/W) (B is the first to play) : ");
    scanf(" %c", &color);
    if(color == 'B'){
        AIFirst = true;
    }
    
    char board[n][26];
    
    generateStartBoard(n,board);
    printBoard(board,n);
    
    //the game
    
    while(!(boardIsFull(board,n)) && !(AIFinish && playerFinish) &&!stop){
        
        
        
        if(AIFirst){
            if(!stop){
                computerMove(board,n,color, &AIFinish);
            }
            playerMove(board,n,oppositecolor(color), &playerFinish,&stop);
        }
        else{
            playerMove(board,n,oppositecolor(color), &playerFinish,&stop);
            if(!stop && !(boardIsFull(board,n))){
                computerMove(board,n,color, &AIFinish);
            }
        }
        if(boardIsFull(board,n)){
            AIFinish = true;
            playerFinish = true;
        }
        
    }
    
    
    
    //win conditions
    if(stop){
        printf("%c player wins.\n", color);
    }
    else if (getWinner(board, n, color) > getWinner(board, n, oppositecolor(color))){
        printf("%c player wins.\n", color);
    }
    else if(getWinner(board, n, color) < getWinner(board, n, oppositecolor(color))){
        printf("%c player wins.\n", oppositecolor(color));
    }
    else{
        printf("Draw!\n");
    }
}

