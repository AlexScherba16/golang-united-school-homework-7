package coverage

import (
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
