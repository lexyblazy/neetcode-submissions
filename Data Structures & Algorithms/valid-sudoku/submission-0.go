func isValidSudoku(board [][]byte) bool {
   seen := make(map[string]bool)

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            value := board[r][c]

            if value == '.' {
                continue
            }

            rowKey := string(value) + " in row " + string(rune(r))
            colKey := string(value) + " in col " + string(rune(c))
            boxKey := string(value) + " in box " + string(rune((r/3)*3+c/3))

            if seen[rowKey] || seen[colKey] || seen[boxKey] {
                return false
            }

            seen[rowKey] = true
            seen[colKey] = true
            seen[boxKey] = true
        }
    }

    fmt.Println(seen)

    return true
}
