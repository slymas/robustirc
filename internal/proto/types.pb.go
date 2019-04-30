// Code generated by protoc-gen-gogo.
// source: types.proto
// DO NOT EDIT!

/*
	Package proto is a generated protocol buffer package.

	It is generated from these files:
		types.proto
		snapshot.proto

	It has these top-level messages:
		RobustId
		RobustMessage
		RaftLog
		Timestamp
		Snapshot
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto1.ProtoPackageIsVersion1

type RobustMessage_RobustType int32

const (
	RobustMessage_CREATE_SESSION   RobustMessage_RobustType = 0
	RobustMessage_DELETE_SESSION   RobustMessage_RobustType = 1
	RobustMessage_IRC_FROM_CLIENT  RobustMessage_RobustType = 2
	RobustMessage_IRC_TO_CLIENT    RobustMessage_RobustType = 3
	RobustMessage_PING             RobustMessage_RobustType = 4
	RobustMessage_MESSAGE_OF_DEATH RobustMessage_RobustType = 5
	RobustMessage_CONFIG           RobustMessage_RobustType = 6
	RobustMessage_STATE            RobustMessage_RobustType = 7
	RobustMessage_ANY              RobustMessage_RobustType = 8
)

var RobustMessage_RobustType_name = map[int32]string{
	0: "CREATE_SESSION",
	1: "DELETE_SESSION",
	2: "IRC_FROM_CLIENT",
	3: "IRC_TO_CLIENT",
	4: "PING",
	5: "MESSAGE_OF_DEATH",
	6: "CONFIG",
	7: "STATE",
	8: "ANY",
}
var RobustMessage_RobustType_value = map[string]int32{
	"CREATE_SESSION":   0,
	"DELETE_SESSION":   1,
	"IRC_FROM_CLIENT":  2,
	"IRC_TO_CLIENT":    3,
	"PING":             4,
	"MESSAGE_OF_DEATH": 5,
	"CONFIG":           6,
	"STATE":            7,
	"ANY":              8,
}

func (x RobustMessage_RobustType) String() string {
	return proto1.EnumName(RobustMessage_RobustType_name, int32(x))
}
func (RobustMessage_RobustType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorTypes, []int{1, 0}
}

type RaftLog_LogType int32

const (
	RaftLog_COMMAND    RaftLog_LogType = 0
	RaftLog_NOOP       RaftLog_LogType = 1
	RaftLog_ADDPEER    RaftLog_LogType = 2
	RaftLog_REMOVEPEER RaftLog_LogType = 3
	RaftLog_BARRIER    RaftLog_LogType = 4
)

var RaftLog_LogType_name = map[int32]string{
	0: "COMMAND",
	1: "NOOP",
	2: "ADDPEER",
	3: "REMOVEPEER",
	4: "BARRIER",
}
var RaftLog_LogType_value = map[string]int32{
	"COMMAND":    0,
	"NOOP":       1,
	"ADDPEER":    2,
	"REMOVEPEER": 3,
	"BARRIER":    4,
}

func (x RaftLog_LogType) String() string {
	return proto1.EnumName(RaftLog_LogType_name, int32(x))
}
func (RaftLog_LogType) EnumDescriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2, 0} }

type RobustId struct {
	Id    uint64 `protobuf:"fixed64,1,opt,name=id,proto3" json:"id,omitempty"`
	Reply uint64 `protobuf:"fixed64,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (m *RobustId) Reset()                    { *m = RobustId{} }
func (m *RobustId) String() string            { return proto1.CompactTextString(m) }
func (*RobustId) ProtoMessage()               {}
func (*RobustId) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

type RobustMessage struct {
	Id      *RobustId                `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Session *RobustId                `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Type    RobustMessage_RobustType `protobuf:"varint,3,opt,name=type,proto3,enum=proto.RobustMessage_RobustType" json:"type,omitempty"`
	Data    string                   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// TODO: use oneof for the following to save space?
	UnixNano        int64    `protobuf:"varint,5,opt,name=unix_nano,json=unixNano,proto3" json:"unix_nano,omitempty"`
	Servers         []string `protobuf:"bytes,6,rep,name=servers" json:"servers,omitempty"`
	CurrentMaster   string   `protobuf:"bytes,7,opt,name=current_master,json=currentMaster,proto3" json:"current_master,omitempty"`
	ClientMessageId uint64   `protobuf:"varint,8,opt,name=client_message_id,json=clientMessageId,proto3" json:"client_message_id,omitempty"`
	Revision        uint64   `protobuf:"varint,9,opt,name=revision,proto3" json:"revision,omitempty"`
	RemoteAddr      string   `protobuf:"bytes,10,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
}

func (m *RobustMessage) Reset()                    { *m = RobustMessage{} }
func (m *RobustMessage) String() string            { return proto1.CompactTextString(m) }
func (*RobustMessage) ProtoMessage()               {}
func (*RobustMessage) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *RobustMessage) GetId() *RobustId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *RobustMessage) GetSession() *RobustId {
	if m != nil {
		return m.Session
	}
	return nil
}

type RaftLog struct {
	Index      uint64          `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Term       uint64          `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	Type       RaftLog_LogType `protobuf:"varint,3,opt,name=type,proto3,enum=proto.RaftLog_LogType" json:"type,omitempty"`
	Data       []byte          `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Extensions []byte          `protobuf:"bytes,5,opt,name=extensions,proto3" json:"extensions,omitempty"`
}

func (m *RaftLog) Reset()                    { *m = RaftLog{} }
func (m *RaftLog) String() string            { return proto1.CompactTextString(m) }
func (*RaftLog) ProtoMessage()               {}
func (*RaftLog) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

func init() {
	proto1.RegisterType((*RobustId)(nil), "proto.RobustId")
	proto1.RegisterType((*RobustMessage)(nil), "proto.RobustMessage")
	proto1.RegisterType((*RaftLog)(nil), "proto.RaftLog")
	proto1.RegisterEnum("proto.RobustMessage_RobustType", RobustMessage_RobustType_name, RobustMessage_RobustType_value)
	proto1.RegisterEnum("proto.RaftLog_LogType", RaftLog_LogType_name, RaftLog_LogType_value)
}
func (m *RobustId) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RobustId) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		data[i] = 0x9
		i++
		i = encodeFixed64Types(data, i, uint64(m.Id))
	}
	if m.Reply != 0 {
		data[i] = 0x11
		i++
		i = encodeFixed64Types(data, i, uint64(m.Reply))
	}
	return i, nil
}

func (m *RobustMessage) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RobustMessage) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != nil {
		data[i] = 0xa
		i++
		i = encodeVarintTypes(data, i, uint64(m.Id.Size()))
		n1, err := m.Id.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Session != nil {
		data[i] = 0x12
		i++
		i = encodeVarintTypes(data, i, uint64(m.Session.Size()))
		n2, err := m.Session.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Type != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintTypes(data, i, uint64(m.Type))
	}
	if len(m.Data) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintTypes(data, i, uint64(len(m.Data)))
		i += copy(data[i:], m.Data)
	}
	if m.UnixNano != 0 {
		data[i] = 0x28
		i++
		i = encodeVarintTypes(data, i, uint64(m.UnixNano))
	}
	if len(m.Servers) > 0 {
		for _, s := range m.Servers {
			data[i] = 0x32
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if len(m.CurrentMaster) > 0 {
		data[i] = 0x3a
		i++
		i = encodeVarintTypes(data, i, uint64(len(m.CurrentMaster)))
		i += copy(data[i:], m.CurrentMaster)
	}
	if m.ClientMessageId != 0 {
		data[i] = 0x40
		i++
		i = encodeVarintTypes(data, i, uint64(m.ClientMessageId))
	}
	if m.Revision != 0 {
		data[i] = 0x48
		i++
		i = encodeVarintTypes(data, i, uint64(m.Revision))
	}
	if len(m.RemoteAddr) > 0 {
		data[i] = 0x52
		i++
		i = encodeVarintTypes(data, i, uint64(len(m.RemoteAddr)))
		i += copy(data[i:], m.RemoteAddr)
	}
	return i, nil
}

func (m *RaftLog) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftLog) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintTypes(data, i, uint64(m.Index))
	}
	if m.Term != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintTypes(data, i, uint64(m.Term))
	}
	if m.Type != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintTypes(data, i, uint64(m.Type))
	}
	if len(m.Data) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintTypes(data, i, uint64(len(m.Data)))
		i += copy(data[i:], m.Data)
	}
	if len(m.Extensions) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintTypes(data, i, uint64(len(m.Extensions)))
		i += copy(data[i:], m.Extensions)
	}
	return i, nil
}

func encodeFixed64Types(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Types(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTypes(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *RobustId) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 9
	}
	if m.Reply != 0 {
		n += 9
	}
	return n
}

func (m *RobustMessage) Size() (n int) {
	var l int
	_ = l
	if m.Id != nil {
		l = m.Id.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Session != nil {
		l = m.Session.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.UnixNano != 0 {
		n += 1 + sovTypes(uint64(m.UnixNano))
	}
	if len(m.Servers) > 0 {
		for _, s := range m.Servers {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.CurrentMaster)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ClientMessageId != 0 {
		n += 1 + sovTypes(uint64(m.ClientMessageId))
	}
	if m.Revision != 0 {
		n += 1 + sovTypes(uint64(m.Revision))
	}
	l = len(m.RemoteAddr)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *RaftLog) Size() (n int) {
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovTypes(uint64(m.Index))
	}
	if m.Term != 0 {
		n += 1 + sovTypes(uint64(m.Term))
	}
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Extensions)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RobustId) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RobustId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RobustId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			m.Id = uint64(data[iNdEx-8])
			m.Id |= uint64(data[iNdEx-7]) << 8
			m.Id |= uint64(data[iNdEx-6]) << 16
			m.Id |= uint64(data[iNdEx-5]) << 24
			m.Id |= uint64(data[iNdEx-4]) << 32
			m.Id |= uint64(data[iNdEx-3]) << 40
			m.Id |= uint64(data[iNdEx-2]) << 48
			m.Id |= uint64(data[iNdEx-1]) << 56
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reply", wireType)
			}
			m.Reply = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			m.Reply = uint64(data[iNdEx-8])
			m.Reply |= uint64(data[iNdEx-7]) << 8
			m.Reply |= uint64(data[iNdEx-6]) << 16
			m.Reply |= uint64(data[iNdEx-5]) << 24
			m.Reply |= uint64(data[iNdEx-4]) << 32
			m.Reply |= uint64(data[iNdEx-3]) << 40
			m.Reply |= uint64(data[iNdEx-2]) << 48
			m.Reply |= uint64(data[iNdEx-1]) << 56
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RobustMessage) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RobustMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RobustMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Id == nil {
				m.Id = &RobustId{}
			}
			if err := m.Id.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Session == nil {
				m.Session = &RobustId{}
			}
			if err := m.Session.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (RobustMessage_RobustType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnixNano", wireType)
			}
			m.UnixNano = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.UnixNano |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Servers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Servers = append(m.Servers, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentMaster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentMaster = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientMessageId", wireType)
			}
			m.ClientMessageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ClientMessageId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Revision |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteAddr = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RaftLog) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RaftLog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftLog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Index |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (RaftLog_LogType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], data[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extensions", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extensions = append(m.Extensions[:0], data[iNdEx:postIndex]...)
			if m.Extensions == nil {
				m.Extensions = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTypes(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorTypes = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xd1, 0x6a, 0x9c, 0x40,
	0x14, 0x86, 0x33, 0xab, 0xbb, 0xba, 0x67, 0x93, 0xcd, 0xe4, 0x34, 0x14, 0x69, 0x61, 0x23, 0x0b,
	0x05, 0x9b, 0x8b, 0x50, 0x92, 0x27, 0x30, 0xbb, 0x93, 0xd4, 0xb2, 0x6a, 0x18, 0xa5, 0xd0, 0x2b,
	0x31, 0x71, 0x1a, 0x84, 0x44, 0x97, 0xd1, 0x84, 0xec, 0x2b, 0xf4, 0x09, 0x7a, 0xd1, 0x07, 0xea,
	0x65, 0xdf, 0xa0, 0x65, 0xfb, 0x22, 0xc5, 0xd1, 0x4d, 0x13, 0x68, 0xaf, 0x9c, 0xff, 0x9b, 0x33,
	0xff, 0x9c, 0xe3, 0x3f, 0x30, 0xaa, 0x57, 0x4b, 0x51, 0x1d, 0x2d, 0x65, 0x59, 0x97, 0xd8, 0x57,
	0x9f, 0xe9, 0x3b, 0x30, 0x79, 0x79, 0x79, 0x57, 0xd5, 0x5e, 0x86, 0x63, 0xe8, 0xe5, 0x99, 0x45,
	0x6c, 0xe2, 0x0c, 0x78, 0x2f, 0xcf, 0x70, 0x1f, 0xfa, 0x52, 0x2c, 0x6f, 0x56, 0x56, 0x4f, 0xa1,
	0x56, 0x4c, 0xbf, 0xe8, 0xb0, 0xd3, 0x1e, 0xf1, 0x45, 0x55, 0xa5, 0xd7, 0x02, 0x0f, 0x1e, 0xcf,
	0x8d, 0x8e, 0x77, 0x5b, 0xfb, 0xa3, 0x8d, 0xa9, 0x32, 0x7a, 0x0b, 0x46, 0x25, 0xaa, 0x2a, 0x2f,
	0x0b, 0x65, 0xf5, 0x8f, 0xaa, 0xcd, 0x3e, 0x9e, 0x80, 0xde, 0x74, 0x69, 0x69, 0x36, 0x71, 0xc6,
	0xc7, 0x07, 0xcf, 0xea, 0xba, 0xfb, 0x3a, 0x15, 0xaf, 0x96, 0x82, 0xab, 0x62, 0x44, 0xd0, 0xb3,
	0xb4, 0x4e, 0x2d, 0xdd, 0x26, 0xce, 0x90, 0xab, 0x35, 0xbe, 0x86, 0xe1, 0x5d, 0x91, 0x3f, 0x24,
	0x45, 0x5a, 0x94, 0x56, 0xdf, 0x26, 0x8e, 0xc6, 0xcd, 0x06, 0x04, 0x69, 0x51, 0xa2, 0xd5, 0x34,
	0x24, 0xef, 0x85, 0xac, 0xac, 0x81, 0xad, 0x39, 0x43, 0xbe, 0x91, 0xf8, 0x06, 0xc6, 0x57, 0x77,
	0x52, 0x8a, 0xa2, 0x4e, 0x6e, 0xd3, 0xaa, 0x16, 0xd2, 0x32, 0x94, 0xe9, 0x4e, 0x47, 0x7d, 0x05,
	0xf1, 0x10, 0xf6, 0xae, 0x6e, 0x72, 0x55, 0xd5, 0x36, 0x95, 0xe4, 0x99, 0x65, 0xda, 0xc4, 0xd1,
	0xf9, 0x6e, 0xbb, 0xd1, 0x35, 0xeb, 0x65, 0xf8, 0x0a, 0x4c, 0x29, 0xee, 0x73, 0x35, 0xfe, 0x50,
	0x95, 0x3c, 0x6a, 0x3c, 0x80, 0x91, 0x14, 0xb7, 0x65, 0x2d, 0x92, 0x34, 0xcb, 0xa4, 0x05, 0xea,
	0x2e, 0x68, 0x91, 0x9b, 0x65, 0x72, 0xfa, 0x8d, 0x00, 0xfc, 0x9d, 0x17, 0x11, 0xc6, 0x33, 0xce,
	0xdc, 0x98, 0x25, 0x11, 0x8b, 0x22, 0x2f, 0x0c, 0xe8, 0x56, 0xc3, 0xe6, 0x6c, 0xc1, 0x9e, 0x30,
	0x82, 0x2f, 0x60, 0xd7, 0xe3, 0xb3, 0xe4, 0x8c, 0x87, 0x7e, 0x32, 0x5b, 0x78, 0x2c, 0x88, 0x69,
	0x0f, 0xf7, 0x60, 0xa7, 0x81, 0x71, 0xb8, 0x41, 0x1a, 0x9a, 0xa0, 0x5f, 0x78, 0xc1, 0x39, 0xd5,
	0x71, 0x1f, 0xa8, 0xcf, 0xa2, 0xc8, 0x3d, 0x67, 0x49, 0x78, 0x96, 0xcc, 0x99, 0x1b, 0xbf, 0xa7,
	0x7d, 0x04, 0x18, 0xcc, 0xc2, 0xe0, 0xcc, 0x3b, 0xa7, 0x03, 0x1c, 0x42, 0x3f, 0x8a, 0xdd, 0x98,
	0x51, 0x03, 0x0d, 0xd0, 0xdc, 0xe0, 0x13, 0x35, 0xa7, 0x3f, 0x09, 0x18, 0x3c, 0xfd, 0x5c, 0x2f,
	0xca, 0xeb, 0xe6, 0xb9, 0xe4, 0x45, 0x26, 0x1e, 0xd4, 0x4b, 0xd0, 0x79, 0x2b, 0x9a, 0x6c, 0x6a,
	0x21, 0x6f, 0x55, 0xf0, 0x3a, 0x57, 0x6b, 0x3c, 0x7c, 0x16, 0xf2, 0xcb, 0x4d, 0xc8, 0xad, 0xcf,
	0xd1, 0xa2, 0xbc, 0xfe, 0x4f, 0xb6, 0xdb, 0x5d, 0xb6, 0x13, 0x00, 0xf1, 0x50, 0x8b, 0xa2, 0xf9,
	0x85, 0x95, 0x0a, 0x77, 0x9b, 0x3f, 0x21, 0xd3, 0x0f, 0x60, 0x74, 0x26, 0x38, 0x02, 0x63, 0x16,
	0xfa, 0xbe, 0x1b, 0xcc, 0xe9, 0x56, 0x33, 0x6d, 0x10, 0x86, 0x17, 0x94, 0x34, 0xd8, 0x9d, 0xcf,
	0x2f, 0x18, 0xe3, 0xb4, 0x87, 0x63, 0x00, 0xce, 0xfc, 0xf0, 0x23, 0x53, 0x5a, 0x6b, 0x36, 0x4f,
	0x5d, 0xce, 0x3d, 0xc6, 0xa9, 0x7e, 0x4a, 0xbf, 0xaf, 0x27, 0xe4, 0xc7, 0x7a, 0x42, 0x7e, 0xad,
	0x27, 0xe4, 0xeb, 0xef, 0xc9, 0xd6, 0xe5, 0x40, 0xb5, 0x7b, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x6e, 0xa3, 0xf1, 0xfb, 0x4f, 0x03, 0x00, 0x00,
}
