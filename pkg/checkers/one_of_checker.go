package checkers

type OneOfChecker struct {
	expectedList  []int
	successResult string
}

func NewOneOfChecker(expectedList []int, successResult string) *OneOfChecker {
	return &OneOfChecker{expectedList: expectedList, successResult: successResult}
}

func (c *OneOfChecker) Check(number int) string {
	for _, expected := range c.expectedList {
		if number == expected {
			return c.successResult
		}
	}

	return ""
}
