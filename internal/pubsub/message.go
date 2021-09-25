package pubsub

type Message string

const (
	MessageOpen  Message = "openpls"
	MessageClose Message = "close"
)
