package ipfsx

import (
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr-net"
	"log"
	"os/exec"
	"strings"
)

func GetPeers() []string  {

	peers, err := exec.Command("ipfs", "swarm", "peers").Output()
	if err != nil {
		log.Fatal(err)
	}

	peersStrs := strings.Split(string(peers), "\n")

	var ips []string

	for _, p := range peersStrs {
		if p == "" {
			continue
		}
		multiaddr, err := ma.NewMultiaddr(p)
		if err != nil {
			log.Fatal(err)
			return nil
		}

		ip, err := manet.ToIP(multiaddr)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		ips = append(ips, ip.String())
	}

	return ips
}