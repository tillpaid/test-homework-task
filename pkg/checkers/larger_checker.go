package checkers

type LargerChecker struct {
	target        int
	successResult string
}

func NewLargerChecker(target int, successResult string) *LargerChecker {
	return &LargerChecker{target: target, successResult: successResult}
}

func (c *LargerChecker) Check(number int) string {
	if number > c.target {
		return c.successResult
	}

	return ""
}
