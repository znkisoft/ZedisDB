package server

import (
	"io"
	"strings"
	"sync"

	"github.com/znkisoft/zedisDB/server/handler"

	"github.com/znkisoft/zedisDB/parser"
	"github.com/znkisoft/zedisDB/pkg/logger"

	"net"
)

type Server struct {
	mu       sync.RWMutex
	handlers map[string]handler.CmdFunc
	accept   func(conn *parser.RESPConn) bool
	lruClock int64 /* Clock for LRU eviction */

	// TODO add stats fields
	// Stat stat
	// // Fields used only for stats
	// // 服务器启动时间
	// time_t stat_starttime;          /* Server start time */
	// // 已处理命令的数量
	// long long stat_numcommands;     /* Number of processed commands */
	// // 服务器接到的连接请求数量
	// long long stat_numconnections;  /* Number of connections received */
	//
	// // 已过期的键数量
	// long long stat_expiredkeys;     /* Number of expired keys */
	//
	// // 因为回收内存而被释放的过期键的数量
	// long long stat_evictedkeys;     /* Number of evicted keys (maxmemory) */
	//
	// // 成功查找键的次数
	// long long stat_keyspace_hits;   /* Number of successful lookups of keys */
	//
	// // 查找键失败的次数
	// long long stat_keyspace_misses; /* Number of failed lookups of keys */
	//
	// // 已使用内存峰值
	// size_t stat_peak_memory;        /* Max used memory record */
	//
	// // 最后一次执行 fork() 时消耗的时间
	// long long stat_fork_time;       /* Time needed to perform latest fork() */
	//
	// // 服务器因为客户端数量过多而拒绝客户端连接的次数
	// long long stat_rejected_conn;   /* Clients rejected because of maxclients */
	//
	// // 执行 full sync 的次数
	// long long stat_sync_full;       /* Number of full resyncs with slaves. */
	//
	// // PSYNC 成功执行的次数
	// long long stat_sync_partial_ok; /* Number of accepted PSYNC requests. */
	//
	// // PSYNC 执行失败的次数
	// long long stat_sync_partial_err;/

	// TODO add pub/sub fields
	// /* Pubsub */
	// // 字典，键为频道，值为链表
	// // 链表中保存了所有订阅某个频道的客户端
	// // 新客户端总是被添加到链表的表尾
	// dict *pubsub_channels;  /* Map channels to list of subscribed clients */
	//
	// // 这个链表记录了客户端订阅的所有模式的名字
	// list *pubsub_patterns;  /* A list of pubsub_patterns */
	// typedef struct pubsubPattern {
	// 	// 订阅模式的客户端
	// 	redisClient *client;
	// 	// 被订阅的模式
	// 	robj *pattern;
	// } pubsubPattern;
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
				if _, ok := err.(parser.ErrProtocol); ok {
					io.WriteString(conn, err.Error())
				} else {
					io.WriteString(conn, "-ERR unknown error\r\n")
				}
			}
			conn.Close()
		}()
	}
}

func (s *Server) handleConnection(c net.Conn) error {
	conn := parser.NewRESPConn(c)

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
