package main

import (
	checkers2 "app/pkg/checkers"
	"app/pkg/models"
	"fmt"
)

func main() {
	fmt.Println("Task v1:")
	fmt.Println(getFirstTask().Process())

	fmt.Println("Task v2:")
	fmt.Println(getSecondTask().Process())

	fmt.Println("Task v3:")
	fmt.Println(getThirdTask().Process())
}

func getFirstTask() *models.Task {
	checkers := []checkers2.Checker{
		checkers2.NewDivisionChecker(3, "pa"),
		checkers2.NewDivisionChecker(5, "pow"),
	}
	rule := models.NewRule(checkers)

	return models.NewTask(1, 20, " ", rule)
}

func getSecondTask() *models.Task {
	checkers := []checkers2.Checker{
		checkers2.NewDivisionChecker(2, "hatee"),
		checkers2.NewDivisionChecker(7, "ho"),
	}
	rule := models.NewRule(checkers)

	return models.NewTask(1, 15, "-", rule)
}

func getThirdTask() *models.Task {
	checkers := []checkers2.Checker{
		checkers2.NewOneOfChecker([]int{1, 4, 9}, "joff"),
		checkers2.NewLargerChecker(5, "tchoff"),
	}
	rule := models.NewRule(checkers)

	return models.NewTask(1, 10, "-", rule)
}
