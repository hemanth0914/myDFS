package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established connection
type TCPPeer struct {
	//conn is the underlying connection of the peer
	conn net.Conn
	// if we dial and retrive a conn => outbound
	// if we accept and retrive a conn => inbound
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress	string
	listener		net.Listener
	shakeHands		HandshakeFunc
	decoder			Decoder

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}


func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		shakeHands: NOPHandshakeFunc, 
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)

	if err != nil {
		return err
	}
	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()

		if err != nil {
			fmt.Printf("TCP accept error: %s", err)
		}
		fmt.Printf("New incoming Connection %+v\n", conn)
		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err:= t.shakeHands(peer); err != nil {

	}

	//Read loop
	msg:= &Temp{}
	for{
		if err:= t.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("tcp error: %s\n", err)
			continue 
			 
		}
	}

	
}
