// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	//"enode://1e236bf17d8994a690b9bc3feba8f265927eff30c7056ae5f84aa8f818758895f4eec594fd42038340d5e9bbfb8cc1465cf0a1f94315ab0fc26b14369c2ac184@127.0.0.1:32668",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	"enode://924543a43d18bc5759a8bdcd17fa9c7c35df63968e9333640b80b58dab94b17a012371c9d46bed10ce7508a607cac76828ca04685893958eee44ade83b856dc2@47.242.237.63:32668",
	"enode://ebad898d980b520ef6adb54ffb6a68117686e7332f1ea01f7551b7a296a34dd945445a078d7cad019d864c5ef0e0b7f2b5777d94f93adf7dc59f798af72609ac@47.242.235.121:32668",
}

var V5Bootnodes = []string{}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
