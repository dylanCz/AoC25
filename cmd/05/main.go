package main

import (
	helper "aoc25/internal"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func p1(input string) int {
	freshIngredients, ingredients := helper.P5Parse(input)
	ingredientList := newIngredients(freshIngredients)
	freshIngredientCount := 0
	for _, ingredient := range ingredients {
		ingredient, _ := strconv.Atoi(ingredient)
		if ingredientList.isFresh(ingredient) {
			freshIngredientCount += 1
		}
	}
	return freshIngredientCount
}

func p2(input string) int {
	freshIngredients, _ := helper.P5Parse(input)
	ingredientList := newIngredients(freshIngredients)
	return ingredientList.total()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	inputFile := os.Getenv("input_file")
	slog.Info("AoC Day 5", "Ingredients P1", p1(helper.LoadInput(inputFile)))
	slog.Info("AoC Day 5", "Ingredients P2", p2(helper.LoadInput(inputFile)))
}
