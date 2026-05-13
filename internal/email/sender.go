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
	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<body style="margin:0; padding:0; background:#f7f7f5; font-family:Arial, sans-serif; color:#2f2f2f;">
		  <div style="max-width:560px; margin:40px auto; background:#ffffff; border-radius:24px; padding:36px; box-shadow:0 8px 24px rgba(0,0,0,0.06);">
		    <h1 style="margin:0 0 12px; font-size:28px; color:#2f2f2f;">
		      PPlaner 🌿
		    </h1>

		    <p style="font-size:16px; line-height:1.6; margin:0 0 24px;">
		      Привіт! На зв’язку команда PPlaner ✨
		    </p>

		    <p style="font-size:16px; line-height:1.6; margin:0 0 24px;">
		      Дякуємо за довіру та реєстрацію в нашому додатку.
		      Щоб підтвердити вашу електронну пошту й завершити створення акаунту, введіть цей код:
		    </p>

		    <div style="background:#eef4ef; border:1px solid #9caf9f; border-radius:18px; padding:22px; text-align:center; margin:28px 0;">
		      <div style="font-size:34px; font-weight:700; letter-spacing:8px; color:#8fa68f;">
		        %s
		      </div>
		    </div>

		    <p style="font-size:14px; line-height:1.6; color:#777; margin:0 0 16px;">
		      Код дійсний протягом 10 хвилин.
		    </p>

		    <p style="font-size:14px; line-height:1.6; color:#777; margin:0;">
		      Якщо ви не створювали акаунт у PPlaner, просто проігноруйте цей лист.
		    </p>

		    <hr style="border:none; border-top:1px solid #eeeeee; margin:28px 0;">

		    <p style="font-size:14px; color:#8fa68f; margin:0;">
		      З турботою,<br>
		      команда PPlaner 🌿
		    </p>
		  </div>
		</body>
		</html>
	`, code)

	message := []byte(
		"From: " + s.cfg.From + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
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
