package iface

import (
	"github.com/vishvananda/netlink"
	"strings"
	"log"
)

/* Set interface named @iface up (if not already running) */
func SetUp(iface string) {
	dev, err := netlink.LinkByName(iface)
	if err != nil {
		log.Fatalf("failed to look up interface %s: %s", iface, err)
	}
	log.Printf("bringing up interface %s", iface)
	if err := netlink.LinkSetUp(dev); err != nil {
		log.Fatalf("failed to set up interface %s: %s", iface, err)
	}
}

/* Configure @iface with address @cidr. Deletes any existing addresses */
func SetAddr(iface string, cidr string) {
	if strings.Index(cidr, "/") == -1 {
		/* No scope given: assume full netmask */
		cidr = cidr + "/32"
	}
	addr, err := netlink.ParseAddr(cidr)
	if err != nil {
		log.Fatalf("failed to parse %s interface address %s: %s", iface, cidr, err)
	}
	link, err := netlink.LinkByName(iface)
	if err != nil {
		log.Fatalf("failed to look up interface %s: %s", iface, err)
	}

	addrs, err := netlink.AddrList(link, 0)
	if err != nil {
		log.Fatalf("failed to get %s interface addresses: %s", iface, err)
	}

	already_configured := false
	for _, old_addr := range addrs {
		/* Test if already configured, otherwise EEXISTS */
		if old_addr.Equal(*addr) {
			already_configured = true
		} else {
			log.Printf("removing %s from %s", old_addr.IPNet, iface)
			if err := netlink.AddrDel(link, &old_addr); err != nil {
				log.Fatalf("failed to remove address from %s: %s", iface, err)
			}
		}
	}

	log.Printf("configuring %s on %s", addr.IPNet, iface)
	if ! already_configured {
		if err := netlink.AddrAdd(link, addr); err != nil {
			log.Fatalf("failed to add address to %s: %s", iface, err)
		}
	}
}

/* Return true if @iface name exists */
func Exists(iface string) bool {
	_, err := netlink.LinkByName(iface)
	return err == nil
}

/* Delete interface */
func Del(iface string) {
	dev, err := netlink.LinkByName(iface)
	if err != nil {
		log.Fatalf("failed to look up interface %s: %s", iface, err)
	}
	log.Printf("removing interface %s", iface)
	if err := netlink.LinkDel(dev); err != nil {
		log.Fatalf("failed to remove interface %s: %s", iface, err)
	}
}
