package checkers

type OneOfChecker struct {
	ExpectedList  []int
	SuccessResult string
}

func NewOneOfChecker(expectedList []int, successResult string) *OneOfChecker {
	return &OneOfChecker{ExpectedList: expectedList, SuccessResult: successResult}
}

func (c *OneOfChecker) Check(number int) string {
	for _, expected := range c.ExpectedList {
		if number == expected {
			return c.SuccessResult
		}
	}

	return ""
}
