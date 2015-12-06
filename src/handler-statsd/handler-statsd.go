// Take well-formed input in graphie format and send it to statsd
//
// LICENSE:
//   Copyright 2015 Yieldbot. <devops@yieldbot.com>
//   Released under the MIT License; see LICENSE
//   for details.

package main

import (
	"flag"
	dracky "github.com/yieldbot/sensu-yieldbot-library/src"
  "github.com/quipo/statsd"
)

func main() {

  // set commandline flags
	HostPtr := flag.String("host", "localhost" "the statsd host")
	PortPtr := flag.String("port", "8125", "the statsd port")
  ProtocolPtr := flag.String("protocol", "udp", "the protocol to connect with")

	flag.Parse()
	statsdHost := *HostPtr
	statsdPort := *PortPtr
  statsdProtocol := *ProtocolPtr

   statsdclient := statsd.NewStatsdClient(statsdHost + ":" + statsdPort)
   statsdclient.CreateSocket()
   defer stats.Close()

   sensuEvent := new(dracky.SensuEvent)
   sensuEvent = sensuEvent.AcquireSensuEvent()

   metricData := sensuEvent.Check.Output

   for i, m := range metricData {
     statsdclient.SesdEvent(m)
   }
 }
