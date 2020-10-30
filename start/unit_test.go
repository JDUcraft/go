package main

import (
	"testing"
)

func TestComputeCircle(t *testing.T) {
	circle := Circle{
		ShapeType: "Circle",
		Radius:    3,
	}

	result := compute(circle)

	expectedArea := 88.83
	expectedShapeType := "Circle"
	if result.Area != expectedArea {
		t.Errorf("expected: %f, received: %f", expectedArea, result.Area)
	}
	if result.ShapeType != expectedShapeType {
		t.Errorf("expected: %s, receved: %s", expectedShapeType, result.ShapeType)
	}
}

func TestComputeSquare(t *testing.T) {
	square := Square{
		ShapeType: "Circle",
		Side:    3,
	}

	result := compute(square)

	expectedArea := 9.00
	expectedShapeType := "Square"
	if result.Area != expectedArea {
		t.Errorf("expected: %f, received: %f", expectedArea, result.Area)
	}
	if result.ShapeType != expectedShapeType {
		t.Errorf("expected: %s, receved: %s", expectedShapeType, result.ShapeType)
	}
}
