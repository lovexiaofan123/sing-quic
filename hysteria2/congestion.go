package hysteria2

import "github.com/sagernet/quic-go"

var SetCongestionController = func(quicConn quic.Connection, cc string, cwnd int) {
	// do nothing
	// clash.meta will replace this function after init
}
