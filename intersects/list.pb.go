// Code generated by protoc-gen-gogo.
// source: list.proto
// DO NOT EDIT!

/*
	Package intersect is a generated protocol buffer package.

	It is generated from these files:
		list.proto

	It has these top-level messages:
		DeltaList
		FixedList
*/
package intersect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeltaList struct {
	BucketSize int32    `protobuf:"varint,1,opt,name=bucket_size,json=bucketSize,proto3" json:"bucket_size,omitempty"`
	Buckets    []uint64 `protobuf:"fixed64,2,rep,packed,name=buckets" json:"buckets,omitempty"`
	Uids       []uint64 `protobuf:"varint,3,rep,packed,name=uids" json:"uids,omitempty"`
}

func (m *DeltaList) Reset()                    { *m = DeltaList{} }
func (m *DeltaList) String() string            { return proto.CompactTextString(m) }
func (*DeltaList) ProtoMessage()               {}
func (*DeltaList) Descriptor() ([]byte, []int) { return fileDescriptorList, []int{0} }

func (m *DeltaList) GetBucketSize() int32 {
	if m != nil {
		return m.BucketSize
	}
	return 0
}

func (m *DeltaList) GetBuckets() []uint64 {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *DeltaList) GetUids() []uint64 {
	if m != nil {
		return m.Uids
	}
	return nil
}

type FixedList struct {
	Uids []uint64 `protobuf:"fixed64,1,rep,packed,name=uids" json:"uids,omitempty"`
}

func (m *FixedList) Reset()                    { *m = FixedList{} }
func (m *FixedList) String() string            { return proto.CompactTextString(m) }
func (*FixedList) ProtoMessage()               {}
func (*FixedList) Descriptor() ([]byte, []int) { return fileDescriptorList, []int{1} }

func (m *FixedList) GetUids() []uint64 {
	if m != nil {
		return m.Uids
	}
	return nil
}

func init() {
	proto.RegisterType((*DeltaList)(nil), "intersect.DeltaList")
	proto.RegisterType((*FixedList)(nil), "intersect.FixedList")
}
func (m *DeltaList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeltaList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BucketSize != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintList(dAtA, i, uint64(m.BucketSize))
	}
	if len(m.Buckets) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintList(dAtA, i, uint64(len(m.Buckets)*8))
		for _, num := range m.Buckets {
			dAtA[i] = uint8(num)
			i++
			dAtA[i] = uint8(num >> 8)
			i++
			dAtA[i] = uint8(num >> 16)
			i++
			dAtA[i] = uint8(num >> 24)
			i++
			dAtA[i] = uint8(num >> 32)
			i++
			dAtA[i] = uint8(num >> 40)
			i++
			dAtA[i] = uint8(num >> 48)
			i++
			dAtA[i] = uint8(num >> 56)
			i++
		}
	}
	if len(m.Uids) > 0 {
		dAtA2 := make([]byte, len(m.Uids)*10)
		var j1 int
		for _, num := range m.Uids {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintList(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	return i, nil
}

func (m *FixedList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FixedList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uids) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintList(dAtA, i, uint64(len(m.Uids)*8))
		for _, num := range m.Uids {
			dAtA[i] = uint8(num)
			i++
			dAtA[i] = uint8(num >> 8)
			i++
			dAtA[i] = uint8(num >> 16)
			i++
			dAtA[i] = uint8(num >> 24)
			i++
			dAtA[i] = uint8(num >> 32)
			i++
			dAtA[i] = uint8(num >> 40)
			i++
			dAtA[i] = uint8(num >> 48)
			i++
			dAtA[i] = uint8(num >> 56)
			i++
		}
	}
	return i, nil
}

func encodeFixed64List(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32List(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintList(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DeltaList) Size() (n int) {
	var l int
	_ = l
	if m.BucketSize != 0 {
		n += 1 + sovList(uint64(m.BucketSize))
	}
	if len(m.Buckets) > 0 {
		n += 1 + sovList(uint64(len(m.Buckets)*8)) + len(m.Buckets)*8
	}
	if len(m.Uids) > 0 {
		l = 0
		for _, e := range m.Uids {
			l += sovList(uint64(e))
		}
		n += 1 + sovList(uint64(l)) + l
	}
	return n
}

func (m *FixedList) Size() (n int) {
	var l int
	_ = l
	if len(m.Uids) > 0 {
		n += 1 + sovList(uint64(len(m.Uids)*8)) + len(m.Uids)*8
	}
	return n
}

func sovList(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozList(x uint64) (n int) {
	return sovList(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeltaList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowList
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
			return fmt.Errorf("proto: DeltaList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeltaList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BucketSize", wireType)
			}
			m.BucketSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowList
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BucketSize |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				iNdEx += 8
				v = uint64(dAtA[iNdEx-8])
				v |= uint64(dAtA[iNdEx-7]) << 8
				v |= uint64(dAtA[iNdEx-6]) << 16
				v |= uint64(dAtA[iNdEx-5]) << 24
				v |= uint64(dAtA[iNdEx-4]) << 32
				v |= uint64(dAtA[iNdEx-3]) << 40
				v |= uint64(dAtA[iNdEx-2]) << 48
				v |= uint64(dAtA[iNdEx-1]) << 56
				m.Buckets = append(m.Buckets, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowList
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthList
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					iNdEx += 8
					v = uint64(dAtA[iNdEx-8])
					v |= uint64(dAtA[iNdEx-7]) << 8
					v |= uint64(dAtA[iNdEx-6]) << 16
					v |= uint64(dAtA[iNdEx-5]) << 24
					v |= uint64(dAtA[iNdEx-4]) << 32
					v |= uint64(dAtA[iNdEx-3]) << 40
					v |= uint64(dAtA[iNdEx-2]) << 48
					v |= uint64(dAtA[iNdEx-1]) << 56
					m.Buckets = append(m.Buckets, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Buckets", wireType)
			}
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowList
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Uids = append(m.Uids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowList
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthList
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowList
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Uids = append(m.Uids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Uids", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipList(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthList
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
func (m *FixedList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowList
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
			return fmt.Errorf("proto: FixedList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FixedList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				iNdEx += 8
				v = uint64(dAtA[iNdEx-8])
				v |= uint64(dAtA[iNdEx-7]) << 8
				v |= uint64(dAtA[iNdEx-6]) << 16
				v |= uint64(dAtA[iNdEx-5]) << 24
				v |= uint64(dAtA[iNdEx-4]) << 32
				v |= uint64(dAtA[iNdEx-3]) << 40
				v |= uint64(dAtA[iNdEx-2]) << 48
				v |= uint64(dAtA[iNdEx-1]) << 56
				m.Uids = append(m.Uids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowList
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthList
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					iNdEx += 8
					v = uint64(dAtA[iNdEx-8])
					v |= uint64(dAtA[iNdEx-7]) << 8
					v |= uint64(dAtA[iNdEx-6]) << 16
					v |= uint64(dAtA[iNdEx-5]) << 24
					v |= uint64(dAtA[iNdEx-4]) << 32
					v |= uint64(dAtA[iNdEx-3]) << 40
					v |= uint64(dAtA[iNdEx-2]) << 48
					v |= uint64(dAtA[iNdEx-1]) << 56
					m.Uids = append(m.Uids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Uids", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipList(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthList
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
func skipList(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowList
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
					return 0, ErrIntOverflowList
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
					return 0, ErrIntOverflowList
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
				return 0, ErrInvalidLengthList
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowList
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
				next, err := skipList(dAtA[start:])
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
	ErrInvalidLengthList = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowList   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("list.proto", fileDescriptorList) }

var fileDescriptorList = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xc9, 0x2c, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0xcc, 0x2b, 0x49, 0x2d, 0x2a, 0x4e, 0x4d,
	0x2e, 0x51, 0x8a, 0xe2, 0xe2, 0x74, 0x49, 0xcd, 0x29, 0x49, 0xf4, 0xc9, 0x2c, 0x2e, 0x11, 0x92,
	0xe7, 0xe2, 0x4e, 0x2a, 0x4d, 0xce, 0x4e, 0x2d, 0x89, 0x2f, 0xce, 0xac, 0x4a, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0x60, 0x0d, 0xe2, 0x82, 0x08, 0x05, 0x67, 0x56, 0xa5, 0x0a, 0x49, 0x70, 0xb1, 0x43,
	0x78, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0x6c, 0x41, 0x30, 0xae, 0x90, 0x10, 0x17, 0x4b, 0x69,
	0x66, 0x4a, 0xb1, 0x04, 0xb3, 0x02, 0xb3, 0x06, 0x4b, 0x10, 0x98, 0xad, 0x24, 0xcf, 0xc5, 0xe9,
	0x96, 0x59, 0x91, 0x9a, 0x02, 0x36, 0x1b, 0xa6, 0x80, 0x11, 0xac, 0x0f, 0xcc, 0x76, 0x12, 0x38,
	0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63,
	0x48, 0x62, 0x03, 0x3b, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x59, 0x44, 0xd8, 0xae,
	0x00, 0x00, 0x00,
}