package main

import (
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/pkg/errors"
)

func handlePacket(pkt gopacket.Packet, db database) error {
	dnsLayer := pkt.Layer(layers.LayerTypeDNS)
	if dnsLayer == nil { // No DNS layer
		return nil
	}

	dns, _ := dnsLayer.(*layers.DNS)
	if dns.QR == false { // means QR is 0 (Query). boolean is confusing
		return nil
	}

	recTypeMap := map[layers.DNSType]string{
		layers.DNSTypeA:    "A",
		layers.DNSTypeAAAA: "AAAA",
	}

	for _, answer := range dns.Answers {
		recType, ok := recTypeMap[answer.Type]
		if !ok {
			continue
		}

		name := string(answer.Name)
		addr := answer.IP.String()
		if err := db.put(name, addr, recType, time.Now().UTC()); err != nil {
			return errors.Wrap(err, "Fail database.put")
		}
	}

	return nil
}
