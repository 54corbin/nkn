// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/node.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SyncState int32

const (
	SyncStarted     SyncState = 0
	SyncFinished    SyncState = 1
	PersistFinished SyncState = 2
	WaitForSyncing  SyncState = 3
)

var SyncState_name = map[int32]string{
	0: "SyncStarted",
	1: "SyncFinished",
	2: "PersistFinished",
	3: "WaitForSyncing",
}
var SyncState_value = map[string]int32{
	"SyncStarted":     0,
	"SyncFinished":    1,
	"PersistFinished": 2,
	"WaitForSyncing":  3,
}

func (SyncState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_node_cfbb661efc0650a7, []int{0}
}

type NodeData struct {
	PublicKey       []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	WebsocketPort   uint32 `protobuf:"varint,2,opt,name=websocket_port,json=websocketPort,proto3" json:"websocket_port,omitempty"`
	JsonRpcPort     uint32 `protobuf:"varint,3,opt,name=json_rpc_port,json=jsonRpcPort,proto3" json:"json_rpc_port,omitempty"`
	ProtocolVersion uint32 `protobuf:"varint,5,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
}

func (m *NodeData) Reset()      { *m = NodeData{} }
func (*NodeData) ProtoMessage() {}
func (*NodeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_cfbb661efc0650a7, []int{0}
}
func (m *NodeData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *NodeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeData.Merge(dst, src)
}
func (m *NodeData) XXX_Size() int {
	return m.Size()
}
func (m *NodeData) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeData.DiscardUnknown(m)
}

var xxx_messageInfo_NodeData proto.InternalMessageInfo

func (m *NodeData) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *NodeData) GetWebsocketPort() uint32 {
	if m != nil {
		return m.WebsocketPort
	}
	return 0
}

func (m *NodeData) GetJsonRpcPort() uint32 {
	if m != nil {
		return m.JsonRpcPort
	}
	return 0
}

func (m *NodeData) GetProtocolVersion() uint32 {
	if m != nil {
		return m.ProtocolVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*NodeData)(nil), "pb.NodeData")
	proto.RegisterEnum("pb.SyncState", SyncState_name, SyncState_value)
}
func (x SyncState) String() string {
	s, ok := SyncState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *NodeData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NodeData)
	if !ok {
		that2, ok := that.(NodeData)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.PublicKey, that1.PublicKey) {
		return false
	}
	if this.WebsocketPort != that1.WebsocketPort {
		return false
	}
	if this.JsonRpcPort != that1.JsonRpcPort {
		return false
	}
	if this.ProtocolVersion != that1.ProtocolVersion {
		return false
	}
	return true
}
func (this *NodeData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&pb.NodeData{")
	s = append(s, "PublicKey: "+fmt.Sprintf("%#v", this.PublicKey)+",\n")
	s = append(s, "WebsocketPort: "+fmt.Sprintf("%#v", this.WebsocketPort)+",\n")
	s = append(s, "JsonRpcPort: "+fmt.Sprintf("%#v", this.JsonRpcPort)+",\n")
	s = append(s, "ProtocolVersion: "+fmt.Sprintf("%#v", this.ProtocolVersion)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringNode(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *NodeData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PublicKey) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNode(dAtA, i, uint64(len(m.PublicKey)))
		i += copy(dAtA[i:], m.PublicKey)
	}
	if m.WebsocketPort != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.WebsocketPort))
	}
	if m.JsonRpcPort != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.JsonRpcPort))
	}
	if m.ProtocolVersion != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.ProtocolVersion))
	}
	return i, nil
}

func encodeVarintNode(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedNodeData(r randyNode, easy bool) *NodeData {
	this := &NodeData{}
	v1 := r.Intn(100)
	this.PublicKey = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.PublicKey[i] = byte(r.Intn(256))
	}
	this.WebsocketPort = uint32(r.Uint32())
	this.JsonRpcPort = uint32(r.Uint32())
	this.ProtocolVersion = uint32(r.Uint32())
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyNode interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneNode(r randyNode) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringNode(r randyNode) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneNode(r)
	}
	return string(tmps)
}
func randUnrecognizedNode(r randyNode, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldNode(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldNode(dAtA []byte, r randyNode, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateNode(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateNode(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateNode(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateNode(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateNode(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateNode(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateNode(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *NodeData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if m.WebsocketPort != 0 {
		n += 1 + sovNode(uint64(m.WebsocketPort))
	}
	if m.JsonRpcPort != 0 {
		n += 1 + sovNode(uint64(m.JsonRpcPort))
	}
	if m.ProtocolVersion != 0 {
		n += 1 + sovNode(uint64(m.ProtocolVersion))
	}
	return n
}

func sovNode(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNode(x uint64) (n int) {
	return sovNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *NodeData) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NodeData{`,
		`PublicKey:` + fmt.Sprintf("%v", this.PublicKey) + `,`,
		`WebsocketPort:` + fmt.Sprintf("%v", this.WebsocketPort) + `,`,
		`JsonRpcPort:` + fmt.Sprintf("%v", this.JsonRpcPort) + `,`,
		`ProtocolVersion:` + fmt.Sprintf("%v", this.ProtocolVersion) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringNode(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *NodeData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = append(m.PublicKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PublicKey == nil {
				m.PublicKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WebsocketPort", wireType)
			}
			m.WebsocketPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WebsocketPort |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JsonRpcPort", wireType)
			}
			m.JsonRpcPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JsonRpcPort |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolVersion", wireType)
			}
			m.ProtocolVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProtocolVersion |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
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
func skipNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNode
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowNode
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowNode
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthNode
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNode
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipNode(dAtA[start:])
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
	ErrInvalidLengthNode = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNode   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pb/node.proto", fileDescriptor_node_cfbb661efc0650a7) }

var fileDescriptor_node_cfbb661efc0650a7 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xbf, 0x4e, 0x2a, 0x41,
	0x14, 0xc6, 0xf7, 0x40, 0xee, 0xcd, 0x65, 0x60, 0x61, 0x33, 0xb7, 0x21, 0x26, 0x9e, 0x10, 0x12,
	0x13, 0x34, 0x11, 0x0a, 0x5b, 0x2b, 0x63, 0x68, 0x4c, 0x0c, 0x81, 0x44, 0xca, 0xcd, 0xfe, 0x19,
	0x97, 0x15, 0xdc, 0xb3, 0x99, 0x1d, 0x34, 0x74, 0x3e, 0x82, 0x6f, 0x60, 0xeb, 0x23, 0xf8, 0x08,
	0x96, 0x94, 0x94, 0xee, 0xd0, 0x58, 0x52, 0x5a, 0x9a, 0x9d, 0x55, 0xba, 0xf3, 0xfd, 0xce, 0xef,
	0xcb, 0x39, 0xcc, 0x4e, 0xfd, 0x41, 0x42, 0xa1, 0xe8, 0xa7, 0x92, 0x14, 0xf1, 0x4a, 0xea, 0x1f,
	0x9c, 0x46, 0xb1, 0x9a, 0x2d, 0xfd, 0x7e, 0x40, 0xf7, 0x83, 0x88, 0x22, 0x1a, 0x98, 0x95, 0xbf,
	0xbc, 0x35, 0xc9, 0x04, 0x33, 0x95, 0x95, 0xee, 0x0b, 0xb0, 0x7f, 0xd7, 0x14, 0x8a, 0x4b, 0x4f,
	0x79, 0xfc, 0x90, 0xb1, 0x74, 0xe9, 0x2f, 0xe2, 0xc0, 0x9d, 0x8b, 0x55, 0x1b, 0x3a, 0xd0, 0x6b,
	0x8c, 0x6b, 0x25, 0xb9, 0x12, 0x2b, 0x7e, 0xc4, 0x9a, 0x8f, 0xc2, 0xcf, 0x28, 0x98, 0x0b, 0xe5,
	0xa6, 0x24, 0x55, 0xbb, 0xd2, 0x81, 0x9e, 0x3d, 0xb6, 0xf7, 0x74, 0x44, 0x52, 0xf1, 0x2e, 0xb3,
	0xef, 0x32, 0x4a, 0x5c, 0x99, 0x06, 0xa5, 0x55, 0x35, 0x56, 0xbd, 0x80, 0xe3, 0x34, 0x30, 0xce,
	0x31, 0x73, 0xcc, 0xfd, 0x80, 0x16, 0xee, 0x83, 0x90, 0x59, 0x4c, 0x49, 0xfb, 0x8f, 0xd1, 0x5a,
	0xbf, 0xfc, 0xa6, 0xc4, 0x27, 0x53, 0x56, 0x9b, 0xac, 0x92, 0x60, 0xa2, 0x3c, 0x25, 0x78, 0x8b,
	0xd5, 0x7f, 0x82, 0x54, 0x22, 0x74, 0x2c, 0xee, 0xb0, 0x46, 0x01, 0x86, 0x71, 0x12, 0x67, 0x33,
	0x11, 0x3a, 0xc0, 0xff, 0xb3, 0xd6, 0xa8, 0xa8, 0x66, 0x6a, 0x0f, 0x2b, 0x9c, 0xb3, 0xe6, 0xd4,
	0x8b, 0xd5, 0x90, 0x64, 0x61, 0xc7, 0x49, 0xe4, 0x54, 0x2f, 0xce, 0xd7, 0x39, 0x5a, 0x9b, 0x1c,
	0xad, 0x5d, 0x8e, 0xf0, 0x95, 0x23, 0x3c, 0x69, 0x84, 0x57, 0x8d, 0xf0, 0xa6, 0x11, 0xde, 0x35,
	0xc2, 0x5a, 0x23, 0x7c, 0x68, 0x84, 0x4f, 0x8d, 0xd6, 0x4e, 0x23, 0x3c, 0x6f, 0xd1, 0x5a, 0x6f,
	0xd1, 0xda, 0x6c, 0xd1, 0xf2, 0xff, 0x9a, 0x3f, 0xcf, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xaa,
	0xdd, 0x8e, 0x37, 0x83, 0x01, 0x00, 0x00,
}
