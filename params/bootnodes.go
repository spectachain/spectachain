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
	"enode://63be1492a791d253b5de3491367cf9b1de4efa223f097e38668d5868ce6583430ac372ae55bc9362b82385ef78a552bf2a566d5b78aefb16bec4f9a6223124a5@139.180.132.5:31288",
	"enode://37a0c1a99b46d1ebb01ee3efa600489d95e4b7c7c567ad7cf337ac853a802722bc1bc53af41998baf50c8e6addcd0eb7c0aa7681e2d3641896aa1b6e40d122b1@45.32.120.23:31288",
	"enode://444a96b12917c96e07c09830a00a9e17e269785815e0577bd43a534d7abce1697aad429befeac06402f00e053535649795528baec47840549edd84f8222316c9@45.77.39.166:31288",
	"enode://4bbea063c01dd15d75e3ce5d80084b5a75b1ca276708134b5fc16a73bdfb7203385bd49007886cfeafb3483a9bbe9b6ec00df458bbacb0b639721ee357857f4d@45.76.187.54:31288",
	"enode://a2529943818dba5c017d72178b4199599ee3e2953cb27c22cf77e6d745f1a8c94700bdf38c04fecb86681cc3a576b416a6b2f83cfbf5ec7bfe7a811d2a2d1abc@45.76.158.98:31288",
	"enode://3d5a57f1e940457ec9fb9de26a99620ab0e1d99ff310e89f68ecd614510e2eb7471192f3410aaa545e6efe823cdd57cc4d73d2edeafc7d336b7d92702769adf9@45.77.174.174:31288",
	"enode://40facdba9dcbe5e1c3409eac741eff72f0037b3385173c9bb9880e722f4baca38dbf58dc5ea0b3049f8de85831c98968cc505d5bf847c956695160de4e349df1@66.42.56.76:31288",
	"enode://4f419492740f38dc3ac24060b64f23e539e67319fea7ef21a5662872bda9a40e557e817bcc69e209cb1774087b3245cd033ae347fdd5cfa52fe61f46750ad5e5@139.180.186.151:31288",
	"enode://8027a5c74822b948ec169330913cf2f695715af4a041cc0e839d102da8e9646598d0ea087c973fbbdd6671d05b8076ba7bda5d7db78fb37f03c179266b042de6@149.28.136.78:31288",
}
var TestnetBootnodes = []string{

	"enode://9e9eacf00fe30afe145f7f034eaff0ef230c8722c83bd4e49cab2ac720b1099b1bc957294cc8442797a9d6a95263e225a18b7b3308ef2ead8b0e45f7298ab7a5@34.128.125.238:31288",
	"enode://2e7063b628a102428c4b274ad12e71239fc43708a860e0ffc55174ff5d84e4cedafebd91eeae086668c558f407d276e219a39b0bcefdbe9f5d1337e6de7f287d@34.101.51.196:31288",
	"enode://eb871c92bee71f9a82864940c337dda71befd49ff9048009ad65eda36b2b98975fe9414275c85aa9154abc244f8c3a4568c05d30d98accacb4d8c007aa0f61f3@34.128.81.110:31288",
	"enode://837db97395a7e742033dc78fb29b8d9620cbe5cc48b7bae492d70c7e2efa2d2e4b647dbe5f56d68bde0f7e19fd0d9b9c1fa09cc507a4e0ba10d26a72e95576f6@34.101.54.29:31288",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
