package client

import (
	"golang.org/x/net/context"

	"github.com/keybase/cli"
	"github.com/keybase/client/go/libcmdline"
	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	rpc "github.com/keybase/go-framed-msgpack-rpc"
)

func NewCmdTestPassphrase(cl *libcmdline.CommandLine, g *libkb.GlobalContext) cli.Command {
	return cli.Command{
		Name:  "test-passphrase",
		Usage: "Test the GetPassphrase protocol",
		Action: func(c *cli.Context) {
			cl.ChooseCommand(&CmdTestPassphrase{Contextified: libkb.NewContextified(g)}, "test-passphrase", c)
		},
	}
}

type CmdTestPassphrase struct {
	libkb.Contextified
}

func (s *CmdTestPassphrase) ParseArgv(ctx *cli.Context) error {
	return nil
}

func (s *CmdTestPassphrase) GetUsage() libkb.Usage {
	return libkb.Usage{}
}

func (s *CmdTestPassphrase) Run() (err error) {
	defer libkb.Trace(s.G().Log, "CmdTestPassphrase.Run", func() error { return err })()

	protocols := []rpc.Protocol{
		NewSecretUIProtocol(s.G()),
	}
	if err = RegisterProtocols(protocols); err != nil {
		return err
	}

	cli, err := GetAccountClient(s.G())
	if err != nil {
		return err
	}

	arg := keybase1.PassphrasePromptArg{
		GuiArg: libkb.DefaultPassphraseArg(s.G(), s.G().SecretStoreAll != nil),
	}
	res, err := cli.PassphrasePrompt(context.TODO(), arg)
	if err != nil {
		return err
	}
	res.Passphrase = "[passphrase redacted]"
	s.G().Log.Debug("passphrase prompt result: %+v", res)
	return nil
}
