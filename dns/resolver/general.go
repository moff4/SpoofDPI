package client

import (
	"context"
	"fmt"
	"net"

	"github.com/miekg/dns"
)

type DNSResult struct {
	msg *dns.Msg
	err error
}

type GeneralResolver struct {
	client *dns.Client
	server string
}

func NewGeneralClient(server string) *GeneralResolver {
	return &GeneralResolver{
		client: &dns.Client{},
		server: server,
	}
}

func (r *GeneralResolver) Resolve(ctx context.Context, host string, qTypes []uint16) ([]net.IPAddr, error) {
	sendMsg := func(ctx context.Context, msg *dns.Msg) (*dns.Msg, error) {
		resp, _, err := r.client.Exchange(msg, r.server)
		return resp, err
	}

	resultCh := lookup(ctx, host, qTypes, sendMsg)
	addrs, err := processResults(ctx, resultCh)
	return addrs, err
}

func (c *GeneralResolver) String() string {
	return fmt.Sprintf("general resolver(%s)", c.server)
}
