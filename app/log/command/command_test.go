package command_test

import (
	"context"
	"testing"

	"github.com/nokaka/Xray-core/app/dispatcher"
	"github.com/nokaka/Xray-core/app/log"
	. "github.com/nokaka/Xray-core/app/log/command"
	"github.com/nokaka/Xray-core/app/proxyman"
	_ "github.com/nokaka/Xray-core/app/proxyman/inbound"
	_ "github.com/nokaka/Xray-core/app/proxyman/outbound"
	"github.com/nokaka/Xray-core/common"
	"github.com/nokaka/Xray-core/common/serial"
	"github.com/nokaka/Xray-core/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
