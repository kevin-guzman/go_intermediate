package main

import "fmt"

// SMS EMAIL

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}
type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// SMS
type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Se está enviando la notif via SMS")
}
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// EMAIL
type EMAILNotification struct {
}

func (EMAILNotification) SendNotification() {
	fmt.Println("Se está enviando la notif via SMS")
}
func (EMAILNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// ---- Sender ----- //
// Isender SMS
type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "twilo"
}

// Isender EMAIL
type EMAILNotificationSender struct {
}

func (EMAILNotificationSender) GetSenderMethod() string {
	return "EMAIL"
}
func (EMAILNotificationSender) GetSenderChannel() string {
	return "SES"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "EMAIL" {
		return &EMAILNotification{}, nil
	}
	return nil, fmt.Errorf("No notification type")
}

func SendNotification(f INotificationFactory) {
	f.SendNotification()
}

func GetMethod(f INotificationFactory) {
	fmt.Println("Method->", f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("EMAIL")
	SendNotification(smsFactory)
	SendNotification(emailFactory)

	GetMethod(smsFactory)
	GetMethod(emailFactory)
}
