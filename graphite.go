package graphite

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

// Mertic defines a graphite metric of name, value and timestamp
type Metric struct {
	Name      string
	Value     string
	Timestamp int64
}

// Client defines a graphite client
type Client struct {
	Host       string
	Port       int
	Protocol   string
	Timeout    time.Duration
	Connection net.Conn
}

const defaultTimeout = 5

// Connect is used to connect to graphite server
func (client *Client) Connect() error {
	if client.Connection != nil {
		client.Connection.Close()
	}

	var err error

	address := fmt.Sprintf("%s:%d", client.Host, client.Port)

	if client.Timeout == 0 {
		client.Timeout = defaultTimeout * time.Second
	}

	if client.Protocol == "udp" {
		udpAddr, err := net.ResolveUDPAddr("udp", address)
		if err != nil {
			return err
		}
		client.Connection, err = net.DialUDP(client.Protocol, nil, udpAddr)
	} else {
		client.Connection, err = net.DialTimeout(client.Protocol, address, client.Timeout)
	}

	if err != nil {
		return err
	}

	return nil
}

// Disconnect is used to disconnect to graphite server
func (client *Client) Disconnect() error {
	err := client.Connection.Close()
	client.Connection = nil
	return err
}

func (client *Client) SendMetric(metric Metric) error {
	metrics := make([]Metric, 1)
	metrics[0] = metric

	return client.sendMetrics(metrics)
}

func (client *Client) sendMetrics(metrics []Metric) error {

	zeroed_metric := Metric{} // ignore unintialized metrics
	buf := bytes.NewBufferString("")
	for _, metric := range metrics {
		if metric == zeroed_metric {
			continue // ignore unintialized metrics
		}
		if metric.Timestamp == 0 {
			metric.Timestamp = time.Now().Unix()
		}
		if client.Protocol == "udp" {
			fmt.Fprintf(client.Connection, "%s %s %d\n", metric.Name, metric.Value, metric.Timestamp)
			continue
		}
		buf.WriteString(fmt.Sprintf("%s %s %d\n", metric.Name, metric.Value, metric.Timestamp))
	}
	if client.Protocol == "tcp" {
		_, err := client.Connection.Write(buf.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}

func (graphite *Client) Send(stat string, value string) error {
	metrics := make([]Metric, 1)
	metrics[0] = Metric{
		Name:      stat,
		Value:     value,
		Timestamp: time.Now().Unix(),
	}
	err := graphite.sendMetrics(metrics)
	if err != nil {
		return err
	}
	return nil
}
