package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"strconv"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

type SimpleHook struct {
}

func (s *SimpleHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.WarnLevel,
	}
}

func (s *SimpleHook) Fire(entry *logrus.Entry) error {
	entry.Data["message"] = entry.Message
	entry.Message = "Sensitive data was scrubbed"
	return nil
}

func TestLoggerSimpleHook(t *testing.T) {
	log := logrus.New()            // Create a new instance of the Logger type
	log.Hooks.Add(new(SimpleHook)) // log.AddHook(new(SimpleHook)) // log.AddHook(&SimpleHook{})
	log.Error("This is an error")  // Log with Hook
	log.Warn("This is a warning")  // Log with Hook
	log.Info("This is an info")    // Log without Hook
	log.Debug("This is a debug")   // Log without Hook
}

const (
	format = "20060102 15:04:05"
)

// MailHook to sends logs by email without authentication.
type MailHook struct {
	AppName string
	c       *smtp.Client
}

// MailAuthHook to sends logs by email with authentication.
type MailAuthHook struct {
	AppName  string
	Host     string
	Port     int
	From     *mail.Address
	To       *mail.Address
	Username string
	Password string
}

// NewMailHook creates a hook to be added to an instance of logger.
func NewMailHook(appname string, host string, port int, from string, to string) (*MailHook, error) {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial(host + ":" + strconv.Itoa(port))
	if err != nil {
		return nil, err
	}

	// Validate sender and recipient
	sender, err := mail.ParseAddress(from)
	if err != nil {
		return nil, err
	}
	recipient, err := mail.ParseAddress(to)
	if err != nil {
		return nil, err
	}

	// Set the sender and recipient.
	if err := c.Mail(sender.Address); err != nil {
		return nil, err
	}
	if err := c.Rcpt(recipient.Address); err != nil {
		return nil, err
	}

	return &MailHook{
		AppName: appname,
		c:       c,
	}, nil

}

// NewMailAuthHook creates a hook to be added to an instance of logger.
func NewMailAuthHook(appname string, host string, port int, from string, to string, username string, password string) (*MailAuthHook, error) {
	// Check if server listens on that port.
	conn, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), 3*time.Second)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Validate sender and recipient
	sender, err := mail.ParseAddress(from)
	if err != nil {
		return nil, err
	}
	receiver, err := mail.ParseAddress(to)
	if err != nil {
		return nil, err
	}

	return &MailAuthHook{
		AppName:  appname,
		Host:     host,
		Port:     port,
		From:     sender,
		To:       receiver,
		Username: username,
		Password: password}, nil
}

// Fire is called when a log event is fired.
func (hook *MailHook) Fire(entry *logrus.Entry) error {
	wc, err := hook.c.Data()
	if err != nil {
		return err
	}
	defer wc.Close()
	message := createMessage(entry, hook.AppName)
	if _, err = message.WriteTo(wc); err != nil {
		return err
	}
	return nil
}

// Fire is called when a log event is fired.
func (hook *MailAuthHook) Fire(entry *logrus.Entry) error {
	auth := smtp.PlainAuth("", hook.Username, hook.Password, hook.Host)

	message := createMessage(entry, hook.AppName)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		hook.Host+":"+strconv.Itoa(hook.Port),
		auth,
		hook.From.Address,
		[]string{hook.To.Address},
		message.Bytes(),
	)
	if err != nil {
		return err
	}
	return nil
}

// Levels returns the available logging levels.
func (hook *MailAuthHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
}

// Levels returns the available logging levels.
func (hook *MailHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
}

func createMessage(entry *logrus.Entry, appname string) *bytes.Buffer {
	body := entry.Time.Format(format) + " - " + entry.Message
	subject := appname + " - " + entry.Level.String()
	fields, _ := json.MarshalIndent(entry.Data, "", "\t")
	contents := fmt.Sprintf("Subject: %s\r\n\r\n%s\r\n\r\n%s", subject, body, fields)
	message := bytes.NewBufferString(contents)
	return message
}

func TestNewMailAuthHookFailed(t *testing.T) {

	// invalid port
	_, err := NewMailAuthHook("testapp", "smtp.gmail.com", 10, "user.name@gmail.com", "user.name@gmail.com", "user.name", "password")
	if err == nil {
		t.Errorf("no error on invalid port")
	}

	// invalid mail host
	_, err = NewMailAuthHook("testapp", "www.gmail.com", 587, "user.name@gmail.com", "user.name@gmail.com", "user.name", "password")
	if err == nil {
		t.Errorf("no error on invalid hostname")
	}

	// invalid email address
	_, err = NewMailAuthHook("testapp", "smtp.gmail.com", 587, "user.name", "user.name@gmail.com", "user.name", "password")
	if err == nil {
		t.Errorf("no error on invalid email address")
	}

}

func TestHookAuthMailSuccess(t *testing.T) {
	log := logrus.New()
	// if you do not need authentication for your smtp host
	hook, err := NewMailAuthHook("Test Logrus", "smtp.gmail.com", 587, "from@anakdesa.id", "to@anakdesa.id", "username@anakdesa.id", "password")

	if err == nil {
		log.Hooks.Add(hook)
		log.WithFields(logrus.Fields{
			"message": "test",
			"level":   "info",
		}).Error("test")
	} else {
		t.Errorf("error on creating hook")
	}
}
