package session

import (
	"log"
	"net"
	"sync"
)
var sessionLock sync.RWMutex


type TCPSession struct {
	conn *net.TCPConn
	pid uint32
	msgChan chan []byte
}
/*
创建一个session
*/
func CreateTCPSession(c *net.TCPConn,pid uint32) *TCPSession {
	sessionLock.Lock()
	defer sessionLock.Unlock()

	session:= &TCPSession{
		conn:c,
		pid :pid,
	}
	session.Init()
	return session
}

func DestoryTCPSession(session *TCPSession) {
	sessionLock.Lock()
	defer sessionLock.Unlock()
	session.Destory()
}

func (t *TCPSession) Destory() {
	if t.conn != nil {
		t.conn.Close()
	}
	t.conn = nil
}

func (this *TCPSession) Init() {
	//t.StartReader()
	//this.StartAuth()
}

func(this *TCPSession)StartAuth ()  {

}

func (t *TCPSession) StartReader() {
	log.Println("[Reader Goroutine is running]")
	defer log.Println(t.conn.RemoteAddr().String(), "[conn Reader exit!]")
}

