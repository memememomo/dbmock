package dbmock

import (
	"testing"
	"fmt"
)

var PseudoDB []*User

type User struct {
	ID uint64
	Name string
}

func (u *User) ToDB() error {
	PseudoDB = append(PseudoDB, u)
	return nil
}

func user(i uint64) DBMapper {
	return Mock(&User{
		ID: i + 1,
		Name: fmt.Sprintf("Name_%d", i + 1),
	})
}

func UserMock() *Generator {
	return NewGenerator(user)
}



func TestGenerator(t *testing.T) {
	generator := UserMock()

	// SingleM
	mock1 := generator.SingleM(0, nil).(*User)
	if mock1.ID != 1 || mock1.Name != "Name_1" {
		t.Fatal("Failed to mock by SingleM")
	}
	if len(PseudoDB) != 0 {
		t.Fatal("Invalid save by SingleM")
	}

	// Single
	mock2 := generator.Single(1, nil).(*User)
	if mock2.ID != 2 || mock2.Name != "Name_2" {
		t.Fatal("Failed to mock by Single")
	}
	if len(PseudoDB) != 1 {
		t.Fatal("Failed to save by Single")
	}
	PseudoDB = []*User{}

	// MultiM
	mocks1 := generator.MultiM(5, nil)
	if len(mocks1) != 5 {
		t.Fatal("Failed to mock by MultiM")
	}
	for i, m := range mocks1 {
		u := m.(*User)
		if u.ID != (uint64)(i+1) || u.Name != fmt.Sprintf("Name_%d", i+1) {
			t.Fatalf("Failed to mock by MultiM: %d", i + 1)
		}
	}
	if len(PseudoDB) != 0 {
		t.Fatal("Invalid save by MultiM")
	}

	// Multi
	mocks2 := generator.Multi(5, nil)
	for i, m := range mocks2 {
		u := m.(*User)
		if u.ID != (uint64)(i+1) || u.Name != fmt.Sprintf("Name_%d", i+1) {
			t.Fatalf("Failed to mock by Multi: %d", i + 1)
		}
	}
	if len(PseudoDB) != 5 {
		t.Fatalf("Failed to save by Multi")
	}
}