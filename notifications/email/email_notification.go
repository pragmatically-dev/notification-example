package email

import (
	"fmt"

	i "github.com/pragmatically-dev/notification-example/notifications"
)

// Implementación para Email
type EmailReceiver struct {
	emailAddress string
}

func (recv *EmailReceiver) SetContactInfo(email string) error {
	recv.emailAddress = email
	return nil
}

func (recv *EmailReceiver) GetContactInfo() (string, error) {
	return recv.emailAddress, nil
}

// EmailNotification es una implementación de notificación para correo electrónico
type EmailNotification struct {
	from     string
	to       i.Receiver
	priority i.Priority
	body     interface{}
	media    []interface{}
}

func (n *EmailNotification) SetPriority(p i.Priority) error {
	n.priority = p
	return nil
}

func (n *EmailNotification) GetPriority() (i.Priority, error) {
	return n.priority, nil
}

func (n *EmailNotification) SetMedia(m []interface{}) error {
	n.media = m
	return nil
}

func (n *EmailNotification) GetMedia() ([]interface{}, error) {
	return n.media, nil
}

func (n *EmailNotification) SetBody(b interface{}) error {
	n.body = b
	return nil
}

func (n *EmailNotification) GetBody() (interface{}, error) {
	return n.body, nil
}

func (n *EmailNotification) Send() error {
	contact, err := n.to.GetContactInfo()
	if err != nil {
		return err
	}
	fmt.Printf("Enviando correo a %s con prioridad %s y contenido: %v\n", contact, n.priority, n.body)
	return nil
}

func NewEmailNotification(from string, to i.Receiver) *EmailNotification {
	return &EmailNotification{
		from: from,
		to:   to,
	}
}
