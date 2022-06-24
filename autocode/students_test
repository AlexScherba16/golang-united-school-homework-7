package coverage

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

const (
	Matrix1x3 = "1 2 3"
	Matrix2x3 = "1 2 3 \n 4 5 6 "
	Matrix3x3 = "1 2 3 \n 4 5 6 \n 7 8 9"
	SetValue  = 100500
)

// Len Unit Tests
func TestLenEmptyList(t *testing.T) {
	var p People
	if p.Len() != 0 {
		t.Error("the length of the empty list must be zero, not ", p.Len())
	}
}

func TestLen(t *testing.T) {
	var p People
	p = append(p, Person{firstName: "firstName", lastName: "lastName", birthDay: time.Now()})
	p = append(p, Person{firstName: "firstName_1", lastName: "lastName_1", birthDay: time.Now()})

	err := p.Len()
	if err != 2 {
		t.Error("the length of list must be 2, not", p.Len())
	}
}

// Less Unit Tests
func TestLess_Birthday(t *testing.T) {
	lessTestData := map[string]struct {
		first  Person
		second Person
		result bool
	}{
		"birthday ok": {
			first:  Person{firstName: "", lastName: "", birthDay: time.Now().Add(time.Second * 10)},
			second: Person{firstName: "", lastName: "", birthDay: time.Now()},
			result: true,
		},
		"birthday error": {
			first:  Person{firstName: "", lastName: "", birthDay: time.Now()},
			second: Person{firstName: "", lastName: "", birthDay: time.Now().Add(time.Second * 10)},
			result: false,
		},
	}

	var p People
	for name, data := range lessTestData {
		p = append(p, data.first, data.second)
		result := p.Less(0, 1)

		if result != data.result {
			t.Fatalf("[%s] expected %t, Less function returns %t", name, data.result, result)
		}
		p = p[:0]
	}
}

func TestLess_FirstName(t *testing.T) {
	lessTestData := map[string]struct {
		first  Person
		second Person
		result bool
	}{
		"first name ok": {
			first:  Person{firstName: "GFG", lastName: "", birthDay: time.Now()},
			second: Person{firstName: "Geeks", lastName: "", birthDay: time.Now()},
			result: true,
		},
		"birthday error": {
			first:  Person{firstName: "Geeks", lastName: "", birthDay: time.Now()},
			second: Person{firstName: "GFG", lastName: "", birthDay: time.Now().Add(time.Second * 10)},
			result: false,
		},
	}

	var p People
	for name, data := range lessTestData {
		p = append(p, data.first, data.second)
		result := p.Less(0, 1)

		if result != data.result {
			t.Fatalf("[%s] expected %t, Less function returns %t", name, data.result, result)
		}
		p = p[:0]
	}
}

func TestLess_LastName(t *testing.T) {
	lessTestData := map[string]struct {
		first  Person
		second Person
		result bool
	}{
		"last name ok": {
			first:  Person{firstName: "", lastName: "GFG", birthDay: time.Now()},
			second: Person{firstName: "", lastName: "Geeks", birthDay: time.Now()},
			result: true,
		},
		"last name error": {
			first:  Person{firstName: "", lastName: "Geeks", birthDay: time.Now()},
			second: Person{firstName: "", lastName: "GFG", birthDay: time.Now().Add(time.Second * 10)},
			result: false,
		},
	}

	var p People
	for name, data := range lessTestData {
		p = append(p, data.first, data.second)
		result := p.Less(0, 1)

		if result != data.result {
			t.Fatalf("[%s] expected %t, Less function returns %t", name, data.result, result)
		}
		p = p[:0]
	}
}

// Swap Unit Tests
func TestSwap(t *testing.T) {
	var p People
	first := Person{firstName: "firstName", lastName: "lastName", birthDay: time.Now().Add(time.Second * 10)}
	second := Person{firstName: "firstName_1", lastName: "lastName_1", birthDay: time.Now()}

	p = append(p, first, second)
	p.Swap(0, 1)

	if p[0] == first {
		t.Fatal("same elements after Swap")
	}
}

// New Unit Tests
func TestNew_MatrixValidStrings(t *testing.T) {
	matrixValidData := map[string]struct {
		matrixString string
	}{
		"1 line": {matrixString: Matrix1x3},
		"2 line": {matrixString: Matrix2x3},
		"3 line": {matrixString: Matrix3x3},
	}

	for name, data := range matrixValidData {
		matrix, err := New(data.matrixString)

		if matrix == nil || err != nil {
			t.Fatalf("[%s] should be matrix and empty error", name)
		}
	}
}

func TestNew_MatrixInvalidStrings(t *testing.T) {
	matrixErrorData := map[string]struct {
		matrixString string
	}{
		"different length":  {matrixString: "1 2 3 \n 4 5 6 7 \n"},
		"empty line":        {matrixString: ""},
		"space line":        {matrixString: " "},
		"atoi error":        {matrixString: "1 2 3 \n a 4 3, \n"},
		"additional spaces": {matrixString: "1 2    3"},
	}

	for name, data := range matrixErrorData {
		matrix, err := New(data.matrixString)

		if matrix != nil || err == nil {
			t.Fatalf("[%s] should be error and empty matrix", name)
		}
	}
}

// Rows Unit Tests
func TestRows(t *testing.T) {
	matrixRowsValidData := map[string]struct {
		matrixString string
		expectedRows [][]int
	}{
		"1 line": {matrixString: Matrix1x3, expectedRows: [][]int{{1, 2, 3}}},
		"2 line": {matrixString: Matrix2x3, expectedRows: [][]int{{1, 2, 3}, {4, 5, 6}}},
		"3 line": {matrixString: Matrix3x3, expectedRows: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
	}

	for name, data := range matrixRowsValidData {
		matrix, err := New(data.matrixString)
		if matrix == nil || err != nil {
			t.Fatalf("[%s] should be matrix and empty error, not : %s", name, err)
		}

		err = compareMatrices(matrix.Rows(), data.expectedRows)
		if err != nil {
			t.Fatalf("[%s] %s", name, err)
		}
	}
}

// Columns Unit Tests
func TestColumns(t *testing.T) {
	matrixColumnsValidData := map[string]struct {
		matrixString    string
		expectedColumns [][]int
	}{
		"1 column":  {matrixString: Matrix1x3, expectedColumns: [][]int{{1}, {2}, {3}}},
		"2 columns": {matrixString: Matrix2x3, expectedColumns: [][]int{{1, 4}, {2, 5}, {3, 6}}},
		"3 columns": {matrixString: Matrix3x3, expectedColumns: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
	}

	for name, data := range matrixColumnsValidData {
		matrix, err := New(data.matrixString)
		if matrix == nil || err != nil {
			t.Fatalf("[%s] should be matrix and empty error, not : %s", name, err)
		}

		err = compareMatrices(matrix.Cols(), data.expectedColumns)
		if err != nil {
			t.Fatalf("[%s] %s", name, err)
		}
	}
}

// Set Unit Tests
func TestSet(t *testing.T) {
	matrixSetData := map[string]struct {
		matrixString string
		row          int
		column       int
		value        int
		result       bool
	}{
		"negative row position":    {matrixString: Matrix1x3, row: -1, column: 0, value: SetValue, result: false},
		"overflow row position":    {matrixString: Matrix1x3, row: 100, column: 0, value: SetValue, result: false},
		"negative column position": {matrixString: Matrix1x3, row: 0, column: -1, value: SetValue, result: false},
		"overflow column position": {matrixString: Matrix1x3, row: 0, column: 100, value: SetValue, result: false},

		"set value at [0,0] in 1 line matrix": {matrixString: Matrix1x3, row: 0, column: 0, value: SetValue, result: true},
		"set value at [0,2] in 1 line matrix": {matrixString: Matrix1x3, row: 0, column: 2, value: SetValue, result: true},

		"set value at [0,0] in 2 line matrix": {matrixString: Matrix2x3, row: 0, column: 0, value: SetValue, result: true},
		"set value at [1,2] in 2 line matrix": {matrixString: Matrix2x3, row: 1, column: 2, value: SetValue, result: true},

		"set value at [0,0] in 3 line matrix": {matrixString: Matrix3x3, row: 0, column: 0, value: SetValue, result: true},
		"set value at [1,1] in 3 line matrix": {matrixString: Matrix3x3, row: 1, column: 1, value: SetValue, result: true},
		"set value at [2,2] in 3 line matrix": {matrixString: Matrix3x3, row: 2, column: 2, value: SetValue, result: true},
	}

	for name, data := range matrixSetData {
		matrix, err := New(data.matrixString)
		if matrix == nil || err != nil {
			t.Fatalf("[%s] should be matrix and empty error, not : %s", name, err)
		}

		result := matrix.Set(data.row, data.column, data.value)
		if result != data.result {
			t.Fatalf("[%s] incorrect set result %t", name, result)
		}
	}
}

// Util functions
func compareMatrices(a, b [][]int) error {
	if len(a) != len(b) {
		return fmt.Errorf("ivalid length; len(a) %d, len(b) %d", len(a), len(b))
	}

	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return fmt.Errorf("ivalid length; len(a[i]) %d, len(b[i]) %d", len(a[i]), len(b[i]))
		}

		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return fmt.Errorf("ivalid element %d, expected %d", a[i][j], b[i][j])
			}
		}
	}
	return nil
}
