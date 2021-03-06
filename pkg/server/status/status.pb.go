// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/status/status.proto

/*
	Package status is a generated protocol buffer package.

	It is generated from these files:
		server/status/status.proto

	It has these top-level messages:
		StoreStatus
		NodeStatus
*/
package status

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_build "github.com/cockroachdb/cockroach/pkg/build"

import github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
import encoding_binary "encoding/binary"

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

// StoreStatus records the most recent values of metrics for a store.
type StoreStatus struct {
	Desc    cockroach_roachpb.StoreDescriptor `protobuf:"bytes,1,opt,name=desc" json:"desc"`
	Metrics map[string]float64                `protobuf:"bytes,2,rep,name=metrics" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (m *StoreStatus) Reset()                    { *m = StoreStatus{} }
func (m *StoreStatus) String() string            { return proto.CompactTextString(m) }
func (*StoreStatus) ProtoMessage()               {}
func (*StoreStatus) Descriptor() ([]byte, []int) { return fileDescriptorStatus, []int{0} }

// NodeStatus records the most recent values of metrics for a node.
type NodeStatus struct {
	Desc          cockroach_roachpb.NodeDescriptor `protobuf:"bytes,1,opt,name=desc" json:"desc"`
	BuildInfo     cockroach_build.Info             `protobuf:"bytes,2,opt,name=build_info,json=buildInfo" json:"build_info"`
	StartedAt     int64                            `protobuf:"varint,3,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	UpdatedAt     int64                            `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Metrics       map[string]float64               `protobuf:"bytes,5,rep,name=metrics" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	StoreStatuses []StoreStatus                    `protobuf:"bytes,6,rep,name=store_statuses,json=storeStatuses" json:"store_statuses"`
	Args          []string                         `protobuf:"bytes,7,rep,name=args" json:"args,omitempty"`
	Env           []string                         `protobuf:"bytes,8,rep,name=env" json:"env,omitempty"`
	// latencies is a map of nodeIDs to nanoseconds which is the latency between
	// this node and the other node.
	Latencies map[github_com_cockroachdb_cockroach_pkg_roachpb.NodeID]int64 `protobuf:"bytes,9,rep,name=latencies,castkey=github.com/cockroachdb/cockroach/pkg/roachpb.NodeID" json:"latencies" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *NodeStatus) Reset()                    { *m = NodeStatus{} }
func (m *NodeStatus) String() string            { return proto.CompactTextString(m) }
func (*NodeStatus) ProtoMessage()               {}
func (*NodeStatus) Descriptor() ([]byte, []int) { return fileDescriptorStatus, []int{1} }

func init() {
	proto.RegisterType((*StoreStatus)(nil), "cockroach.server.status.StoreStatus")
	proto.RegisterType((*NodeStatus)(nil), "cockroach.server.status.NodeStatus")
}
func (m *StoreStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintStatus(dAtA, i, uint64(m.Desc.Size()))
	n1, err := m.Desc.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Metrics) > 0 {
		keysForMetrics := make([]string, 0, len(m.Metrics))
		for k := range m.Metrics {
			keysForMetrics = append(keysForMetrics, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForMetrics)
		for _, k := range keysForMetrics {
			dAtA[i] = 0x12
			i++
			v := m.Metrics[string(k)]
			mapSize := 1 + len(k) + sovStatus(uint64(len(k))) + 1 + 8
			i = encodeVarintStatus(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintStatus(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x11
			i++
			encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(v))))
			i += 8
		}
	}
	return i, nil
}

func (m *NodeStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintStatus(dAtA, i, uint64(m.Desc.Size()))
	n2, err := m.Desc.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x12
	i++
	i = encodeVarintStatus(dAtA, i, uint64(m.BuildInfo.Size()))
	n3, err := m.BuildInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if m.StartedAt != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintStatus(dAtA, i, uint64(m.StartedAt))
	}
	if m.UpdatedAt != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintStatus(dAtA, i, uint64(m.UpdatedAt))
	}
	if len(m.Metrics) > 0 {
		keysForMetrics := make([]string, 0, len(m.Metrics))
		for k := range m.Metrics {
			keysForMetrics = append(keysForMetrics, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForMetrics)
		for _, k := range keysForMetrics {
			dAtA[i] = 0x2a
			i++
			v := m.Metrics[string(k)]
			mapSize := 1 + len(k) + sovStatus(uint64(len(k))) + 1 + 8
			i = encodeVarintStatus(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintStatus(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x11
			i++
			encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(v))))
			i += 8
		}
	}
	if len(m.StoreStatuses) > 0 {
		for _, msg := range m.StoreStatuses {
			dAtA[i] = 0x32
			i++
			i = encodeVarintStatus(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			dAtA[i] = 0x3a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Env) > 0 {
		for _, s := range m.Env {
			dAtA[i] = 0x42
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Latencies) > 0 {
		keysForLatencies := make([]int32, 0, len(m.Latencies))
		for k := range m.Latencies {
			keysForLatencies = append(keysForLatencies, int32(k))
		}
		github_com_gogo_protobuf_sortkeys.Int32s(keysForLatencies)
		for _, k := range keysForLatencies {
			dAtA[i] = 0x4a
			i++
			v := m.Latencies[github_com_cockroachdb_cockroach_pkg_roachpb.NodeID(k)]
			mapSize := 1 + sovStatus(uint64(k)) + 1 + sovStatus(uint64(v))
			i = encodeVarintStatus(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintStatus(dAtA, i, uint64(k))
			dAtA[i] = 0x10
			i++
			i = encodeVarintStatus(dAtA, i, uint64(v))
		}
	}
	return i, nil
}

func encodeVarintStatus(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StoreStatus) Size() (n int) {
	var l int
	_ = l
	l = m.Desc.Size()
	n += 1 + l + sovStatus(uint64(l))
	if len(m.Metrics) > 0 {
		for k, v := range m.Metrics {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovStatus(uint64(len(k))) + 1 + 8
			n += mapEntrySize + 1 + sovStatus(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *NodeStatus) Size() (n int) {
	var l int
	_ = l
	l = m.Desc.Size()
	n += 1 + l + sovStatus(uint64(l))
	l = m.BuildInfo.Size()
	n += 1 + l + sovStatus(uint64(l))
	if m.StartedAt != 0 {
		n += 1 + sovStatus(uint64(m.StartedAt))
	}
	if m.UpdatedAt != 0 {
		n += 1 + sovStatus(uint64(m.UpdatedAt))
	}
	if len(m.Metrics) > 0 {
		for k, v := range m.Metrics {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovStatus(uint64(len(k))) + 1 + 8
			n += mapEntrySize + 1 + sovStatus(uint64(mapEntrySize))
		}
	}
	if len(m.StoreStatuses) > 0 {
		for _, e := range m.StoreStatuses {
			l = e.Size()
			n += 1 + l + sovStatus(uint64(l))
		}
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			l = len(s)
			n += 1 + l + sovStatus(uint64(l))
		}
	}
	if len(m.Env) > 0 {
		for _, s := range m.Env {
			l = len(s)
			n += 1 + l + sovStatus(uint64(l))
		}
	}
	if len(m.Latencies) > 0 {
		for k, v := range m.Latencies {
			_ = k
			_ = v
			mapEntrySize := 1 + sovStatus(uint64(k)) + 1 + sovStatus(uint64(v))
			n += mapEntrySize + 1 + sovStatus(uint64(mapEntrySize))
		}
	}
	return n
}

func sovStatus(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStatus(x uint64) (n int) {
	return sovStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoreStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatus
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
			return fmt.Errorf("proto: StoreStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Desc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metrics == nil {
				m.Metrics = make(map[string]float64)
			}
			var mapkey string
			var mapvalue float64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStatus
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowStatus
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthStatus
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapvaluetemp uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					mapvaluetemp = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					mapvalue = math.Float64frombits(mapvaluetemp)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipStatus(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthStatus
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metrics[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatus
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
func (m *NodeStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatus
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
			return fmt.Errorf("proto: NodeStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Desc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuildInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BuildInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartedAt", wireType)
			}
			m.StartedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartedAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metrics == nil {
				m.Metrics = make(map[string]float64)
			}
			var mapkey string
			var mapvalue float64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStatus
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowStatus
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthStatus
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapvaluetemp uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					mapvaluetemp = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					mapvalue = math.Float64frombits(mapvaluetemp)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipStatus(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthStatus
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metrics[mapkey] = mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreStatuses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoreStatuses = append(m.StoreStatuses, StoreStatus{})
			if err := m.StoreStatuses[len(m.StoreStatuses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Env", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Env = append(m.Env, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latencies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Latencies == nil {
				m.Latencies = make(map[github_com_cockroachdb_cockroach_pkg_roachpb.NodeID]int64)
			}
			var mapkey int32
			var mapvalue int64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStatus
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowStatus
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowStatus
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipStatus(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthStatus
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Latencies[github_com_cockroachdb_cockroach_pkg_roachpb.NodeID(mapkey)] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatus
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
func skipStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStatus
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
					return 0, ErrIntOverflowStatus
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
					return 0, ErrIntOverflowStatus
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
				return 0, ErrInvalidLengthStatus
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStatus
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
				next, err := skipStatus(dAtA[start:])
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
	ErrInvalidLengthStatus = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStatus   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("server/status/status.proto", fileDescriptorStatus) }

var fileDescriptorStatus = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xbf, 0x6f, 0xd3, 0x40,
	0x18, 0xcd, 0xc5, 0x49, 0x8a, 0x2f, 0x50, 0x55, 0xa7, 0x02, 0x96, 0x25, 0x5c, 0x13, 0x31, 0x78,
	0x3a, 0x43, 0xba, 0xa0, 0xb4, 0x4b, 0xa3, 0x32, 0x94, 0x5f, 0x12, 0xee, 0xc6, 0x12, 0x9d, 0xed,
	0xab, 0x6b, 0x25, 0xf1, 0x59, 0x77, 0xe7, 0x48, 0x5d, 0xd9, 0x91, 0x10, 0x7f, 0x55, 0x46, 0x06,
	0x06, 0x26, 0x7e, 0x84, 0x7f, 0x04, 0xdd, 0xd9, 0xc1, 0xae, 0x94, 0x48, 0x91, 0x3a, 0xf9, 0xf3,
	0x7b, 0xf7, 0xbe, 0xef, 0x7b, 0xef, 0x6c, 0x68, 0x0b, 0xca, 0x17, 0x94, 0xfb, 0x42, 0x12, 0x59,
	0x88, 0xea, 0x81, 0x73, 0xce, 0x24, 0x43, 0x8f, 0x23, 0x16, 0x4d, 0x39, 0x23, 0xd1, 0x35, 0x2e,
	0x4f, 0xe1, 0x92, 0xb6, 0x1f, 0x69, 0x30, 0x0f, 0xfd, 0x39, 0x95, 0x24, 0x26, 0x92, 0x94, 0x02,
	0xfb, 0x20, 0x2c, 0xd2, 0x59, 0xec, 0xa7, 0xd9, 0x15, 0xab, 0x90, 0xc3, 0x84, 0x25, 0x4c, 0x97,
	0xbe, 0xaa, 0x4a, 0x74, 0xf0, 0x1d, 0xc0, 0xfe, 0xa5, 0x64, 0x9c, 0x5e, 0xea, 0x7e, 0xe8, 0x14,
	0x76, 0x62, 0x2a, 0x22, 0x0b, 0xb8, 0xc0, 0xeb, 0x0f, 0x07, 0xb8, 0x9e, 0x5b, 0x0d, 0xc2, 0xfa,
	0xf4, 0x39, 0x15, 0x11, 0x4f, 0x73, 0xc9, 0xf8, 0xb8, 0xb3, 0xfc, 0x79, 0xd4, 0x0a, 0xb4, 0x0a,
	0xbd, 0x81, 0x7b, 0x73, 0x2a, 0x79, 0x1a, 0x09, 0xab, 0xed, 0x1a, 0x5e, 0x7f, 0xf8, 0x02, 0x6f,
	0x59, 0x1c, 0x37, 0x86, 0xe2, 0x77, 0xa5, 0xe6, 0x55, 0x26, 0xf9, 0x4d, 0xb0, 0xee, 0x60, 0x8f,
	0xe0, 0xfd, 0x26, 0x81, 0x0e, 0xa0, 0x31, 0xa5, 0x37, 0x7a, 0x33, 0x33, 0x50, 0x25, 0x3a, 0x84,
	0xdd, 0x05, 0x99, 0x15, 0xd4, 0x6a, 0xbb, 0xc0, 0x03, 0x41, 0xf9, 0x32, 0x6a, 0xbf, 0x04, 0x83,
	0xaf, 0x5d, 0x08, 0xdf, 0xb3, 0x78, 0xed, 0xea, 0xe4, 0x96, 0xab, 0xa7, 0x1b, 0x5c, 0xa9, 0xc3,
	0x5b, 0x4c, 0x8d, 0x20, 0xd4, 0x61, 0x4e, 0x54, 0x98, 0x7a, 0x54, 0x7f, 0xf8, 0xb0, 0xd1, 0x42,
	0x93, 0xf8, 0x22, 0xbb, 0x62, 0x95, 0xcc, 0xd4, 0x88, 0x02, 0xd0, 0x13, 0x08, 0x85, 0x24, 0x5c,
	0xd2, 0x78, 0x42, 0xa4, 0x65, 0xb8, 0xc0, 0x33, 0x02, 0xb3, 0x42, 0xce, 0xa4, 0xa2, 0x8b, 0x3c,
	0x26, 0x15, 0xdd, 0x29, 0xe9, 0x0a, 0x39, 0x93, 0xe8, 0x75, 0x1d, 0x67, 0x57, 0xc7, 0xf9, 0x7c,
	0x6b, 0x9c, 0xb5, 0xd9, 0xcd, 0x69, 0xa2, 0x0f, 0x70, 0x5f, 0xa8, 0xc8, 0x27, 0xa5, 0x80, 0x0a,
	0xab, 0xa7, 0x5b, 0x3e, 0xdb, 0xe5, 0x86, 0x2a, 0x63, 0x0f, 0x44, 0x0d, 0x51, 0x81, 0x10, 0xec,
	0x10, 0x9e, 0x08, 0x6b, 0xcf, 0x35, 0x3c, 0x33, 0xd0, 0xb5, 0xba, 0x24, 0x9a, 0x2d, 0xac, 0x7b,
	0x1a, 0x52, 0x25, 0xfa, 0x0c, 0xa0, 0x39, 0x23, 0x92, 0x66, 0x51, 0x4a, 0x85, 0x65, 0xea, 0xa1,
	0xc3, 0x5d, 0x7c, 0xbc, 0x5d, 0x8b, 0xb4, 0x93, 0xf1, 0x89, 0x5a, 0xe1, 0xd3, 0xaf, 0xa3, 0xe3,
	0x24, 0x95, 0xd7, 0x45, 0x88, 0x23, 0x36, 0xf7, 0xff, 0x77, 0x89, 0xc3, 0xba, 0xf6, 0xf3, 0x69,
	0xe2, 0x37, 0xef, 0xf5, 0xe2, 0x3c, 0xa8, 0x37, 0xb8, 0xcb, 0x67, 0x65, 0x9f, 0xc2, 0xfd, 0xdb,
	0x5b, 0x35, 0xd5, 0xdd, 0x0d, 0x6a, 0xa3, 0xa1, 0x1e, 0xbb, 0xcb, 0x3f, 0x4e, 0x6b, 0xb9, 0x72,
	0xc0, 0xb7, 0x95, 0x03, 0x7e, 0xac, 0x1c, 0xf0, 0x7b, 0xe5, 0x80, 0x2f, 0x7f, 0x9d, 0xd6, 0xc7,
	0x5e, 0xe9, 0x3e, 0xec, 0xe9, 0x9f, 0xf2, 0xf8, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0xed,
	0x10, 0x1a, 0x0b, 0x04, 0x00, 0x00,
}
