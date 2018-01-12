go-graphite-client
===============

Package go-graphite-client is a simple client used to send mertics to graphite server using TCP or UDP.
It could be used to send a single metric or batch of mertics.

## Installation

To install go-graphite-client:
```
go get github.com/almariah/go-graphite-client
```
## Examples

To send single metric:
```go
```

To send batch of metrics:
```go
```

Example:

```go
package main

import (
  "github.com/almariah/go-graphite-client"
  "fmt"
)

func main() {

  // craete client
  client := &Client{
    Host: graphiteHost,
    Port: graphitePort,
    Protocol: TCP,
  }

  // connect a graphite server
  err := client.Connect()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  // send Metric
  client.Send("stats.test.metric", "1")         

}
```
