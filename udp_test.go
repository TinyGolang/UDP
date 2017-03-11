package udp_test

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	pUDPAddr, err := net.ResolveUDPAddr("udp", ":7070")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}

	pUDPConn, err := net.ListenUDP("udp", pUDPAddr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}

	defer pUDPConn.Close()

	buf := make([]byte, 256)
	for {

		n, pUDPAddr, err := pUDPConn.ReadFromUDP(buf)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "recv: %s", string(buf[0:n]))

		_, err = pUDPConn.WriteToUDP(buf[0:n], pUDPAddr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			return
		}

	}
}
func TestClient(t *testing.T) {
	pUDPAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:7070")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ResolveUDPAddr: %s", err.Error())
		return
	}

	pUDPConn, err := net.DialUDP("udp", nil, pUDPAddr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error DialUDP: %s", err.Error())
		return
	}

	n, err := pUDPConn.Write([]byte("你好啊！！！"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error WriteToUDP: %s", err.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "writed: %d", n)

	buf := make([]byte, 1024)
	n, _, err = pUDPConn.ReadFromUDP(buf)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ReadFromUDP: %s", err.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "readed: %d  %s", n, string(buf[:n]))
}
func TestPacketServer(t *testing.T) {

	fmt.Println("TestListenPacketServer")

	packetConn, err := net.ListenPacket("udp", ":7070")

	if err != nil {
		fmt.Fprintf(os.Stdout, "Error: %s", err.Error())
		return
	}
	defer packetConn.Close()

	var buf [512]byte
	for {
		n, addr, err := packetConn.ReadFrom(buf[0:])
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error: %s", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "recv: %s", string(buf[0:n]))

		_, err = packetConn.WriteTo(buf[0:n], addr)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error: %s", err.Error())
			return
		}
	}

}

func TestPacketClient(t *testing.T) {
	conn, err := net.Dial("udp", "127.0.0.1:7070")

	if err != nil {
		fmt.Fprintf(os.Stdout, "Error: %s", err.Error())
		return
	}

	defer conn.Close()

	_, err = conn.Write([]byte("你好啊UDP"))
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error: %s", err.Error())
		return
	}

	var buf [512]byte
	conn.SetReadDeadline(time.Now().Add(time.Second * 1)) // 阻塞，直到接收到消息,设置阻塞时间1秒
	n, err := conn.Read(buf[0:])
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error: %s", err.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "recv: %s", string(buf[0:n]))
}
