package interfaces

func GetExpenseReport(e expense) (string, float64) {
	// we have an expense interface object.
	s, ok := e.(sms)
	if ok {
		return s.toPhoneNumber, s.cost()
	}

	return "", 0.0
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * 0.1
	}
	return float64(len(s.body)) * 0.03
}
