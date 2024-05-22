package checkers

type DivisionChecker struct {
	Divisor       int
	SuccessResult string
}

func NewDivisionChecker(divisor int, successResult string) *DivisionChecker {
	return &DivisionChecker{Divisor: divisor, SuccessResult: successResult}
}

func (c *DivisionChecker) Check(number int) string {
	if number%c.Divisor == 0 {
		return c.SuccessResult
	}

	return ""
}
