// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: square/eds/byzantine/pb/share.proto

package share_eds_byzantine_pb

import (
	fmt "fmt"
	pb "github.com/celestiaorg/nmt/pb"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Axis int32

const (
	Axis_ROW Axis = 0
	Axis_COL Axis = 1
)

var Axis_name = map[int32]string{
	0: "ROW",
	1: "COL",
}

var Axis_value = map[string]int32{
	"ROW": 0,
	"COL": 1,
}

func (x Axis) String() string {
	return proto.EnumName(Axis_name, int32(x))
}

func (Axis) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5c467604cd602917, []int{0}
}

type Share struct {
	Data      []byte    `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	Proof     *pb.Proof `protobuf:"bytes,2,opt,name=Proof,proto3" json:"Proof,omitempty"`
	ProofAxis Axis      `protobuf:"varint,3,opt,name=ProofAxis,proto3,enum=share.eds.byzantine.pb.Axis" json:"ProofAxis,omitempty"`
}

func (m *Share) Reset()         { *m = Share{} }
func (m *Share) String() string { return proto.CompactTextString(m) }
func (*Share) ProtoMessage()    {}
func (*Share) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c467604cd602917, []int{0}
}
func (m *Share) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Share) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Share.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Share) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Share.Merge(m, src)
}
func (m *Share) XXX_Size() int {
	return m.Size()
}
func (m *Share) XXX_DiscardUnknown() {
	xxx_messageInfo_Share.DiscardUnknown(m)
}

var xxx_messageInfo_Share proto.InternalMessageInfo

func (m *Share) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Share) GetProof() *pb.Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *Share) GetProofAxis() Axis {
	if m != nil {
		return m.ProofAxis
	}
	return Axis_ROW
}

type BadEncoding struct {
	HeaderHash []byte   `protobuf:"bytes,1,opt,name=HeaderHash,proto3" json:"HeaderHash,omitempty"`
	Height     uint64   `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	Shares     []*Share `protobuf:"bytes,3,rep,name=Shares,proto3" json:"Shares,omitempty"`
	Index      uint32   `protobuf:"varint,4,opt,name=Index,proto3" json:"Index,omitempty"`
	Axis       Axis     `protobuf:"varint,5,opt,name=Axis,proto3,enum=share.eds.byzantine.pb.Axis" json:"Axis,omitempty"`
}

func (m *BadEncoding) Reset()         { *m = BadEncoding{} }
func (m *BadEncoding) String() string { return proto.CompactTextString(m) }
func (*BadEncoding) ProtoMessage()    {}
func (*BadEncoding) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c467604cd602917, []int{1}
}
func (m *BadEncoding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BadEncoding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BadEncoding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BadEncoding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadEncoding.Merge(m, src)
}
func (m *BadEncoding) XXX_Size() int {
	return m.Size()
}
func (m *BadEncoding) XXX_DiscardUnknown() {
	xxx_messageInfo_BadEncoding.DiscardUnknown(m)
}

var xxx_messageInfo_BadEncoding proto.InternalMessageInfo

func (m *BadEncoding) GetHeaderHash() []byte {
	if m != nil {
		return m.HeaderHash
	}
	return nil
}

func (m *BadEncoding) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BadEncoding) GetShares() []*Share {
	if m != nil {
		return m.Shares
	}
	return nil
}

func (m *BadEncoding) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BadEncoding) GetAxis() Axis {
	if m != nil {
		return m.Axis
	}
	return Axis_ROW
}

func init() {
	proto.RegisterEnum("share.eds.byzantine.pb.Axis", Axis_name, Axis_value)
	proto.RegisterType((*Share)(nil), "share.eds.byzantine.pb.Share")
	proto.RegisterType((*BadEncoding)(nil), "share.eds.byzantine.pb.BadEncoding")
}

func init() {
	proto.RegisterFile("square/eds/byzantine/pb/share.proto", fileDescriptor_5c467604cd602917)
}

var fileDescriptor_5c467604cd602917 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4b, 0xfb, 0x30,
	0x1c, 0xc5, 0x9b, 0xdf, 0xda, 0xfd, 0xf0, 0x3b, 0x9d, 0x23, 0xc8, 0x08, 0xa2, 0xa1, 0x4c, 0x84,
	0xe2, 0x21, 0x95, 0x89, 0x17, 0x6f, 0x4e, 0x85, 0x09, 0xc2, 0x24, 0x1e, 0x3c, 0x27, 0x26, 0x6e,
	0xbd, 0xb4, 0xb5, 0xa9, 0x50, 0x3d, 0xf8, 0x37, 0xf8, 0x47, 0x79, 0xf0, 0xb8, 0xa3, 0x47, 0x69,
	0xff, 0x11, 0x69, 0x5a, 0xd4, 0x83, 0x82, 0xb7, 0xf7, 0x5e, 0x3e, 0x21, 0xef, 0x11, 0xd8, 0x31,
	0x77, 0xf7, 0x22, 0xd3, 0xa1, 0x56, 0x26, 0x94, 0x0f, 0x8f, 0x22, 0xce, 0xa3, 0x58, 0x87, 0xa9,
	0x0c, 0xcd, 0x42, 0x64, 0x9a, 0xa5, 0x59, 0x92, 0x27, 0x78, 0xd8, 0x18, 0xad, 0x0c, 0xfb, 0x64,
	0x58, 0x2a, 0x37, 0xfb, 0xa9, 0x0c, 0xd3, 0x2c, 0x49, 0x6e, 0x1b, 0x6e, 0xf4, 0x04, 0xde, 0x55,
	0x4d, 0x62, 0x0c, 0xee, 0xa9, 0xc8, 0x05, 0x41, 0x3e, 0x0a, 0x56, 0xb9, 0xd5, 0x78, 0x17, 0xbc,
	0xcb, 0x9a, 0x25, 0xff, 0x7c, 0x14, 0xf4, 0xc6, 0xeb, 0xac, 0xbd, 0x29, 0x99, 0x8d, 0x79, 0x73,
	0x8a, 0x8f, 0x60, 0xc5, 0x8a, 0xe3, 0x22, 0x32, 0xa4, 0xe3, 0xa3, 0xa0, 0x3f, 0xde, 0x62, 0x3f,
	0xbf, 0xcf, 0x44, 0x11, 0x19, 0xfe, 0x85, 0x8f, 0x5e, 0x10, 0xf4, 0x26, 0x42, 0x9d, 0xc5, 0x37,
	0x89, 0x8a, 0xe2, 0x39, 0xa6, 0x00, 0x53, 0x2d, 0x94, 0xce, 0xa6, 0xc2, 0x2c, 0xda, 0x32, 0xdf,
	0x12, 0x3c, 0x84, 0xee, 0x54, 0x47, 0xf3, 0x45, 0x6e, 0x3b, 0xb9, 0xbc, 0x75, 0xf8, 0x10, 0xba,
	0x76, 0x47, 0x5d, 0xa0, 0x13, 0xf4, 0xc6, 0xdb, 0xbf, 0x15, 0xb0, 0x14, 0x6f, 0x61, 0xbc, 0x01,
	0xde, 0x79, 0xac, 0x74, 0x41, 0x5c, 0x1f, 0x05, 0x6b, 0xbc, 0x31, 0x78, 0x1f, 0x5c, 0xbb, 0xc5,
	0xfb, 0xc3, 0x16, 0x4b, 0xee, 0x11, 0x70, 0x6b, 0x87, 0xff, 0x43, 0x87, 0xcf, 0xae, 0x07, 0x4e,
	0x2d, 0x4e, 0x66, 0x17, 0x03, 0x34, 0x21, 0xaf, 0x25, 0x45, 0xcb, 0x92, 0xa2, 0xf7, 0x92, 0xa2,
	0xe7, 0x8a, 0x3a, 0xcb, 0x8a, 0x3a, 0x6f, 0x15, 0x75, 0x64, 0xd7, 0xfe, 0xc0, 0xc1, 0x47, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x7c, 0x07, 0xc6, 0x14, 0xd0, 0x01, 0x00, 0x00,
}

func (m *Share) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Share) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Share) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProofAxis != 0 {
		i = encodeVarintShare(dAtA, i, uint64(m.ProofAxis))
		i--
		dAtA[i] = 0x18
	}
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintShare(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintShare(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BadEncoding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BadEncoding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BadEncoding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Axis != 0 {
		i = encodeVarintShare(dAtA, i, uint64(m.Axis))
		i--
		dAtA[i] = 0x28
	}
	if m.Index != 0 {
		i = encodeVarintShare(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Shares) > 0 {
		for iNdEx := len(m.Shares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Shares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintShare(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Height != 0 {
		i = encodeVarintShare(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.HeaderHash) > 0 {
		i -= len(m.HeaderHash)
		copy(dAtA[i:], m.HeaderHash)
		i = encodeVarintShare(dAtA, i, uint64(len(m.HeaderHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintShare(dAtA []byte, offset int, v uint64) int {
	offset -= sovShare(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Share) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovShare(uint64(l))
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovShare(uint64(l))
	}
	if m.ProofAxis != 0 {
		n += 1 + sovShare(uint64(m.ProofAxis))
	}
	return n
}

func (m *BadEncoding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HeaderHash)
	if l > 0 {
		n += 1 + l + sovShare(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovShare(uint64(m.Height))
	}
	if len(m.Shares) > 0 {
		for _, e := range m.Shares {
			l = e.Size()
			n += 1 + l + sovShare(uint64(l))
		}
	}
	if m.Index != 0 {
		n += 1 + sovShare(uint64(m.Index))
	}
	if m.Axis != 0 {
		n += 1 + sovShare(uint64(m.Axis))
	}
	return n
}

func sovShare(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShare(x uint64) (n int) {
	return sovShare(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Share) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Share: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Share: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &pb.Proof{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofAxis", wireType)
			}
			m.ProofAxis = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProofAxis |= Axis(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShare
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
func (m *BadEncoding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BadEncoding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BadEncoding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderHash = append(m.HeaderHash[:0], dAtA[iNdEx:postIndex]...)
			if m.HeaderHash == nil {
				m.HeaderHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Shares = append(m.Shares, &Share{})
			if err := m.Shares[len(m.Shares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Axis", wireType)
			}
			m.Axis = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Axis |= Axis(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShare
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
func skipShare(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShare
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
					return 0, ErrIntOverflowShare
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowShare
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
			if length < 0 {
				return 0, ErrInvalidLengthShare
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupShare
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthShare
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthShare        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShare          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupShare = fmt.Errorf("proto: unexpected end of group")
)