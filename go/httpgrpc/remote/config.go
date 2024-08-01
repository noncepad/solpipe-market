package remote

import (
	"errors"

	"github.com/noncepad/solpipe-market/go/httpgrpc"
)

type Configuration struct {
	Remote *httpgrpc.Address `json:"remote"`
}

func (c *Configuration) Check() error {
	if c.Remote == nil {
		return errors.New("blank remote")
	}
	return nil
}
