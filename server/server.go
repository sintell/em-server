package server

import (
	"github.com/sintell/em-server/server/controller"
	"github.com/sintell/em-server/server/middleware"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	STOPPED = 1 << iota
	STARTED
	STOPPING
	STARTING
)

type LogLevel uint8

const (
	ERRORS  LogLevel = 1 << iota
	VERBOSE LogLevel = 1 << iota
	DEBUG   LogLevel = 1 << iota
)

type Server struct {
	host     string
	port     string
	status   int
	logger   *log.Logger
	logLevel LogLevel
}

func New(addr string, loggingLevel LogLevel) *Server {
	host, port := strings.Split(addr, ":")[0], strings.Split(addr, ":")[1]

	if port == "" {
		port = "80"
	}
	var prefix string = "[HTTP_SERVER] "

	switch {
	case loggingLevel^DEBUG == 0:
		prefix = "[DEBUG] "
	case loggingLevel^VERBOSE == 0:
		prefix = "[VERBOSE] "
	case loggingLevel^ERRORS == 0:
		prefix = "[ERROR] "
	default:
		prefix = "[ERROR] "
		loggingLevel = ERRORS
	}

	return &Server{
		host,
		port,
		STOPPED,
		log.New(os.Stdout, prefix, log.Ldate|log.Ltime|log.Lmicroseconds),
		loggingLevel,
	}
}

func (s *Server) Start() error {
	s.status = STARTING
	if s.logLevel^DEBUG == 0 {
		s.logger.Printf("Start serving at %s:%s", s.host, s.port)
	}

	if s.logLevel^DEBUG == 0 {
		s.logger.Printf("Initialising controllers\n")
	}

	logging := middleware.NewLogging(s.logger)

	controllers := []controller.Controller{
		controller.Default(nil, nil),
		controller.User(),
	}

	for _, controller := range controllers {
		if s.logLevel^DEBUG == 0 {
			s.logger.Printf("Initialising controller for [%s]", controller.Resourse())
		}
		resourse, handler := controller.Bind()
		go http.Handle(resourse, logging(http.HandlerFunc(handler)))
	}

	http.ListenAndServe(s.host+":"+s.port, nil)
	s.status = STARTED
	return nil
}
