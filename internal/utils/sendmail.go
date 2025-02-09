package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendMail(to, subject, body, attachmentPath string) error {

	//Step 1: Load Environment Variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	//Step 2: Retrieve SMTP Configuration
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	//Step 3: Check for Missing SMTP Configuration
	if smtpHost == "" || smtpPortStr == "" || smtpUser == "" || smtpPass == "" {
		return fmt.Errorf("missing SMTP configuration")
	}

	//Step 4: Convert SMTP Port to Integer
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		return fmt.Errorf("invalid SMTP_PORT: %v", err)
	}

	//Step 5: Create the Email Message
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", smtpUser)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/plain", body)

	//Step 6: Attach File (Optional)
	if attachmentPath != "" {
		mailer.Attach(attachmentPath)
	}

	//Step 7: Set Up the SMTP Connection and Send Email
	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	if err := dialer.DialAndSend(mailer); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	log.Println("Email sent successfully to:", to)
	return nil
}
