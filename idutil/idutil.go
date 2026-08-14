package idutil

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"net"
	"sync"
	"time"
)

// ---------------------------------------------------------------------------
// UUID
// ---------------------------------------------------------------------------

// UUID generates a UUID v4 string (e.g. "550e8400-e29b-41d4-a716-446655440000").
func UUID() string {
	b := UUIDBytes()
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// UUIDBytes generates a UUID v4 as a 16-byte slice.
func UUIDBytes() []byte {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	// Set version (4) and variant (RFC 4122)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return b
}

// SimpleUUID generates a UUID v4 without hyphens (32 hex characters).
func SimpleUUID() string {
	b := UUIDBytes()
	dst := make([]byte, 32)
	hex.Encode(dst, b)
	return string(dst)
}

// ---------------------------------------------------------------------------
// Snowflake ID
// ---------------------------------------------------------------------------

const (
	snowflakeEpoch        int64 = 1288834974657 // Twitter epoch (2010-11-04 01:42:54 UTC)
	snowflakeMachineBits  int64 = 10
	snowflakeSequenceBits int64 = 12
	snowflakeMaxMachine   int64 = (1 << snowflakeMachineBits) - 1
	snowflakeMaxSequence  int64 = (1 << snowflakeSequenceBits) - 1
)

// SnowflakeID is a distributed unique ID generator inspired by Twitter Snowflake.
type SnowflakeID struct {
	mu        sync.Mutex
	epoch     int64
	machineID int64
	sequence  int64
	lastTime  int64
}

// NewSnowflake creates a new snowflake ID generator.
// machineID must be in range [0, 1023].
func NewSnowflake(machineID int64) *SnowflakeID {
	if machineID < 0 || machineID > snowflakeMaxMachine {
		panic(fmt.Sprintf("idutil: machineID must be in range [0, %d]", snowflakeMaxMachine))
	}
	return &SnowflakeID{
		epoch:     snowflakeEpoch,
		machineID: machineID,
	}
}

// NextID generates the next snowflake ID.
func (s *SnowflakeID) NextID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli() - s.epoch
	if now < 0 {
		return 0, errors.New("idutil: clock moved backwards")
	}

	if now == s.lastTime {
		s.sequence = (s.sequence + 1) & snowflakeMaxSequence
		if s.sequence == 0 {
			// Sequence exhausted for this millisecond, wait
			for now <= s.lastTime {
				now = time.Now().UnixMilli() - s.epoch
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTime = now

	id := (now << (snowflakeMachineBits + snowflakeSequenceBits)) |
		(s.machineID << snowflakeSequenceBits) |
		s.sequence
	return id, nil
}

// ---------------------------------------------------------------------------
// ULID
// ---------------------------------------------------------------------------

// ulidEncoding is Crockford's Base32 used by ULID.
var ulidEncoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// ULID generates a Universally Unique Lexicographically Sortable Identifier.
// Format: 26 characters, Crockford Base32 encoded.
func ULID() string {
	var b [16]byte

	// 48-bit timestamp (milliseconds since Unix epoch)
	ts := time.Now().UnixMilli()
	b[0] = byte(ts >> 40)
	b[1] = byte(ts >> 32)
	b[2] = byte(ts >> 24)
	b[3] = byte(ts >> 16)
	b[4] = byte(ts >> 8)
	b[5] = byte(ts)

	// 80-bit randomness
	_, _ = rand.Read(b[6:])

	// Encode as 26-character Crockford Base32
	var out [26]byte
	// Time component (10 chars)
	out[0] = ulidEncoding[(b[0]&0xE0)>>5]
	out[1] = ulidEncoding[b[0]&0x1F]
	out[2] = ulidEncoding[(b[1]&0xF8)>>3]
	out[3] = ulidEncoding[((b[1]&0x07)<<2)|((b[2]&0xC0)>>6)]
	out[4] = ulidEncoding[(b[2]&0x3E)>>1]
	out[5] = ulidEncoding[((b[2]&0x01)<<4)|((b[3]&0xF0)>>4)]
	out[6] = ulidEncoding[((b[3]&0x0F)<<1)|((b[4]&0x80)>>7)]
	out[7] = ulidEncoding[(b[4]&0x7C)>>2]
	out[8] = ulidEncoding[((b[4]&0x03)<<3)|((b[5]&0xE0)>>5)]
	out[9] = ulidEncoding[b[5]&0x1F]
	// Random component (16 chars)
	out[10] = ulidEncoding[(b[6]&0xF8)>>3]
	out[11] = ulidEncoding[((b[6]&0x07)<<2)|((b[7]&0xC0)>>6)]
	out[12] = ulidEncoding[(b[7]&0x3E)>>1]
	out[13] = ulidEncoding[((b[7]&0x01)<<4)|((b[8]&0xF0)>>4)]
	out[14] = ulidEncoding[((b[8]&0x0F)<<1)|((b[9]&0x80)>>7)]
	out[15] = ulidEncoding[(b[9]&0x7C)>>2]
	out[16] = ulidEncoding[((b[9]&0x03)<<3)|((b[10]&0xE0)>>5)]
	out[17] = ulidEncoding[b[10]&0x1F]
	out[18] = ulidEncoding[(b[11]&0xF8)>>3]
	out[19] = ulidEncoding[((b[11]&0x07)<<2)|((b[12]&0xC0)>>6)]
	out[20] = ulidEncoding[(b[12]&0x3E)>>1]
	out[21] = ulidEncoding[((b[12]&0x01)<<4)|((b[13]&0xF0)>>4)]
	out[22] = ulidEncoding[((b[13]&0x0F)<<1)|((b[14]&0x80)>>7)]
	out[23] = ulidEncoding[(b[14]&0x7C)>>2]
	out[24] = ulidEncoding[((b[14]&0x03)<<3)|((b[15]&0xE0)>>5)]
	out[25] = ulidEncoding[b[15]&0x1F]

	return string(out[:])
}

// ---------------------------------------------------------------------------
// NanoID
// ---------------------------------------------------------------------------

const nanoIDAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-"

// NanoID generates a NanoID with the default size of 21 characters.
func NanoID() string {
	return NanoIDWithSize(21)
}

// NanoIDWithSize generates a NanoID with the specified size.
func NanoIDWithSize(size int) string {
	if size <= 0 {
		return ""
	}
	return RandomIDWithAlphabet(size, nanoIDAlphabet)
}

// ---------------------------------------------------------------------------
// Short / Random ID
// ---------------------------------------------------------------------------

// ShortID generates an 8-character short ID (alphanumeric).
func ShortID() string {
	return RandomIDWithAlphabet(8, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
}

// RandomID generates a random alphanumeric ID of the specified length.
func RandomID(length int) string {
	return RandomIDWithAlphabet(length, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
}

// RandomIDWithAlphabet generates a random ID using the given alphabet string.
func RandomIDWithAlphabet(length int, alphabet string) string {
	if length <= 0 || len(alphabet) == 0 {
		return ""
	}
	b := make([]byte, length)
	alphLen := big.NewInt(int64(len(alphabet)))
	for i := range b {
		n, _ := rand.Int(rand.Reader, alphLen)
		b[i] = alphabet[n.Int64()]
	}
	return string(b)
}

// ---------------------------------------------------------------------------
// Sequence
// ---------------------------------------------------------------------------

// Sequence is a thread-safe auto-incrementing sequence generator.
type Sequence struct {
	mu       sync.Mutex
	current  int64
	step     int64
	minValue int64
	maxValue int64
	cycled   bool
}

// NewSequence creates a new sequence starting at start, incrementing by step.
// The sequence has no upper bound and will not cycle.
func NewSequence(start, step int64) *Sequence {
	if step <= 0 {
		panic("idutil: sequence step must be positive")
	}
	return &Sequence{
		current:  start,
		step:     step,
		minValue: start,
		maxValue: 0, // 0 means unbounded
		cycled:   false,
	}
}

// NewSequenceWithRange creates a new sequence with bounds.
// start is the initial value, end is the upper bound (inclusive), step is the increment.
// If cycled is true, the sequence wraps around to start when end is exceeded.
func NewSequenceWithRange(start, end, step int64, cycled bool) *Sequence {
	if step <= 0 {
		panic("idutil: sequence step must be positive")
	}
	if end < start {
		panic("idutil: end must be >= start")
	}
	return &Sequence{
		current:  start,
		step:     step,
		minValue: start,
		maxValue: end,
		cycled:   cycled,
	}
}

// Next returns the next value in the sequence.
func (s *Sequence) Next() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	val := s.current
	s.current += s.step

	if s.maxValue > 0 && s.current > s.maxValue {
		if s.cycled {
			s.current = s.minValue
		} else {
			s.current = s.maxValue // clamp
		}
	}
	return val
}

// NextN returns the next n values in the sequence.
func (s *Sequence) NextN(n int) []int64 {
	if n <= 0 {
		return nil
	}
	result := make([]int64, n)
	for i := 0; i < n; i++ {
		result[i] = s.Next()
	}
	return result
}

// Reset resets the sequence to its initial value.
func (s *Sequence) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.current = s.minValue
}

// Current returns the current value without advancing the sequence.
func (s *Sequence) Current() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.current
}

// ---------------------------------------------------------------------------
// ObjectID (MongoDB-style)
// ---------------------------------------------------------------------------

var (
	objectIDCounter     uint32
	objectIDCounterOnce sync.Once
	objectIDMachine     [3]byte
	objectIDProcess     uint16
)

func initObjectID() {
	objectIDCounterOnce.Do(func() {
		// Machine identifier: 3 bytes from MAC address or random
		if mac := getMAC(); mac != nil {
			copy(objectIDMachine[:], hashMAC(mac))
		} else {
			_, _ = rand.Read(objectIDMachine[:])
		}
		// Process ID
		objectIDProcess = uint16(time.Now().UnixNano() & 0xFFFF)
		// Random counter start
		n, _ := rand.Int(rand.Reader, big.NewInt(0xFFFFFF))
		objectIDCounter = uint32(n.Int64())
	})
}

func getMAC() net.HardwareAddr {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	for _, iface := range ifaces {
		if len(iface.HardwareAddr) >= 6 {
			return iface.HardwareAddr
		}
	}
	return nil
}

func hashMAC(mac net.HardwareAddr) []byte {
	// Simple hash of MAC to 3 bytes
	h := make([]byte, 3)
	copy(h, mac[:3])
	for i := 3; i < len(mac); i++ {
		h[i%3] ^= mac[i]
	}
	return h
}

// ObjectID generates a MongoDB-style ObjectID (24 hex characters).
// Format: 4-byte timestamp + 3-byte machine + 2-byte process + 3-byte counter.
func ObjectID() string {
	initObjectID()

	var b [12]byte
	// Timestamp (seconds since Unix epoch)
	ts := uint32(time.Now().Unix())
	binary.BigEndian.PutUint32(b[0:4], ts)
	// Machine
	copy(b[4:7], objectIDMachine[:])
	// Process
	binary.BigEndian.PutUint16(b[7:9], objectIDProcess)
	// Counter (3 bytes, big-endian)
	c := objectIDCounter
	objectIDCounter = (objectIDCounter + 1) & 0xFFFFFF
	b[9] = byte(c >> 16)
	b[10] = byte(c >> 8)
	b[11] = byte(c)

	dst := make([]byte, 24)
	hex.Encode(dst, b[:])
	return string(dst)
}

// ---------------------------------------------------------------------------
// Machine / Worker helpers
// ---------------------------------------------------------------------------

// MachineID returns a machine identifier based on the MAC address.
func MachineID() (string, error) {
	mac := getMAC()
	if mac == nil {
		return "", errors.New("idutil: no network interface found")
	}
	return hex.EncodeToString(hashMAC(mac)), nil
}

// WorkerID generates a worker ID derived from the first non-loopback IP address.
// The IP bytes are summed modulo 1024 to produce an ID suitable for Snowflake.
func WorkerID() (int64, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return 0, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			sum := int64(ip[0]) + int64(ip[1]) + int64(ip[2]) + int64(ip[3])
			return sum % 1024, nil
		}
	}
	return 0, errors.New("idutil: no suitable network interface found")
}
