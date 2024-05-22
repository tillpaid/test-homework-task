package models

import (
	"app/pkg/checkers"
	"strconv"
)

type Rule struct {
	checkers []checkers.Checker
}

func NewRule(checkers []checkers.Checker) *Rule {
	return &Rule{checkers: checkers}
}

func (r *Rule) GetCheckedValue(number int) string {
	var result string

	for _, checker := range r.checkers {
		result += checker.Check(number)
	}

	if result == "" {
		result = strconv.Itoa(number)
	}

	return result
}
