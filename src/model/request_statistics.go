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
	ReadResponseDoneTime     time.Time

	ClientTrace *httptrace.ClientTrace
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
}

func (h *RequestStatistics) dnsDone(dnsDoneInfo httptrace.DNSDoneInfo) {
	h.DNSDoneTime = time.Now()
}

func (h *RequestStatistics) connectStart(network, addr string) {
	h.ConnectStartTime = time.Now()
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
}

func (h *RequestStatistics) tlsHandshakeDone(cs tls.ConnectionState, err error) {
	h.TLSHandshakeDoneTime = time.Now()
}

func (h *RequestStatistics) gotFirstResponseByte() {
	h.GotFirstResponseByteTime = time.Now()
}
