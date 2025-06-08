package main

import "fmt"

// 0 - nothing
// 9 - bomb
// 10 - flag
// 11 - clicked
func main() {
	fmt.Printf("Test")
	size := 9
	grid := get_grid(size) // create 9x9 array
	enter_grid(grid, size)
	checks:=rule1(grid)
	fmt.Println(checks)
}
func enter_grid(grid [][]int, size int) {
	for i:=0;i<size;i++ {
		for j:=0;j<size;j++ {
			var v int
			fmt.Scan(&v)
			grid[i][j] = v
		}
		fmt.Println("row ",i)
	}
}
// returns 0 if all squares are safe
// returns 1 if all empty squares are bombs
// returns 3 if something else happens
type Triple struct {
	A int
	B int
	C int
}
func rule1(grid [][]int) []Triple {
	var checks []Triple
	for i:=0;i<len(grid);i++ {
		for j:=0;j<len(grid);j++ {
			element := grid[i][j];
			if element != 0 && element != 11{
				_, empty, bombs := check_around(grid, i,j)
				if bombs ==  element{
					fmt.Print("Squares around are safe", i,j)
					checks = append(checks, Triple{i,j,0})
				} else if element - bombs == empty {
					fmt.Print("All empty squares around are bombs", i,j)
					checks = append(checks, Triple{i,j,1})
				}

			}
		}
	}
	return checks
}
func check_around(grid [][]int,i int, j int) ([]int, int, int) {
	around := []int {}
	touching := [][]int{{-1,-1}, {-1,0}, {-1, 1}, {0, 1}, {0,0},{0,-1},{1,1},{1,0},{1,-1}}
	bombs:= 0
	empty := 0
	rows := len(grid)
	cols := len(grid[0])
	for _, offset := range touching {
		ni := i + offset[0]
		nj := j + offset[1]

		if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
			val := grid[ni][nj]
			around = append(around, val)
			if val == 9 || val == 10 {
				bombs++
			} else if val == 0 {
				empty++
			}
		}
	}
	return around, empty, bombs
}
func get_grid(size int) [][]int {
	grid := make([][]int, size)
	for i:=0;i<size;i++ {
			grid[i]= make([]int,size)
	}
	return grid
}
