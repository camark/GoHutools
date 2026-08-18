package netutil

import (
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// privateIPBlocks returns the IPv4 and IPv6 CIDR blocks considered "internal".
func privateIPBlocks() []*net.IPNet {
	cidrs := []string{
		"0.0.0.0/8",       // "this" network
		"10.0.0.0/8",      // private
		"100.64.0.0/10",   // shared address space (CGNAT)
		"127.0.0.0/8",     // loopback
		"169.254.0.0/16",  // link-local
		"172.16.0.0/12",   // private
		"192.0.0.0/24",    // IETF protocol assignments
		"192.168.0.0/16",  // private
		"198.18.0.0/15",   // benchmarking
		"224.0.0.0/4",     // multicast
		"240.0.0.0/4",     // reserved
		"::1/128",         // loopback v6
		"fc00::/7",        // unique local v6
		"fe80::/10",       // link-local v6
		"ff00::/8",        // multicast v6
		"::/128",          // unspecified v6
	}
	blocks := make([]*net.IPNet, 0, len(cidrs))
	for _, c := range cidrs {
		if _, n, err := net.ParseCIDR(c); err == nil {
			blocks = append(blocks, n)
		}
	}
	return blocks
}

// IsInnerIP reports whether ip is a private/internal address.
func IsInnerIP(ip string) bool {
	parsed := net.ParseIP(strings.TrimSpace(ip))
	if parsed == nil {
		return false
	}
	for _, block := range privateIPBlocks() {
		if block.Contains(parsed) {
			return true
		}
	}
	return false
}

// GetLocalhost returns a non-loopback IP of the local machine.
// It prefers the first IPv4 address of the first up-and-running interface.
func GetLocalhost() net.IP {
	ips := GetLocalIPs()
	for _, ip := range ips {
		if ip4 := ip.To4(); ip4 != nil {
			return ip4
		}
	}
	if len(ips) > 0 {
		return ips[0]
	}
	// fall back to loopback
	return net.ParseIP("127.0.0.1")
}

// GetLocalhostStr returns the local IP address as a string.
func GetLocalhostStr() string {
	ip := GetLocalhost()
	if ip == nil {
		return ""
	}
	return ip.String()
}

// GetLocalIPs returns all non-loopback, non-unspecified IPs of local interfaces.
func GetLocalIPs() []net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	ips := make([]net.IP, 0, len(addrs))
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		default:
			continue
		}
		if ip == nil || ip.IsLoopback() || ip.IsUnspecified() {
			continue
		}
		ips = append(ips, ip)
	}
	return ips
}

// HideIpPart masks an IPv4 address, keeping only the first two octets visible,
// e.g. "192.168.1.100" -> "192.168.*.*". Returns empty string for invalid input.
func HideIpPart(ip string) string {
	parsed := net.ParseIP(strings.TrimSpace(ip))
	if parsed == nil {
		return ""
	}
	ip4 := parsed.To4()
	if ip4 == nil {
		return ""
	}
	return strconv.Itoa(int(ip4[0])) + "." + strconv.Itoa(int(ip4[1])) + ".*.*"
}

// IsIPv4 reports whether s is a valid IPv4 address.
func IsIPv4(s string) bool {
	ip := net.ParseIP(strings.TrimSpace(s))
	return ip != nil && ip.To4() != nil
}

// IsIPv6 reports whether s is a valid IPv6 address.
func IsIPv6(s string) bool {
	ip := net.ParseIP(strings.TrimSpace(s))
	return ip != nil && ip.To4() == nil
}

// ParseIP parses s as an IP address, returning nil when invalid.
func ParseIP(s string) net.IP {
	return net.ParseIP(strings.TrimSpace(s))
}

// GetHostname returns the OS hostname (strips trailing dot if present).
func GetHostname() string {
	h, err := os.Hostname()
	if err != nil {
		return ""
	}
	return strings.TrimSuffix(h, ".")
}

// IsValidPort reports whether p is a valid TCP/UDP port (1-65535).
func IsValidPort(p int) bool {
	return p >= 1 && p <= 65535
}

// IsReachable checks whether host:port is reachable within timeoutMilliseconds.
func IsReachable(host string, timeoutMilliseconds int64) (bool, error) {
	if host == "" {
		return false, &net.AddrError{Err: "empty host", Addr: ""}
	}
	timeout := time.Duration(timeoutMilliseconds) * time.Millisecond
	conn, err := net.DialTimeout("tcp", host, timeout)
	if err != nil {
		return false, err
	}
	_ = conn.Close()
	return true, nil
}

// IsReachableURL checks whether the given URL is reachable via HTTP HEAD request
// within timeoutMilliseconds.
func IsReachableURL(url string, timeoutMilliseconds int64) (bool, error) {
	client := &http.Client{Timeout: time.Duration(timeoutMilliseconds) * time.Millisecond}
	resp, err := client.Head(url)
	if err != nil {
		return false, err
	}
	_ = resp.Body.Close()
	return true, nil
}

// GetPort checks whether a port is in use by attempting to listen on it.
// Returns true when the port is free.
func GetPort(port int) bool {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return false
	}
	_ = ln.Close()
	return true
}

