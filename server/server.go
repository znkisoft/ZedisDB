package server

import (
	"io"
	"strings"
	"sync"

	"github.com/znkisoft/zedisDB/server/handler"

	"github.com/znkisoft/zedisDB/lib/logger"
	"github.com/znkisoft/zedisDB/parser"

	"net"
)

type Server struct {
	mu       sync.RWMutex
	handlers map[string]handler.CmdFunc
	accept   func(conn *parser.RESPConn) bool
}

func NewServer() *Server {
	return &Server{
		handlers: handler.Router,
	}
}

func (s *Server) HandleFunc(command string, handler func(conn *parser.RESPConn, args []parser.Value) error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[strings.ToUpper(command)] = handler
}

func (s *Server) AcceptFunc(accept func(conn *parser.RESPConn) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.accept = accept
}

func (s *Server) ListenAndServe(addr string) error {
	// resolve tcp addr
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}

	// bind listener addr
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}
	defer l.Close()
	logger.CommonLog.Printf("(connected) ZedisDB is bounding to %s", addr)

	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			return err
		}
		go func() {
			err := s.handleConnection(conn)
			if err != nil {
				if _, ok := err.(*parser.ErrProtocol); ok {
					io.WriteString(conn, err.Error()+"\r\n")
				} else {
					io.WriteString(conn, "-ERR unknown error\r\n")
				}
			}
		}()
	}
}

func (s *Server) handleConnection(c net.Conn) error {
	conn := parser.NewRESPConn(c)
	defer conn.Conn.Close()

	s.mu.RLock()
	accept := s.accept
	s.mu.RUnlock()
	if accept != nil && !accept(conn) {
		return nil
	}

	for {
		v, _, err := conn.ReadValue()
		if err != nil {
			if err == io.EOF {
				logger.CommonLog.Printf("(disconnected) %s", c.RemoteAddr())
				return nil
			}
		}
		values := v.Array()
		if len(values) < 1 {
			return nil
		}
		command := strings.ToUpper(values[0].String())

		s.mu.RLock()
		h := s.handlers[command]
		s.mu.RUnlock()

		if h == nil {
			if err := conn.WriteError(
				parser.ErrProtocol{Type: parser.Client, Message: "unknown command '" + command + "'"},
			); err != nil {
				return err
			}
		} else {
			err := h(conn, values)
			if err != nil {
				return err
			}
		}
	}
}
