package utility

import "net/smtp"

func StudySmtp() {
	auth := smtp.PlainAuth("", "spade_ace@naver.com", "***", "smtp.naver.com")

	from := "spade_ace@naver.com"
	to := []string{"spadeace.net@gmail.com"}

	headerSubject := "Subject: 테스트11\r\n"
	headerFrom := "From: spade_ace@naver.com\r\n"
	headerBlank := "\r\n"
	body := "메일 테스트올시다.\r\n"
	msg := []byte(headerSubject + headerFrom + headerBlank + body)

	err := smtp.SendMail("smtp.naver.com:587", auth, from, to, msg)
	if err != nil {
		panic(err)
	}
}
