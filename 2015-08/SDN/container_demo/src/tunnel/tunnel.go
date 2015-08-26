package tunnel

import (
	"os/exec"
	"strings"
	"iface"
	"log"
)

/**
 * Add tunnel named @name from @local -> @remote. Mode is one of 'ipip' (for IP-in-IP) or 'gre'
 * NOTE: shelling out here, since tunnel commands are based on ioctls with nested structs.
 *       It is not the purpose of this presentation to document such programming.
 */
func Add(name, mode, local, remote string) {
	if mode != "ipip" && mode != "gre" {
		log.Fatalf("unsupported tunnel %s mode %q", name, mode)
	}
	/* If tunnel already exists, tear it down, since the mode may have changed */
	if iface.Exists(name) {
		log.Printf("tunnel %q exists already", name)
		Del(name)
	}
	/* Address may contain CIDR netmask, hence need to parse */
	local  = extract_ip(local)
	remote = extract_ip(remote)

	log.Printf("adding %s tunnel %s: %s -> %s", mode, name, local, remote)
	err := exec.Command("ip", "tunnel", "add", name,
	                    "mode", mode, "local", local, "remote", remote).Run()
	if err != nil {
		log.Fatalf("failed to add tunnel %s %s -> %s: %s", name, local, remote, err)
	}
}

/* Shell out here, too. iface.Del() only works for GRE tunnels */
func Del(name string) {
	log.Printf("removing tunnel %s", name)
	err := exec.Command("ip", "tunnel", "del", name).Run()
	if err != nil {
		log.Fatalf("failed to remove tunnel: %s", err)
	}
}

/* Helper routine to extract dotted-quad IP address from "a.b.c.d" or "a.b.c.d/mask-length" */
func extract_ip(ip_or_cidr string) string {
	idx := strings.Index(ip_or_cidr, "/")
	if idx > -1 {
		return ip_or_cidr[:idx]
	}
	return ip_or_cidr
}
