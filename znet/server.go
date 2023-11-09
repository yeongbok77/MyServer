package znet

import (
	"fmt"
	"github.com/yeongbok77/MyServer/ziface"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

// 启动
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP :%s, Port: %d, is starting", s.IP, s.Port)
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("resolve tcp listen error: ", err)
			return
		}
		fmt.Println("Start Server success, ", s.Name, "Listenning...")
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err: ", err)
				continue
			}
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err: ", err)
						continue
					}
					// 回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err: ", err)
						continue
					}
				}
			}()
		}
	}()

}

// 停止
func (s *Server) Stop() {
	// TODO: 释放资源
}

// 运行
func (s *Server) Serve() {
	s.Start()
	select {}
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
