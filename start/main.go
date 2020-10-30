package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
)

type Geometry interface {
	compute() ComputedArea
}

type ComputedArea struct {
	ShapeType string
	Area      float64
}

type Circle struct {
	ShapeType string
	Radius    float64
}

type Square struct {
	ShapeType string
	Side      float64
}

func (circle Circle) compute() ComputedArea {
	return ComputedArea{ShapeType: "Circle", Area: math.Round(math.Pow(math.Pi*circle.Radius, 2)*100) / 100}
}

func (square Square) compute() ComputedArea {
	return ComputedArea{ShapeType: "Square", Area: math.Round(math.Pow(square.Side, 2)*100) / 100}
}

func compute(shape Geometry) ComputedArea {
	return shape.compute()
}

func homeController(w http.ResponseWriter, _ *http.Request) {
	circle := Circle{Radius: 3.425}
	square := Square{Side: 4}

	var outputJson []ComputedArea
	outputJson = append(outputJson, compute(circle))
	outputJson = append(outputJson, compute(square))

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(outputJson)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", homeController)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
