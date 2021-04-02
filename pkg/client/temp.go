package client

import (
	"github.com/antoninbas/p4runtime-go-client/pkg/p4info"
)

func (c *Client) GetP4Info() *p4info.P4Info {
	return p4info.GetNewP4Info(c.p4Info)
}
