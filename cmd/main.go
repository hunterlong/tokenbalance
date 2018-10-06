// Package main inside the cmd folder contains TokenBalance binary file for your operating system
package main

import (
	"fmt"
	tb "github.com/hunterlong/tokenbalance"
	"github.com/mkideal/cli"
	"os"
)

var (
	configs *tb.Config
)

func main() {
	err := runCli()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runCli() error {
	return cli.Root(rootCommand,
		cli.Tree(help),
		cli.Tree(startCommand),
		cli.Tree(versionCommand),
	).Run(os.Args[1:])
}

type commandArgs struct {
	cli.Helper
	Geth string `cli:"*g,geth" usage:"attach geth IPC or HTTP location" dft:"https://eth.coinapp.io"`
	IP   string `cli:"ip" usage:"Bind to IP Address" dft:"0.0.0.0"`
	Port int    `cli:"p,port" usage:"HTTP server port for token information in JSON" dft:"8080"`
	Logs bool   `cli:"l,logs" usage:"Output logs on each transaction" dft:"true"`
}

var help = cli.HelpCommand("display help information")

type rootT struct {
	cli.Helper
}

var rootCommand = &cli.Command{
	Name: "tokenbalance",
	Desc: "\n      #######################\n" +
		"           tokenBalance\n" +
		"      #######################\n\n" +
		"tokenBalance is an easy to use server that \n" +
		"give you your ERC20 token balance without \n" +
		"any troubles. Connects to your local geth \n" +
		"IPC and prints out a simple JSON response \n" +
		"for ethereum token balances.",
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		ctx.String("To start the tokenbalance server, use command:\ntokenbalance start --geth \"/rootCommand/ethereum/geth.ipc\" --port 8080 --ip 0.0.0.0\n * replace geth location with your own *\n")
		return nil
	},
}

var startCommand = &cli.Command{
	Name: "start",
	Desc: "run the tokenbalance http server",
	Argv: func() interface{} { return new(commandArgs) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*commandArgs)
		configs = &tb.Config{
			GethLocation: argv.Geth,
			Logs:         argv.Logs,
		}
		err := configs.Connect()
		if err != nil {
			return err
		}
		return startServer(argv.IP, argv.Port)
	},
}

var versionCommand = &cli.Command{
	Name: "version",
	Desc: "get the version of tokenbalance server",
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		ctx.String(tb.VERSION + "\n")
		return nil
	},
}

type errorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
