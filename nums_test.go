package main

import "testing"

func TestToWord_NotAWord(t *testing.T) {
	// Arrange
	n := Num{"asdf"}

	// Act
	s, err := n.ToWord()

	// Assert
	if s != "" {
		t.Errorf(s + "was expected to be empty")
	}
	if err == nil {
		t.Errorf("err was expected")
	}
}

func TestToWord_PositiveNumber(t *testing.T) {
	// Arrange
	n := Num{"1234"}

	// Act
	s, err := n.ToWord()

	// Assert
	if s != "OneTwoThreeFour" {
		t.Errorf(s + "was expected to be OneTwoThreeFour")
	}
	if err != nil {
		t.Errorf("no err was expected")
	}
}

func TestToWord_NegativeNumber(t *testing.T) {
	// Arrange
	n := Num{"-1234"}

	// Act
	s, err := n.ToWord()

	// Assert
	if s != "NegativeOneTwoThreeFour" {
		t.Errorf(s + "was expected to be NegativeOneTwoThreeFour")
	}
	if err != nil {
		t.Errorf("no err was expected")
	}
}

func TestToWord_EmptyString(t *testing.T) {
	// Arrange
	n := Num{""}

	// Act
	s, err := n.ToWord()

	// Assert
	if s != "" {
		t.Errorf(s + "was expected to be empty")
	}
	if err == nil {
		t.Errorf("err was expected")
	}
}
