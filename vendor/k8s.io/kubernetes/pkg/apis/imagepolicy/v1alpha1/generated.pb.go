/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1/generated.proto
// DO NOT EDIT!

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1/generated.proto

	It has these top-level messages:
		ImageReview
		ImageReviewContainerSpec
		ImageReviewSpec
		ImageReviewStatus
*/
package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

func (m *ImageReview) Reset()                    { *m = ImageReview{} }
func (*ImageReview) ProtoMessage()               {}
func (*ImageReview) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *ImageReviewContainerSpec) Reset()      { *m = ImageReviewContainerSpec{} }
func (*ImageReviewContainerSpec) ProtoMessage() {}
func (*ImageReviewContainerSpec) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{1}
}

func (m *ImageReviewSpec) Reset()                    { *m = ImageReviewSpec{} }
func (*ImageReviewSpec) ProtoMessage()               {}
func (*ImageReviewSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *ImageReviewStatus) Reset()                    { *m = ImageReviewStatus{} }
func (*ImageReviewStatus) ProtoMessage()               {}
func (*ImageReviewStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func init() {
	proto.RegisterType((*ImageReview)(nil), "k8s.io.kubernetes.pkg.apis.imagepolicy.v1alpha1.ImageReview")
	proto.RegisterType((*ImageReviewContainerSpec)(nil), "k8s.io.kubernetes.pkg.apis.imagepolicy.v1alpha1.ImageReviewContainerSpec")
	proto.RegisterType((*ImageReviewSpec)(nil), "k8s.io.kubernetes.pkg.apis.imagepolicy.v1alpha1.ImageReviewSpec")
	proto.RegisterType((*ImageReviewStatus)(nil), "k8s.io.kubernetes.pkg.apis.imagepolicy.v1alpha1.ImageReviewStatus")
}
func (m *ImageReview) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ImageReview) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Status.Size()))
	n3, err := m.Status.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *ImageReviewContainerSpec) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ImageReviewContainerSpec) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Image)))
	i += copy(data[i:], m.Image)
	return i, nil
}

func (m *ImageReviewSpec) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ImageReviewSpec) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Containers) > 0 {
		for _, msg := range m.Containers {
			data[i] = 0xa
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Annotations) > 0 {
		keysForAnnotations := make([]string, 0, len(m.Annotations))
		for k := range m.Annotations {
			keysForAnnotations = append(keysForAnnotations, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForAnnotations)
		for _, k := range keysForAnnotations {
			data[i] = 0x12
			i++
			v := m.Annotations[string(k)]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintGenerated(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(len(v)))
			i += copy(data[i:], v)
		}
	}
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Namespace)))
	i += copy(data[i:], m.Namespace)
	return i, nil
}

func (m *ImageReviewStatus) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ImageReviewStatus) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	if m.Allowed {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Reason)))
	i += copy(data[i:], m.Reason)
	return i, nil
}

func encodeFixed64Generated(data []byte, offset int, v uint64) int {
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
func encodeFixed32Generated(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ImageReview) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ImageReviewContainerSpec) Size() (n int) {
	var l int
	_ = l
	l = len(m.Image)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ImageReviewSpec) Size() (n int) {
	var l int
	_ = l
	if len(m.Containers) > 0 {
		for _, e := range m.Containers {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Annotations) > 0 {
		for k, v := range m.Annotations {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	l = len(m.Namespace)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ImageReviewStatus) Size() (n int) {
	var l int
	_ = l
	n += 2
	l = len(m.Reason)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ImageReview) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ImageReview{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "ImageReviewSpec", "ImageReviewSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "ImageReviewStatus", "ImageReviewStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ImageReviewContainerSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ImageReviewContainerSpec{`,
		`Image:` + fmt.Sprintf("%v", this.Image) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ImageReviewSpec) String() string {
	if this == nil {
		return "nil"
	}
	keysForAnnotations := make([]string, 0, len(this.Annotations))
	for k := range this.Annotations {
		keysForAnnotations = append(keysForAnnotations, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForAnnotations)
	mapStringForAnnotations := "map[string]string{"
	for _, k := range keysForAnnotations {
		mapStringForAnnotations += fmt.Sprintf("%v: %v,", k, this.Annotations[k])
	}
	mapStringForAnnotations += "}"
	s := strings.Join([]string{`&ImageReviewSpec{`,
		`Containers:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Containers), "ImageReviewContainerSpec", "ImageReviewContainerSpec", 1), `&`, ``, 1) + `,`,
		`Annotations:` + mapStringForAnnotations + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ImageReviewStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ImageReviewStatus{`,
		`Allowed:` + fmt.Sprintf("%v", this.Allowed) + `,`,
		`Reason:` + fmt.Sprintf("%v", this.Reason) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ImageReview) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ImageReview: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageReview: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *ImageReviewContainerSpec) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ImageReviewContainerSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageReviewContainerSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *ImageReviewSpec) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ImageReviewSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageReviewSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Containers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Containers = append(m.Containers, ImageReviewContainerSpec{})
			if err := m.Containers[len(m.Containers)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Annotations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapvalue uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapvalue |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapvalue := int(stringLenmapvalue)
			if intStringLenmapvalue < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapvalue := iNdEx + intStringLenmapvalue
			if postStringIndexmapvalue > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := string(data[iNdEx:postStringIndexmapvalue])
			iNdEx = postStringIndexmapvalue
			if m.Annotations == nil {
				m.Annotations = make(map[string]string)
			}
			m.Annotations[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *ImageReviewStatus) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ImageReviewStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImageReviewStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Allowed = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(data[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorGenerated = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0xbf, 0x6f, 0xd3, 0x4e,
	0x18, 0xc6, 0xe3, 0xa4, 0xbf, 0x72, 0xf9, 0x7e, 0x69, 0x7b, 0x30, 0x44, 0x19, 0xdc, 0x2a, 0x48,
	0xa8, 0x0c, 0xdc, 0xd1, 0x82, 0x50, 0xc5, 0x40, 0xa9, 0x11, 0x43, 0x07, 0x40, 0x1c, 0x1b, 0x13,
	0x17, 0xf7, 0xad, 0x73, 0x8d, 0x7d, 0x67, 0xf9, 0xce, 0x29, 0x19, 0x90, 0x18, 0x19, 0x18, 0xf8,
	0x8f, 0x58, 0x3b, 0x76, 0xec, 0x54, 0x51, 0xf3, 0x8f, 0x20, 0x9f, 0x9d, 0xda, 0x34, 0x54, 0x08,
	0x65, 0xf3, 0x7b, 0xef, 0xbd, 0x9f, 0xe7, 0x79, 0xef, 0x31, 0xda, 0x1b, 0xed, 0x6a, 0x22, 0x14,
	0x1d, 0xa5, 0x03, 0x48, 0x24, 0x18, 0xd0, 0x34, 0x1e, 0x05, 0x94, 0xc7, 0x42, 0x53, 0x11, 0xf1,
	0x00, 0x62, 0x15, 0x0a, 0x7f, 0x42, 0xc7, 0xdb, 0x3c, 0x8c, 0x87, 0x7c, 0x9b, 0x06, 0x20, 0x21,
	0xe1, 0x06, 0x0e, 0x49, 0x9c, 0x28, 0xa3, 0x30, 0x2d, 0x00, 0xa4, 0x02, 0x90, 0x78, 0x14, 0x90,
	0x1c, 0x40, 0x6a, 0x00, 0x32, 0x05, 0xf4, 0x1e, 0x04, 0xc2, 0x0c, 0xd3, 0x01, 0xf1, 0x55, 0x44,
	0x03, 0x15, 0x28, 0x6a, 0x39, 0x83, 0xf4, 0xc8, 0x56, 0xb6, 0xb0, 0x5f, 0x05, 0xbf, 0xf7, 0xb8,
	0x34, 0xc8, 0x63, 0x11, 0x71, 0x7f, 0x28, 0x24, 0x24, 0x93, 0xca, 0x62, 0x04, 0x86, 0xd3, 0xf1,
	0x8c, 0xab, 0x1e, 0xbd, 0x69, 0x2a, 0x49, 0xa5, 0x11, 0x11, 0xcc, 0x0c, 0x3c, 0xf9, 0xdb, 0x80,
	0xf6, 0x87, 0x10, 0xf1, 0x99, 0xb9, 0x47, 0x37, 0xcd, 0xa5, 0x46, 0x84, 0x54, 0x48, 0xa3, 0x4d,
	0x32, 0x33, 0x54, 0xdb, 0x49, 0x43, 0x32, 0x86, 0xa4, 0x5a, 0x08, 0x3e, 0xf2, 0x28, 0x0e, 0xe1,
	0x0f, 0x3b, 0xf5, 0xbf, 0x37, 0x51, 0xe7, 0x20, 0x7f, 0x51, 0x06, 0x63, 0x01, 0x27, 0xf8, 0x03,
	0x5a, 0xc9, 0xd7, 0x3f, 0xe4, 0x86, 0x77, 0x9d, 0x4d, 0x67, 0xab, 0xb3, 0xf3, 0x90, 0x94, 0x61,
	0xd4, 0xdd, 0x54, 0x71, 0xe4, 0xb7, 0xc9, 0x78, 0x9b, 0xbc, 0x19, 0x1c, 0x83, 0x6f, 0x5e, 0x81,
	0xe1, 0x1e, 0x3e, 0xbd, 0xd8, 0x68, 0x64, 0x17, 0x1b, 0xa8, 0x3a, 0x63, 0x57, 0x54, 0x3c, 0x40,
	0x0b, 0x3a, 0x06, 0xbf, 0xdb, 0xb4, 0xf4, 0xe7, 0xe4, 0x1f, 0xa3, 0x26, 0x35, 0xb7, 0xef, 0x62,
	0xf0, 0xbd, 0xff, 0x4a, 0xb5, 0x85, 0xbc, 0x62, 0x96, 0x8d, 0x8f, 0xd1, 0x92, 0x36, 0xdc, 0xa4,
	0xba, 0xdb, 0xb2, 0x2a, 0xde, 0x5c, 0x2a, 0x96, 0xe4, 0xdd, 0x2a, 0x75, 0x96, 0x8a, 0x9a, 0x95,
	0x0a, 0xfd, 0x3d, 0xd4, 0xad, 0x5d, 0x7e, 0xa1, 0xa4, 0xe1, 0xf9, 0x13, 0xe5, 0x6e, 0xf0, 0x5d,
	0xb4, 0x68, 0xe9, 0xf6, 0x29, 0xdb, 0xde, 0xff, 0x25, 0x62, 0xb1, 0x18, 0x28, 0x7a, 0xfd, 0xaf,
	0x2d, 0xb4, 0x7a, 0x6d, 0x29, 0xfc, 0x09, 0x21, 0x7f, 0x4a, 0xd2, 0x5d, 0x67, 0xb3, 0xb5, 0xd5,
	0xd9, 0x39, 0x98, 0x67, 0x89, 0xdf, 0x7c, 0x55, 0x09, 0x5d, 0x1d, 0x6b, 0x56, 0x13, 0xc4, 0x5f,
	0x1c, 0xd4, 0xe1, 0x52, 0x2a, 0xc3, 0x8d, 0x50, 0x52, 0x77, 0x9b, 0xd6, 0xc0, 0xdb, 0x79, 0xb3,
	0x22, 0xfb, 0x15, 0xf3, 0xa5, 0x34, 0xc9, 0xc4, 0xbb, 0x5d, 0x1a, 0xe9, 0xd4, 0x3a, 0xac, 0x2e,
	0x8d, 0x29, 0x6a, 0x4b, 0x1e, 0x81, 0x8e, 0xb9, 0x0f, 0x36, 0xcd, 0xb6, 0xb7, 0x5e, 0x0e, 0xb5,
	0x5f, 0x4f, 0x1b, 0xac, 0xba, 0xd3, 0x7b, 0x86, 0xd6, 0xae, 0xcb, 0xe0, 0x35, 0xd4, 0x1a, 0xc1,
	0xa4, 0x48, 0x81, 0xe5, 0x9f, 0xf8, 0x0e, 0x5a, 0x1c, 0xf3, 0x30, 0x05, 0xfb, 0x1b, 0xb6, 0x59,
	0x51, 0x3c, 0x6d, 0xee, 0x3a, 0xfd, 0x23, 0xb4, 0x3e, 0x13, 0x3e, 0xbe, 0x8f, 0x96, 0x79, 0x18,
	0xaa, 0x13, 0x38, 0xb4, 0x90, 0x15, 0x6f, 0xb5, 0xf4, 0xb0, 0xbc, 0x5f, 0x1c, 0xb3, 0x69, 0x1f,
	0xdf, 0x43, 0x4b, 0x09, 0x70, 0xad, 0x64, 0x81, 0xae, 0xfe, 0x1b, 0x66, 0x4f, 0x59, 0xd9, 0xf5,
	0xc8, 0xe9, 0xa5, 0xdb, 0x38, 0xbb, 0x74, 0x1b, 0xe7, 0x97, 0x6e, 0xe3, 0x73, 0xe6, 0x3a, 0xa7,
	0x99, 0xeb, 0x9c, 0x65, 0xae, 0x73, 0x9e, 0xb9, 0xce, 0x8f, 0xcc, 0x75, 0xbe, 0xfd, 0x74, 0x1b,
	0xef, 0x57, 0xa6, 0x6f, 0xf9, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x29, 0x7f, 0x0a, 0x29, 0x55, 0x05,
	0x00, 0x00,
}
