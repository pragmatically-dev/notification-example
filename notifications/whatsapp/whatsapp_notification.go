package whatsapp

import (
	"errors"
	"fmt"

	i "github.com/pragmatically-dev/notification-example/notifications"
)

// Implementación para WhatsApp
type WhatsAppReceiver struct {
	phoneNumber string
}

func (recv *WhatsAppReceiver) SetContactInfo(phone string) error {
	if phone == "" {
		return fmt.Errorf("El número está vacío")
	}
	recv.phoneNumber = phone
	return nil
}

func (recv *WhatsAppReceiver) GetContactInfo() (string, error) {
	if recv.phoneNumber == "" {
		return "", errors.New("El número de teléfono no está seteado")
	}
	return recv.phoneNumber, nil
}

// WhatsAppNotification es una implementación de notificación para WhatsApp
type WhatsAppNotification struct {
	from     string
	to       i.Receiver
	priority i.Priority
	body     interface{}
	media    []interface{}
}

func (n *WhatsAppNotification) SetPriority(p i.Priority) error {
	n.priority = p
	return nil
}

func (n *WhatsAppNotification) GetPriority() (i.Priority, error) {
	return n.priority, nil
}

func (n *WhatsAppNotification) SetMedia(m []interface{}) error {
	n.media = m
	return nil
}

func (n *WhatsAppNotification) GetMedia() ([]interface{}, error) {
	return n.media, nil
}

func (n *WhatsAppNotification) SetBody(b interface{}) error {
	n.body = b
	return nil
}

func (n *WhatsAppNotification) GetBody() (interface{}, error) {
	return n.body, nil
}

func (n *WhatsAppNotification) Send() error {
	contact, err := n.to.GetContactInfo()
	if err != nil {
		return err
	}
	fmt.Printf("Enviando WhatsApp a %s con prioridad %s y contenido: %v\n", contact, n.priority, n.body)
	if len(n.media) > 0 {
		fmt.Printf("Adjuntos: %v\n", n.media)
	}
	return nil
}

func NewWhatsAppNotification(from string, to i.Receiver) *WhatsAppNotification {
	return &WhatsAppNotification{

		from: from,
		to:   to,
	}
}
