package main

import (
	"context"
	"os"

	"github.com/elliottpolk/super/config"
	"github.com/elliottpolk/super/grpc"
	"github.com/elliottpolk/super/rest"

	log "github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v2"
	altsrc "gopkg.in/urfave/cli.v2/altsrc"
)

var (
	RpcPortFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "rpc-port",
		Value:   "7000",
		Usage:   "RPC port to listen on",
		EnvVars: []string{"DOPE_RPC_PORT"},
	})

	HttpPortFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "http-port",
		Value:   "8080",
		Usage:   "HTTP port to listen on",
		EnvVars: []string{"DOPE_HTTP_PORT"},
	})

	HttpsPortFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "tls-port",
		Value:   "8443",
		Usage:   "HTTPS port to listen on",
		EnvVars: []string{"DOPE_HTTPS_PORT"},
	})

	TlsCertFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "tls-cert",
		Usage:   "TLS certificate file for HTTPS",
		EnvVars: []string{"DOPE_TLS_CERT"},
	})

	TlsKeyFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "tls-key",
		Usage:   "TLS key file for HTTPS",
		EnvVars: []string{"DOPE_TLS_KEY"},
	})

	DatastoreAddrFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "datastore-addr",
		Aliases: []string{"ds-addr", "dsa"},
		Usage:   "Database address",
	})

	DatastorePortFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "datastore-port",
		Aliases: []string{"ds-port", "dsp"},
		Value:   "27017",
		Usage:   "Database port",
	})

	DatastoreNameFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "datastore-name",
		Aliases: []string{"ds-name", "dsn"},
		Value:   "super",
		Usage:   "Database name",
	})

	DatastoreUserFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "datastore-user",
		Aliases: []string{"ds-user", "dsu"},
		Usage:   "Database user",
	})

	DatastorePasswordFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "datastore-password",
		Aliases: []string{"ds-password", "dspwd"},
		Usage:   "Database password",
	})

	CfgFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "config",
		Aliases: []string{"c", "cfg", "confg"},
		Usage:   "optional path to config file",
	})

	flags = []cli.Flag{
		CfgFlag,
		RpcPortFlag,
		HttpPortFlag,
		HttpsPortFlag,
		TlsCertFlag,
		TlsKeyFlag,
		DatastoreAddrFlag,
		DatastorePortFlag,
		DatastoreNameFlag,
		DatastoreUserFlag,
		DatastorePasswordFlag,
	}
)

func main() {
	app := cli.App{
		Name:  "dope",
		Flags: flags,
		Before: func(ctx *cli.Context) error {
			if len(ctx.String(CfgFlag.Name)) > 0 {
				return altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config"))(ctx)
			}
			return nil
		},
		Action: func(ctx *cli.Context) error {
			// read in the configuration
			comp := &config.Composition{
				Server: &config.ServerCfg{
					RpcPort:   ctx.String(RpcPortFlag.Name),
					HttpPort:  ctx.String(HttpPortFlag.Name),
					HttpsPort: ctx.String(HttpsPortFlag.Name),
					TlsCert:   ctx.String(TlsCertFlag.Name),
					TlsKey:    ctx.String(TlsKeyFlag.Name),
				},
				Db: &config.DbCfg{
					Addr:     ctx.String(DatastoreAddrFlag.Name),
					Port:     ctx.String(DatastorePortFlag.Name),
					DbName:   ctx.String(DatastoreNameFlag.Name),
					User:     ctx.String(DatastoreUserFlag.Name),
					Password: ctx.String(DatastorePasswordFlag.Name),
				},
			}

			// run in a non-blocking goroutine since it is blocking
			go func() {
				if err := rest.Serve(context.Background(), comp); err != nil {
					log.Fatal(err)
				}
			}()

			// use this one to block and prevent exiting
			if err := grpc.Serve(context.Background(), comp); err != nil {
				return cli.Exit(err, 1)
			}

			return nil
		},
	}

	app.Run(os.Args)
}
