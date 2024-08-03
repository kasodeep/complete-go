package functions

func Reformat(message string, formatter func(string) string) string {
	message = formatter(message)
	message = formatter(message)
	message = formatter(message)
	message = "TEXTIO: " + message
	return message
}
