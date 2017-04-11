package mailapi

import (
	"net/http"

	"gopkg.in/gomail.v2"
)

type SendMailRequest struct {
	Message string `json:"message"`
}

type service struct {
	d       *gomail.Dialer
	from    string
	to      string
	subject string
}

func NewService(d *gomail.Dialer, to string, from string, subject string) *service {
	return &service{d: d, to: to, from: from}
}

func (s *service) SendMail(r *http.Request, req *SendMailRequest) (bool, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", s.to)
	m.SetHeader("Subject", s.subject)
	m.SetBody("text/html", req.Message)

	if err := s.d.DialAndSend(m); err != nil {
		return false, err
	}
	return true, nil
}
