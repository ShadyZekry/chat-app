package services

func ValidateMessage(application_token string, chat_number string, message_number string) bool {
	max_message_number, err := Get(application_token, chat_number)

	if err != nil {
		return false
	}
	return message_number <= max_message_number
}
