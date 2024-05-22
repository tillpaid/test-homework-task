package main

import (
	checkers2 "app/pkg/checkers"
	"app/pkg/models"
	"fmt"
)

func main() {
	tasks := []*models.Task{
		getFirstTask(),
		getSecondTask(),
		getThirdTask(),
	}

	for i, task := range tasks {
		fmt.Printf("Task v%d:\n", i+1)
		fmt.Println(task.Process())
	}
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
