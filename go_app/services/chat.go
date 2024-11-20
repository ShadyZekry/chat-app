package services

func ValidateChat(application_token string, number string) bool {
	_, err := Get(application_token, number)

	if err != nil {
		return false
	}
	return true
}
