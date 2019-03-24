package main

import (
	"testing"
)

func TestSquareRoot(t *testing.T) {
	squared := 25
	subject := SquareRoot(subject)
	expected := 5
	if subject != expected {
		t.Errorf("Square root was incorrect, got %d, expected %d.", subject, expected)
	}
}

func TestSalesByMonth(t *testing.T) {
	salesByMonth := SalesByMonth()
	subject1 := salesByMonth["Jan"]
	expected1 := 1284.20
	subject2 := salesByMonth["Aug"]
	expected2 := 3420.10
	if subject1 != expected1 {
		t.Errorf("Sales by month was incorrect, got %f, expected %f.", subject1, expected1)
	}

	if subject2 != expected2 {
		t.Errorf("Sales by month was incorrect, got %f, expected %f.", subject2, expected2)
	}
}

func TestHelloWorld(t *testing.T) {
	subject := 123
	expected := HelloWorld()
	if subject != expected {
		t.Errorf("Hello world was incorrect, got %d, expected %d.", subject, expected)
	}
}
