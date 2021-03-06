package servers

import "github.com/e-travel/message-dispatcher/dispatchers"

const (
	UNIXGRAM = "unixgram"
	UDP      = "udp"
)

type Server interface {
	Serve(buffer dispatchers.Dispatcher, running chan<- bool)
}

func CreateServer(serverType string, serverAddress string) Server {
	var server Server
	switch serverType {
	case UNIXGRAM:
		server = &UnixDatagram{
			Address: serverAddress,
		}
	case UDP:
		server = &Udp{
			Address: serverAddress,
		}
	}
	return server
}
