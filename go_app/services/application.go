package services

func ValidateToken(token string) bool {
	_, err := Get(token, "chats")

	if err != nil {
		return false
	}
	return true
}
