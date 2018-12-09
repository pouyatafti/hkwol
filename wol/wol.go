// based on the following blog post:
// https://sabhiram.com/development/2015/02/16/sending_wol_packets_with_golang.html

package wol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

type macAddr [6]byte

type magicPkt struct {
	header [6]byte
	payload [16]macAddr
}

func newMagicPkt(macStr string) (*magicPkt, error) {
	var pkt magicPkt
	var mac macAddr

	hwAddr, err := net.ParseMAC(macStr)
	if err != nil {
		return nil, err
	}

	for i := range mac {
		mac[i] = hwAddr[i]
	}
	for i := range pkt.header {
		pkt.header[i] = 0xFF
	}
	for i := range pkt.payload {
		pkt.payload[i] = mac
	}

	return &pkt, nil
}

func Broadcast(macStr string) error {
	pkt, err := newMagicPkt(macStr)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, pkt)

	udpAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:9")
	if err != nil {
        	return errors.New("cannot get UDP address")
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
        	return errors.New("cannot dial UDP address")
	}
	defer conn.Close()

	count, err := conn.Write(buf.Bytes())
	if err != nil {
		return err
	} else if count != 102 {
        	return errors.New("incorrect number of bytes written")
	}

	return nil
}
