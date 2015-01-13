package fbp

// This is a basic flow based programming package which implements Information
// Packets (IPs) and connections (using Go's channels). The IPs are by default
// data agnostic, and can store any type of data by casting the values to an
// `interface{}`. This package also provides convenience functions for creating
// IPs from the basic Go types, and back from IPs to the basic types (with some
// general error checking).
//
// You can access the raw interface{} values through the structure itself, or
// through the `Void()` method on the IP type.
//
// We also implement list begin and end packets, as well as a ConnEnd packet,
// which should be used to indicate that no further values will be coming
// across a particular connection.
//
// With the provided channels and IP objects, it is very simple to implement
// the [Go Pipeline](http://blog.golang.org/pipelines) concurrency patern and
// create a complete Flow Based Programming graph.

// Represents a single Information Packet
type IP struct {
	ConnEnd   bool
	ListStart bool
	ListEnd   bool
	Value     interface{}
}

// Represents a FBP connection which can be used to send and receive IPs. Note:
// This is implemented as a `chan *IP`, which lets us use existing go channel
// semantics for sending and receiving IPs.
type Connection chan *IP

// Creates a new FBP connection. This is represented by a Go channel of IP
// pointers.
func NewConnection() Connection {
	ipchan := make(Connection, 5)
	return ipchan
}

// Create a new, empty Information Packet
func NewIP() *IP {
	ip := new(IP)
	ip.ConnEnd = false
	ip.ListStart = false
	ip.ListEnd = false
	return ip
}

// Creates and returns a list start IP
func NewIPListStart() *IP {
	ip := NewIP()
	ip.ListStart = true
	return ip
}

// Returns true if the current IP is a list start IP.
func (ip *IP) IsListStart() bool {
	return ip.ListStart
}

// Creates and returns a list end IP
func NewIPListEnd() *IP {
	ip := NewIP()
	ip.ListEnd = true
	return ip
}

// Returns true if the current IP is a list end IP.
func (ip *IP) IsListEnd() bool {
	return ip.ListEnd
}

// Create a new IP which indicates a connection end.
func NewIPConnEnd() *IP {
	ip := NewIP()
	ip.ConnEnd = true
	return ip
}

// Returns true if the current IP indicates the end of a connection
func (ip *IP) IsConnEnd() bool {
	return ip.ConnEnd
}

// Helper function for generating an IP from a byte
func NewIPByte(val byte) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving a byte from an IP
func (ip *IP) Byte() (val byte) {
	val = ip.Value.(byte)
	return
}

// Helper function for generating an IP from a complex128
func NewIPComplex128(val complex128) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving a complex128 from an IP
func (ip *IP) Complex128() (val complex128) {
	val = ip.Value.(complex128)
	return
}

// Helper function for generating an IP from a complex64
func NewIPComplex64(val complex64) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving a complex64 from an IP
func (ip *IP) Complex64() (val complex64) {
	val = ip.Value.(complex64)
	return
}

// Helper function for generating an IP from an error
func NewIPError(val error) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving an error from an IP
func (ip *IP) Error() (val error) {
	val = ip.Value.(error)
	return
}

// Helper function for generating an IP from a float32
func NewIPFloat32(val float32) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving a float32 from an IP
func (ip *IP) Float32() (val float32) {
	val = ip.Value.(float32)
	return
}

// Helper function for generating an IP from a float64
func NewIPFloat64(val float64) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving a float64 from an IP
func (ip *IP) Float64() (val float64) {
	val = ip.Value.(float64)
	return
}

// Helper function for generating an IP from an int
func NewIPInt(val int) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving an int from an IP
func (ip *IP) Int() (val int) {
	val = ip.Value.(int)
	return
}

// Helper function for generating an IP from an int16
func NewIPInt16(val int16) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving an int16 from an IP
func (ip *IP) Int16() (val int16) {
	val = ip.Value.(int16)
	return
}

// Helper function for generating an IP from an int32
func NewIPInt32(val int32) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving an int32 from an IP
func (ip *IP) Int32() (val int32) {
	val = ip.Value.(int32)
	return
}

// Helper function for generating an IP from an int64
func NewIPInt64(val int64) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving an int64 from an IP
func (ip *IP) Int64() (val int64) {
	val = ip.Value.(int64)
	return
}

// Helper function for generating an IP from an int8
func NewIPInt8(val int8) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving an int8 from an IP
func (ip *IP) Int8() (val int8) {
	val = ip.Value.(int8)
	return
}

// Helper function for generating an IP from a rune
func NewIPRune(val rune) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving a rune from an IP
func (ip *IP) Rune() (val rune) {
	val = ip.Value.(rune)
	return
}

// Helper function for generating an IP from a string
func NewIPString(val string) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving a string from an IP
func (ip *IP) String() (val string) {
	val = ip.Value.(string)
	return
}

// Helper function for generating an IP from an uint
func NewIPUint(val uint) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving an uint from an IP
func (ip *IP) Uint() (val uint) {
	val = ip.Value.(uint)
	return
}

// Helper function for generating an IP from an uint16
func NewIPUint16(val uint16) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving an uint16 from an IP
func (ip *IP) Uint16() (val uint16) {
	val = ip.Value.(uint16)
	return
}

// Helper function for generating an IP from an uint32
func NewIPUint32(val uint32) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving an uint32 from an IP
func (ip *IP) Uint32() (val uint32) {
	val = ip.Value.(uint32)
	return
}

// Helper function for generating an IP from an uint64
func NewIPUint64(val uint64) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving an uint64 from an IP
func (ip *IP) Uint64() (val uint64) {
	val = ip.Value.(uint64)
	return
}

// Helper function for generating an IP from an uint8
func NewIPUint8(val uint8) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving an uint8 from an IP
func (ip *IP) Uint8() (val uint8) {
	val = ip.Value.(uint8)
	return
}

// Helper function for generating an IP from an uintptr
func NewIPUintptr(val uintptr) *IP {
	ip := NewIP()
	ip.Value = interface{}(val)
	return ip
}

// Helper function for retrieving an uintptr from an IP
func (ip *IP) Uintptr() (val uintptr) {
	val = ip.Value.(uintptr)
	return
}

// Helper to create an IP from a interface{} object
func NewIPUVoid(val interface{}) *IP {
	ip := NewIP()
	ip.Value = val
	return ip
}

// Helper to retrieve an interface{} object from an IP
func (ip *IP) Void() (val interface{}) {
	val = ip.Value
	return
}
