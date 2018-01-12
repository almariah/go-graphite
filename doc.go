// Package go-graphite-client is a simple client used to send mertics to graphite server using TCP or UDP.
// It could be used to send a single metric or batch of mertics.
//
// Example:
//
//     package main
//
//     import (
//         "github.com/almariah/go-graphite-client"
//         "fmt"
//     )
//
//     func main() {
//
//         // craete client
//         client := &Client{Host: graphiteHost, Port: graphitePort, Protocol: TCP}
//
//         // connect a graphite server
//         err := client.Connect()
//         if err != nil {
//             fmt.Println(err)
//             exit
//         }
//
//         client.Send("stats.test.metric", "1")
//     }
//
package graphite
