package controller

import (
  "github.com/gin-gonic/gin"
  "log"
  "net/http"
  "net/smtp"
  "os"
)

type Email struct {
  To      string `json:"to"`
  Subject string `json:"subject"`
  Body    string `json:"body"`
}

func MailSend(c *gin.Context) {
  var email Email

  if err := c.BindJSON(&email); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  // Set up authentication information.
  auth := smtp.PlainAuth(
    "",
    os.Getenv("MAIL_USER"),     // Your email
    os.Getenv("MAIL_PASSWORD"), // Your password
    "smtp.gmail.com",           // SMTP server
  )

  // Set up email content.
  msg := []byte("To: " + email.To + "\r\n" +
    "Subject: " + email.Subject + "\r\n" +
    "\r\n" +
    email.Body + "\r\n")

  // Send email.
  err := smtp.SendMail(
    "smtp.gmail.com:587", // SMTP server and port
    auth,
    os.Getenv("MAIL_USER"), // From
    []string{email.To},     // To
    msg,
  )
  if err != nil {
    log.Println(err)
    c.JSON(http.StatusOK, gin.H{
      "message": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Email sent successfully!",
  })
}
