package znet

import (
	"github.com/yeongbok77/MyServer/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

// 启动
func (s *Server) Start() {

}

// 停止
func (s *Server) Stop() {

}

// 运行
func (s *Server) Serve() {

}

// new a server
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
