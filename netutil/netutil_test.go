package netutil

import (
	"net"
	"strings"
	"testing"
)

func TestIsInnerIP(t *testing.T) {
	inner := []string{
		"10.0.0.1", "10.255.255.255",
		"172.16.0.0", "172.31.255.255",
		"192.168.1.1", "192.168.255.255",
		"127.0.0.1", "127.255.255.255",
		"169.254.1.1", // link-local
		"0.0.0.0",
		"::1",              // loopback v6
		"fc00::1",          // unique local
		"fd12:3456::1",     // unique local
		"fe80::1",          // link-local v6
		"::",               // unspecified
	}
	outer := []string{
		"8.8.8.8", "1.1.1.1", "114.114.114.114",
		"172.32.0.1", "172.15.255.255", // just outside private 172 range
		"11.0.0.1", "192.169.1.1", "128.0.0.1",
		"2606:4700:4700::1111", // public v6
		"2001:4860:4860::8888",
	}

	for _, ip := range inner {
		if !IsInnerIP(ip) {
			t.Errorf("IsInnerIP(%q) = false, want true", ip)
		}
	}
	for _, ip := range outer {
		if IsInnerIP(ip) {
			t.Errorf("IsInnerIP(%q) = true, want false", ip)
		}
	}
}

func TestIsInnerIPInvalid(t *testing.T) {
	if IsInnerIP("not-an-ip") {
		t.Error("IsInnerIP(invalid) should be false")
	}
	if IsInnerIP("") {
		t.Error("IsInnerIP(empty) should be false")
	}
	if IsInnerIP("999.999.999.999") {
		t.Error("IsInnerIP(out of range) should be false")
	}
}

func TestGetLocalhost(t *testing.T) {
	ip := GetLocalhost()
	if ip == nil || ip.IsUnspecified() {
		t.Fatalf("GetLocalhost() returned nil/unspecified: %v", ip)
	}
}

func TestGetLocalhostStr(t *testing.T) {
	s := GetLocalhostStr()
	if s == "" {
		t.Fatal("GetLocalhostStr() returned empty")
	}
	// returned value should be a valid IP string
	ip := net.ParseIP(s)
	if ip == nil {
		t.Errorf("GetLocalhostStr() = %q is not a valid IP", s)
	}
	if IsInnerIP(s) && s != "127.0.0.1" && s != "::1" {
		t.Logf("local IP %q is an internal address (expected on most machines)", s)
	}
}

func TestGetLocalIPs(t *testing.T) {
	ips := GetLocalIPs()
	if len(ips) == 0 {
		t.Fatal("GetLocalIPs() returned no IPs")
	}
	for _, ip := range ips {
		if ip == nil {
			t.Error("GetLocalIPs() returned nil entry")
			continue
		}
		if ip.IsUnspecified() {
			t.Errorf("GetLocalIPs() returned unspecified IP %v", ip)
		}
	}
}

func TestHideIpPart(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"192.168.1.100", "192.168.*.*"},
		{"10.0.0.1", "10.0.*.*"},
		{"1.2.3.4", "1.2.*.*"},
	}
	for _, tt := range tests {
		got := HideIpPart(tt.in)
		if got != tt.want {
			t.Errorf("HideIpPart(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestHideIpPartInvalid(t *testing.T) {
	if got := HideIpPart("not-an-ip"); got != "" {
		t.Errorf("HideIpPart(invalid) = %q, want empty", got)
	}
	if got := HideIpPart(""); got != "" {
		t.Errorf("HideIpPart(empty) = %q, want empty", got)
	}
}

func TestIsIPv4(t *testing.T) {
	if !IsIPv4("192.168.1.1") {
		t.Error("IsIPv4(192.168.1.1) should be true")
	}
	if IsIPv4("2001:db8::1") {
		t.Error("IsIPv4(ipv6) should be false")
	}
	if IsIPv4("999.1.1.1") {
		t.Error("IsIPv4(out of range) should be false")
	}
}

func TestIsIPv6(t *testing.T) {
	if !IsIPv6("2001:db8::1") {
		t.Error("IsIPv6(2001:db8::1) should be true")
	}
	if !IsIPv6("::1") {
		t.Error("IsIPv6(::1) should be true")
	}
	if IsIPv6("192.168.1.1") {
		t.Error("IsIPv6(ipv4) should be false")
	}
}

func TestGetHostname(t *testing.T) {
	h := GetHostname()
	if h == "" {
		t.Fatal("GetHostname() returned empty")
	}
	// should usually contain letters/digits/dashes/dots
	for _, r := range h {
		if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '-' || r == '.' || r == '_') {
			t.Errorf("GetHostname() = %q contains unexpected char %q", h, r)
			break
		}
	}
}

func TestIsValidPort(t *testing.T) {
	if !IsValidPort(80) || !IsValidPort(1) || !IsValidPort(65535) {
		t.Error("valid ports rejected")
	}
	if IsValidPort(0) || IsValidPort(65536) || IsValidPort(-1) {
		t.Error("invalid ports accepted")
	}
}

func TestIsReachable(t *testing.T) {
	// 8.8.8.8:53 is almost always reachable; but tests must not depend on the network.
	// Only assert that the function returns without error on invalid inputs.
	if _, err := IsReachable("", 1*1000); err == nil {
		t.Error("IsReachable(empty) should error")
	}
}

func TestParseIP(t *testing.T) {
	if ip := ParseIP("192.168.1.1"); ip == nil || ip.String() != "192.168.1.1" {
		t.Errorf("ParseIP(v4) = %v", ip)
	}
	if ip := ParseIP("2001:db8::1"); ip == nil || ip.String() != "2001:db8::1" {
		t.Errorf("ParseIP(v6) = %v", ip)
	}
	if ip := ParseIP("nope"); ip != nil {
		t.Errorf("ParseIP(invalid) = %v, want nil", ip)
	}
}

func TestLocalIPContainsNonLoopback(t *testing.T) {
	// On typical machines the primary NIC IP is non-loopback.
	found := false
	for _, ip := range GetLocalIPs() {
		if strings.HasPrefix(ip.String(), "127.") || ip.String() == "::1" {
			continue
		}
		found = true
		break
	}
	if !found {
		t.Log("No non-loopback IP found — machine may be offline or headless")
	}
}