package functions

func SplitEmail(email string) (string, string) {
	username, domain := "", ""

	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}
