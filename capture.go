package main

import (
	"fmt"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/pkg/errors"
)

const (
	snapshotLen int32         = 1024
	promiscuous bool          = true
	timeout     time.Duration = 30 * time.Second
)

func chooseInterface() (*string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, errors.Wrap(err, "Fail net.Interfaces")
	}

	for _, i := range ifaces {
		if i.Flags&net.FlagUp == 0 || i.HardwareAddr == nil {
			continue
		}

		addrs, err := i.Addrs()
		if err != nil {
			return nil, errors.Wrap(err, "Fail Interface.Addrs()")
		}

		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPNet:
				if len(v.Mask) == 4 {
					return &i.Name, nil
				}
			}
		}

	}

	return nil, nil
}

func capture(interfaceName string, db database) error {
	if interfaceName == "" {
		ifName, err := chooseInterface()
		if err != nil {
			return errors.Wrap(err, "Fail selectInterface")
		}
		if ifName == nil {
			return fmt.Errorf("No avaiable network interface is found. Use -i option")
		}
		interfaceName = *ifName
		logger.Infof("Interface name is not set, then %s is chosen", ifName)
	}

	logger.Infow("Starting packet capture", "interface", interfaceName)

	handle, err := pcap.OpenLive(interfaceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		return errors.Wrap(err, "Fail to OpenLive")
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		if err := handlePacket(packet, db); err != nil {
			return errors.Wrap(err, "Fail handlePacket")
		}
	}

	return nil
}
