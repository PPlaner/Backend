package email

import (
	"fmt"
	"net/smtp"

	"github.com/PPlaner/Backend/internal/config"
)

type Sender struct {
	cfg config.SMTP
}

func NewSender(cfg config.SMTP) *Sender {
	return &Sender{
		cfg: cfg,
	}
}

func (s *Sender) SendVerificationCode(to string, code string) error {
	auth := smtp.PlainAuth(
		"",
		s.cfg.User,
		s.cfg.Password,
		s.cfg.Host,
	)

	subject := "Код підтвердження реєстрації"
	body := fmt.Sprintf(
		"Привіт!\n\nНа зв’язку команда PPlaner 💌\n\nДякуємо за довіру та реєстрацію в нашому додатку. Щоб підтвердити вашу електронну пошту й завершити створення акаунта, введіть цей код:\n\n%s\n\nКод дійсний протягом 10 хвилин.\n\nЯкщо ви не створювали акаунт у PPlaner, просто проігноруйте цей лист.\n\nЗ турботою,\nкоманда PPlaner",
		code,
	)

	message := []byte(
		"From: " + s.cfg.From + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
			"\r\n" +
			body,
	)

	addr := s.cfg.Host + ":" + s.cfg.Port

	return smtp.SendMail(
		addr,
		auth,
		s.cfg.User,
		[]string{to},
		message,
	)
}
