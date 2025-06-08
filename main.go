package main

import "fmt"

// 0 - nothing
// 9 - bomb
// 10 - flag
func main() {
	fmt.Printf("Test")
	grid := get_grid(9) // create 9x9 array
	rule1(grid)
}
func rule1(grid [][]int) {
	for i:=0;i<len(grid);i++ {
		for j:=0;j<len(grid);j++ {
			if grid[i][j] != 0 {
				around, check := check_around(grid, i,j)
				if check == grid[i][j] {
					fmt.Print("Squares around are safe", i,j)
				}
			}
		}
	}
}
func check_around(grid [][]int,i int, j int) ([]int, int) {
	around := make([]int,8)
	touching := [][]int{{-1,-1}, {-1,0}, {-1, 1}, {0, 1}, {0,0},{0,-1},{1,1},{1,0},{1,-1}}
	empty_squares_around:= 0
	for v:=0;v<len(touching);v++ {
		element := grid[touching[v][0]][touching[v][1]];
		if element == 9 || element == 10 {
			empty_squares_around += 1;
		}
		around = append(around, element)
	}
	return around, empty_squares_around
}
func get_grid(size int) [][]int {
	grid := make([][]int, size)
	for i:=0;i<size-1;i++ {
		for j:=0;j<size-1;i++ {
			grid[i][j] = 0
		}
	}
	return grid
}
