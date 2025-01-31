package notifications

// aca lo que estoy haciendo es crear una especie de ENUM para las prioridades
type Priority int

const (
	Low Priority = iota
	Medium
	High
)

// Aca defino que comportamiento debe tener el destinatario de la notificacion, al hacerlo con una interfaz, me permite aceptar cualquier destinatario que implemente los metodos
type Receiver interface {
	SetContactInfo(string) error
	GetContactInfo() (string, error)
}

// Esta interfaz va definir el comportamiento base de cada notificacion uwu
type Notification interface {
	Send() error

	SetPriority(Priority) error
	GetPriority() (Priority, error)

	SetMedia([]interface{}) error
	GetMedia() ([]interface{}, error)

	SetBody(interface{}) error
	GetBody() (interface{}, error)
}
