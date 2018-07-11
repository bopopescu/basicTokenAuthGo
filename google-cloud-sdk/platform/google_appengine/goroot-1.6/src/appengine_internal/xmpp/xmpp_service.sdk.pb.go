// Code generated by protoc_gen_go. DO NOT EDIT.

/*
Package appengine is a generated protocol buffer package.

It is generated from these files:
	appengine_internal/xmpp

It has these top-level messages:
	XmppServiceError
	PresenceRequest
	PresenceResponse
	BulkPresenceRequest
	BulkPresenceResponse
	XmppMessageRequest
	XmppMessageResponse
	XmppSendPresenceRequest
	XmppSendPresenceResponse
	XmppInviteRequest
	XmppInviteResponse
*/
package appengine

import proto "appengine_internal/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type XmppServiceError_ErrorCode int32

const (
	XmppServiceError_UNSPECIFIED_ERROR    XmppServiceError_ErrorCode = 1
	XmppServiceError_INVALID_JID          XmppServiceError_ErrorCode = 2
	XmppServiceError_NO_BODY              XmppServiceError_ErrorCode = 3
	XmppServiceError_INVALID_XML          XmppServiceError_ErrorCode = 4
	XmppServiceError_INVALID_TYPE         XmppServiceError_ErrorCode = 5
	XmppServiceError_INVALID_SHOW         XmppServiceError_ErrorCode = 6
	XmppServiceError_EXCEEDED_MAX_SIZE    XmppServiceError_ErrorCode = 7
	XmppServiceError_APPID_ALIAS_REQUIRED XmppServiceError_ErrorCode = 8
	XmppServiceError_NONDEFAULT_MODULE    XmppServiceError_ErrorCode = 9
)

var XmppServiceError_ErrorCode_name = map[int32]string{
	1: "UNSPECIFIED_ERROR",
	2: "INVALID_JID",
	3: "NO_BODY",
	4: "INVALID_XML",
	5: "INVALID_TYPE",
	6: "INVALID_SHOW",
	7: "EXCEEDED_MAX_SIZE",
	8: "APPID_ALIAS_REQUIRED",
	9: "NONDEFAULT_MODULE",
}
var XmppServiceError_ErrorCode_value = map[string]int32{
	"UNSPECIFIED_ERROR":    1,
	"INVALID_JID":          2,
	"NO_BODY":              3,
	"INVALID_XML":          4,
	"INVALID_TYPE":         5,
	"INVALID_SHOW":         6,
	"EXCEEDED_MAX_SIZE":    7,
	"APPID_ALIAS_REQUIRED": 8,
	"NONDEFAULT_MODULE":    9,
}

func (x XmppServiceError_ErrorCode) Enum() *XmppServiceError_ErrorCode {
	p := new(XmppServiceError_ErrorCode)
	*p = x
	return p
}
func (x XmppServiceError_ErrorCode) String() string {
	return proto.EnumName(XmppServiceError_ErrorCode_name, int32(x))
}
func (x *XmppServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(XmppServiceError_ErrorCode_value, data, "XmppServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = XmppServiceError_ErrorCode(value)
	return nil
}

type PresenceResponse_SHOW int32

const (
	PresenceResponse_NORMAL         PresenceResponse_SHOW = 0
	PresenceResponse_AWAY           PresenceResponse_SHOW = 1
	PresenceResponse_DO_NOT_DISTURB PresenceResponse_SHOW = 2
	PresenceResponse_CHAT           PresenceResponse_SHOW = 3
	PresenceResponse_EXTENDED_AWAY  PresenceResponse_SHOW = 4
)

var PresenceResponse_SHOW_name = map[int32]string{
	0: "NORMAL",
	1: "AWAY",
	2: "DO_NOT_DISTURB",
	3: "CHAT",
	4: "EXTENDED_AWAY",
}
var PresenceResponse_SHOW_value = map[string]int32{
	"NORMAL":         0,
	"AWAY":           1,
	"DO_NOT_DISTURB": 2,
	"CHAT":           3,
	"EXTENDED_AWAY":  4,
}

func (x PresenceResponse_SHOW) Enum() *PresenceResponse_SHOW {
	p := new(PresenceResponse_SHOW)
	*p = x
	return p
}
func (x PresenceResponse_SHOW) String() string {
	return proto.EnumName(PresenceResponse_SHOW_name, int32(x))
}
func (x *PresenceResponse_SHOW) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PresenceResponse_SHOW_value, data, "PresenceResponse_SHOW")
	if err != nil {
		return err
	}
	*x = PresenceResponse_SHOW(value)
	return nil
}

type XmppMessageResponse_XmppMessageStatus int32

const (
	XmppMessageResponse_NO_ERROR     XmppMessageResponse_XmppMessageStatus = 0
	XmppMessageResponse_INVALID_JID  XmppMessageResponse_XmppMessageStatus = 1
	XmppMessageResponse_OTHER_ERROR  XmppMessageResponse_XmppMessageStatus = 2
	XmppMessageResponse_INVALID_XML  XmppMessageResponse_XmppMessageStatus = 3
	XmppMessageResponse_INVALID_TYPE XmppMessageResponse_XmppMessageStatus = 4
)

var XmppMessageResponse_XmppMessageStatus_name = map[int32]string{
	0: "NO_ERROR",
	1: "INVALID_JID",
	2: "OTHER_ERROR",
	3: "INVALID_XML",
	4: "INVALID_TYPE",
}
var XmppMessageResponse_XmppMessageStatus_value = map[string]int32{
	"NO_ERROR":     0,
	"INVALID_JID":  1,
	"OTHER_ERROR":  2,
	"INVALID_XML":  3,
	"INVALID_TYPE": 4,
}

func (x XmppMessageResponse_XmppMessageStatus) Enum() *XmppMessageResponse_XmppMessageStatus {
	p := new(XmppMessageResponse_XmppMessageStatus)
	*p = x
	return p
}
func (x XmppMessageResponse_XmppMessageStatus) String() string {
	return proto.EnumName(XmppMessageResponse_XmppMessageStatus_name, int32(x))
}
func (x *XmppMessageResponse_XmppMessageStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(XmppMessageResponse_XmppMessageStatus_value, data, "XmppMessageResponse_XmppMessageStatus")
	if err != nil {
		return err
	}
	*x = XmppMessageResponse_XmppMessageStatus(value)
	return nil
}

type XmppServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *XmppServiceError) Reset()         { *m = XmppServiceError{} }
func (m *XmppServiceError) String() string { return proto.CompactTextString(m) }
func (*XmppServiceError) ProtoMessage()    {}

type PresenceRequest struct {
	Jid              *string `protobuf:"bytes,1,req,name=jid" json:"jid,omitempty"`
	FromJid          *string `protobuf:"bytes,2,opt,name=from_jid,json=fromJid" json:"from_jid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PresenceRequest) Reset()         { *m = PresenceRequest{} }
func (m *PresenceRequest) String() string { return proto.CompactTextString(m) }
func (*PresenceRequest) ProtoMessage()    {}

func (m *PresenceRequest) GetJid() string {
	if m != nil && m.Jid != nil {
		return *m.Jid
	}
	return ""
}

func (m *PresenceRequest) GetFromJid() string {
	if m != nil && m.FromJid != nil {
		return *m.FromJid
	}
	return ""
}

type PresenceResponse struct {
	IsAvailable      *bool                  `protobuf:"varint,1,req,name=is_available,json=isAvailable" json:"is_available,omitempty"`
	Presence         *PresenceResponse_SHOW `protobuf:"varint,2,opt,name=presence,enum=appengine.PresenceResponse_SHOW" json:"presence,omitempty"`
	Valid            *bool                  `protobuf:"varint,3,opt,name=valid" json:"valid,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *PresenceResponse) Reset()         { *m = PresenceResponse{} }
func (m *PresenceResponse) String() string { return proto.CompactTextString(m) }
func (*PresenceResponse) ProtoMessage()    {}

func (m *PresenceResponse) GetIsAvailable() bool {
	if m != nil && m.IsAvailable != nil {
		return *m.IsAvailable
	}
	return false
}

func (m *PresenceResponse) GetPresence() PresenceResponse_SHOW {
	if m != nil && m.Presence != nil {
		return *m.Presence
	}
	return PresenceResponse_NORMAL
}

func (m *PresenceResponse) GetValid() bool {
	if m != nil && m.Valid != nil {
		return *m.Valid
	}
	return false
}

type BulkPresenceRequest struct {
	Jid              []string `protobuf:"bytes,1,rep,name=jid" json:"jid,omitempty"`
	FromJid          *string  `protobuf:"bytes,2,opt,name=from_jid,json=fromJid" json:"from_jid,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BulkPresenceRequest) Reset()         { *m = BulkPresenceRequest{} }
func (m *BulkPresenceRequest) String() string { return proto.CompactTextString(m) }
func (*BulkPresenceRequest) ProtoMessage()    {}

func (m *BulkPresenceRequest) GetJid() []string {
	if m != nil {
		return m.Jid
	}
	return nil
}

func (m *BulkPresenceRequest) GetFromJid() string {
	if m != nil && m.FromJid != nil {
		return *m.FromJid
	}
	return ""
}

type BulkPresenceResponse struct {
	PresenceResponse []*PresenceResponse `protobuf:"bytes,1,rep,name=presence_response,json=presenceResponse" json:"presence_response,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *BulkPresenceResponse) Reset()         { *m = BulkPresenceResponse{} }
func (m *BulkPresenceResponse) String() string { return proto.CompactTextString(m) }
func (*BulkPresenceResponse) ProtoMessage()    {}

func (m *BulkPresenceResponse) GetPresenceResponse() []*PresenceResponse {
	if m != nil {
		return m.PresenceResponse
	}
	return nil
}

type XmppMessageRequest struct {
	Jid              []string `protobuf:"bytes,1,rep,name=jid" json:"jid,omitempty"`
	Body             *string  `protobuf:"bytes,2,req,name=body" json:"body,omitempty"`
	RawXml           *bool    `protobuf:"varint,3,opt,name=raw_xml,json=rawXml,def=0" json:"raw_xml,omitempty"`
	Type             *string  `protobuf:"bytes,4,opt,name=type,def=chat" json:"type,omitempty"`
	FromJid          *string  `protobuf:"bytes,5,opt,name=from_jid,json=fromJid" json:"from_jid,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *XmppMessageRequest) Reset()         { *m = XmppMessageRequest{} }
func (m *XmppMessageRequest) String() string { return proto.CompactTextString(m) }
func (*XmppMessageRequest) ProtoMessage()    {}

const Default_XmppMessageRequest_RawXml bool = false
const Default_XmppMessageRequest_Type string = "chat"

func (m *XmppMessageRequest) GetJid() []string {
	if m != nil {
		return m.Jid
	}
	return nil
}

func (m *XmppMessageRequest) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func (m *XmppMessageRequest) GetRawXml() bool {
	if m != nil && m.RawXml != nil {
		return *m.RawXml
	}
	return Default_XmppMessageRequest_RawXml
}

func (m *XmppMessageRequest) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_XmppMessageRequest_Type
}

func (m *XmppMessageRequest) GetFromJid() string {
	if m != nil && m.FromJid != nil {
		return *m.FromJid
	}
	return ""
}

type XmppMessageResponse struct {
	Status           []XmppMessageResponse_XmppMessageStatus `protobuf:"varint,1,rep,name=status,enum=appengine.XmppMessageResponse_XmppMessageStatus" json:"status,omitempty"`
	XXX_unrecognized []byte                                  `json:"-"`
}

func (m *XmppMessageResponse) Reset()         { *m = XmppMessageResponse{} }
func (m *XmppMessageResponse) String() string { return proto.CompactTextString(m) }
func (*XmppMessageResponse) ProtoMessage()    {}

func (m *XmppMessageResponse) GetStatus() []XmppMessageResponse_XmppMessageStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type XmppSendPresenceRequest struct {
	Jid              *string `protobuf:"bytes,1,req,name=jid" json:"jid,omitempty"`
	Type             *string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Show             *string `protobuf:"bytes,3,opt,name=show" json:"show,omitempty"`
	Status           *string `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	FromJid          *string `protobuf:"bytes,5,opt,name=from_jid,json=fromJid" json:"from_jid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *XmppSendPresenceRequest) Reset()         { *m = XmppSendPresenceRequest{} }
func (m *XmppSendPresenceRequest) String() string { return proto.CompactTextString(m) }
func (*XmppSendPresenceRequest) ProtoMessage()    {}

func (m *XmppSendPresenceRequest) GetJid() string {
	if m != nil && m.Jid != nil {
		return *m.Jid
	}
	return ""
}

func (m *XmppSendPresenceRequest) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *XmppSendPresenceRequest) GetShow() string {
	if m != nil && m.Show != nil {
		return *m.Show
	}
	return ""
}

func (m *XmppSendPresenceRequest) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *XmppSendPresenceRequest) GetFromJid() string {
	if m != nil && m.FromJid != nil {
		return *m.FromJid
	}
	return ""
}

type XmppSendPresenceResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *XmppSendPresenceResponse) Reset()         { *m = XmppSendPresenceResponse{} }
func (m *XmppSendPresenceResponse) String() string { return proto.CompactTextString(m) }
func (*XmppSendPresenceResponse) ProtoMessage()    {}

type XmppInviteRequest struct {
	Jid              *string `protobuf:"bytes,1,req,name=jid" json:"jid,omitempty"`
	FromJid          *string `protobuf:"bytes,2,opt,name=from_jid,json=fromJid" json:"from_jid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *XmppInviteRequest) Reset()         { *m = XmppInviteRequest{} }
func (m *XmppInviteRequest) String() string { return proto.CompactTextString(m) }
func (*XmppInviteRequest) ProtoMessage()    {}

func (m *XmppInviteRequest) GetJid() string {
	if m != nil && m.Jid != nil {
		return *m.Jid
	}
	return ""
}

func (m *XmppInviteRequest) GetFromJid() string {
	if m != nil && m.FromJid != nil {
		return *m.FromJid
	}
	return ""
}

type XmppInviteResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *XmppInviteResponse) Reset()         { *m = XmppInviteResponse{} }
func (m *XmppInviteResponse) String() string { return proto.CompactTextString(m) }
func (*XmppInviteResponse) ProtoMessage()    {}

func init() {
	proto.RegisterType((*XmppServiceError)(nil), "appengine.XmppServiceError")
	proto.RegisterType((*PresenceRequest)(nil), "appengine.PresenceRequest")
	proto.RegisterType((*PresenceResponse)(nil), "appengine.PresenceResponse")
	proto.RegisterType((*BulkPresenceRequest)(nil), "appengine.BulkPresenceRequest")
	proto.RegisterType((*BulkPresenceResponse)(nil), "appengine.BulkPresenceResponse")
	proto.RegisterType((*XmppMessageRequest)(nil), "appengine.XmppMessageRequest")
	proto.RegisterType((*XmppMessageResponse)(nil), "appengine.XmppMessageResponse")
	proto.RegisterType((*XmppSendPresenceRequest)(nil), "appengine.XmppSendPresenceRequest")
	proto.RegisterType((*XmppSendPresenceResponse)(nil), "appengine.XmppSendPresenceResponse")
	proto.RegisterType((*XmppInviteRequest)(nil), "appengine.XmppInviteRequest")
	proto.RegisterType((*XmppInviteResponse)(nil), "appengine.XmppInviteResponse")
	proto.RegisterEnum("appengine.XmppServiceError_ErrorCode", XmppServiceError_ErrorCode_name, XmppServiceError_ErrorCode_value)
	proto.RegisterEnum("appengine.PresenceResponse_SHOW", PresenceResponse_SHOW_name, PresenceResponse_SHOW_value)
	proto.RegisterEnum("appengine.XmppMessageResponse_XmppMessageStatus", XmppMessageResponse_XmppMessageStatus_name, XmppMessageResponse_XmppMessageStatus_value)
}

func init() {
}
