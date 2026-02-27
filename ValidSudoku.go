func isValidSudoku(board [][]byte) bool {
    //for row check
    for i:=0; i<9;i++ {
        arr := make([]bool, 9)
        for j:=0; j<9; j++ {
            if board[i][j] != '.'{
                num := board[i][j] - '1'
                if arr[num] {
                    return false
                }
                arr[num] = true
            }
            
            
        }
    }
    //column check
    for i:=0; i<9;i++ {
        arr := make([]bool, 9)
        for j:=0; j<9; j++ {
            if board[j][i] != '.'{
                num := board[j][i] - '1'
                if arr[num] {
                    return false
                }
                arr[num] = true
            }
        }
    }

    for k:=0; k< 9;k++{
        arr := make([]bool, 9)
        rowStart := (k /3) * 3
        colStart := (k %3) * 3
        for i:=0; i<3;i++ {
            for j:=0; j<3; j++ {
                if board[i + rowStart][j + colStart] != '.'{
                    num := board[i + rowStart][j + colStart] - '1'
                    if arr[num] {
                        return false
                    }
                    arr[num] = true
                }
            }
        }
    }
    return true
}
