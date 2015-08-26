package route

import (
	"github.com/vishvananda/netlink"
	"strings"
	"log"
	"net"
)

/* List routes belonging to interface name(s) @ifs */
func List(ifs ...string) {
	var ifmap = make(map[int]string)
	for _, iface := range ifs {
		link, err := netlink.LinkByName(iface)
		if err != nil {
			log.Fatalf("failed to look up interface %s: %s", iface, err)
		}
		ifmap[link.Attrs().Index] = iface
	}

	routes, err := netlink.RouteList(nil, netlink.FAMILY_V4)
	if err != nil {
		log.Fatalf("failed to list routes: %s", err)
	}
	log.Printf("routes through %s:", strings.Join(ifs, ", "))
        for _, rt := range routes {
		/*
                 *  FIXME: netlink.RouteList() has a bug - the 'link' argument has no effect.
                 *         Hence filtering manually here; maybe open an issue on github.
                 */
		if dev, ok := ifmap[rt.LinkIndex]; ok {
			if rt.Src != nil {
				log.Printf("%-10s %s -> %s\n", dev + ":", rt.Src, rt.Dst)
			} else {
				log.Printf("%-10s %s\n", dev + ":", rt.Dst)
			}
		}
        }
}

/* Check if route @src -> @dst exists with the same destination netmask as @r */
func Exists(r *netlink.Route) bool {
	/* FIXME: see bug of RouteList mentioned in above List() command */
	routes, err := netlink.RouteList(nil, 0)
	if err != nil {
		log.Fatalf("failed to list routes: %s", err)
	}
        for _, rt := range routes {
		if rt.Src.Equal(r.Src) && rt.Dst.IP.Equal(r.Dst.IP) {
			sizea, _ := rt.Dst.Mask.Size()
			sizeb, _ := r.Dst.Mask.Size()
			if sizea == sizeb {
				return true
			}
		}
        }
	return false
}

/**
 * Add a route @src -> @dst through interface @dev.
 * The gateway @gw is optional and may be an emtpy string.
 */
func Add(src, dst, dev, gw string) {
	route := mkRoute(src, dst, dev, gw)
	/* Adding an existing route leads to error */
	if Exists(route) {
		log.Printf("route %s -> %s exists already, not adding it", src, dst)
		return
	}
	err := netlink.RouteAdd(route)
        if err != nil {
		log.Fatalf("failed to add route %s -> %s: %s", src, dst, err)
        }
}

/* Helper routine to generate a netlink.Route object from string descriptions */
func mkRoute(src, dst, dev, gw string) *netlink.Route {
	_, dst_net, err := net.ParseCIDR(dst)
	if err != nil {
		log.Fatalf("failed to parse route destination %s: %s", dst, err)
	}

	route := netlink.Route{ Src: net.ParseIP(extract_ip(src)), Dst: dst_net }

	if dev != "" {
		link, err := netlink.LinkByName(dev)
		if err != nil {
			log.Fatalf("failed to look up outgoing interface %s: %s", dev, err)
		}
		route.LinkIndex = link.Attrs().Index
	}

	if gw != "" {
		route.Gw = net.ParseIP(gw)
	}
	return &route
}

/* Helper routine to extract dotted-quad IP address from "a.b.c.d" or "a.b.c.d/mask-length" */
func extract_ip(ip_or_cidr string) string {
	idx := strings.Index(ip_or_cidr, "/")
	if idx > -1 {
		return ip_or_cidr[:idx]
	}
	return ip_or_cidr
}
