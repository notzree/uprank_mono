// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: proto/inference.proto

package inference_backend

import (
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

type EmbedTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text     string            `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EmbedTextRequest) Reset() {
	*x = EmbedTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmbedTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmbedTextRequest) ProtoMessage() {}

func (x *EmbedTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmbedTextRequest.ProtoReflect.Descriptor instead.
func (*EmbedTextRequest) Descriptor() ([]byte, []int) {
	return file_proto_inference_proto_rawDescGZIP(), []int{0}
}

func (x *EmbedTextRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EmbedTextRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *EmbedTextRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type EmbedTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vector *Vector `protobuf:"bytes,1,opt,name=vector,proto3" json:"vector,omitempty"`
	Error  string  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *EmbedTextResponse) Reset() {
	*x = EmbedTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inference_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmbedTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmbedTextResponse) ProtoMessage() {}

func (x *EmbedTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inference_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmbedTextResponse.ProtoReflect.Descriptor instead.
func (*EmbedTextResponse) Descriptor() ([]byte, []int) {
	return file_proto_inference_proto_rawDescGZIP(), []int{1}
}

func (x *EmbedTextResponse) GetVector() *Vector {
	if x != nil {
		return x.Vector
	}
	return nil
}

func (x *EmbedTextResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpsertVectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string    `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Vectors   []*Vector `protobuf:"bytes,2,rep,name=vectors,proto3" json:"vectors,omitempty"`
}

func (x *UpsertVectorRequest) Reset() {
	*x = UpsertVectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inference_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertVectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertVectorRequest) ProtoMessage() {}

func (x *UpsertVectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inference_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertVectorRequest.ProtoReflect.Descriptor instead.
func (*UpsertVectorRequest) Descriptor() ([]byte, []int) {
	return file_proto_inference_proto_rawDescGZIP(), []int{2}
}

func (x *UpsertVectorRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *UpsertVectorRequest) GetVectors() []*Vector {
	if x != nil {
		return x.Vectors
	}
	return nil
}

type Vector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Vector   []float32         `protobuf:"fixed32,2,rep,packed,name=vector,proto3" json:"vector,omitempty"`
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Vector) Reset() {
	*x = Vector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inference_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector) ProtoMessage() {}

func (x *Vector) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inference_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector.ProtoReflect.Descriptor instead.
func (*Vector) Descriptor() ([]byte, []int) {
	return file_proto_inference_proto_rawDescGZIP(), []int{3}
}

func (x *Vector) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Vector) GetVector() []float32 {
	if x != nil {
		return x.Vector
	}
	return nil
}

func (x *Vector) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type UpsertVectorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpsertVectorResponse) Reset() {
	*x = UpsertVectorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inference_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertVectorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertVectorResponse) ProtoMessage() {}

func (x *UpsertVectorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inference_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertVectorResponse.ProtoReflect.Descriptor instead.
func (*UpsertVectorResponse) Descriptor() ([]byte, []int) {
	return file_proto_inference_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertVectorResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type QueryVectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Vector    []float32         `protobuf:"fixed32,2,rep,packed,name=vector,proto3" json:"vector,omitempty"`
	TopK      int32             `protobuf:"varint,3,opt,name=top_k,json=topK,proto3" json:"top_k,omitempty"`
	Filter    map[string]string `protobuf:"bytes,4,rep,name=filter,proto3" json:"filter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *QueryVectorRequest) Reset() {
	*x = QueryVectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inference_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryVectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryVectorRequest) ProtoMessage() {}

func (x *QueryVectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inference_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryVectorRequest.ProtoReflect.Descriptor instead.
func (*QueryVectorRequest) Descriptor() ([]byte, []int) {
	return file_proto_inference_proto_rawDescGZIP(), []int{5}
}

func (x *QueryVectorRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryVectorRequest) GetVector() []float32 {
	if x != nil {
		return x.Vector
	}
	return nil
}

func (x *QueryVectorRequest) GetTopK() int32 {
	if x != nil {
		return x.TopK
	}
	return 0
}

func (x *QueryVectorRequest) GetFilter() map[string]string {
	if x != nil {
		return x.Filter
	}
	return nil
}

// Define the response message
type QueryVectorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matches   []*Match `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
	Namespace string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Usage     *Usage   `protobuf:"bytes,3,opt,name=usage,proto3" json:"usage,omitempty"`
}

func (x *QueryVectorResponse) Reset() {
	*x = QueryVectorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inference_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryVectorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryVectorResponse) ProtoMessage() {}

func (x *QueryVectorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inference_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryVectorResponse.ProtoReflect.Descriptor instead.
func (*QueryVectorResponse) Descriptor() ([]byte, []int) {
	return file_proto_inference_proto_rawDescGZIP(), []int{6}
}

func (x *QueryVectorResponse) GetMatches() []*Match {
	if x != nil {
		return x.Matches
	}
	return nil
}

func (x *QueryVectorResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryVectorResponse) GetUsage() *Usage {
	if x != nil {
		return x.Usage
	}
	return nil
}

// Define the Match message
type Match struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Score  float32   `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	Values []float32 `protobuf:"fixed32,3,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *Match) Reset() {
	*x = Match{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inference_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Match) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Match) ProtoMessage() {}

func (x *Match) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inference_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Match.ProtoReflect.Descriptor instead.
func (*Match) Descriptor() ([]byte, []int) {
	return file_proto_inference_proto_rawDescGZIP(), []int{7}
}

func (x *Match) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Match) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Match) GetValues() []float32 {
	if x != nil {
		return x.Values
	}
	return nil
}

// Define the Usage message
type Usage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReadUnits int32 `protobuf:"varint,1,opt,name=read_units,json=readUnits,proto3" json:"read_units,omitempty"`
}

func (x *Usage) Reset() {
	*x = Usage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inference_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Usage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Usage) ProtoMessage() {}

func (x *Usage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inference_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Usage.ProtoReflect.Descriptor instead.
func (*Usage) Descriptor() ([]byte, []int) {
	return file_proto_inference_proto_rawDescGZIP(), []int{8}
}

func (x *Usage) GetReadUnits() int32 {
	if x != nil {
		return x.ReadUnits
	}
	return 0
}

type DeleteVectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids       []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	Namespace string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *DeleteVectorRequest) Reset() {
	*x = DeleteVectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inference_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVectorRequest) ProtoMessage() {}

func (x *DeleteVectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inference_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVectorRequest.ProtoReflect.Descriptor instead.
func (*DeleteVectorRequest) Descriptor() ([]byte, []int) {
	return file_proto_inference_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteVectorRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *DeleteVectorRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type DeleteVectorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids   []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	Ok    bool     `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
	Error string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteVectorResponse) Reset() {
	*x = DeleteVectorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inference_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVectorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVectorResponse) ProtoMessage() {}

func (x *DeleteVectorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inference_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVectorResponse.ProtoReflect.Descriptor instead.
func (*DeleteVectorResponse) Descriptor() ([]byte, []int) {
	return file_proto_inference_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteVectorResponse) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *DeleteVectorResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *DeleteVectorResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_inference_proto protoreflect.FileDescriptor

var file_proto_inference_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x10, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x45, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x54,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x54, 0x0a, 0x11, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x60, 0x0a, 0x13, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x69, 0x6e,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07,
	0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x06, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x02, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69,
	0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x2c, 0x0a, 0x14, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0xdd, 0x01, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x02, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x13, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x5f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x74, 0x6f, 0x70, 0x4b, 0x12, 0x41, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x87, 0x01, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6e,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x07, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x22, 0x45, 0x0a, 0x05,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x02, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x61, 0x64, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x72, 0x65, 0x61, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x22, 0x45, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x22, 0x4e, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0xc3, 0x02, 0x0a, 0x09, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x46, 0x0a, 0x09, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1b, 0x2e,
	0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x54,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x55, 0x70, 0x73, 0x65,
	0x72, 0x74, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x74, 0x7a, 0x72, 0x65, 0x65, 0x2f, 0x75,
	0x70, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x61, 0x69, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_inference_proto_rawDescOnce sync.Once
	file_proto_inference_proto_rawDescData = file_proto_inference_proto_rawDesc
)

func file_proto_inference_proto_rawDescGZIP() []byte {
	file_proto_inference_proto_rawDescOnce.Do(func() {
		file_proto_inference_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_inference_proto_rawDescData)
	})
	return file_proto_inference_proto_rawDescData
}

var file_proto_inference_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_inference_proto_goTypes = []interface{}{
	(*EmbedTextRequest)(nil),     // 0: inference.EmbedTextRequest
	(*EmbedTextResponse)(nil),    // 1: inference.EmbedTextResponse
	(*UpsertVectorRequest)(nil),  // 2: inference.UpsertVectorRequest
	(*Vector)(nil),               // 3: inference.Vector
	(*UpsertVectorResponse)(nil), // 4: inference.UpsertVectorResponse
	(*QueryVectorRequest)(nil),   // 5: inference.QueryVectorRequest
	(*QueryVectorResponse)(nil),  // 6: inference.QueryVectorResponse
	(*Match)(nil),                // 7: inference.Match
	(*Usage)(nil),                // 8: inference.Usage
	(*DeleteVectorRequest)(nil),  // 9: inference.DeleteVectorRequest
	(*DeleteVectorResponse)(nil), // 10: inference.DeleteVectorResponse
	nil,                          // 11: inference.EmbedTextRequest.MetadataEntry
	nil,                          // 12: inference.Vector.MetadataEntry
	nil,                          // 13: inference.QueryVectorRequest.FilterEntry
}
var file_proto_inference_proto_depIdxs = []int32{
	11, // 0: inference.EmbedTextRequest.metadata:type_name -> inference.EmbedTextRequest.MetadataEntry
	3,  // 1: inference.EmbedTextResponse.vector:type_name -> inference.Vector
	3,  // 2: inference.UpsertVectorRequest.vectors:type_name -> inference.Vector
	12, // 3: inference.Vector.metadata:type_name -> inference.Vector.MetadataEntry
	13, // 4: inference.QueryVectorRequest.filter:type_name -> inference.QueryVectorRequest.FilterEntry
	7,  // 5: inference.QueryVectorResponse.matches:type_name -> inference.Match
	8,  // 6: inference.QueryVectorResponse.usage:type_name -> inference.Usage
	0,  // 7: inference.Inference.EmbedText:input_type -> inference.EmbedTextRequest
	2,  // 8: inference.Inference.UpsertVector:input_type -> inference.UpsertVectorRequest
	5,  // 9: inference.Inference.QueryVector:input_type -> inference.QueryVectorRequest
	9,  // 10: inference.Inference.DeleteVector:input_type -> inference.DeleteVectorRequest
	1,  // 11: inference.Inference.EmbedText:output_type -> inference.EmbedTextResponse
	4,  // 12: inference.Inference.UpsertVector:output_type -> inference.UpsertVectorResponse
	6,  // 13: inference.Inference.QueryVector:output_type -> inference.QueryVectorResponse
	10, // 14: inference.Inference.DeleteVector:output_type -> inference.DeleteVectorResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_inference_proto_init() }
func file_proto_inference_proto_init() {
	if File_proto_inference_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_inference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmbedTextRequest); i {
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
		file_proto_inference_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmbedTextResponse); i {
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
		file_proto_inference_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertVectorRequest); i {
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
		file_proto_inference_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vector); i {
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
		file_proto_inference_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertVectorResponse); i {
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
		file_proto_inference_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryVectorRequest); i {
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
		file_proto_inference_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryVectorResponse); i {
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
		file_proto_inference_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Match); i {
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
		file_proto_inference_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Usage); i {
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
		file_proto_inference_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVectorRequest); i {
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
		file_proto_inference_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVectorResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_inference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_inference_proto_goTypes,
		DependencyIndexes: file_proto_inference_proto_depIdxs,
		MessageInfos:      file_proto_inference_proto_msgTypes,
	}.Build()
	File_proto_inference_proto = out.File
	file_proto_inference_proto_rawDesc = nil
	file_proto_inference_proto_goTypes = nil
	file_proto_inference_proto_depIdxs = nil
}
