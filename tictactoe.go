package main

import (
    "bufio"
    "fmt"
    "os"
)

func initTicTacToe(m map[string]map[string]string) int{
    k_arr := [3]string{"A", "B", "C"}
    for i := range(k_arr) {
        row := k_arr[i]
        m[row] = map[string]string{}
        for j := range(k_arr) {
            col := k_arr[j]
            m[row][col] = " "
        }
    }
    //m["A"]["B"] = "C"
    //m["A"]["C"] = "C"
    return 1
}

func displayTicTacToe(m map[string]map[string]string, comp string, player string)int {
    fmt.Printf("  A B C \n")
    k_arr := [3]string{"A", "B", "C"}
    for i := range(k_arr) {
        row := k_arr[i]
        fmt.Printf("%s", row)
        //m[row] = map[string]string{}
        for j := range(k_arr) {
            col := k_arr[j]
            //m[row][col] = "K"
            fmt.Printf("|%s",m[row][col])
        }
        fmt.Printf("|\n")
    }
    fmt.Printf("Computer: %v\n", comp)
    fmt.Printf("Player: %v\n", player)
    //fmt.Printf("Entire DSC: %v\n", m)
    return 1
}

func readInputChar1() string {
    reader := bufio.NewReader(os.Stdin)
    char, _, _ := reader.ReadRune()
    input := string(char)
    return input
}

func getMoves()[]string{
    pos_arr := make([]string, 0)
    fmt.Println("Enter Row:")
    row_reader := bufio.NewReader(os.Stdin)
    row_char, _, _ := row_reader.ReadRune()
    row := string(row_char)
    //fmt.Println(row)
    
    fmt.Println("Enter Col:")
    col_reader := bufio.NewReader(os.Stdin)
    col_char, _, _ := col_reader.ReadRune()
    col := string(col_char)
    //fmt.Println(col)
    pos_arr = append(pos_arr, row)
    pos_arr = append(pos_arr, col)
    return pos_arr
}

func checkValidMove(mom_ttt map[string]map[string]string, row string, col string)int{
    val := mom_ttt[row][col]
    if val != " " {
        return 0
    }
    return 1
}

func applyMove(mom_ttt map[string]map[string]string, player string, row string, col string)int{
    mom_ttt[row][col] = player
    return 1
}

func getRowWiseCounts(m map[string]map[string]string) string{
    cntr := make(map[string]int)
    var v_arr [] string
    k_arr := [3]string{"A", "B", "C"}
    for i := range(k_arr) {
        row := k_arr[i]
        for j := range(k_arr) {
            col := k_arr[j]
            v_arr = append(v_arr, m[row][col])
        }
            getCountOfMoves(cntr, v_arr)
            if cntr["C"] == 3 {
                return "Computer"
            } else if cntr["P"] == 3 {
                return "Player"
            }
        //fmt.Printf("Entire DSC: %v\n", cntr)
            cntr = make(map[string]int)
        //fmt.Printf("Entire DSC: %v\n", v_arr)
        v_arr = nil
    }
    return "N"
}

func getColWiseCounts(m map[string]map[string]string) string{
    cntr := make(map[string]int)
    var v_arr [] string
    k_arr := [3]string{"A", "B", "C"}
    for i := range(k_arr) {
        col := k_arr[i]
        for j := range(k_arr) {
            row := k_arr[j]
            v_arr = append(v_arr, m[row][col])
        }
            getCountOfMoves(cntr, v_arr)
            if cntr["C"] == 3 {
                return "Computer"
            } else if cntr["P"] == 3 {
                return "Player"
            }
        //fmt.Printf("Entire DSC: %v\n", cntr)
            cntr = make(map[string]int)
        //fmt.Printf("Entire DSC: %v\n", v_arr)
        v_arr = nil
    }
    return "N"
}

func getdiagWiseCounts(m map[string]map[string]string) string{
    cntr := make(map[string]int)
    var v_arr1 [] string
    v_arr1 = append(v_arr1, m["A"]["A"], m["B"]["B"], m["C"]["C"])
    getCountOfMoves(cntr, v_arr1)
    if cntr["C"] == 3 {
        return "Computer"
    } else if cntr["P"] == 3 {
        return "Player"
    }
            cntr = make(map[string]int)

    var v_arr2 [] string
    v_arr2 = append(v_arr2, m["A"]["C"], m["B"]["B"], m["C"]["A"])
    getCountOfMoves(cntr, v_arr2)
    if cntr["C"] == 3 {
        return "Computer"
    } else if cntr["P"] == 3 {
        return "Player"
    }
    return "N"
}

func getCountOfMoves(cntr map[string]int, v_arr []string) int{
    for i := range(v_arr) {
        val := v_arr[i]
        cntr[val]++
    }
    return 1
}

func checkWinStatus(m map[string]map[string]string) string {
    win_status_col := getColWiseCounts(m)
    win_status_row := getRowWiseCounts(m)
    win_status_diag := getdiagWiseCounts(m)
    if win_status_row != "N" {
        return win_status_row
    } else if win_status_col != "N" {
        return win_status_col
    } else if win_status_diag != "N" {
        return win_status_diag
    } else {
        return "N"
    }
    
}

func checkDrawStatus(m map[string]map[string]string) int {
    var empty int
    k_arr := [3]string{"A", "B", "C"}
    for i := range(k_arr) {
        row := k_arr[i]
        for j := range(k_arr) {
            col := k_arr[j]
            if m[row][col] == " " {
                empty++
            }
        }
    }
    return empty
}

func generateMove(m map[string]map[string]string){
    k_arr := [3]string{"A", "B", "C"}
    for i := range(k_arr) {
        row := k_arr[i]
        for j := range(k_arr) {
            col := k_arr[j]
            if m[row][col] == " " {
                m[row][col] = "C"
                return
            }
        }
    }
    return
}

func main() {
    mom_ttt := make(map[string]map[string]string)
    var pos_arr [] string
    comp := "C"
    player := "P"
    initTicTacToe(mom_ttt)
    displayTicTacToe(mom_ttt, comp, player)
    for {
        pos_arr = getMoves()
        valid_move := checkValidMove(mom_ttt, pos_arr[0], pos_arr[1])
        if valid_move == 1 {
            fmt.Println("Valid Move")
            applyMove(mom_ttt, player, pos_arr[0], pos_arr[1])
        } else {
            fmt.Println("Invalid Move, Please Re-Enter")
        }
        win_status := checkWinStatus(mom_ttt)
        if win_status != "N" {
            fmt.Printf("Won By: %v\n", win_status)
            return
        }
        draw_status := checkDrawStatus(mom_ttt)
        if draw_status == 0 {
            fmt.Printf("Match is Draw\n")
            return
        }
        generateMove(mom_ttt)
        win := checkWinStatus(mom_ttt)
        if win != "N" {
            fmt.Printf("Won By: %v\n", win)
            return
        } else {
            fmt.Printf("Match is Draw\n")
        }
        displayTicTacToe(mom_ttt, comp, player)
    }
}
