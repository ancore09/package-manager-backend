// Code generated by protoc-gen-go. DO NOT EDIT
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/package-service.proto

package package_service

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_package_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_api_package_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_api_package_service_proto_rawDescGZIP(), []int{0}
}

func (x *Package) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetPackageByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPackageByNameRequest) Reset() {
	*x = GetPackageByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_package_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackageByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackageByNameRequest) ProtoMessage() {}

func (x *GetPackageByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_package_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackageByNameRequest.ProtoReflect.Descriptor instead.
func (*GetPackageByNameRequest) Descriptor() ([]byte, []int) {
	return file_api_package_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetPackageByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetPackageByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *Package `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *GetPackageByNameResponse) Reset() {
	*x = GetPackageByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_package_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackageByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackageByNameResponse) ProtoMessage() {}

func (x *GetPackageByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_package_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackageByNameResponse.ProtoReflect.Descriptor instead.
func (*GetPackageByNameResponse) Descriptor() ([]byte, []int) {
	return file_api_package_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetPackageByNameResponse) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type GetPackagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPackagesRequest) Reset() {
	*x = GetPackagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_package_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackagesRequest) ProtoMessage() {}

func (x *GetPackagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_package_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackagesRequest.ProtoReflect.Descriptor instead.
func (*GetPackagesRequest) Descriptor() ([]byte, []int) {
	return file_api_package_service_proto_rawDescGZIP(), []int{3}
}

type GetPackagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package []*Package `protobuf:"bytes,1,rep,name=package,proto3" json:"package,omitempty"`
}

func (x *GetPackagesResponse) Reset() {
	*x = GetPackagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_package_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackagesResponse) ProtoMessage() {}

func (x *GetPackagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_package_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackagesResponse.ProtoReflect.Descriptor instead.
func (*GetPackagesResponse) Descriptor() ([]byte, []int) {
	return file_api_package_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetPackagesResponse) GetPackage() []*Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type CreatePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreatePackageRequest) Reset() {
	*x = CreatePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_package_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePackageRequest) ProtoMessage() {}

func (x *CreatePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_package_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePackageRequest.ProtoReflect.Descriptor instead.
func (*CreatePackageRequest) Descriptor() ([]byte, []int) {
	return file_api_package_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreatePackageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreatePackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *Package `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *CreatePackageResponse) Reset() {
	*x = CreatePackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_package_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePackageResponse) ProtoMessage() {}

func (x *CreatePackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_package_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePackageResponse.ProtoReflect.Descriptor instead.
func (*CreatePackageResponse) Descriptor() ([]byte, []int) {
	return file_api_package_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreatePackageResponse) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type UpdatePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *Package `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *UpdatePackageRequest) Reset() {
	*x = UpdatePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_package_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePackageRequest) ProtoMessage() {}

func (x *UpdatePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_package_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePackageRequest.ProtoReflect.Descriptor instead.
func (*UpdatePackageRequest) Descriptor() ([]byte, []int) {
	return file_api_package_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePackageRequest) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type UpdatePackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *Package `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *UpdatePackageResponse) Reset() {
	*x = UpdatePackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_package_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePackageResponse) ProtoMessage() {}

func (x *UpdatePackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_package_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePackageResponse.ProtoReflect.Descriptor instead.
func (*UpdatePackageResponse) Descriptor() ([]byte, []int) {
	return file_api_package_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatePackageResponse) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type DeletePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePackageRequest) Reset() {
	*x = DeletePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_package_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePackageRequest) ProtoMessage() {}

func (x *DeletePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_package_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePackageRequest.ProtoReflect.Descriptor instead.
func (*DeletePackageRequest) Descriptor() ([]byte, []int) {
	return file_api_package_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeletePackageRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeletePackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Completed bool `protobuf:"varint,1,opt,name=completed,proto3" json:"completed,omitempty"`
}

func (x *DeletePackageResponse) Reset() {
	*x = DeletePackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_package_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePackageResponse) ProtoMessage() {}

func (x *DeletePackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_package_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePackageResponse.ProtoReflect.Descriptor instead.
func (*DeletePackageResponse) Descriptor() ([]byte, []int) {
	return file_api_package_service_proto_rawDescGZIP(), []int{10}
}

func (x *DeletePackageResponse) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

var File_api_package_service_proto protoreflect.FileDescriptor

var file_api_package_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30, 0x61, 0x6e, 0x63,
	0x6f, 0x72, 0x65, 0x30, 0x39, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x6f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53,
	0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x39, 0x2e, 0x61, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x30, 0x39, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6a, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x53, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x39, 0x2e, 0x61, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x30, 0x39, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x6c, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x07, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x61, 0x6e,
	0x63, 0x6f, 0x72, 0x65, 0x30, 0x39, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22,
	0x6b, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x61, 0x6e, 0x63, 0x6f, 0x72,
	0x65, 0x30, 0x39, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x6c, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x61, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x30,
	0x39, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x35, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x32, 0xc2, 0x06, 0x0a, 0x0e, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa9, 0x01, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x49, 0x2e, 0x61, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x30, 0x39, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4a, 0x2e, 0x61,
	0x6e, 0x63, 0x6f, 0x72, 0x65, 0x30, 0x39, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9a, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x44, 0x2e, 0x61, 0x6e, 0x63, 0x6f, 0x72,
	0x65, 0x30, 0x39, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45,
	0x2e, 0x61, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x30, 0x39, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x46, 0x2e, 0x61, 0x6e, 0x63, 0x6f, 0x72, 0x65,
	0x30, 0x39, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x47, 0x2e, 0x61, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x30, 0x39, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x46, 0x2e, 0x61, 0x6e, 0x63,
	0x6f, 0x72, 0x65, 0x30, 0x39, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x47, 0x2e, 0x61, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x30, 0x39, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x46, 0x2e,
	0x61, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x30, 0x39, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x61, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x30, 0x39,
	0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27,
	0x5a, 0x25, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_package_service_proto_rawDescOnce sync.Once
	file_api_package_service_proto_rawDescData = file_api_package_service_proto_rawDesc
)

func file_api_package_service_proto_rawDescGZIP() []byte {
	file_api_package_service_proto_rawDescOnce.Do(func() {
		file_api_package_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_package_service_proto_rawDescData)
	})
	return file_api_package_service_proto_rawDescData
}

var file_api_package_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_package_service_proto_goTypes = []interface{}{
	(*Package)(nil),                  // 0: ancore09.package_manager_backend.package_service.Package
	(*GetPackageByNameRequest)(nil),  // 1: ancore09.package_manager_backend.package_service.GetPackageByNameRequest
	(*GetPackageByNameResponse)(nil), // 2: ancore09.package_manager_backend.package_service.GetPackageByNameResponse
	(*GetPackagesRequest)(nil),       // 3: ancore09.package_manager_backend.package_service.GetPackagesRequest
	(*GetPackagesResponse)(nil),      // 4: ancore09.package_manager_backend.package_service.GetPackagesResponse
	(*CreatePackageRequest)(nil),     // 5: ancore09.package_manager_backend.package_service.CreatePackageRequest
	(*CreatePackageResponse)(nil),    // 6: ancore09.package_manager_backend.package_service.CreatePackageResponse
	(*UpdatePackageRequest)(nil),     // 7: ancore09.package_manager_backend.package_service.UpdatePackageRequest
	(*UpdatePackageResponse)(nil),    // 8: ancore09.package_manager_backend.package_service.UpdatePackageResponse
	(*DeletePackageRequest)(nil),     // 9: ancore09.package_manager_backend.package_service.DeletePackageRequest
	(*DeletePackageResponse)(nil),    // 10: ancore09.package_manager_backend.package_service.DeletePackageResponse
}
var file_api_package_service_proto_depIdxs = []int32{
	0,  // 0: ancore09.package_manager_backend.package_service.GetPackageByNameResponse.package:type_name -> ancore09.package_manager_backend.package_service.Package
	0,  // 1: ancore09.package_manager_backend.package_service.GetPackagesResponse.package:type_name -> ancore09.package_manager_backend.package_service.Package
	0,  // 2: ancore09.package_manager_backend.package_service.CreatePackageResponse.package:type_name -> ancore09.package_manager_backend.package_service.Package
	0,  // 3: ancore09.package_manager_backend.package_service.UpdatePackageRequest.package:type_name -> ancore09.package_manager_backend.package_service.Package
	0,  // 4: ancore09.package_manager_backend.package_service.UpdatePackageResponse.package:type_name -> ancore09.package_manager_backend.package_service.Package
	1,  // 5: ancore09.package_manager_backend.package_service.PackageService.GetPackageByName:input_type -> ancore09.package_manager_backend.package_service.GetPackageByNameRequest
	3,  // 6: ancore09.package_manager_backend.package_service.PackageService.GetPackages:input_type -> ancore09.package_manager_backend.package_service.GetPackagesRequest
	5,  // 7: ancore09.package_manager_backend.package_service.PackageService.CreatePackage:input_type -> ancore09.package_manager_backend.package_service.CreatePackageRequest
	7,  // 8: ancore09.package_manager_backend.package_service.PackageService.UpdatePackage:input_type -> ancore09.package_manager_backend.package_service.UpdatePackageRequest
	9,  // 9: ancore09.package_manager_backend.package_service.PackageService.DeletePackage:input_type -> ancore09.package_manager_backend.package_service.DeletePackageRequest
	2,  // 10: ancore09.package_manager_backend.package_service.PackageService.GetPackageByName:output_type -> ancore09.package_manager_backend.package_service.GetPackageByNameResponse
	4,  // 11: ancore09.package_manager_backend.package_service.PackageService.GetPackages:output_type -> ancore09.package_manager_backend.package_service.GetPackagesResponse
	6,  // 12: ancore09.package_manager_backend.package_service.PackageService.CreatePackage:output_type -> ancore09.package_manager_backend.package_service.CreatePackageResponse
	8,  // 13: ancore09.package_manager_backend.package_service.PackageService.UpdatePackage:output_type -> ancore09.package_manager_backend.package_service.UpdatePackageResponse
	10, // 14: ancore09.package_manager_backend.package_service.PackageService.DeletePackage:output_type -> ancore09.package_manager_backend.package_service.DeletePackageResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_package_service_proto_init() }
func file_api_package_service_proto_init() {
	if File_api_package_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_package_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_api_package_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackageByNameRequest); i {
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
		file_api_package_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackageByNameResponse); i {
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
		file_api_package_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackagesRequest); i {
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
		file_api_package_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackagesResponse); i {
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
		file_api_package_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePackageRequest); i {
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
		file_api_package_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePackageResponse); i {
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
		file_api_package_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePackageRequest); i {
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
		file_api_package_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePackageResponse); i {
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
		file_api_package_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePackageRequest); i {
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
		file_api_package_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePackageResponse); i {
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
			RawDescriptor: file_api_package_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_package_service_proto_goTypes,
		DependencyIndexes: file_api_package_service_proto_depIdxs,
		MessageInfos:      file_api_package_service_proto_msgTypes,
	}.Build()
	File_api_package_service_proto = out.File
	file_api_package_service_proto_rawDesc = nil
	file_api_package_service_proto_goTypes = nil
	file_api_package_service_proto_depIdxs = nil
}
