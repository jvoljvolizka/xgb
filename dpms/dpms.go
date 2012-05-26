// Package dpms is the X client API for the DPMS extension.
package dpms

/*
	This file was generated by dpms.xml on May 26 2012 6:23:12pm EDT.
	This file is automatically generated. Edit at your peril!
*/

import (
	"github.com/BurntSushi/xgb"

	"github.com/BurntSushi/xgb/xproto"
)

// Init must be called before using the DPMS extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 4, "DPMS").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named DPMS could be found on on the server.")
	}

	xgb.ExtLock.Lock()
	c.Extensions["DPMS"] = reply.MajorOpcode
	for evNum, fun := range xgb.NewExtEventFuncs["DPMS"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["DPMS"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	xgb.ExtLock.Unlock()

	return nil
}

func init() {
	xgb.NewExtEventFuncs["DPMS"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["DPMS"] = make(map[int]xgb.NewErrorFun)
}

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Card32'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

const (
	DPMSModeOn      = 0
	DPMSModeStandby = 1
	DPMSModeSuspend = 2
	DPMSModeOff     = 3
)

// GetVersionCookie is a cookie used only for GetVersion requests.
type GetVersionCookie struct {
	*xgb.Cookie
}

// GetVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling GetVersionCookie.Reply()
func GetVersion(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) GetVersionCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'GetVersion' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(getVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return GetVersionCookie{cookie}
}

// GetVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func GetVersionUnchecked(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) GetVersionCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'GetVersion' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(getVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return GetVersionCookie{cookie}
}

// GetVersionReply represents the data returned from a GetVersion request.
type GetVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ServerMajorVersion uint16
	ServerMinorVersion uint16
}

// Reply blocks and returns the reply data for a GetVersion request.
func (cook GetVersionCookie) Reply() (*GetVersionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getVersionReply(buf), nil
}

// getVersionReply reads a byte slice into a GetVersionReply value.
func getVersionReply(buf []byte) *GetVersionReply {
	v := new(GetVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.ServerMajorVersion = xgb.Get16(buf[b:])
	b += 2

	v.ServerMinorVersion = xgb.Get16(buf[b:])
	b += 2

	return v
}

// Write request to wire for GetVersion
// getVersionRequest writes a GetVersion request to a byte slice.
func getVersionRequest(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put16(buf[b:], ClientMajorVersion)
	b += 2

	xgb.Put16(buf[b:], ClientMinorVersion)
	b += 2

	return buf
}

// CapableCookie is a cookie used only for Capable requests.
type CapableCookie struct {
	*xgb.Cookie
}

// Capable sends a checked request.
// If an error occurs, it will be returned with the reply by calling CapableCookie.Reply()
func Capable(c *xgb.Conn) CapableCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'Capable' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(capableRequest(c), cookie)
	return CapableCookie{cookie}
}

// CapableUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func CapableUnchecked(c *xgb.Conn) CapableCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'Capable' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(capableRequest(c), cookie)
	return CapableCookie{cookie}
}

// CapableReply represents the data returned from a Capable request.
type CapableReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Capable bool
	// padding: 23 bytes
}

// Reply blocks and returns the reply data for a Capable request.
func (cook CapableCookie) Reply() (*CapableReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return capableReply(buf), nil
}

// capableReply reads a byte slice into a CapableReply value.
func capableReply(buf []byte) *CapableReply {
	v := new(CapableReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	if buf[b] == 1 {
		v.Capable = true
	} else {
		v.Capable = false
	}
	b += 1

	b += 23 // padding

	return v
}

// Write request to wire for Capable
// capableRequest writes a Capable request to a byte slice.
func capableRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetTimeoutsCookie is a cookie used only for GetTimeouts requests.
type GetTimeoutsCookie struct {
	*xgb.Cookie
}

// GetTimeouts sends a checked request.
// If an error occurs, it will be returned with the reply by calling GetTimeoutsCookie.Reply()
func GetTimeouts(c *xgb.Conn) GetTimeoutsCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'GetTimeouts' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(getTimeoutsRequest(c), cookie)
	return GetTimeoutsCookie{cookie}
}

// GetTimeoutsUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func GetTimeoutsUnchecked(c *xgb.Conn) GetTimeoutsCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'GetTimeouts' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(getTimeoutsRequest(c), cookie)
	return GetTimeoutsCookie{cookie}
}

// GetTimeoutsReply represents the data returned from a GetTimeouts request.
type GetTimeoutsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	StandbyTimeout uint16
	SuspendTimeout uint16
	OffTimeout     uint16
	// padding: 18 bytes
}

// Reply blocks and returns the reply data for a GetTimeouts request.
func (cook GetTimeoutsCookie) Reply() (*GetTimeoutsReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getTimeoutsReply(buf), nil
}

// getTimeoutsReply reads a byte slice into a GetTimeoutsReply value.
func getTimeoutsReply(buf []byte) *GetTimeoutsReply {
	v := new(GetTimeoutsReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.StandbyTimeout = xgb.Get16(buf[b:])
	b += 2

	v.SuspendTimeout = xgb.Get16(buf[b:])
	b += 2

	v.OffTimeout = xgb.Get16(buf[b:])
	b += 2

	b += 18 // padding

	return v
}

// Write request to wire for GetTimeouts
// getTimeoutsRequest writes a GetTimeouts request to a byte slice.
func getTimeoutsRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// SetTimeoutsCookie is a cookie used only for SetTimeouts requests.
type SetTimeoutsCookie struct {
	*xgb.Cookie
}

// SetTimeouts sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func SetTimeouts(c *xgb.Conn, StandbyTimeout uint16, SuspendTimeout uint16, OffTimeout uint16) SetTimeoutsCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'SetTimeouts' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(setTimeoutsRequest(c, StandbyTimeout, SuspendTimeout, OffTimeout), cookie)
	return SetTimeoutsCookie{cookie}
}

// SetTimeoutsChecked sends a checked request.
// If an error occurs, it can be retrieved using SetTimeoutsCookie.Check()
func SetTimeoutsChecked(c *xgb.Conn, StandbyTimeout uint16, SuspendTimeout uint16, OffTimeout uint16) SetTimeoutsCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'SetTimeouts' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(setTimeoutsRequest(c, StandbyTimeout, SuspendTimeout, OffTimeout), cookie)
	return SetTimeoutsCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook SetTimeoutsCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for SetTimeouts
// setTimeoutsRequest writes a SetTimeouts request to a byte slice.
func setTimeoutsRequest(c *xgb.Conn, StandbyTimeout uint16, SuspendTimeout uint16, OffTimeout uint16) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put16(buf[b:], StandbyTimeout)
	b += 2

	xgb.Put16(buf[b:], SuspendTimeout)
	b += 2

	xgb.Put16(buf[b:], OffTimeout)
	b += 2

	return buf
}

// EnableCookie is a cookie used only for Enable requests.
type EnableCookie struct {
	*xgb.Cookie
}

// Enable sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func Enable(c *xgb.Conn) EnableCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'Enable' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(enableRequest(c), cookie)
	return EnableCookie{cookie}
}

// EnableChecked sends a checked request.
// If an error occurs, it can be retrieved using EnableCookie.Check()
func EnableChecked(c *xgb.Conn) EnableCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'Enable' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(enableRequest(c), cookie)
	return EnableCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook EnableCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Enable
// enableRequest writes a Enable request to a byte slice.
func enableRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// DisableCookie is a cookie used only for Disable requests.
type DisableCookie struct {
	*xgb.Cookie
}

// Disable sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func Disable(c *xgb.Conn) DisableCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'Disable' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(disableRequest(c), cookie)
	return DisableCookie{cookie}
}

// DisableChecked sends a checked request.
// If an error occurs, it can be retrieved using DisableCookie.Check()
func DisableChecked(c *xgb.Conn) DisableCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'Disable' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(disableRequest(c), cookie)
	return DisableCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook DisableCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Disable
// disableRequest writes a Disable request to a byte slice.
func disableRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// ForceLevelCookie is a cookie used only for ForceLevel requests.
type ForceLevelCookie struct {
	*xgb.Cookie
}

// ForceLevel sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func ForceLevel(c *xgb.Conn, PowerLevel uint16) ForceLevelCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'ForceLevel' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(forceLevelRequest(c, PowerLevel), cookie)
	return ForceLevelCookie{cookie}
}

// ForceLevelChecked sends a checked request.
// If an error occurs, it can be retrieved using ForceLevelCookie.Check()
func ForceLevelChecked(c *xgb.Conn, PowerLevel uint16) ForceLevelCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'ForceLevel' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(forceLevelRequest(c, PowerLevel), cookie)
	return ForceLevelCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook ForceLevelCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for ForceLevel
// forceLevelRequest writes a ForceLevel request to a byte slice.
func forceLevelRequest(c *xgb.Conn, PowerLevel uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put16(buf[b:], PowerLevel)
	b += 2

	return buf
}

// InfoCookie is a cookie used only for Info requests.
type InfoCookie struct {
	*xgb.Cookie
}

// Info sends a checked request.
// If an error occurs, it will be returned with the reply by calling InfoCookie.Reply()
func Info(c *xgb.Conn) InfoCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'Info' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(infoRequest(c), cookie)
	return InfoCookie{cookie}
}

// InfoUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func InfoUnchecked(c *xgb.Conn) InfoCookie {
	if _, ok := c.Extensions["DPMS"]; !ok {
		panic("Cannot issue request 'Info' using the uninitialized extension 'DPMS'. dpms.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(infoRequest(c), cookie)
	return InfoCookie{cookie}
}

// InfoReply represents the data returned from a Info request.
type InfoReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	PowerLevel uint16
	State      bool
	// padding: 21 bytes
}

// Reply blocks and returns the reply data for a Info request.
func (cook InfoCookie) Reply() (*InfoReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return infoReply(buf), nil
}

// infoReply reads a byte slice into a InfoReply value.
func infoReply(buf []byte) *InfoReply {
	v := new(InfoReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.PowerLevel = xgb.Get16(buf[b:])
	b += 2

	if buf[b] == 1 {
		v.State = true
	} else {
		v.State = false
	}
	b += 1

	b += 21 // padding

	return v
}

// Write request to wire for Info
// infoRequest writes a Info request to a byte slice.
func infoRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DPMS"]
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}
