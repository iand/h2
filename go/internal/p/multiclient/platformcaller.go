package multiclient

import (
	"github.com/hailocab/protobuf/proto"

	"github.com/hailocab/h2/go/internal/p/client"
	"github.com/hailocab/h2/go/internal/p/errors"
)

// PlatformCaller is the default caller and makes requests via the platform layer
// RPC mechanism (eg: RabbitMQ)
func PlatformCaller() Caller {
	return func(req *client.Request, rsp proto.Message) errors.Error {
		return client.Req(req, rsp)
	}
}
