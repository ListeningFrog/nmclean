package main

import (
  "os"
  // "fmt"
  "gopkg.in/alecthomas/kingpin.v2"
  logging "github.com/op/go-logging"
)

var version = "none"
var log = logging.MustGetLogger("nmclean")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color} %{level:.4s} %{id:03x}%{color:reset}> %{message}`,
)

var (
  debug   = kingpin.Flag("debug", "Enable debug mode.").Bool()
  timeout = kingpin.Flag("timeout", "Timeout waiting for ping.").Default("5s").OverrideDefaultFromEnvar("PING_TIMEOUT").Short('t').Duration()
  ip      = kingpin.Arg("ip", "IP address to ping.").Required().IP()
  count   = kingpin.Arg("count", "Number of packets to send").Int()
)

func main() {
  backend := logging.NewLogBackend(os.Stdout, "", 0)
  logging.SetBackend(logging.NewBackendFormatter(backend, format))

  kingpin.Version(version)
  kingpin.Parse()
  log.Debugf("Would ping: %s with timeout %s and count %d\n", *ip, *timeout, *count)
}
