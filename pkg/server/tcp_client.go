// Implement a TCP client to comunicate from victim machine
package server

// Server config
type ServerConfig struct {
	Addr string // The address server
	Port string // The port to listen server
}

// Function to reate a new server
func NewServer(Addr, Port string) *ServerConfig {
	s := &ServerConfig{
		Addr: Addr, // Add address to server config
		Port: Port, // Add port to server config
	}
	return s
}

// Listen server function
func (s *ServerConfig) Listen() {

}

// Send data to server
