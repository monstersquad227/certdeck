package service

import (
	"certdeck/config"
	"certdeck/repository"
	"certdeck/utils"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
)

type EmailService struct {
	EmailRepo *repository.EmailRepository
}

func (e *EmailService) SendEmail(email string) error {

	htmlBody := fmt.Sprintf(`
		<div style="max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #e0e0e0; border-radius: 10px; box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1); font-family: sans-serif;">
			<h2 style="color: #333;">🔔 认证提醒</h2>
			<p style="font-size: 16px; color: #555;">
				您正在尝试登录 <strong>CertDeck</strong>，验证码如下：
			</p>
			<div style="text-align: center; font-size: 28px; color: #1976d2; font-weight: bold; margin: 20px 0;">
				%s
			</div>
			<p style="font-size: 14px; color: #999;">验证码有效期为 5 分钟，请勿泄露。</p>
			<hr />
			<p style="font-size: 12px; color: #bbb; text-align: center;">MonsterSquad © 2025</p>
		</div>
	`, utils.GenerateVerificationCode())
	message := gomail.NewMessage()
	//message.SetHeader("From", config.GlobalConfig.Mail.Username)
	message.SetAddressHeader("From", config.GlobalConfig.Mail.Username, config.GlobalConfig.Application.Name)
	message.SetHeader("To", email)
	message.SetHeader("Subject", "请查收您的验证码")
	message.SetBody("text/html", htmlBody)

	d := gomail.NewDialer("smtp.163.com", 465, config.GlobalConfig.Mail.Username, config.GlobalConfig.Mail.Password)
	d.SSL = true

	if err := d.DialAndSend(message); err != nil {
		log.Fatalf("发送邮件失败: %v", err)
	}
	log.Println("发送成功")
	return nil
}
