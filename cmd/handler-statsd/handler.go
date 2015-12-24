// Take well-formed input in graphie format and send it to statsd
//
// LICENSE:
//   Copyright 2015 Yieldbot. <devops@yieldbot.com>
//   Released under the MIT License; see LICENSE
//   for details.

package main

import (
	"flag"
	"github.com/quipo/statsd"
	"github.com/yieldbot/sensuplugin/sensuhandler"
)

func main() {

	// set commandline flags
	HostPtr := flag.String("host", "localhost", "the statsd host")
	PortPtr := flag.String("port", "8125", "the statsd port")
	ProtocolPtr := flag.String("protocol", "udp", "the protocol to connect with")

	flag.Parse()
	statsdHost := *HostPtr
	statsdPort := *PortPtr
	statsdProtocol := *ProtocolPtr

	statsdClient := statsd.NewStatsdClient(statsdHost+":"+statsdPort, "")
	statsdClient.CreateSocket()
	defer statsdClient.Close()

	sensuEvent := new(sensuhandler.SensuEvent)
	sensuEvent = sensuEvent.AcquireSensuEvent()

	metricData := sensuEvent.Check.Output

	for i, m := range metricData {
		// I need to break up the metric into the string and the value
		statsdClient.FAbsolute(m, 0)
	}
}
