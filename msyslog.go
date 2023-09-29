package msyslog

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

type Severity uint8

const (
	EMERGENCY     Severity = 0
	ALERT         Severity = 1
	CRITICAL      Severity = 2
	ERROR         Severity = 3
	WARNING       Severity = 4
	NOTICE        Severity = 5
	INFORMATIONAL Severity = 6
	DEBUG         Severity = 7
)

type MSyslog struct {
	Facility uint8
	Severity Severity
	Hostname string
	AppName  string
	ProcId   string
	MsgId    string
	conn     net.Conn
}

func (m MSyslog) Write(p []byte) (int, error) {
	level := m.Facility*8 + uint8(m.Severity)
	fmt.Println(level, m.Facility, m.Severity)
	version := 1
	timestamp := time.Now().Format(time.RFC3339)
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "<%d>%d %s %s %s %d ", level, version, timestamp, m.Hostname, m.AppName, os.Getpid())
	buf.Write(p)
	return m.conn.Write(buf.Bytes())
}

func New(address *net.IP, port uint16) (MSyslog, error) {
	var logger MSyslog
	//dst := net.UDPAddr{IP: *address, Port: int(port)}
	//conn, err := net.DialUDP("udp", nil, &dst)
	conn, err := net.Dial("udp", net.JoinHostPort(address.String(), strconv.Itoa(int(port))))
	if err != nil {
		return logger, err
	}
	logger.conn = conn
	logger.Facility = 23
	logger.Severity = DEBUG
	logger.Hostname, _ = os.Hostname()
	logger.AppName = "-"
	logger.ProcId = "-"
	return logger, nil
}

func (m MSyslog) Close() error {
	return m.conn.Close()
}
