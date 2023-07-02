package model

import (
	"crypto/tls"
	"fmt"
	"net/http/httptrace"
	"time"
)

type RequestStatistics struct {
	RequestStartTime         time.Time
	GotFirstResponseByteTime time.Time
	DNSStartTime             time.Time
	DNSDoneTime              time.Time
	ConnectStartTime         time.Time
	ConnectDoneTime          time.Time
	TLSHandshakeStartTime    time.Time
	TLSHandshakeDoneTime     time.Time

	ClientTrace *httptrace.ClientTrace
}

type clientTrace struct {
	stats RequestStatistics
}

func InitRequestStatistics(requestStats *RequestStatistics) {
	clientTrace := &httptrace.ClientTrace{
		DNSStart:             requestStats.dnsStart,
		DNSDone:              requestStats.dnsDone,
		ConnectStart:         requestStats.connectStart,
		ConnectDone:          requestStats.connectDone,
		TLSHandshakeStart:    requestStats.tlsHandshakeStart,
		TLSHandshakeDone:     requestStats.tlsHandshakeDone,
		GotFirstResponseByte: requestStats.gotFirstResponseByte,
	}
	requestStats.RequestStartTime = time.Now()
	requestStats.ClientTrace = clientTrace
}

func (h *RequestStatistics) dnsStart(dnsStartInfo httptrace.DNSStartInfo) {
	h.DNSStartTime = time.Now()
	fmt.Printf("DNS Start Info: %+v\n", dnsStartInfo)
}

func (h *RequestStatistics) dnsDone(dnsDoneInfo httptrace.DNSDoneInfo) {
	h.DNSDoneTime = time.Now()
	fmt.Printf("DNS Done Info: %+v\n", dnsDoneInfo)
}

func (h *RequestStatistics) connectStart(network, addr string) {
	h.ConnectStartTime = time.Now()
	fmt.Printf("Connecting to: %s : %s\n", network, addr)
}

func (h *RequestStatistics) connectDone(network, addr string, err error) {
	h.ConnectDoneTime = time.Now()
	if err != nil {
		fmt.Printf("Error while Connecting to: %s : %s\n; Error: %+v\n", network, addr, err)
	} else {
		fmt.Printf("Successfully connected to: %s : %s\n", network, addr)
	}
}

func (h *RequestStatistics) tlsHandshakeStart() {
	h.TLSHandshakeStartTime = time.Now()
	fmt.Printf("Starting TLS handshake...\n")
}

func (h *RequestStatistics) tlsHandshakeDone(cs tls.ConnectionState, err error) {
	h.TLSHandshakeDoneTime = time.Now()
	fmt.Printf("TLS handshake done...\n")
}

func (h *RequestStatistics) gotFirstResponseByte() {
	h.GotFirstResponseByteTime = time.Now()
	fmt.Printf("Received first response byte...\n")
}
