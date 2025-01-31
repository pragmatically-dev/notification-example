package main

import (
	"fmt"

	"github.com/pragmatically-dev/notification-example/notifications"
	"github.com/pragmatically-dev/notification-example/notifications/email"
	"github.com/pragmatically-dev/notification-example/notifications/whatsapp"
)

// Estructura (i
// Con esta funcion se ve el poder de las interfaces en golang:
// esta es una funcion varidica, o sea que no tiene una aridad (arity )definida
func SendAllNotifications(notification ...notifications.Notification) error {
	// Fijate que estoy iterando sobre entidades que implementan la interfaz [Notification]
	// Por lo tanto se van a enviar ambos tipos de notificaciones, de whatsapp y de email
	//Se podria agragar concurrencia etc
	for _, n := range notification {
		err := n.Send()
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// creo un receptor de email
	receiverEmail := &email.EmailReceiver{}
	receiverEmail.SetContactInfo("usuario@example.com")

	// Creo la notificación
	emailNotif := email.NewEmailNotification("noreply@example.com", receiverEmail)
	emailNotif.SetPriority(notifications.High)
	emailNotif.SetBody("Este es el contenido del correo")
	emailNotif.SetMedia([]interface{}{"imagen.jpg", "documento.pdf"})

	//Creo receptor para whastapp
	receiverWsp := &whatsapp.WhatsAppReceiver{}
	err := receiverWsp.SetContactInfo("+5493513393200")
	if err != nil {
		fmt.Println("Error configurando receptor:", err)
		return
	}

	// Crear notificación de WhatsApp
	whatsAppNotif := whatsapp.NewWhatsAppNotification("tunumero", receiverWsp)
	whatsAppNotif.SetPriority(notifications.High)
	whatsAppNotif.SetBody("Este es un mensaje de prueba desde golang.")
	whatsAppNotif.SetMedia([]interface{}{"test.png", "test.mp4"})

	//Ahora tenemos listos dos mensajes paar enviar, uno por mail y otro por whatsaap
	//usamos el poder de las interfaces para hacerlo todo desde un solo lugar:
	SendAllNotifications(emailNotif, whatsAppNotif)
}
