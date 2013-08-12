// Package shm is the X client API for the MIT-SHM extension.
package shm

// This file is automatically generated from shm.xml. Edit at your peril!

import (
	"github.com/BurntSushi/xgb"

	"github.com/BurntSushi/xgb/xproto"
)

// Init must be called before using the MIT-SHM extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 7, "MIT-SHM").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named MIT-SHM could be found on on the server.")
	}

	xgb.ExtLock.Lock()
	c.Extensions["MIT-SHM"] = reply.MajorOpcode
	for evNum, fun := range xgb.NewExtEventFuncs["MIT-SHM"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["MIT-SHM"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	xgb.ExtLock.Unlock()

	return nil
}

func init() {
	xgb.NewExtEventFuncs["MIT-SHM"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["MIT-SHM"] = make(map[int]xgb.NewErrorFun)
}

// BadBadSeg is the error number for a BadBadSeg.
const BadBadSeg = 0

type BadSegError xproto.ValueError

// BadSegErrorNew constructs a BadSegError value that implements xgb.Error from a byte slice.
func BadSegErrorNew(buf []byte) xgb.Error {
	v := BadSegError(xproto.ValueErrorNew(buf).(xproto.ValueError))
	v.NiceName = "BadSeg"
	return v
}

// SequenceId returns the sequence id attached to the BadBadSeg error.
// This is mostly used internally.
func (err BadSegError) SequenceId() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadBadSeg error. If no bad value exists, 0 is returned.
func (err BadSegError) BadId() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadSeg error.
func (err BadSegError) Error() string {
	fieldVals := make([]string, 0, 4)
	fieldVals = append(fieldVals, "NiceName: "+err.NiceName)
	fieldVals = append(fieldVals, xgb.Sprintf("Sequence: %d", err.Sequence))
	fieldVals = append(fieldVals, xgb.Sprintf("BadValue: %d", err.BadValue))
	fieldVals = append(fieldVals, xgb.Sprintf("MinorOpcode: %d", err.MinorOpcode))
	fieldVals = append(fieldVals, xgb.Sprintf("MajorOpcode: %d", err.MajorOpcode))
	return "BadBadSeg {" + xgb.StringsJoin(fieldVals, ", ") + "}"
}

func init() {
	xgb.NewExtErrorFuncs["MIT-SHM"][0] = BadSegErrorNew
}

// Completion is the event number for a CompletionEvent.
const Completion = 0

type CompletionEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Drawable   xproto.Drawable
	MinorEvent uint16
	MajorEvent byte
	// padding: 1 bytes
	Shmseg Seg
	Offset uint32
}

// CompletionEventNew constructs a CompletionEvent value that implements xgb.Event from a byte slice.
func CompletionEventNew(buf []byte) xgb.Event {
	v := CompletionEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Drawable = xproto.Drawable(xgb.Get32(buf[b:]))
	b += 4

	v.MinorEvent = xgb.Get16(buf[b:])
	b += 2

	v.MajorEvent = buf[b]
	b += 1

	b += 1 // padding

	v.Shmseg = Seg(xgb.Get32(buf[b:]))
	b += 4

	v.Offset = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Bytes writes a CompletionEvent value to a byte slice.
func (v CompletionEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	xgb.Put32(buf[b:], uint32(v.Drawable))
	b += 4

	xgb.Put16(buf[b:], v.MinorEvent)
	b += 2

	buf[b] = v.MajorEvent
	b += 1

	b += 1 // padding

	xgb.Put32(buf[b:], uint32(v.Shmseg))
	b += 4

	xgb.Put32(buf[b:], v.Offset)
	b += 4

	return buf
}

// SequenceId returns the sequence id attached to the Completion event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v CompletionEvent) SequenceId() uint16 {
	return v.Sequence
}

// String is a rudimentary string representation of CompletionEvent.
func (v CompletionEvent) String() string {
	fieldVals := make([]string, 0, 7)
	fieldVals = append(fieldVals, xgb.Sprintf("Sequence: %d", v.Sequence))
	fieldVals = append(fieldVals, xgb.Sprintf("Drawable: %d", v.Drawable))
	fieldVals = append(fieldVals, xgb.Sprintf("MinorEvent: %d", v.MinorEvent))
	fieldVals = append(fieldVals, xgb.Sprintf("MajorEvent: %d", v.MajorEvent))
	fieldVals = append(fieldVals, xgb.Sprintf("Shmseg: %d", v.Shmseg))
	fieldVals = append(fieldVals, xgb.Sprintf("Offset: %d", v.Offset))
	return "Completion {" + xgb.StringsJoin(fieldVals, ", ") + "}"
}

func init() {
	xgb.NewExtEventFuncs["MIT-SHM"][0] = CompletionEventNew
}

type Seg uint32

func NewSegId(c *xgb.Conn) (Seg, error) {
	id, err := c.NewId()
	if err != nil {
		return 0, err
	}
	return Seg(id), nil
}

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Card32'

// AttachCookie is a cookie used only for Attach requests.
type AttachCookie struct {
	*xgb.Cookie
}

// Attach sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func Attach(c *xgb.Conn, Shmseg Seg, Shmid uint32, ReadOnly bool) AttachCookie {
	if _, ok := c.Extensions["MIT-SHM"]; !ok {
		panic("Cannot issue request 'Attach' using the uninitialized extension 'MIT-SHM'. shm.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(attachRequest(c, Shmseg, Shmid, ReadOnly), cookie)
	return AttachCookie{cookie}
}

// AttachChecked sends a checked request.
// If an error occurs, it can be retrieved using AttachCookie.Check()
func AttachChecked(c *xgb.Conn, Shmseg Seg, Shmid uint32, ReadOnly bool) AttachCookie {
	if _, ok := c.Extensions["MIT-SHM"]; !ok {
		panic("Cannot issue request 'Attach' using the uninitialized extension 'MIT-SHM'. shm.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(attachRequest(c, Shmseg, Shmid, ReadOnly), cookie)
	return AttachCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook AttachCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Attach
// attachRequest writes a Attach request to a byte slice.
func attachRequest(c *xgb.Conn, Shmseg Seg, Shmid uint32, ReadOnly bool) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["MIT-SHM"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Shmseg))
	b += 4

	xgb.Put32(buf[b:], Shmid)
	b += 4

	if ReadOnly {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}

// CreatePixmapCookie is a cookie used only for CreatePixmap requests.
type CreatePixmapCookie struct {
	*xgb.Cookie
}

// CreatePixmap sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func CreatePixmap(c *xgb.Conn, Pid xproto.Pixmap, Drawable xproto.Drawable, Width uint16, Height uint16, Depth byte, Shmseg Seg, Offset uint32) CreatePixmapCookie {
	if _, ok := c.Extensions["MIT-SHM"]; !ok {
		panic("Cannot issue request 'CreatePixmap' using the uninitialized extension 'MIT-SHM'. shm.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(createPixmapRequest(c, Pid, Drawable, Width, Height, Depth, Shmseg, Offset), cookie)
	return CreatePixmapCookie{cookie}
}

// CreatePixmapChecked sends a checked request.
// If an error occurs, it can be retrieved using CreatePixmapCookie.Check()
func CreatePixmapChecked(c *xgb.Conn, Pid xproto.Pixmap, Drawable xproto.Drawable, Width uint16, Height uint16, Depth byte, Shmseg Seg, Offset uint32) CreatePixmapCookie {
	if _, ok := c.Extensions["MIT-SHM"]; !ok {
		panic("Cannot issue request 'CreatePixmap' using the uninitialized extension 'MIT-SHM'. shm.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(createPixmapRequest(c, Pid, Drawable, Width, Height, Depth, Shmseg, Offset), cookie)
	return CreatePixmapCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook CreatePixmapCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for CreatePixmap
// createPixmapRequest writes a CreatePixmap request to a byte slice.
func createPixmapRequest(c *xgb.Conn, Pid xproto.Pixmap, Drawable xproto.Drawable, Width uint16, Height uint16, Depth byte, Shmseg Seg, Offset uint32) []byte {
	size := 28
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["MIT-SHM"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Pid))
	b += 4

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	xgb.Put16(buf[b:], Width)
	b += 2

	xgb.Put16(buf[b:], Height)
	b += 2

	buf[b] = Depth
	b += 1

	b += 3 // padding

	xgb.Put32(buf[b:], uint32(Shmseg))
	b += 4

	xgb.Put32(buf[b:], Offset)
	b += 4

	return buf
}

// DetachCookie is a cookie used only for Detach requests.
type DetachCookie struct {
	*xgb.Cookie
}

// Detach sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func Detach(c *xgb.Conn, Shmseg Seg) DetachCookie {
	if _, ok := c.Extensions["MIT-SHM"]; !ok {
		panic("Cannot issue request 'Detach' using the uninitialized extension 'MIT-SHM'. shm.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(detachRequest(c, Shmseg), cookie)
	return DetachCookie{cookie}
}

// DetachChecked sends a checked request.
// If an error occurs, it can be retrieved using DetachCookie.Check()
func DetachChecked(c *xgb.Conn, Shmseg Seg) DetachCookie {
	if _, ok := c.Extensions["MIT-SHM"]; !ok {
		panic("Cannot issue request 'Detach' using the uninitialized extension 'MIT-SHM'. shm.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(detachRequest(c, Shmseg), cookie)
	return DetachCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook DetachCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Detach
// detachRequest writes a Detach request to a byte slice.
func detachRequest(c *xgb.Conn, Shmseg Seg) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["MIT-SHM"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Shmseg))
	b += 4

	return buf
}

// GetImageCookie is a cookie used only for GetImage requests.
type GetImageCookie struct {
	*xgb.Cookie
}

// GetImage sends a checked request.
// If an error occurs, it will be returned with the reply by calling GetImageCookie.Reply()
func GetImage(c *xgb.Conn, Drawable xproto.Drawable, X int16, Y int16, Width uint16, Height uint16, PlaneMask uint32, Format byte, Shmseg Seg, Offset uint32) GetImageCookie {
	if _, ok := c.Extensions["MIT-SHM"]; !ok {
		panic("Cannot issue request 'GetImage' using the uninitialized extension 'MIT-SHM'. shm.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(getImageRequest(c, Drawable, X, Y, Width, Height, PlaneMask, Format, Shmseg, Offset), cookie)
	return GetImageCookie{cookie}
}

// GetImageUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func GetImageUnchecked(c *xgb.Conn, Drawable xproto.Drawable, X int16, Y int16, Width uint16, Height uint16, PlaneMask uint32, Format byte, Shmseg Seg, Offset uint32) GetImageCookie {
	if _, ok := c.Extensions["MIT-SHM"]; !ok {
		panic("Cannot issue request 'GetImage' using the uninitialized extension 'MIT-SHM'. shm.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(getImageRequest(c, Drawable, X, Y, Width, Height, PlaneMask, Format, Shmseg, Offset), cookie)
	return GetImageCookie{cookie}
}

// GetImageReply represents the data returned from a GetImage request.
type GetImageReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Depth    byte
	Visual   xproto.Visualid
	Size     uint32
}

// Reply blocks and returns the reply data for a GetImage request.
func (cook GetImageCookie) Reply() (*GetImageReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getImageReply(buf), nil
}

// getImageReply reads a byte slice into a GetImageReply value.
func getImageReply(buf []byte) *GetImageReply {
	v := new(GetImageReply)
	b := 1 // skip reply determinant

	v.Depth = buf[b]
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Visual = xproto.Visualid(xgb.Get32(buf[b:]))
	b += 4

	v.Size = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for GetImage
// getImageRequest writes a GetImage request to a byte slice.
func getImageRequest(c *xgb.Conn, Drawable xproto.Drawable, X int16, Y int16, Width uint16, Height uint16, PlaneMask uint32, Format byte, Shmseg Seg, Offset uint32) []byte {
	size := 32
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["MIT-SHM"]
	b += 1

	buf[b] = 4 // request opcode
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

	xgb.Put32(buf[b:], PlaneMask)
	b += 4

	buf[b] = Format
	b += 1

	b += 3 // padding

	xgb.Put32(buf[b:], uint32(Shmseg))
	b += 4

	xgb.Put32(buf[b:], Offset)
	b += 4

	return buf
}

// PutImageCookie is a cookie used only for PutImage requests.
type PutImageCookie struct {
	*xgb.Cookie
}

// PutImage sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func PutImage(c *xgb.Conn, Drawable xproto.Drawable, Gc xproto.Gcontext, TotalWidth uint16, TotalHeight uint16, SrcX uint16, SrcY uint16, SrcWidth uint16, SrcHeight uint16, DstX int16, DstY int16, Depth byte, Format byte, SendEvent byte, Shmseg Seg, Offset uint32) PutImageCookie {
	if _, ok := c.Extensions["MIT-SHM"]; !ok {
		panic("Cannot issue request 'PutImage' using the uninitialized extension 'MIT-SHM'. shm.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(putImageRequest(c, Drawable, Gc, TotalWidth, TotalHeight, SrcX, SrcY, SrcWidth, SrcHeight, DstX, DstY, Depth, Format, SendEvent, Shmseg, Offset), cookie)
	return PutImageCookie{cookie}
}

// PutImageChecked sends a checked request.
// If an error occurs, it can be retrieved using PutImageCookie.Check()
func PutImageChecked(c *xgb.Conn, Drawable xproto.Drawable, Gc xproto.Gcontext, TotalWidth uint16, TotalHeight uint16, SrcX uint16, SrcY uint16, SrcWidth uint16, SrcHeight uint16, DstX int16, DstY int16, Depth byte, Format byte, SendEvent byte, Shmseg Seg, Offset uint32) PutImageCookie {
	if _, ok := c.Extensions["MIT-SHM"]; !ok {
		panic("Cannot issue request 'PutImage' using the uninitialized extension 'MIT-SHM'. shm.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(putImageRequest(c, Drawable, Gc, TotalWidth, TotalHeight, SrcX, SrcY, SrcWidth, SrcHeight, DstX, DstY, Depth, Format, SendEvent, Shmseg, Offset), cookie)
	return PutImageCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook PutImageCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for PutImage
// putImageRequest writes a PutImage request to a byte slice.
func putImageRequest(c *xgb.Conn, Drawable xproto.Drawable, Gc xproto.Gcontext, TotalWidth uint16, TotalHeight uint16, SrcX uint16, SrcY uint16, SrcWidth uint16, SrcHeight uint16, DstX int16, DstY int16, Depth byte, Format byte, SendEvent byte, Shmseg Seg, Offset uint32) []byte {
	size := 40
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["MIT-SHM"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	xgb.Put32(buf[b:], uint32(Gc))
	b += 4

	xgb.Put16(buf[b:], TotalWidth)
	b += 2

	xgb.Put16(buf[b:], TotalHeight)
	b += 2

	xgb.Put16(buf[b:], SrcX)
	b += 2

	xgb.Put16(buf[b:], SrcY)
	b += 2

	xgb.Put16(buf[b:], SrcWidth)
	b += 2

	xgb.Put16(buf[b:], SrcHeight)
	b += 2

	xgb.Put16(buf[b:], uint16(DstX))
	b += 2

	xgb.Put16(buf[b:], uint16(DstY))
	b += 2

	buf[b] = Depth
	b += 1

	buf[b] = Format
	b += 1

	buf[b] = SendEvent
	b += 1

	b += 1 // padding

	xgb.Put32(buf[b:], uint32(Shmseg))
	b += 4

	xgb.Put32(buf[b:], Offset)
	b += 4

	return buf
}

// QueryVersionCookie is a cookie used only for QueryVersion requests.
type QueryVersionCookie struct {
	*xgb.Cookie
}

// QueryVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryVersionCookie.Reply()
func QueryVersion(c *xgb.Conn) QueryVersionCookie {
	if _, ok := c.Extensions["MIT-SHM"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'MIT-SHM'. shm.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryVersionUnchecked(c *xgb.Conn) QueryVersionCookie {
	if _, ok := c.Extensions["MIT-SHM"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'MIT-SHM'. shm.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence      uint16 // sequence number of the request for this reply
	Length        uint32 // number of bytes in this reply
	SharedPixmaps bool
	MajorVersion  uint16
	MinorVersion  uint16
	Uid           uint16
	Gid           uint16
	PixmapFormat  byte
	// padding: 15 bytes
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

	if buf[b] == 1 {
		v.SharedPixmaps = true
	} else {
		v.SharedPixmaps = false
	}
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.MajorVersion = xgb.Get16(buf[b:])
	b += 2

	v.MinorVersion = xgb.Get16(buf[b:])
	b += 2

	v.Uid = xgb.Get16(buf[b:])
	b += 2

	v.Gid = xgb.Get16(buf[b:])
	b += 2

	v.PixmapFormat = buf[b]
	b += 1

	b += 15 // padding

	return v
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["MIT-SHM"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}
