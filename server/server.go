package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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
		prefix = "[DEBUG] [HTTP_SERVER] "
	case loggingLevel^VERBOSE == 0:
		prefix = "[VERBOSE] [HTTP_SERVER] "
	case loggingLevel^ERRORS == 0:
		prefix = "[ERROR] [HTTP_SERVER] "
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			s.logger.Fatalf("Can't parse request. Reason: %s", err.Error())
		}

		id, err := strconv.Atoi(r.Form.Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			s.logger.Fatalf("Can't parse request. Reason: %s", err.Error())
		}

		action := r.Form.Get("action")

		if s.logLevel^DEBUG == 0 {
			s.logger.Printf("id: %d, action: %s\n", id, action)
		}
		fmt.Fprintf(w, "%d", 200)
	})

	http.ListenAndServe(s.host+":"+s.port, nil)
	s.status = STARTED
	return nil
}
