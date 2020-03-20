package webserver

type Server struct {
    Timeout time.Duration
    Port int
    ListenAddress string
}

type Option func(*Server)

func Timeout(d time.Duration) Option {
    return func(s *Server) {
        s.Timeout = d
    }
}
func Port(p int) Option {
    return func(s *Server) {
        s.Port = p
    }
}

func New(opts ...Option) *Server {
    // initialise the server with the default values
    s := &Server{
        Timeout: 500*time.Millisecond,
        Port: 0, // uses a random port on the host
        ListenAddres: "http://localhost",
    }
    // apply all options
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// usage examples from outside the package:
s := webserver.New() // uses all the default options
s := webserver.New(webserver.Timeout(time.Second), webserver.Port(8080))
