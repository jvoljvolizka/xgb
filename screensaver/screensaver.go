// Package screensaver is the X client API for the MIT-SCREEN-SAVER extension.
package screensaver

/*
	This file was generated by screensaver.xml on May 26 2012 6:23:13pm EDT.
	This file is automatically generated. Edit at your peril!
*/

import (
	"github.com/BurntSushi/xgb"

	"github.com/BurntSushi/xgb/xproto"
)

// Init must be called before using the MIT-SCREEN-SAVER extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 16, "MIT-SCREEN-SAVER").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named MIT-SCREEN-SAVER could be found on on the server.")
	}

	xgb.ExtLock.Lock()
	c.Extensions["MIT-SCREEN-SAVER"] = reply.MajorOpcode
	for evNum, fun := range xgb.NewExtEventFuncs["MIT-SCREEN-SAVER"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["MIT-SCREEN-SAVER"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	xgb.ExtLock.Unlock()

	return nil
}

func init() {
	xgb.NewExtEventFuncs["MIT-SCREEN-SAVER"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["MIT-SCREEN-SAVER"] = make(map[int]xgb.NewErrorFun)
}

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Card32'

const (
	KindBlanked  = 0
	KindInternal = 1
	KindExternal = 2
)

const (
	EventNotifyMask = 1
	EventCycleMask  = 2
)

const (
	StateOff      = 0
	StateOn       = 1
	StateCycle    = 2
	StateDisabled = 3
)

// Notify is the event number for a NotifyEvent.
const Notify = 0

type NotifyEvent struct {
	Sequence uint16
	Code     byte
	State    byte
	// padding: 1 bytes
	SequenceNumber uint16
	Time           xproto.Timestamp
	Root           xproto.Window
	Window         xproto.Window
	Kind           byte
	Forced         bool
	// padding: 14 bytes
}

// NotifyEventNew constructs a NotifyEvent value that implements xgb.Event from a byte slice.
func NotifyEventNew(buf []byte) xgb.Event {
	v := NotifyEvent{}
	b := 1 // don't read event number

	v.Code = buf[b]
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.State = buf[b]
	b += 1

	b += 1 // padding

	v.SequenceNumber = xgb.Get16(buf[b:])
	b += 2

	v.Time = xproto.Timestamp(xgb.Get32(buf[b:]))
	b += 4

	v.Root = xproto.Window(xgb.Get32(buf[b:]))
	b += 4

	v.Window = xproto.Window(xgb.Get32(buf[b:]))
	b += 4

	v.Kind = buf[b]
	b += 1

	if buf[b] == 1 {
		v.Forced = true
	} else {
		v.Forced = false
	}
	b += 1

	b += 14 // padding

	return v
}

// Bytes writes a NotifyEvent value to a byte slice.
func (v NotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	buf[b] = v.Code
	b += 1

	b += 2 // skip sequence number

	buf[b] = v.State
	b += 1

	b += 1 // padding

	xgb.Put16(buf[b:], v.SequenceNumber)
	b += 2

	xgb.Put32(buf[b:], uint32(v.Time))
	b += 4

	xgb.Put32(buf[b:], uint32(v.Root))
	b += 4

	xgb.Put32(buf[b:], uint32(v.Window))
	b += 4

	buf[b] = v.Kind
	b += 1

	if v.Forced {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 14 // padding

	return buf
}

// SequenceId returns the sequence id attached to the Notify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v NotifyEvent) SequenceId() uint16 {
	return v.Sequence
}

// String is a rudimentary string representation of NotifyEvent.
func (v NotifyEvent) String() string {
	fieldVals := make([]string, 0, 10)
	fieldVals = append(fieldVals, xgb.Sprintf("Sequence: %d", v.Sequence))
	fieldVals = append(fieldVals, xgb.Sprintf("Code: %d", v.Code))
	fieldVals = append(fieldVals, xgb.Sprintf("State: %d", v.State))
	fieldVals = append(fieldVals, xgb.Sprintf("SequenceNumber: %d", v.SequenceNumber))
	fieldVals = append(fieldVals, xgb.Sprintf("Time: %d", v.Time))
	fieldVals = append(fieldVals, xgb.Sprintf("Root: %d", v.Root))
	fieldVals = append(fieldVals, xgb.Sprintf("Window: %d", v.Window))
	fieldVals = append(fieldVals, xgb.Sprintf("Kind: %d", v.Kind))
	fieldVals = append(fieldVals, xgb.Sprintf("Forced: %t", v.Forced))
	return "Notify {" + xgb.StringsJoin(fieldVals, ", ") + "}"
}

func init() {
	xgb.NewExtEventFuncs["MIT-SCREEN-SAVER"][0] = NotifyEventNew
}

// QueryVersionCookie is a cookie used only for QueryVersion requests.
type QueryVersionCookie struct {
	*xgb.Cookie
}

// QueryVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryVersionCookie.Reply()
func QueryVersion(c *xgb.Conn, ClientMajorVersion byte, ClientMinorVersion byte) QueryVersionCookie {
	if _, ok := c.Extensions["MIT-SCREEN-SAVER"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'MIT-SCREEN-SAVER'. screensaver.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryVersionUnchecked(c *xgb.Conn, ClientMajorVersion byte, ClientMinorVersion byte) QueryVersionCookie {
	if _, ok := c.Extensions["MIT-SCREEN-SAVER"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'MIT-SCREEN-SAVER'. screensaver.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ServerMajorVersion uint16
	ServerMinorVersion uint16
	// padding: 20 bytes
}

// Reply blocks and returns the reply data for a QueryVersion request.
func (cook QueryVersionCookie) Reply() (*QueryVersionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryVersionReply(buf), nil
}

// queryVersionReply reads a byte slice into a QueryVersionReply value.
func queryVersionReply(buf []byte) *QueryVersionReply {
	v := new(QueryVersionReply)
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

	b += 20 // padding

	return v
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(c *xgb.Conn, ClientMajorVersion byte, ClientMinorVersion byte) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["MIT-SCREEN-SAVER"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = ClientMajorVersion
	b += 1

	buf[b] = ClientMinorVersion
	b += 1

	b += 2 // padding

	return buf
}

// QueryInfoCookie is a cookie used only for QueryInfo requests.
type QueryInfoCookie struct {
	*xgb.Cookie
}

// QueryInfo sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryInfoCookie.Reply()
func QueryInfo(c *xgb.Conn, Drawable xproto.Drawable) QueryInfoCookie {
	if _, ok := c.Extensions["MIT-SCREEN-SAVER"]; !ok {
		panic("Cannot issue request 'QueryInfo' using the uninitialized extension 'MIT-SCREEN-SAVER'. screensaver.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryInfoRequest(c, Drawable), cookie)
	return QueryInfoCookie{cookie}
}

// QueryInfoUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryInfoUnchecked(c *xgb.Conn, Drawable xproto.Drawable) QueryInfoCookie {
	if _, ok := c.Extensions["MIT-SCREEN-SAVER"]; !ok {
		panic("Cannot issue request 'QueryInfo' using the uninitialized extension 'MIT-SCREEN-SAVER'. screensaver.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryInfoRequest(c, Drawable), cookie)
	return QueryInfoCookie{cookie}
}

// QueryInfoReply represents the data returned from a QueryInfo request.
type QueryInfoReply struct {
	Sequence         uint16 // sequence number of the request for this reply
	Length           uint32 // number of bytes in this reply
	State            byte
	SaverWindow      xproto.Window
	MsUntilServer    uint32
	MsSinceUserInput uint32
	EventMask        uint32
	Kind             byte
	// padding: 7 bytes
}

// Reply blocks and returns the reply data for a QueryInfo request.
func (cook QueryInfoCookie) Reply() (*QueryInfoReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryInfoReply(buf), nil
}

// queryInfoReply reads a byte slice into a QueryInfoReply value.
func queryInfoReply(buf []byte) *QueryInfoReply {
	v := new(QueryInfoReply)
	b := 1 // skip reply determinant

	v.State = buf[b]
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.SaverWindow = xproto.Window(xgb.Get32(buf[b:]))
	b += 4

	v.MsUntilServer = xgb.Get32(buf[b:])
	b += 4

	v.MsSinceUserInput = xgb.Get32(buf[b:])
	b += 4

	v.EventMask = xgb.Get32(buf[b:])
	b += 4

	v.Kind = buf[b]
	b += 1

	b += 7 // padding

	return v
}

// Write request to wire for QueryInfo
// queryInfoRequest writes a QueryInfo request to a byte slice.
func queryInfoRequest(c *xgb.Conn, Drawable xproto.Drawable) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["MIT-SCREEN-SAVER"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	return buf
}

// SelectInputCookie is a cookie used only for SelectInput requests.
type SelectInputCookie struct {
	*xgb.Cookie
}

// SelectInput sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func SelectInput(c *xgb.Conn, Drawable xproto.Drawable, EventMask uint32) SelectInputCookie {
	if _, ok := c.Extensions["MIT-SCREEN-SAVER"]; !ok {
		panic("Cannot issue request 'SelectInput' using the uninitialized extension 'MIT-SCREEN-SAVER'. screensaver.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(selectInputRequest(c, Drawable, EventMask), cookie)
	return SelectInputCookie{cookie}
}

// SelectInputChecked sends a checked request.
// If an error occurs, it can be retrieved using SelectInputCookie.Check()
func SelectInputChecked(c *xgb.Conn, Drawable xproto.Drawable, EventMask uint32) SelectInputCookie {
	if _, ok := c.Extensions["MIT-SCREEN-SAVER"]; !ok {
		panic("Cannot issue request 'SelectInput' using the uninitialized extension 'MIT-SCREEN-SAVER'. screensaver.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(selectInputRequest(c, Drawable, EventMask), cookie)
	return SelectInputCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook SelectInputCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for SelectInput
// selectInputRequest writes a SelectInput request to a byte slice.
func selectInputRequest(c *xgb.Conn, Drawable xproto.Drawable, EventMask uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["MIT-SCREEN-SAVER"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	xgb.Put32(buf[b:], EventMask)
	b += 4

	return buf
}

// SetAttributesCookie is a cookie used only for SetAttributes requests.
type SetAttributesCookie struct {
	*xgb.Cookie
}

// SetAttributes sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func SetAttributes(c *xgb.Conn, Drawable xproto.Drawable, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class byte, Depth byte, Visual xproto.Visualid, ValueMask uint32, ValueList []uint32) SetAttributesCookie {
	if _, ok := c.Extensions["MIT-SCREEN-SAVER"]; !ok {
		panic("Cannot issue request 'SetAttributes' using the uninitialized extension 'MIT-SCREEN-SAVER'. screensaver.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(setAttributesRequest(c, Drawable, X, Y, Width, Height, BorderWidth, Class, Depth, Visual, ValueMask, ValueList), cookie)
	return SetAttributesCookie{cookie}
}

// SetAttributesChecked sends a checked request.
// If an error occurs, it can be retrieved using SetAttributesCookie.Check()
func SetAttributesChecked(c *xgb.Conn, Drawable xproto.Drawable, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class byte, Depth byte, Visual xproto.Visualid, ValueMask uint32, ValueList []uint32) SetAttributesCookie {
	if _, ok := c.Extensions["MIT-SCREEN-SAVER"]; !ok {
		panic("Cannot issue request 'SetAttributes' using the uninitialized extension 'MIT-SCREEN-SAVER'. screensaver.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(setAttributesRequest(c, Drawable, X, Y, Width, Height, BorderWidth, Class, Depth, Visual, ValueMask, ValueList), cookie)
	return SetAttributesCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook SetAttributesCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for SetAttributes
// setAttributesRequest writes a SetAttributes request to a byte slice.
func setAttributesRequest(c *xgb.Conn, Drawable xproto.Drawable, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class byte, Depth byte, Visual xproto.Visualid, ValueMask uint32, ValueList []uint32) []byte {
	size := xgb.Pad((24 + (4 + xgb.Pad((4 * xgb.PopCount(int(ValueMask)))))))
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["MIT-SCREEN-SAVER"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	xgb.Put16(buf[b:], uint16(X))
	b += 2

	xgb.Put16(buf[b:], uint16(Y))
	b += 2

	xgb.Put16(buf[b:], Width)
	b += 2

	xgb.Put16(buf[b:], Height)
	b += 2

	xgb.Put16(buf[b:], BorderWidth)
	b += 2

	buf[b] = Class
	b += 1

	buf[b] = Depth
	b += 1

	xgb.Put32(buf[b:], uint32(Visual))
	b += 4

	xgb.Put32(buf[b:], ValueMask)
	b += 4
	for i := 0; i < xgb.PopCount(int(ValueMask)); i++ {
		xgb.Put32(buf[b:], ValueList[i])
		b += 4
	}
	b = xgb.Pad(b)

	return buf
}

// UnsetAttributesCookie is a cookie used only for UnsetAttributes requests.
type UnsetAttributesCookie struct {
	*xgb.Cookie
}

// UnsetAttributes sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func UnsetAttributes(c *xgb.Conn, Drawable xproto.Drawable) UnsetAttributesCookie {
	if _, ok := c.Extensions["MIT-SCREEN-SAVER"]; !ok {
		panic("Cannot issue request 'UnsetAttributes' using the uninitialized extension 'MIT-SCREEN-SAVER'. screensaver.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(unsetAttributesRequest(c, Drawable), cookie)
	return UnsetAttributesCookie{cookie}
}

// UnsetAttributesChecked sends a checked request.
// If an error occurs, it can be retrieved using UnsetAttributesCookie.Check()
func UnsetAttributesChecked(c *xgb.Conn, Drawable xproto.Drawable) UnsetAttributesCookie {
	if _, ok := c.Extensions["MIT-SCREEN-SAVER"]; !ok {
		panic("Cannot issue request 'UnsetAttributes' using the uninitialized extension 'MIT-SCREEN-SAVER'. screensaver.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(unsetAttributesRequest(c, Drawable), cookie)
	return UnsetAttributesCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook UnsetAttributesCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for UnsetAttributes
// unsetAttributesRequest writes a UnsetAttributes request to a byte slice.
func unsetAttributesRequest(c *xgb.Conn, Drawable xproto.Drawable) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["MIT-SCREEN-SAVER"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	return buf
}

// SuspendCookie is a cookie used only for Suspend requests.
type SuspendCookie struct {
	*xgb.Cookie
}

// Suspend sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func Suspend(c *xgb.Conn, Suspend bool) SuspendCookie {
	if _, ok := c.Extensions["MIT-SCREEN-SAVER"]; !ok {
		panic("Cannot issue request 'Suspend' using the uninitialized extension 'MIT-SCREEN-SAVER'. screensaver.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(suspendRequest(c, Suspend), cookie)
	return SuspendCookie{cookie}
}

// SuspendChecked sends a checked request.
// If an error occurs, it can be retrieved using SuspendCookie.Check()
func SuspendChecked(c *xgb.Conn, Suspend bool) SuspendCookie {
	if _, ok := c.Extensions["MIT-SCREEN-SAVER"]; !ok {
		panic("Cannot issue request 'Suspend' using the uninitialized extension 'MIT-SCREEN-SAVER'. screensaver.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(suspendRequest(c, Suspend), cookie)
	return SuspendCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook SuspendCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Suspend
// suspendRequest writes a Suspend request to a byte slice.
func suspendRequest(c *xgb.Conn, Suspend bool) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["MIT-SCREEN-SAVER"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	if Suspend {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}
