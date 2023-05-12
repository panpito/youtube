package phone_number_validator

type LengthChecker struct {
	standardLength int
}

func NewLengthChecker(standardLength int) *LengthChecker {
	return &LengthChecker{standardLength: standardLength}
}

func (checker LengthChecker) IsValid(number string) bool {
	if len(number) != checker.standardLength {
		return false
	}

	return true
}
