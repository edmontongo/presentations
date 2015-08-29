/*
 * Set-up for Edmonton Go netlink presentation
 */
package main

import (
	"github.com/Unknwon/goconfig"
	"iface"
	"bridge"
	"tunnel"
	"route"
	"os/exec"
	"log"
	"os"
)


func main() {
	if len(os.Args) < 2 {
		log.Fatalf("need a 'configuration file' argument")
	}
	log.Printf("loading configuration from %s", os.Args[1])
	conf, err := goconfig.LoadConfigFile(os.Args[1])
	if err != nil {
		log.Fatalf("failed to read configuration: %s", err)
	}

	// 1. Configure the bridge to be used by docker
	//    NB: this requires to add the following to /etc/default/docker:
	//        DOCKER_OPTS="-b=bridge0"y
	bridge_name := conf.MustValue("docker", "bridge")
	bridge.Add(bridge_name)
	iface.SetAddr(bridge_name, conf.MustValue("docker", "address"))
	iface.SetUp(bridge_name)

	// 2. Restart docker to pick up the changes
	log.Println("restarting docker")
	if err := exec.Command("service", "docker",  "restart").Run(); err != nil {
		log.Fatalf("failed tor restart docker: %s", err)
	}

	// 3. Configure outgoing interface. Prefer wireless network if configured
	out_iface := conf.MustValue("network", "interface")
	wifi, err := conf.GetSection("wifi")
	if len(wifi) == 3 {
		out_iface = wifi["interface"]
		log.Printf("setting up %s to join %s", wifi["interface"], wifi["essid"])
		if err := exec.Command("iwconfig", wifi["interface"], "mode", wifi["mode"],
			                          "essid", wifi["essid"]).Run(); err != nil {
			log.Fatalf("failed to set up wireless interface: %s", err) 
		}
	}
	src_addr  := conf.MustValue("network", "address")
	iface.SetAddr(out_iface, src_addr)
	iface.SetUp(out_iface)

	// 4. Add tunnel that links the docker-to-docker interfaces
	tun_name := conf.MustValue("tunnel", "name")

	/* The local address of the tunnel is the outgoing network interface */
	tunnel.Add(tun_name, conf.MustValue("tunnel",  "mode"),
		   src_addr, conf.MustValue("tunnel",  "remote"))

	/* Tunnel endpoint talks to the bridge via the tunnel.address */
	iface.SetAddr(tun_name, conf.MustValue("tunnel", "address"))
	iface.SetUp(tun_name)
	
	/* Set up the encapsulated route to tunnel.encap via tun_name, using src_addr */
	route.Add(src_addr, conf.MustValue("tunnel", "encap"), tun_name, "")

	route.List(bridge_name, out_iface, tun_name)
}
