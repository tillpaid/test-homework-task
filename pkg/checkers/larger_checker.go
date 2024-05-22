package checkers

type LargerChecker struct {
	Target        int
	SuccessResult string
}

func NewLargerChecker(target int, successResult string) *LargerChecker {
	return &LargerChecker{Target: target, SuccessResult: successResult}
}

func (c *LargerChecker) Check(number int) string {
	if number > c.Target {
		return c.SuccessResult
	}

	return ""
}
