/*
Copyright 2018-2020, Arm Limited and affiliates.
Copyright The Kubernetes Authors.

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

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/kubernetes/vendor/k8s.io/api/admission/v1beta1/generated.proto

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/api/admission/v1beta1/generated.proto

	It has these top-level messages:
		AdmissionRequest
		AdmissionResponse
		AdmissionReview
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_apimachinery_pkg_apis_meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

import k8s_io_apimachinery_pkg_types "k8s.io/apimachinery/pkg/types"

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *AdmissionRequest) Reset()                    { *m = AdmissionRequest{} }
func (*AdmissionRequest) ProtoMessage()               {}
func (*AdmissionRequest) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *AdmissionResponse) Reset()                    { *m = AdmissionResponse{} }
func (*AdmissionResponse) ProtoMessage()               {}
func (*AdmissionResponse) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *AdmissionReview) Reset()                    { *m = AdmissionReview{} }
func (*AdmissionReview) ProtoMessage()               {}
func (*AdmissionReview) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func init() {
	proto.RegisterType((*AdmissionRequest)(nil), "k8s.io.api.admission.v1beta1.AdmissionRequest")
	proto.RegisterType((*AdmissionResponse)(nil), "k8s.io.api.admission.v1beta1.AdmissionResponse")
	proto.RegisterType((*AdmissionReview)(nil), "k8s.io.api.admission.v1beta1.AdmissionReview")
}
func (m *AdmissionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdmissionRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.UID)))
	i += copy(dAtA[i:], m.UID)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Kind.Size()))
	n1, err := m.Kind.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Resource.Size()))
	n2, err := m.Resource.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.SubResource)))
	i += copy(dAtA[i:], m.SubResource)
	dAtA[i] = 0x2a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	dAtA[i] = 0x32
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Namespace)))
	i += copy(dAtA[i:], m.Namespace)
	dAtA[i] = 0x3a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Operation)))
	i += copy(dAtA[i:], m.Operation)
	dAtA[i] = 0x42
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.UserInfo.Size()))
	n3, err := m.UserInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x4a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Object.Size()))
	n4, err := m.Object.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x52
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.OldObject.Size()))
	n5, err := m.OldObject.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if m.DryRun != nil {
		dAtA[i] = 0x58
		i++
		if *m.DryRun {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	dAtA[i] = 0x62
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.AccountID)))
	i += copy(dAtA[i:], m.AccountID)
	return i, nil
}

func (m *AdmissionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdmissionResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.UID)))
	i += copy(dAtA[i:], m.UID)
	dAtA[i] = 0x10
	i++
	if m.Allowed {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	if m.Result != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Result.Size()))
		n6, err := m.Result.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.Patch != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Patch)))
		i += copy(dAtA[i:], m.Patch)
	}
	if m.PatchType != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.PatchType)))
		i += copy(dAtA[i:], *m.PatchType)
	}
	if len(m.AuditAnnotations) > 0 {
		keysForAuditAnnotations := make([]string, 0, len(m.AuditAnnotations))
		for k := range m.AuditAnnotations {
			keysForAuditAnnotations = append(keysForAuditAnnotations, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForAuditAnnotations)
		for _, k := range keysForAuditAnnotations {
			dAtA[i] = 0x32
			i++
			v := m.AuditAnnotations[string(k)]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func (m *AdmissionReview) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdmissionReview) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Request != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Request.Size()))
		n7, err := m.Request.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.Response != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Response.Size()))
		n8, err := m.Response.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AdmissionRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.UID)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Kind.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Resource.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.SubResource)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Namespace)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Operation)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.UserInfo.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Object.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.OldObject.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.DryRun != nil {
		n += 2
	}
	l = len(m.AccountID)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *AdmissionResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.UID)
	n += 1 + l + sovGenerated(uint64(l))
	n += 2
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Patch != nil {
		l = len(m.Patch)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.PatchType != nil {
		l = len(*m.PatchType)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.AuditAnnotations) > 0 {
		for k, v := range m.AuditAnnotations {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *AdmissionReview) Size() (n int) {
	var l int
	_ = l
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Response != nil {
		l = m.Response.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
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
func (this *AdmissionRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AdmissionRequest{`,
		`UID:` + fmt.Sprintf("%v", this.UID) + `,`,
		`Kind:` + strings.Replace(strings.Replace(this.Kind.String(), "GroupVersionKind", "k8s_io_apimachinery_pkg_apis_meta_v1.GroupVersionKind", 1), `&`, ``, 1) + `,`,
		`Resource:` + strings.Replace(strings.Replace(this.Resource.String(), "GroupVersionResource", "k8s_io_apimachinery_pkg_apis_meta_v1.GroupVersionResource", 1), `&`, ``, 1) + `,`,
		`SubResource:` + fmt.Sprintf("%v", this.SubResource) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Operation:` + fmt.Sprintf("%v", this.Operation) + `,`,
		`UserInfo:` + strings.Replace(strings.Replace(this.UserInfo.String(), "UserInfo", "k8s_io_api_authentication_v1.UserInfo", 1), `&`, ``, 1) + `,`,
		`Object:` + strings.Replace(strings.Replace(this.Object.String(), "RawExtension", "k8s_io_apimachinery_pkg_runtime.RawExtension", 1), `&`, ``, 1) + `,`,
		`OldObject:` + strings.Replace(strings.Replace(this.OldObject.String(), "RawExtension", "k8s_io_apimachinery_pkg_runtime.RawExtension", 1), `&`, ``, 1) + `,`,
		`DryRun:` + valueToStringGenerated(this.DryRun) + `,`,
		`AccountID:` + fmt.Sprintf("%v", this.AccountID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AdmissionResponse) String() string {
	if this == nil {
		return "nil"
	}
	keysForAuditAnnotations := make([]string, 0, len(this.AuditAnnotations))
	for k := range this.AuditAnnotations {
		keysForAuditAnnotations = append(keysForAuditAnnotations, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForAuditAnnotations)
	mapStringForAuditAnnotations := "map[string]string{"
	for _, k := range keysForAuditAnnotations {
		mapStringForAuditAnnotations += fmt.Sprintf("%v: %v,", k, this.AuditAnnotations[k])
	}
	mapStringForAuditAnnotations += "}"
	s := strings.Join([]string{`&AdmissionResponse{`,
		`UID:` + fmt.Sprintf("%v", this.UID) + `,`,
		`Allowed:` + fmt.Sprintf("%v", this.Allowed) + `,`,
		`Result:` + strings.Replace(fmt.Sprintf("%v", this.Result), "Status", "k8s_io_apimachinery_pkg_apis_meta_v1.Status", 1) + `,`,
		`Patch:` + valueToStringGenerated(this.Patch) + `,`,
		`PatchType:` + valueToStringGenerated(this.PatchType) + `,`,
		`AuditAnnotations:` + mapStringForAuditAnnotations + `,`,
		`}`,
	}, "")
	return s
}
func (this *AdmissionReview) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AdmissionReview{`,
		`Request:` + strings.Replace(fmt.Sprintf("%v", this.Request), "AdmissionRequest", "AdmissionRequest", 1) + `,`,
		`Response:` + strings.Replace(fmt.Sprintf("%v", this.Response), "AdmissionResponse", "AdmissionResponse", 1) + `,`,
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
func (m *AdmissionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
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
			return fmt.Errorf("proto: AdmissionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdmissionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = k8s_io_apimachinery_pkg_types.UID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Kind.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Resource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubResource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubResource = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
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
				b := dAtA[iNdEx]
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
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operation = Operation(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UserInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Object", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Object.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldObject", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OldObject.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DryRun", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.DryRun = &b
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
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
func (m *AdmissionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
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
			return fmt.Errorf("proto: AdmissionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdmissionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = k8s_io_apimachinery_pkg_types.UID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Allowed = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Result == nil {
				m.Result = &k8s_io_apimachinery_pkg_apis_meta_v1.Status{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Patch", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Patch = append(m.Patch[:0], dAtA[iNdEx:postIndex]...)
			if m.Patch == nil {
				m.Patch = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PatchType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := PatchType(dAtA[iNdEx:postIndex])
			m.PatchType = &s
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuditAnnotations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuditAnnotations == nil {
				m.AuditAnnotations = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
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
							return ErrIntOverflowGenerated
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
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
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
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.AuditAnnotations[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
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
func (m *AdmissionReview) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
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
			return fmt.Errorf("proto: AdmissionReview: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdmissionReview: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &AdmissionRequest{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Response == nil {
				m.Response = &AdmissionResponse{}
			}
			if err := m.Response.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/admission/v1beta1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 839 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x41, 0x6f, 0xe3, 0x44,
	0x14, 0x8e, 0x37, 0x69, 0x12, 0x4f, 0x2a, 0x36, 0x3b, 0x80, 0x64, 0x45, 0xc8, 0x09, 0x3d, 0xa0,
	0x22, 0x6d, 0xc7, 0xb4, 0x82, 0x55, 0xb5, 0xe2, 0x12, 0xd3, 0x0a, 0x55, 0x48, 0xdb, 0x6a, 0x76,
	0x83, 0x80, 0x03, 0xd2, 0xc4, 0x9e, 0x4d, 0x4c, 0xe2, 0x19, 0xe3, 0x99, 0x49, 0xc9, 0x0d, 0x71,
	0xe5, 0xc2, 0x7f, 0x82, 0x43, 0x8f, 0x7b, 0xdc, 0x53, 0x45, 0xc3, 0xbf, 0xe8, 0x09, 0xcd, 0x78,
	0x1c, 0x67, 0xdb, 0x2d, 0xec, 0x22, 0x4e, 0xf6, 0x7b, 0xef, 0xfb, 0xbe, 0x67, 0x7f, 0x6f, 0xde,
	0x80, 0xe3, 0xd9, 0xa1, 0x40, 0x09, 0x0f, 0x66, 0x6a, 0x4c, 0x73, 0x46, 0x25, 0x15, 0xc1, 0x82,
	0xb2, 0x98, 0xe7, 0x81, 0x2d, 0x90, 0x2c, 0x09, 0x48, 0x9c, 0x26, 0x42, 0x24, 0x9c, 0x05, 0x8b,
	0xfd, 0x31, 0x95, 0x64, 0x3f, 0x98, 0x50, 0x46, 0x73, 0x22, 0x69, 0x8c, 0xb2, 0x9c, 0x4b, 0x0e,
	0x3f, 0x28, 0xd0, 0x88, 0x64, 0x09, 0x5a, 0xa3, 0x91, 0x45, 0xf7, 0xf6, 0x26, 0x89, 0x9c, 0xaa,
	0x31, 0x8a, 0x78, 0x1a, 0x4c, 0xf8, 0x84, 0x07, 0x86, 0x34, 0x56, 0xcf, 0x4d, 0x64, 0x02, 0xf3,
	0x56, 0x88, 0xf5, 0x1e, 0x6e, 0xb6, 0x56, 0x72, 0x4a, 0x99, 0x4c, 0x22, 0x22, 0x8b, 0xfe, 0x37,
	0x5b, 0xf7, 0x3e, 0xad, 0xd0, 0x29, 0x89, 0xa6, 0x09, 0xa3, 0xf9, 0x32, 0xc8, 0x66, 0x13, 0x9d,
	0x10, 0x41, 0x4a, 0x25, 0x79, 0x1d, 0x2b, 0xb8, 0x8b, 0x95, 0x2b, 0x26, 0x93, 0x94, 0xde, 0x22,
	0x3c, 0xfa, 0x37, 0x82, 0x88, 0xa6, 0x34, 0x25, 0x37, 0x79, 0x3b, 0x7f, 0x34, 0x41, 0x77, 0x58,
	0x3a, 0x82, 0xe9, 0x8f, 0x8a, 0x0a, 0x09, 0x43, 0x50, 0x57, 0x49, 0xec, 0x39, 0x03, 0x67, 0xd7,
	0x0d, 0x3f, 0xb9, 0xb8, 0xec, 0xd7, 0x56, 0x97, 0xfd, 0xfa, 0xe8, 0xe4, 0xe8, 0xfa, 0xb2, 0xff,
	0xe1, 0x5d, 0x8d, 0xe4, 0x32, 0xa3, 0x02, 0x8d, 0x4e, 0x8e, 0xb0, 0x26, 0xc3, 0x6f, 0x40, 0x63,
	0x96, 0xb0, 0xd8, 0xbb, 0x37, 0x70, 0x76, 0x3b, 0x07, 0x8f, 0x50, 0x35, 0x81, 0x35, 0x0d, 0x65,
	0xb3, 0x89, 0x4e, 0x08, 0xa4, 0x6d, 0x40, 0x8b, 0x7d, 0xf4, 0x65, 0xce, 0x55, 0xf6, 0x35, 0xcd,
	0xf5, 0xc7, 0x7c, 0x95, 0xb0, 0x38, 0xdc, 0xb6, 0xcd, 0x1b, 0x3a, 0xc2, 0x46, 0x11, 0x4e, 0x41,
	0x3b, 0xa7, 0x82, 0xab, 0x3c, 0xa2, 0x5e, 0xdd, 0xa8, 0x3f, 0x7e, 0x7b, 0x75, 0x6c, 0x15, 0xc2,
	0xae, 0xed, 0xd0, 0x2e, 0x33, 0x78, 0xad, 0x0e, 0x3f, 0x03, 0x1d, 0xa1, 0xc6, 0x65, 0xc1, 0x6b,
	0x18, 0x3f, 0xde, 0xb5, 0x84, 0xce, 0xd3, 0xaa, 0x84, 0x37, 0x71, 0x70, 0x00, 0x1a, 0x8c, 0xa4,
	0xd4, 0xdb, 0x32, 0xf8, 0xf5, 0x2f, 0x3c, 0x21, 0x29, 0xc5, 0xa6, 0x02, 0x03, 0xe0, 0xea, 0xa7,
	0xc8, 0x48, 0x44, 0xbd, 0xa6, 0x81, 0x3d, 0xb0, 0x30, 0xf7, 0x49, 0x59, 0xc0, 0x15, 0x06, 0x7e,
	0x0e, 0x5c, 0x9e, 0xe9, 0xc1, 0x25, 0x9c, 0x79, 0x2d, 0x43, 0xf0, 0x4b, 0xc2, 0x69, 0x59, 0xb8,
	0xde, 0x0c, 0x70, 0x45, 0x80, 0xcf, 0x40, 0x5b, 0x09, 0x9a, 0x9f, 0xb0, 0xe7, 0xdc, 0x6b, 0x1b,
	0xc7, 0x3e, 0x42, 0x9b, 0x1b, 0xf1, 0xca, 0x21, 0xd6, 0x4e, 0x8d, 0x2c, 0xba, 0x72, 0xa7, 0xcc,
	0xe0, 0xb5, 0x12, 0x1c, 0x81, 0x26, 0x1f, 0xff, 0x40, 0x23, 0xe9, 0xb9, 0x46, 0x73, 0xef, 0xce,
	0x29, 0xd8, 0x33, 0x88, 0x30, 0x39, 0x3f, 0xfe, 0x49, 0x52, 0xa6, 0x07, 0x10, 0xbe, 0x63, 0xa5,
	0x9b, 0xa7, 0x46, 0x04, 0x5b, 0x31, 0xf8, 0x3d, 0x70, 0xf9, 0x3c, 0x2e, 0x92, 0x1e, 0xf8, 0x2f,
	0xca, 0x6b, 0x2b, 0x4f, 0x4b, 0x1d, 0x5c, 0x49, 0xc2, 0x1d, 0xd0, 0x8c, 0xf3, 0x25, 0x56, 0xcc,
	0xeb, 0x0c, 0x9c, 0xdd, 0x76, 0x08, 0xf4, 0x37, 0x1c, 0x99, 0x0c, 0xb6, 0x15, 0x3d, 0x1f, 0x12,
	0x45, 0x5c, 0x8b, 0xc7, 0xde, 0xf6, 0xab, 0xf3, 0x19, 0x16, 0x85, 0x93, 0x23, 0x5c, 0x61, 0x76,
	0x7e, 0x69, 0x80, 0x07, 0x1b, 0x6b, 0x24, 0x32, 0xce, 0x04, 0xfd, 0x5f, 0xf6, 0xe8, 0x63, 0xd0,
	0x22, 0xf3, 0x39, 0x3f, 0xa7, 0xc5, 0x2a, 0xb5, 0xc3, 0xfb, 0x56, 0xa7, 0x35, 0x2c, 0xd2, 0xb8,
	0xac, 0xc3, 0x33, 0xd0, 0x14, 0x92, 0x48, 0x25, 0xec, 0x5a, 0x3c, 0x7c, 0xb3, 0xb5, 0x78, 0x6a,
	0x38, 0x85, 0x0f, 0x98, 0x0a, 0x35, 0x97, 0xd8, 0xea, 0xc0, 0x3e, 0xd8, 0xca, 0x88, 0x8c, 0xa6,
	0xe6, 0xe8, 0x6f, 0x87, 0xee, 0xea, 0xb2, 0xbf, 0x75, 0xa6, 0x13, 0xb8, 0xc8, 0xc3, 0x43, 0xe0,
	0x9a, 0x97, 0x67, 0xcb, 0xac, 0x3c, 0xef, 0x3d, 0x6d, 0xd2, 0x59, 0x99, 0xbc, 0xde, 0x0c, 0x70,
	0x05, 0x86, 0xbf, 0x3a, 0xa0, 0x4b, 0x54, 0x9c, 0xc8, 0x21, 0x63, 0x5c, 0x9a, 0x93, 0x27, 0xbc,
	0xe6, 0xa0, 0xbe, 0xdb, 0x39, 0x38, 0x46, 0xff, 0x74, 0x5d, 0xa3, 0x5b, 0x3e, 0xa3, 0xe1, 0x0d,
	0x9d, 0x63, 0x26, 0xf3, 0x65, 0xe8, 0x59, 0xa3, 0xba, 0x37, 0xcb, 0xf8, 0x56, 0xe3, 0xde, 0x17,
	0xe0, 0xfd, 0xd7, 0x8a, 0xc0, 0x2e, 0xa8, 0xcf, 0xe8, 0xb2, 0x18, 0x21, 0xd6, 0xaf, 0xf0, 0x3d,
	0xb0, 0xb5, 0x20, 0x73, 0x45, 0xcd, 0x38, 0x5c, 0x5c, 0x04, 0x8f, 0xef, 0x1d, 0x3a, 0x3b, 0xbf,
	0x3b, 0xe0, 0xfe, 0xc6, 0xc7, 0x2d, 0x12, 0x7a, 0x0e, 0x47, 0xa0, 0x95, 0x17, 0xb7, 0xaa, 0xd1,
	0xe8, 0x1c, 0xa0, 0x37, 0xfe, 0x39, 0xc3, 0x0a, 0x3b, 0x7a, 0xd4, 0x36, 0xc0, 0xa5, 0x16, 0xfc,
	0xd6, 0xdc, 0x81, 0xe6, 0xef, 0xed, 0x0d, 0x1b, 0xbc, 0xa5, 0x69, 0xe1, 0xb6, 0xbd, 0xf4, 0x4c,
	0x84, 0xd7, 0x72, 0xe1, 0xde, 0xc5, 0x95, 0x5f, 0x7b, 0x71, 0xe5, 0xd7, 0x5e, 0x5e, 0xf9, 0xb5,
	0x9f, 0x57, 0xbe, 0x73, 0xb1, 0xf2, 0x9d, 0x17, 0x2b, 0xdf, 0x79, 0xb9, 0xf2, 0x9d, 0x3f, 0x57,
	0xbe, 0xf3, 0xdb, 0x5f, 0x7e, 0xed, 0xbb, 0x96, 0x15, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xc4,
	0x86, 0x52, 0xc5, 0xa2, 0x07, 0x00, 0x00,
}
