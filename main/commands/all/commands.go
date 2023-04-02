package all

import (
	"github.com/nokaka/Xray-core/main/commands/all/api"
	"github.com/nokaka/Xray-core/main/commands/all/tls"
	"github.com/nokaka/Xray-core/main/commands/base"
)

// go:generate go run github.com/nokaka/Xray-core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
	)
}
