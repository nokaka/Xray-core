package tcp

import (
	"github.com/nokaka/Xray-core/common"
	"github.com/nokaka/Xray-core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
