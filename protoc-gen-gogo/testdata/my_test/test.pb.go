// Code generated by protoc-gen-gogo.
// source: my_test/test.proto
// DO NOT EDIT!

package my_test

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import multi2 "multi/multi1.pb"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type HatType int32

const (
	// deliberately skipping 0
	HatType_FEDORA HatType = 1
	HatType_FEZ    HatType = 2
)

var HatType_name = map[int32]string{
	1: "FEDORA",
	2: "FEZ",
}
var HatType_value = map[string]int32{
	"FEDORA": 1,
	"FEZ":    2,
}

func (x HatType) Enum() *HatType {
	p := new(HatType)
	*p = x
	return p
}
func (x HatType) String() string {
	return proto.EnumName(HatType_name, int32(x))
}
func (x HatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *HatType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HatType_value, data, "HatType")
	if err != nil {
		return err
	}
	*x = HatType(value)
	return nil
}

// This enum represents days of the week.
type Days int32

const (
	Days_MONDAY  Days = 1
	Days_TUESDAY Days = 2
	Days_LUNDI   Days = 1
)

var Days_name = map[int32]string{
	1: "MONDAY",
	2: "TUESDAY",
	// Duplicate value: 1: "LUNDI",
}
var Days_value = map[string]int32{
	"MONDAY":  1,
	"TUESDAY": 2,
	"LUNDI":   1,
}

func (x Days) Enum() *Days {
	p := new(Days)
	*p = x
	return p
}
func (x Days) String() string {
	return proto.EnumName(Days_name, int32(x))
}
func (x Days) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Days) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Days_value, data, "Days")
	if err != nil {
		return err
	}
	*x = Days(value)
	return nil
}

type Request_Color int32

const (
	Request_RED   Request_Color = 0
	Request_GREEN Request_Color = 1
	Request_BLUE  Request_Color = 2
)

var Request_Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}
var Request_Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x Request_Color) Enum() *Request_Color {
	p := new(Request_Color)
	*p = x
	return p
}
func (x Request_Color) String() string {
	return proto.EnumName(Request_Color_name, int32(x))
}
func (x Request_Color) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Request_Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Request_Color_value, data, "Request_Color")
	if err != nil {
		return err
	}
	*x = Request_Color(value)
	return nil
}

type Reply_Entry_Game int32

const (
	Reply_Entry_FOOTBALL Reply_Entry_Game = 1
	Reply_Entry_TENNIS   Reply_Entry_Game = 2
)

var Reply_Entry_Game_name = map[int32]string{
	1: "FOOTBALL",
	2: "TENNIS",
}
var Reply_Entry_Game_value = map[string]int32{
	"FOOTBALL": 1,
	"TENNIS":   2,
}

func (x Reply_Entry_Game) Enum() *Reply_Entry_Game {
	p := new(Reply_Entry_Game)
	*p = x
	return p
}
func (x Reply_Entry_Game) String() string {
	return proto.EnumName(Reply_Entry_Game_name, int32(x))
}
func (x Reply_Entry_Game) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Reply_Entry_Game) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Reply_Entry_Game_value, data, "Reply_Entry_Game")
	if err != nil {
		return err
	}
	*x = Reply_Entry_Game(value)
	return nil
}

// This is a message that might be sent somewhere.
type Request struct {
	Key []int64 `protobuf:"varint,1,rep,name=key" json:"key,omitempty"`
	//  optional imp.ImportedMessage imported_message = 2;
	Hue *Request_Color `protobuf:"varint,3,opt,name=hue,enum=my.test.Request_Color" json:"hue,omitempty"`
	Hat *HatType       `protobuf:"varint,4,opt,name=hat,enum=my.test.HatType,def=1" json:"hat,omitempty"`
	//  optional imp.ImportedMessage.Owner owner = 6;
	Deadline         *float32           `protobuf:"fixed32,7,opt,name=deadline,def=inf" json:"deadline,omitempty"`
	Somegroup        *Request_SomeGroup `protobuf:"group,8,opt,name=SomeGroup" json:"somegroup,omitempty"`
	Reset_           *int32             `protobuf:"varint,12,opt,name=reset" json:"reset,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

const Default_Request_Hat HatType = HatType_FEDORA

var Default_Request_Deadline float32 = float32(math.Inf(1))

func (m *Request) GetKey() []int64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request) GetHue() Request_Color {
	if m != nil && m.Hue != nil {
		return *m.Hue
	}
	return 0
}

func (m *Request) GetHat() HatType {
	if m != nil && m.Hat != nil {
		return *m.Hat
	}
	return Default_Request_Hat
}

func (m *Request) GetDeadline() float32 {
	if m != nil && m.Deadline != nil {
		return *m.Deadline
	}
	return Default_Request_Deadline
}

func (m *Request) GetSomegroup() *Request_SomeGroup {
	if m != nil {
		return m.Somegroup
	}
	return nil
}

func (m *Request) GetReset_() int32 {
	if m != nil && m.Reset_ != nil {
		return *m.Reset_
	}
	return 0
}

type Request_SomeGroup struct {
	GroupField       *int32 `protobuf:"varint,9,opt,name=group_field" json:"group_field,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Request_SomeGroup) Reset()         { *m = Request_SomeGroup{} }
func (m *Request_SomeGroup) String() string { return proto.CompactTextString(m) }
func (*Request_SomeGroup) ProtoMessage()    {}

func (m *Request_SomeGroup) GetGroupField() int32 {
	if m != nil && m.GroupField != nil {
		return *m.GroupField
	}
	return 0
}

type Reply struct {
	Found            []*Reply_Entry            `protobuf:"bytes,1,rep,name=found" json:"found,omitempty"`
	CompactKeys      []int32                   `protobuf:"varint,2,rep,packed,name=compact_keys" json:"compact_keys,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}

var extRange_Reply = []proto.ExtensionRange{
	{100, 536870911},
}

func (*Reply) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Reply
}
func (m *Reply) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Reply) GetFound() []*Reply_Entry {
	if m != nil {
		return m.Found
	}
	return nil
}

func (m *Reply) GetCompactKeys() []int32 {
	if m != nil {
		return m.CompactKeys
	}
	return nil
}

type Reply_Entry struct {
	KeyThatNeeds_1234Camel_CasIng *int64 `protobuf:"varint,1,req,name=key_that_needs_1234camel_CasIng" json:"key_that_needs_1234camel_CasIng,omitempty"`
	Value                         *int64 `protobuf:"varint,2,opt,name=value,def=7" json:"value,omitempty"`
	XMyFieldName_2                *int64 `protobuf:"varint,3,opt,name=_my_field_name_2" json:"_my_field_name_2,omitempty"`
	XXX_unrecognized              []byte `json:"-"`
}

func (m *Reply_Entry) Reset()         { *m = Reply_Entry{} }
func (m *Reply_Entry) String() string { return proto.CompactTextString(m) }
func (*Reply_Entry) ProtoMessage()    {}

const Default_Reply_Entry_Value int64 = 7

func (m *Reply_Entry) GetKeyThatNeeds_1234Camel_CasIng() int64 {
	if m != nil && m.KeyThatNeeds_1234Camel_CasIng != nil {
		return *m.KeyThatNeeds_1234Camel_CasIng
	}
	return 0
}

func (m *Reply_Entry) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return Default_Reply_Entry_Value
}

func (m *Reply_Entry) GetXMyFieldName_2() int64 {
	if m != nil && m.XMyFieldName_2 != nil {
		return *m.XMyFieldName_2
	}
	return 0
}

type ReplyExtensions struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ReplyExtensions) Reset()         { *m = ReplyExtensions{} }
func (m *ReplyExtensions) String() string { return proto.CompactTextString(m) }
func (*ReplyExtensions) ProtoMessage()    {}

var E_ReplyExtensions_Time = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*float64)(nil),
	Field:         101,
	Name:          "my.test.ReplyExtensions.time",
	Tag:           "fixed64,101,opt,name=time",
}

type OldReply struct {
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *OldReply) Reset()         { *m = OldReply{} }
func (m *OldReply) String() string { return proto.CompactTextString(m) }
func (*OldReply) ProtoMessage()    {}

func (m *OldReply) Marshal() ([]byte, error) {
	return proto.MarshalMessageSet(m.ExtensionMap())
}
func (m *OldReply) Unmarshal(buf []byte) error {
	return proto.UnmarshalMessageSet(buf, m.ExtensionMap())
}

// ensure OldReply satisfies proto.Marshaler and proto.Unmarshaler
var _ proto.Marshaler = (*OldReply)(nil)
var _ proto.Unmarshaler = (*OldReply)(nil)

var extRange_OldReply = []proto.ExtensionRange{
	{100, 2147483646},
}

func (*OldReply) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_OldReply
}
func (m *OldReply) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

var E_Tag = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*string)(nil),
	Field:         103,
	Name:          "my.test.tag",
	Tag:           "bytes,103,opt,name=tag",
}

func init() {
	proto.RegisterEnum("my.test.HatType", HatType_name, HatType_value)
	proto.RegisterEnum("my.test.Days", Days_name, Days_value)
	proto.RegisterEnum("my.test.Request_Color", Request_Color_name, Request_Color_value)
	proto.RegisterEnum("my.test.Reply_Entry_Game", Reply_Entry_Game_name, Reply_Entry_Game_value)
	proto.RegisterExtension(E_ReplyExtensions_Time)
	proto.RegisterExtension(E_Tag)
}
