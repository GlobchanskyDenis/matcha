package main

import (
    "fmt"
    "net/smtp"
)

func main() {
    // user we are authorizing as
    from := "bsabre.cat@gmail.com"

    // use we are sending email to
    to := "globchansky.denis@gmail.com"

    // server we are authorized to send email through
    host := "smtp.gmail.com"

    // Create the authentication for the SendMail()
    // using PlainText, but other authentication methods are encouraged
    auth := smtp.PlainAuth("", from, "den619392", host)

    // NOTE: Using the backtick here ` works like a heredoc, which is why all the 
    // rest of the lines are forced to the beginning of the line, otherwise the 
    // formatting is wrong for the RFC 822 style
    message := `To: "Some User" <someuser@example.com>
From: "Other User" <otheruser@example.com>
Subject: Testing Email From Go!!

This is the message we are sending. That's it!
<span>comment</span>
`

    if err := smtp.SendMail(host+":587", auth, from, []string{to}, []byte(message)); err != nil {
        fmt.Println("Error SendMail: ", err)
		return
    }
    fmt.Println("Email Sent!")
}