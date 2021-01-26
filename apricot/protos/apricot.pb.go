//
// === This file is part of ALICE O² ===
//
// Copyright 2020 CERN and copyright holders of ALICE O².
// Author: Teo Mrnjavac <teo.mrnjavac@cern.ch>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// In applying this license CERN does not waive the privileges and
// immunities granted to it by virtue of its status as an
// Intergovernmental Organization or submit itself to any jurisdiction.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: protos/apricot.proto

package apricotpb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RunType int32

const (
	RunType_NULL           RunType = 0
	RunType_ANY            RunType = 1
	RunType_PHYSICS        RunType = 2
	RunType_TECHNICAL      RunType = 3
	RunType_COSMIC         RunType = 4
	RunType_PEDESTAL       RunType = 5
	RunType_THRESHOLD_SCAN RunType = 6
	RunType_CALIBRATION    RunType = 7
)

// Enum value maps for RunType.
var (
	RunType_name = map[int32]string{
		0: "NULL",
		1: "ANY",
		2: "PHYSICS",
		3: "TECHNICAL",
		4: "COSMIC",
		5: "PEDESTAL",
		6: "THRESHOLD_SCAN",
		7: "CALIBRATION",
	}
	RunType_value = map[string]int32{
		"NULL":           0,
		"ANY":            1,
		"PHYSICS":        2,
		"TECHNICAL":      3,
		"COSMIC":         4,
		"PEDESTAL":       5,
		"THRESHOLD_SCAN": 6,
		"CALIBRATION":    7,
	}
)

func (x RunType) Enum() *RunType {
	p := new(RunType)
	*p = x
	return p
}

func (x RunType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_apricot_proto_enumTypes[0].Descriptor()
}

func (RunType) Type() protoreflect.EnumType {
	return &file_protos_apricot_proto_enumTypes[0]
}

func (x RunType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunType.Descriptor instead.
func (RunType) EnumDescriptor() ([]byte, []int) {
	return file_protos_apricot_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_apricot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_apricot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protos_apricot_proto_rawDescGZIP(), []int{0}
}

type ComponentQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Component   string  `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
	RunType     RunType `protobuf:"varint,2,opt,name=runType,proto3,enum=apricot.RunType" json:"runType,omitempty"`
	MachineRole string  `protobuf:"bytes,3,opt,name=machineRole,proto3" json:"machineRole,omitempty"`
	Entry       string  `protobuf:"bytes,4,opt,name=entry,proto3" json:"entry,omitempty"`
	Timestamp   string  `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ComponentQuery) Reset() {
	*x = ComponentQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_apricot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentQuery) ProtoMessage() {}

func (x *ComponentQuery) ProtoReflect() protoreflect.Message {
	mi := &file_protos_apricot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentQuery.ProtoReflect.Descriptor instead.
func (*ComponentQuery) Descriptor() ([]byte, []int) {
	return file_protos_apricot_proto_rawDescGZIP(), []int{1}
}

func (x *ComponentQuery) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *ComponentQuery) GetRunType() RunType {
	if x != nil {
		return x.RunType
	}
	return RunType_NULL
}

func (x *ComponentQuery) GetMachineRole() string {
	if x != nil {
		return x.MachineRole
	}
	return ""
}

func (x *ComponentQuery) GetEntry() string {
	if x != nil {
		return x.Entry
	}
	return ""
}

func (x *ComponentQuery) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type ComponentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to QueryPath:
	//	*ComponentRequest_Path
	//	*ComponentRequest_Query
	QueryPath       isComponentRequest_QueryPath `protobuf_oneof:"queryPath"`
	ProcessTemplate bool                         `protobuf:"varint,3,opt,name=processTemplate,proto3" json:"processTemplate,omitempty"`
	VarStack        map[string]string            `protobuf:"bytes,4,rep,name=varStack,proto3" json:"varStack,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ComponentRequest) Reset() {
	*x = ComponentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_apricot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentRequest) ProtoMessage() {}

func (x *ComponentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_apricot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentRequest.ProtoReflect.Descriptor instead.
func (*ComponentRequest) Descriptor() ([]byte, []int) {
	return file_protos_apricot_proto_rawDescGZIP(), []int{2}
}

func (m *ComponentRequest) GetQueryPath() isComponentRequest_QueryPath {
	if m != nil {
		return m.QueryPath
	}
	return nil
}

func (x *ComponentRequest) GetPath() string {
	if x, ok := x.GetQueryPath().(*ComponentRequest_Path); ok {
		return x.Path
	}
	return ""
}

func (x *ComponentRequest) GetQuery() *ComponentQuery {
	if x, ok := x.GetQueryPath().(*ComponentRequest_Query); ok {
		return x.Query
	}
	return nil
}

func (x *ComponentRequest) GetProcessTemplate() bool {
	if x != nil {
		return x.ProcessTemplate
	}
	return false
}

func (x *ComponentRequest) GetVarStack() map[string]string {
	if x != nil {
		return x.VarStack
	}
	return nil
}

type isComponentRequest_QueryPath interface {
	isComponentRequest_QueryPath()
}

type ComponentRequest_Path struct {
	Path string `protobuf:"bytes,1,opt,name=path,proto3,oneof"`
}

type ComponentRequest_Query struct {
	Query *ComponentQuery `protobuf:"bytes,2,opt,name=query,proto3,oneof"`
}

func (*ComponentRequest_Path) isComponentRequest_QueryPath() {}

func (*ComponentRequest_Query) isComponentRequest_QueryPath() {}

type ComponentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ComponentResponse) Reset() {
	*x = ComponentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_apricot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentResponse) ProtoMessage() {}

func (x *ComponentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_apricot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentResponse.ProtoReflect.Descriptor instead.
func (*ComponentResponse) Descriptor() ([]byte, []int) {
	return file_protos_apricot_proto_rawDescGZIP(), []int{3}
}

func (x *ComponentResponse) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type RunNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunNumber uint32 `protobuf:"varint,1,opt,name=runNumber,proto3" json:"runNumber,omitempty"`
}

func (x *RunNumberResponse) Reset() {
	*x = RunNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_apricot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunNumberResponse) ProtoMessage() {}

func (x *RunNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_apricot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunNumberResponse.ProtoReflect.Descriptor instead.
func (*RunNumberResponse) Descriptor() ([]byte, []int) {
	return file_protos_apricot_proto_rawDescGZIP(), []int{4}
}

func (x *RunNumberResponse) GetRunNumber() uint32 {
	if x != nil {
		return x.RunNumber
	}
	return 0
}

type StringMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringMap map[string]string `protobuf:"bytes,1,rep,name=stringMap,proto3" json:"stringMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *StringMap) Reset() {
	*x = StringMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_apricot_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMap) ProtoMessage() {}

func (x *StringMap) ProtoReflect() protoreflect.Message {
	mi := &file_protos_apricot_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMap.ProtoReflect.Descriptor instead.
func (*StringMap) Descriptor() ([]byte, []int) {
	return file_protos_apricot_proto_rawDescGZIP(), []int{5}
}

func (x *StringMap) GetStringMap() map[string]string {
	if x != nil {
		return x.StringMap
	}
	return nil
}

type RawGetRecursiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawPath string `protobuf:"bytes,1,opt,name=rawPath,proto3" json:"rawPath,omitempty"`
}

func (x *RawGetRecursiveRequest) Reset() {
	*x = RawGetRecursiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_apricot_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawGetRecursiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawGetRecursiveRequest) ProtoMessage() {}

func (x *RawGetRecursiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_apricot_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawGetRecursiveRequest.ProtoReflect.Descriptor instead.
func (*RawGetRecursiveRequest) Descriptor() ([]byte, []int) {
	return file_protos_apricot_proto_rawDescGZIP(), []int{6}
}

func (x *RawGetRecursiveRequest) GetRawPath() string {
	if x != nil {
		return x.RawPath
	}
	return ""
}

type GetRuntimeEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Component string `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetRuntimeEntryRequest) Reset() {
	*x = GetRuntimeEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_apricot_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRuntimeEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuntimeEntryRequest) ProtoMessage() {}

func (x *GetRuntimeEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_apricot_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuntimeEntryRequest.ProtoReflect.Descriptor instead.
func (*GetRuntimeEntryRequest) Descriptor() ([]byte, []int) {
	return file_protos_apricot_proto_rawDescGZIP(), []int{7}
}

func (x *GetRuntimeEntryRequest) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *GetRuntimeEntryRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type SetRuntimeEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Component string `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value     string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetRuntimeEntryRequest) Reset() {
	*x = SetRuntimeEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_apricot_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRuntimeEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRuntimeEntryRequest) ProtoMessage() {}

func (x *SetRuntimeEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_apricot_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRuntimeEntryRequest.ProtoReflect.Descriptor instead.
func (*SetRuntimeEntryRequest) Descriptor() ([]byte, []int) {
	return file_protos_apricot_proto_rawDescGZIP(), []int{8}
}

func (x *SetRuntimeEntryRequest) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *SetRuntimeEntryRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetRuntimeEntryRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_protos_apricot_proto protoreflect.FileDescriptor

var file_protos_apricot_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xb0, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x72, 0x75, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x72,
	0x69, 0x63, 0x6f, 0x74, 0x2e, 0x52, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x72, 0x75,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x92, 0x02, 0x0a, 0x10,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x2f, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x43, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56,
	0x61, 0x72, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x76, 0x61,
	0x72, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x1a, 0x3b, 0x0a, 0x0d, 0x56, 0x61, 0x72, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x2d, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x31, 0x0a, 0x11, 0x52, 0x75, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x75, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x75, 0x6e, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x8a, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70,
	0x12, 0x3f, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61,
	0x70, 0x1a, 0x3c, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x32, 0x0a, 0x16, 0x52, 0x61, 0x77, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x77,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x77, 0x50,
	0x61, 0x74, 0x68, 0x22, 0x48, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x5e, 0x0a,
	0x16, 0x53, 0x65, 0x74, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x77, 0x0a,
	0x07, 0x52, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c, 0x4c,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4e, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x48, 0x59, 0x53, 0x49, 0x43, 0x53, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x45, 0x43, 0x48,
	0x4e, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x53, 0x4d, 0x49,
	0x43, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x45, 0x44, 0x45, 0x53, 0x54, 0x41, 0x4c, 0x10,
	0x05, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x53,
	0x43, 0x41, 0x4e, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x41, 0x4c, 0x49, 0x42, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x32, 0xed, 0x03, 0x0a, 0x07, 0x41, 0x70, 0x72, 0x69, 0x63,
	0x6f, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x52, 0x75, 0x6e, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x52, 0x75, 0x6e,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x33, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x0e, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x12, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x4d, 0x61, 0x70, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x56, 0x61, 0x72, 0x73,
	0x12, 0x0e, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x12, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x4d, 0x61, 0x70, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0f,
	0x52, 0x61, 0x77, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x12,
	0x1f, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x52, 0x61, 0x77, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x53, 0x65,
	0x74, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x36, 0x0a, 0x22, 0x63, 0x68, 0x2e, 0x63, 0x65, 0x72,
	0x6e, 0x2e, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x6f, 0x32, 0x2e, 0x61, 0x70, 0x72, 0x69, 0x63,
	0x6f, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5a, 0x10, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x3b, 0x61, 0x70, 0x72, 0x69, 0x63, 0x6f, 0x74, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_apricot_proto_rawDescOnce sync.Once
	file_protos_apricot_proto_rawDescData = file_protos_apricot_proto_rawDesc
)

func file_protos_apricot_proto_rawDescGZIP() []byte {
	file_protos_apricot_proto_rawDescOnce.Do(func() {
		file_protos_apricot_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_apricot_proto_rawDescData)
	})
	return file_protos_apricot_proto_rawDescData
}

var file_protos_apricot_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_apricot_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_protos_apricot_proto_goTypes = []interface{}{
	(RunType)(0),                   // 0: apricot.RunType
	(*Empty)(nil),                  // 1: apricot.Empty
	(*ComponentQuery)(nil),         // 2: apricot.ComponentQuery
	(*ComponentRequest)(nil),       // 3: apricot.ComponentRequest
	(*ComponentResponse)(nil),      // 4: apricot.ComponentResponse
	(*RunNumberResponse)(nil),      // 5: apricot.RunNumberResponse
	(*StringMap)(nil),              // 6: apricot.StringMap
	(*RawGetRecursiveRequest)(nil), // 7: apricot.RawGetRecursiveRequest
	(*GetRuntimeEntryRequest)(nil), // 8: apricot.GetRuntimeEntryRequest
	(*SetRuntimeEntryRequest)(nil), // 9: apricot.SetRuntimeEntryRequest
	nil,                            // 10: apricot.ComponentRequest.VarStackEntry
	nil,                            // 11: apricot.StringMap.StringMapEntry
}
var file_protos_apricot_proto_depIdxs = []int32{
	0,  // 0: apricot.ComponentQuery.runType:type_name -> apricot.RunType
	2,  // 1: apricot.ComponentRequest.query:type_name -> apricot.ComponentQuery
	10, // 2: apricot.ComponentRequest.varStack:type_name -> apricot.ComponentRequest.VarStackEntry
	11, // 3: apricot.StringMap.stringMap:type_name -> apricot.StringMap.StringMapEntry
	1,  // 4: apricot.Apricot.NewRunNumber:input_type -> apricot.Empty
	1,  // 5: apricot.Apricot.GetDefaults:input_type -> apricot.Empty
	1,  // 6: apricot.Apricot.GetVars:input_type -> apricot.Empty
	3,  // 7: apricot.Apricot.GetComponentConfiguration:input_type -> apricot.ComponentRequest
	7,  // 8: apricot.Apricot.RawGetRecursive:input_type -> apricot.RawGetRecursiveRequest
	8,  // 9: apricot.Apricot.GetRuntimeEntry:input_type -> apricot.GetRuntimeEntryRequest
	9,  // 10: apricot.Apricot.SetRuntimeEntry:input_type -> apricot.SetRuntimeEntryRequest
	5,  // 11: apricot.Apricot.NewRunNumber:output_type -> apricot.RunNumberResponse
	6,  // 12: apricot.Apricot.GetDefaults:output_type -> apricot.StringMap
	6,  // 13: apricot.Apricot.GetVars:output_type -> apricot.StringMap
	4,  // 14: apricot.Apricot.GetComponentConfiguration:output_type -> apricot.ComponentResponse
	4,  // 15: apricot.Apricot.RawGetRecursive:output_type -> apricot.ComponentResponse
	4,  // 16: apricot.Apricot.GetRuntimeEntry:output_type -> apricot.ComponentResponse
	1,  // 17: apricot.Apricot.SetRuntimeEntry:output_type -> apricot.Empty
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_protos_apricot_proto_init() }
func file_protos_apricot_proto_init() {
	if File_protos_apricot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_apricot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_apricot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_apricot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_apricot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_apricot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunNumberResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_apricot_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMap); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_apricot_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawGetRecursiveRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_apricot_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRuntimeEntryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_apricot_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRuntimeEntryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_protos_apricot_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ComponentRequest_Path)(nil),
		(*ComponentRequest_Query)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_apricot_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_apricot_proto_goTypes,
		DependencyIndexes: file_protos_apricot_proto_depIdxs,
		EnumInfos:         file_protos_apricot_proto_enumTypes,
		MessageInfos:      file_protos_apricot_proto_msgTypes,
	}.Build()
	File_protos_apricot_proto = out.File
	file_protos_apricot_proto_rawDesc = nil
	file_protos_apricot_proto_goTypes = nil
	file_protos_apricot_proto_depIdxs = nil
}