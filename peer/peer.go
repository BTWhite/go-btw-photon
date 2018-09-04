// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package peer

import (
	"fmt"
	"net"
)

var lp = "Peer:"

type Peer struct {
	Ip   net.IP `json:"ip"`
	Port int    `json:"port"`
}

func NewPeer(ip net.IP, port int) Peer {
	return Peer{
		Ip:   ip,
		Port: port,
	}
}

func (p *Peer) HttpAddr() string {
	return fmt.Sprintf("http://%s:%d/jsonrpc/", p.Ip.String(), p.Port)
}
