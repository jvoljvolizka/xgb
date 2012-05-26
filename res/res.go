// Package res is the X client API for the X-Resource extension.
package res

/*
	This file was generated by res.xml on May 26 2012 6:23:13pm EDT.
	This file is automatically generated. Edit at your peril!
*/

import (
	"github.com/BurntSushi/xgb"

	"github.com/BurntSushi/xgb/xproto"
)

// Init must be called before using the X-Resource extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 10, "X-Resource").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named X-Resource could be found on on the server.")
	}

	xgb.ExtLock.Lock()
	c.Extensions["X-Resource"] = reply.MajorOpcode
	for evNum, fun := range xgb.NewExtEventFuncs["X-Resource"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["X-Resource"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	xgb.ExtLock.Unlock()

	return nil
}

func init() {
	xgb.NewExtEventFuncs["X-Resource"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["X-Resource"] = make(map[int]xgb.NewErrorFun)
}

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

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

type Client struct {
	ResourceBase uint32
	ResourceMask uint32
}

// ClientRead reads a byte slice into a Client value.
func ClientRead(buf []byte, v *Client) int {
	b := 0

	v.ResourceBase = xgb.Get32(buf[b:])
	b += 4

	v.ResourceMask = xgb.Get32(buf[b:])
	b += 4

	return b
}

// ClientReadList reads a byte slice into a list of Client values.
func ClientReadList(buf []byte, dest []Client) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Client{}
		b += ClientRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a Client value to a byte slice.
func (v Client) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	xgb.Put32(buf[b:], v.ResourceBase)
	b += 4

	xgb.Put32(buf[b:], v.ResourceMask)
	b += 4

	return buf
}

// ClientListBytes writes a list of Client values to a byte slice.
func ClientListBytes(buf []byte, list []Client) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}
	return b
}

type Type struct {
	ResourceType xproto.Atom
	Count        uint32
}

// TypeRead reads a byte slice into a Type value.
func TypeRead(buf []byte, v *Type) int {
	b := 0

	v.ResourceType = xproto.Atom(xgb.Get32(buf[b:]))
	b += 4

	v.Count = xgb.Get32(buf[b:])
	b += 4

	return b
}

// TypeReadList reads a byte slice into a list of Type values.
func TypeReadList(buf []byte, dest []Type) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Type{}
		b += TypeRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a Type value to a byte slice.
func (v Type) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	xgb.Put32(buf[b:], uint32(v.ResourceType))
	b += 4

	xgb.Put32(buf[b:], v.Count)
	b += 4

	return buf
}

// TypeListBytes writes a list of Type values to a byte slice.
func TypeListBytes(buf []byte, list []Type) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}
	return b
}

// QueryVersionCookie is a cookie used only for QueryVersion requests.
type QueryVersionCookie struct {
	*xgb.Cookie
}

// QueryVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryVersionCookie.Reply()
func QueryVersion(c *xgb.Conn, ClientMajor byte, ClientMinor byte) QueryVersionCookie {
	if _, ok := c.Extensions["X-RESOURCE"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c, ClientMajor, ClientMinor), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryVersionUnchecked(c *xgb.Conn, ClientMajor byte, ClientMinor byte) QueryVersionCookie {
	if _, ok := c.Extensions["X-RESOURCE"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c, ClientMajor, ClientMinor), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ServerMajor uint16
	ServerMinor uint16
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

	v.ServerMajor = xgb.Get16(buf[b:])
	b += 2

	v.ServerMinor = xgb.Get16(buf[b:])
	b += 2

	return v
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(c *xgb.Conn, ClientMajor byte, ClientMinor byte) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["X-RESOURCE"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = ClientMajor
	b += 1

	buf[b] = ClientMinor
	b += 1

	return buf
}

// QueryClientsCookie is a cookie used only for QueryClients requests.
type QueryClientsCookie struct {
	*xgb.Cookie
}

// QueryClients sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryClientsCookie.Reply()
func QueryClients(c *xgb.Conn) QueryClientsCookie {
	if _, ok := c.Extensions["X-RESOURCE"]; !ok {
		panic("Cannot issue request 'QueryClients' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryClientsRequest(c), cookie)
	return QueryClientsCookie{cookie}
}

// QueryClientsUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryClientsUnchecked(c *xgb.Conn) QueryClientsCookie {
	if _, ok := c.Extensions["X-RESOURCE"]; !ok {
		panic("Cannot issue request 'QueryClients' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryClientsRequest(c), cookie)
	return QueryClientsCookie{cookie}
}

// QueryClientsReply represents the data returned from a QueryClients request.
type QueryClientsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumClients uint32
	// padding: 20 bytes
	Clients []Client // size: xgb.Pad((int(NumClients) * 8))
}

// Reply blocks and returns the reply data for a QueryClients request.
func (cook QueryClientsCookie) Reply() (*QueryClientsReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryClientsReply(buf), nil
}

// queryClientsReply reads a byte slice into a QueryClientsReply value.
func queryClientsReply(buf []byte) *QueryClientsReply {
	v := new(QueryClientsReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.NumClients = xgb.Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Clients = make([]Client, v.NumClients)
	b += ClientReadList(buf[b:], v.Clients)

	return v
}

// Write request to wire for QueryClients
// queryClientsRequest writes a QueryClients request to a byte slice.
func queryClientsRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["X-RESOURCE"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// QueryClientResourcesCookie is a cookie used only for QueryClientResources requests.
type QueryClientResourcesCookie struct {
	*xgb.Cookie
}

// QueryClientResources sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryClientResourcesCookie.Reply()
func QueryClientResources(c *xgb.Conn, Xid uint32) QueryClientResourcesCookie {
	if _, ok := c.Extensions["X-RESOURCE"]; !ok {
		panic("Cannot issue request 'QueryClientResources' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryClientResourcesRequest(c, Xid), cookie)
	return QueryClientResourcesCookie{cookie}
}

// QueryClientResourcesUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryClientResourcesUnchecked(c *xgb.Conn, Xid uint32) QueryClientResourcesCookie {
	if _, ok := c.Extensions["X-RESOURCE"]; !ok {
		panic("Cannot issue request 'QueryClientResources' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryClientResourcesRequest(c, Xid), cookie)
	return QueryClientResourcesCookie{cookie}
}

// QueryClientResourcesReply represents the data returned from a QueryClientResources request.
type QueryClientResourcesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumTypes uint32
	// padding: 20 bytes
	Types []Type // size: xgb.Pad((int(NumTypes) * 8))
}

// Reply blocks and returns the reply data for a QueryClientResources request.
func (cook QueryClientResourcesCookie) Reply() (*QueryClientResourcesReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryClientResourcesReply(buf), nil
}

// queryClientResourcesReply reads a byte slice into a QueryClientResourcesReply value.
func queryClientResourcesReply(buf []byte) *QueryClientResourcesReply {
	v := new(QueryClientResourcesReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.NumTypes = xgb.Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Types = make([]Type, v.NumTypes)
	b += TypeReadList(buf[b:], v.Types)

	return v
}

// Write request to wire for QueryClientResources
// queryClientResourcesRequest writes a QueryClientResources request to a byte slice.
func queryClientResourcesRequest(c *xgb.Conn, Xid uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["X-RESOURCE"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Xid)
	b += 4

	return buf
}

// QueryClientPixmapBytesCookie is a cookie used only for QueryClientPixmapBytes requests.
type QueryClientPixmapBytesCookie struct {
	*xgb.Cookie
}

// QueryClientPixmapBytes sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryClientPixmapBytesCookie.Reply()
func QueryClientPixmapBytes(c *xgb.Conn, Xid uint32) QueryClientPixmapBytesCookie {
	if _, ok := c.Extensions["X-RESOURCE"]; !ok {
		panic("Cannot issue request 'QueryClientPixmapBytes' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryClientPixmapBytesRequest(c, Xid), cookie)
	return QueryClientPixmapBytesCookie{cookie}
}

// QueryClientPixmapBytesUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryClientPixmapBytesUnchecked(c *xgb.Conn, Xid uint32) QueryClientPixmapBytesCookie {
	if _, ok := c.Extensions["X-RESOURCE"]; !ok {
		panic("Cannot issue request 'QueryClientPixmapBytes' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryClientPixmapBytesRequest(c, Xid), cookie)
	return QueryClientPixmapBytesCookie{cookie}
}

// QueryClientPixmapBytesReply represents the data returned from a QueryClientPixmapBytes request.
type QueryClientPixmapBytesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Bytes         uint32
	BytesOverflow uint32
}

// Reply blocks and returns the reply data for a QueryClientPixmapBytes request.
func (cook QueryClientPixmapBytesCookie) Reply() (*QueryClientPixmapBytesReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryClientPixmapBytesReply(buf), nil
}

// queryClientPixmapBytesReply reads a byte slice into a QueryClientPixmapBytesReply value.
func queryClientPixmapBytesReply(buf []byte) *QueryClientPixmapBytesReply {
	v := new(QueryClientPixmapBytesReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Bytes = xgb.Get32(buf[b:])
	b += 4

	v.BytesOverflow = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for QueryClientPixmapBytes
// queryClientPixmapBytesRequest writes a QueryClientPixmapBytes request to a byte slice.
func queryClientPixmapBytesRequest(c *xgb.Conn, Xid uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["X-RESOURCE"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Xid)
	b += 4

	return buf
}
