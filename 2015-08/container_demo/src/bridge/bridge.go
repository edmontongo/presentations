package bridge

import (
	"github.com/vishvananda/netlink"
	"iface"
	"log"
)

func Add(name string) {
	if iface.Exists(name) {
		log.Printf("reusing existing bridge %s", name)
	} else {
		log.Printf("adding bridge %s", name)

		br := netlink.NewLinkAttrs()
		br.Name = name
		if err := netlink.LinkAdd(&netlink.Bridge{br}); err != nil {
			log.Fatalf("failed to add bridge %s: %s", name, err)
		}
	}
}

func Del(name string) {
	br, err := netlink.LinkByName(name)
	if err != nil {
		log.Printf("not removing bridge %s: %s", name, err)
	}
	log.Printf("removing bridge %s", name)
	if err := netlink.LinkDel(&netlink.Bridge{*br.Attrs()}); err != nil {
		log.Fatalf("failed to remove bridge %s: %s", name, err)
	}
}
