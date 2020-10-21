// +build !wasm

package libp2p

import (
	noise "github.com/libp2p/go-libp2p-noise"
	tls "github.com/libp2p/go-libp2p-tls"
	tcp "github.com/libp2p/go-tcp-transport"
	ws "github.com/libp2p/go-ws-transport"
	multiaddr "github.com/multiformats/go-multiaddr"
)

// DefaultSecurity is the default security option.
//
// Useful when you want to extend, but not replace, the supported transport
// security protocols.
var DefaultSecurity = ChainOptions(
	Security(noise.ID, noise.New),
	Security(tls.ID, tls.New),
)

// DefaultTransports are the default libp2p transports.
//
// Use this option when you want to *extend* the set of transports used by
// libp2p instead of replacing them.
var DefaultTransports = ChainOptions(
	Transport(tcp.NewTCPTransport),
	Transport(ws.New),
)

// DefaultListenAddrs configures libp2p to use default listen address.
var DefaultListenAddrs = func(cfg *Config) error {
	defaultIP4ListenAddr, err := multiaddr.NewMultiaddr("/ip4/0.0.0.0/tcp/0")
	if err != nil {
		return err
	}

	defaultIP6ListenAddr, err := multiaddr.NewMultiaddr("/ip6/::/tcp/0")
	if err != nil {
		return err
	}
	return cfg.Apply(ListenAddrs(
		defaultIP4ListenAddr,
		defaultIP6ListenAddr,
	))
}