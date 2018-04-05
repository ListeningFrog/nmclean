package main

import (
  "os"
  // "fmt"
  "gopkg.in/alecthomas/kingpin.v2"
  logging "github.com/op/go-logging"
)

var version = "none"

func getLog(debug bool) * logging.Logger {
  var log = logging.MustGetLogger("nmclean")

  // Example format string. Everything except the message has a custom color
  // which is dependent on the log level. Many fields have a custom output
  // formatting too, eg. the time returns the hour down to the milli second.
  var format = logging.MustStringFormatter(
    `%{color} %{level:.4s} %{id:03x}%{color:reset}> %{message}`,
  )

  backend := logging.NewLogBackend(os.Stdout, "", 0)
  backendFormatted := logging.NewBackendFormatter(backend, format)

  backendLeveled := logging.AddModuleLevel(backendFormatted)

  if debug {
    backendLeveled.SetLevel(logging.DEBUG, "");
  } else {
    backendLeveled.SetLevel(logging.INFO, "");
  }
  logging.SetBackend(backendLeveled)

  return log;
}

var (
  debug   = kingpin.Flag("debug", "Enable debug mode.").Bool()
  timeout = kingpin.Flag("timeout", "Timeout waiting for ping.").Default("5s").OverrideDefaultFromEnvar("PING_TIMEOUT").Short('t').Duration()
  ip      = kingpin.Arg("ip", "IP address to ping.").Required().IP()
  count   = kingpin.Arg("count", "Number of packets to send").Int()
)

func main() {
  kingpin.Version(version)
  kingpin.Parse()

  log := getLog(*debug);

  log.Info("Hello !");
  log.Debugf("Would ping: %s with timeout %s and count %d\n", *ip, *timeout, *count)
}
