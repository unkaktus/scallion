// main.go - proxy 443 tcp port to onion service.
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to this module of scallion, using the creative
// commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package main

import (
	"context"
	"flag"
	"net"
	"os"
	"os/exec"

	"github.com/google/tcpproxy"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/proxy"
)

func ToTor(addr string) *tcpproxy.DialProxy {
	torDialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
	if err != nil {
		log.Fatal().Err(err).Msg("create tor dialer")
	}
	return &tcpproxy.DialProxy{
		Addr: addr,
		DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
			return torDialer.Dial(network, address)
		},
	}
}

func main() {
	log.Printf("started")

	var addr = flag.String("addr", "", "onion address")
	flag.Parse()

	if *addr == "" {
		log.Fatal().Msg("address is not specified")
	}

	// Start tor
	cmd := exec.Command("tor", "--runasdaemon", "0", "-f", "/torrc")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("start tor")
	}
	log.Info().Msg("started tor")

	var proxy tcpproxy.Proxy

	proxy.AddRoute(":443", ToTor(*addr+":443"))
	err = proxy.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("proxy exited")
	}
}
