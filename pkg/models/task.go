package models

import "strings"

type Task struct {
	start     int
	limit     int
	separator string
	rule      *Rule
}

func NewTask(start int, limit int, separator string, rule *Rule) *Task {
	return &Task{start: start, limit: limit, separator: separator, rule: rule}
}

func (t *Task) Process() string {
	var results []string

	for i := t.start; i <= t.limit; i++ {
		results = append(results, t.rule.GetCheckedValue(i))
	}

	return strings.Join(results, t.separator)
}
