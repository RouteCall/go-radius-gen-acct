package main

import (
	"./cdr"
	"./rfc2866"
	"context"
	//"fmt"
	"github.com/urfave/cli"
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
	"log"
	"net"
	"os"
	"sync/atomic"
	"time"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

type Config struct {
	NASPort      int
	NASIPAddress string
	Server       string
	Port         string
	Key          string
	PPS          int
	MaxReq       int
	ShowCount    bool
}

func ParseCdrAttributes(p *radius.Packet, c *cdr.CdrValues, cfg Config) {
	rfc2866.SipAcctStatusType_Add(p, rfc2866.SipAcctStatusType_Value_Stop)
	rfc2866.SipServiceType_Add(p, rfc2866.SipServiceType_Value_SipSession)
	rfc2866.SipResponseCode_AddString(p, c.ResponseCode)
	rfc2866.SipMethod_Add(p, rfc2866.SipMethod_Value_INVITE)
	rfc2866.SipEventTimestamp_Add(p, c.EventTimestamp)
	rfc2866.SipFromTag_AddString(p, c.FromTag)
	rfc2866.SipToTag_AddString(p, c.ToTag)
	rfc2866.SipCallerID_AddString(p, c.CallerId)
	rfc2866.SipCalleeID_AddString(p, c.CalleeId)
	rfc2866.SipDstNumber_AddString(p, c.DstNumber)
	rfc2866.SipAcctSessionID_AddString(p, c.AcctSessionId)
	rfc2866.SipCallMSDuration_Add(p, rfc2866.SipCallMSDuration(c.MsDuration))
	rfc2866.SipCallSetuptime_Add(p, rfc2866.SipCallSetuptime(c.SetupTime))
	rfc2865.NASPort_Add(p, rfc2865.NASPort(cfg.NASPort))
	rfc2865.NASIPAddress_Add(p, net.ParseIP(cfg.NASIPAddress))
	return
}

func SendAcct(c *cdr.CdrValues, cfg Config) {
	packet := radius.New(radius.CodeAccountingRequest, []byte(cfg.Key))
	ParseCdrAttributes(packet, c, cfg)
	_, err := radius.Exchange(context.Background(), packet, cfg.Server+":"+cfg.Port)
	if err != nil {
		os.Exit(1)
	}
}

func CliConfig() Config {
	cfg := Config{}
	cfg.CliCreate()

	return cfg
}

func (cfg *Config) CliCreate() {
	parsed := false
	app := cli.NewApp()
	app.Usage = "A Go (golang) RADIUS client accounting (RFC 2866) implementation for perfomance testing"
	app.Version = "0.10.1"
	app.Compiled = time.Now()

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "pps, p",
			Value:       10,
			Usage:       "packets per second",
			Destination: &cfg.PPS,
		},
		cli.StringFlag{
			Name:        "server, s",
			Usage:       "server to send accts",
			Destination: &cfg.Server,
		},
		cli.StringFlag{
			Name:        "port, P",
			Value:       "1813",
			Usage:       "port to send accts",
			Destination: &cfg.Port,
		},
		cli.StringFlag{
			Name:        "nas-ip",
			Value:       "127.0.0.1",
			Usage:       "NAS-IP-Address on radius packet",
			Destination: &cfg.NASIPAddress,
		},
		cli.IntFlag{
			Name:        "nas-port",
			Value:       5666,
			Usage:       "NAS-Port on radius packet",
			Destination: &cfg.NASPort,
		},
		cli.StringFlag{
			Name:        "key, k",
			Usage:       "key for acct",
			Destination: &cfg.Key,
		},
		cli.IntFlag{
			Name:        "max-req, m",
			Value:       MaxInt,
			Usage:       "stop the test and exit when max-req are reached",
			Destination: &cfg.MaxReq,
		},
		cli.BoolFlag{
			Name:  "c",
			Usage: "show count of requests",
		},
	}

	app.Action = func(c *cli.Context) error {
		if cfg.PPS <= 0 {
			return cli.NewExitError("pps must be greater 0", 1)
		}
		if len(cfg.Server) <= 0 {
			return cli.NewExitError("server not defined", 1)
		}
		if len(cfg.Key) <= 0 {
			return cli.NewExitError("key not defined", 1)
		}
		if c.Bool("c") {
			cfg.ShowCount = true
		}
		parsed = true
		return nil
	}

	err := app.Run(os.Args)
	if err != nil || parsed == false {
		os.Exit(1)
	}
}

func main() {
	cfg := CliConfig()
	sleep := 1000 / cfg.PPS
	var count uint64

	for i := 0; i < cfg.MaxReq; i++ {
		go func() {
			if cfg.ShowCount {
				atomic.AddUint64(&count, 1)
			}
			c := cdr.FillCdr()
			go SendAcct(c, cfg)
		}()
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		if cfg.ShowCount {
			log.Print("total count accounting-request: ", atomic.LoadUint64(&count))
		}
	}

}
