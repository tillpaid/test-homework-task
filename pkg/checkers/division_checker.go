package checkers

type DivisionChecker struct {
	divisor       int
	successResult string
}

func NewDivisionChecker(divisor int, successResult string) *DivisionChecker {
	return &DivisionChecker{divisor: divisor, successResult: successResult}
}

func (c *DivisionChecker) Check(number int) string {
	if number%c.divisor == 0 {
		return c.successResult
	}

	return ""
}
