package graphite

import (
	"net"
	"testing"
)

var graphiteHost = "localhost"
var graphitePort = 2003

func TestClient(t *testing.T) {
	client := &Client{Host: graphiteHost, Port: graphitePort, Protocol: "tcp"}

	client.Connect()

	if _, ok := client.Connection.(*net.TCPConn); !ok {
		t.Error("client.Connection is not TCP connection")
	}
}

func TestClientUPD(t *testing.T) {
	c := &Client{Host: graphiteHost, Port: graphitePort, Protocol: "udp"}

	c.Connect()

	if _, ok := c.Connection.(*net.UDPConn); !ok {
		t.Error("client.Connection is not UDP connection")
	}
}

func TestSendMetric(t *testing.T) {
	client := &Client{Host: graphiteHost, Port: graphitePort, Protocol: "tcp"}

	err := client.Connect()
	if err != nil {
		t.Error(err)
	}

	err = client.Send("stats.test.metric", "1")
	if err != nil {
		t.Error(err)
	}

	client.Disconnect()
}
