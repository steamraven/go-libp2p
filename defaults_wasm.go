// +build wasm

package libp2p

// This file contains all the default configuration options.

import (
	noise "github.com/libp2p/go-libp2p-noise"
	ws "github.com/libp2p/go-ws-transport"
)

// DefaultSecurity is the default security option.
//
// Useful when you want to extend, but not replace, the supported transport
// security protocols.
var DefaultSecurity = ChainOptions(
	Security(noise.ID, noise.New),
)

// DefaultTransports are the default libp2p transports.
//
// Use this option when you want to *extend* the set of transports used by
// libp2p instead of replacing them.
var DefaultTransports = ChainOptions(
	Transport(ws.New),
)

// DefaultListenAddrs configures libp2p to use default listen address.
var DefaultListenAddrs = func(cfg *Config) error {
	return cfg.Apply(ListenAddrs(
	))
}