package main

import (
	"fmt"
	"strings"
)

const (
	rows    = 9
	columns = 9
)

type SudokuErrors []error

func (s SudokuErrors) Error() string {
	var errors []string
	for _, err := range s {
		errors = append(errors, err.Error())
	}
	return strings.Join(errors, ", ")
}

type sudoku struct {
	initialGrid [9][9]int
	grid        [9][9]int
}

func isValidDigit(d int) bool {
	if d < 1 || d > 9 {
		return false
	}
	return true
}

func isInBound(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}

	if column < 0 || column >= columns {
		return false
	}
	return true
}

func (s sudoku) validInRow(row, digit int) bool {
	for i := 0; i < columns; i++ {
		if digit == s.grid[row][i] {
			return false
		}
	}
	return true
}

func (s sudoku) validInColumn(column, digit int) bool {
	for i := 0; i < rows; i++ {
		if digit == s.grid[i][column] {
			return false
		}
	}
	return true
}

func (s sudoku) validInSubSquare(x, y, digit int) bool {
	xSquareIndex := (x / 3) * 3
	ySquareIndex := (y / 3) * 3

	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if s.grid[xSquareIndex+row][ySquareIndex+column] == digit {
				return false
			}
		}
	}
	return true
}

func (s *sudoku) set(x, y, digit int) error {
	var errs SudokuErrors
	if !isInBound(x, y) {
		return fmt.Errorf("The values for x: %v or y %v are outside of the bounds", x, y)
	}
	if s.grid[x][y] == digit {
		errs = append(errs, fmt.Errorf("The digit %v is already set in the x: %v, y: %v position", digit, x, y))
	}
	if !isValidDigit(digit) {
		errs = append(errs, fmt.Errorf("The digit %v isn't a allowed value. It should be between 1 and 9", digit))
	}
	if !s.validInColumn(y, digit) || !s.validInRow(x, digit) || !s.validInSubSquare(x, y, digit) {
		errs = append(errs, fmt.Errorf("The digit %v isn't valid because it's already present in the row, column or sub square", digit))
	}
	if s.initialGrid[x][y] != 0 {
		errs = append(errs, fmt.Errorf("The position x: %v, y: %v is a position with a default value and cannot be overwritten", x, y))
	}
	if len(errs) > 0 {
		return errs
	}

	s.grid[x][y] = digit
	return nil
}

func (s sudoku) show() {
	line := strings.Repeat("-", 31)

	fmt.Println(line)

	for row := 0; row < rows; row++ {
		fmt.Print("|")
		for column := 0; column < columns; column++ {
			val := s.grid[row][column]
			if val == 0 {
				fmt.Print("   ")
			} else {
				fmt.Printf(" %v ", val)
			}
			if (column+1)%3 == 0 {
				fmt.Print("|")
			}
		}
		fmt.Print("\n")
		if (row+1)%3 == 0 {
			fmt.Println(line)
		}
	}
}

func newSudoku(board [9][9]int) sudoku {
	return sudoku{
		grid:        board,
		initialGrid: board,
	}
}

type sudokuPlay struct {
	x, y, digit int
}

func main() {
	s := newSudoku([rows][columns]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	plays := []sudokuPlay{
		{1, 1, 8},  // error, the value already exists in the sub square
		{3, 9, 2},  // error, column out of the bounds
		{6, 8, 4},  // should be success
		{4, 4, 6},  // error, he value already exists in the column
		{8, 3, 9},  // error, he value already exists in the row
		{8, 4, 3},  // error, the cell has a default value and cannot be overwritten
		{0, 8, 15}, // error, the value isn't valid
	}

	s.show()

	for _, play := range plays {
		err := s.set(play.x, play.y, play.digit)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Awesome! The value %v in the posistion x: %v, y: %v is a valid value \n", play.digit, play.x, play.y)
		}
	}

	s.show()
}
