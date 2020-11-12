// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: main_bootstrap_models.proto

package v2

import (
	v21 "github.com/getcouragenow/mod/mod-disco/service/go/rpc/v2"
	v2 "github.com/getcouragenow/sys-share/sys-account/service/go/rpc/v2"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type BSProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project        *v2.ProjectRequest           `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	ProjectDetails *v21.NewDiscoProjectRequest  `protobuf:"bytes,2,opt,name=project_details,json=projectDetails,proto3" json:"project_details,omitempty"`
	SurveySchema   *v21.NewSurveyProjectRequest `protobuf:"bytes,3,opt,name=survey_schema,json=surveySchema,proto3" json:"survey_schema,omitempty"`
}

func (x *BSProject) Reset() {
	*x = BSProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_bootstrap_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BSProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BSProject) ProtoMessage() {}

func (x *BSProject) ProtoReflect() protoreflect.Message {
	mi := &file_main_bootstrap_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BSProject.ProtoReflect.Descriptor instead.
func (*BSProject) Descriptor() ([]byte, []int) {
	return file_main_bootstrap_models_proto_rawDescGZIP(), []int{0}
}

func (x *BSProject) GetProject() *v2.ProjectRequest {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *BSProject) GetProjectDetails() *v21.NewDiscoProjectRequest {
	if x != nil {
		return x.ProjectDetails
	}
	return nil
}

func (x *BSProject) GetSurveySchema() *v21.NewSurveyProjectRequest {
	if x != nil {
		return x.SurveySchema
	}
	return nil
}

type BSOrg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Org *v2.OrgRequest `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty"`
}

func (x *BSOrg) Reset() {
	*x = BSOrg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_bootstrap_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BSOrg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BSOrg) ProtoMessage() {}

func (x *BSOrg) ProtoReflect() protoreflect.Message {
	mi := &file_main_bootstrap_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BSOrg.ProtoReflect.Descriptor instead.
func (*BSOrg) Descriptor() ([]byte, []int) {
	return file_main_bootstrap_models_proto_rawDescGZIP(), []int{1}
}

func (x *BSOrg) GetOrg() *v2.OrgRequest {
	if x != nil {
		return x.Org
	}
	return nil
}

type BSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orgs       []*BSOrg     `protobuf:"bytes,1,rep,name=orgs,proto3" json:"orgs,omitempty"`
	Projects   []*BSProject `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
	Superusers []*BSAccount `protobuf:"bytes,3,rep,name=superusers,proto3" json:"superusers,omitempty"`
}

func (x *BSRequest) Reset() {
	*x = BSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_bootstrap_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BSRequest) ProtoMessage() {}

func (x *BSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_bootstrap_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BSRequest.ProtoReflect.Descriptor instead.
func (*BSRequest) Descriptor() ([]byte, []int) {
	return file_main_bootstrap_models_proto_rawDescGZIP(), []int{2}
}

func (x *BSRequest) GetOrgs() []*BSOrg {
	if x != nil {
		return x.Orgs
	}
	return nil
}

func (x *BSRequest) GetProjects() []*BSProject {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *BSRequest) GetSuperusers() []*BSAccount {
	if x != nil {
		return x.Superusers
	}
	return nil
}

type BSAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitialSuperuser *v2.LoginRequest          `protobuf:"bytes,1,opt,name=initial_superuser,json=initialSuperuser,proto3" json:"initial_superuser,omitempty"`
	SurveyValue      *v21.NewSurveyUserRequest `protobuf:"bytes,2,opt,name=survey_value,json=surveyValue,proto3" json:"survey_value,omitempty"`
}

func (x *BSAccount) Reset() {
	*x = BSAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_bootstrap_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BSAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BSAccount) ProtoMessage() {}

func (x *BSAccount) ProtoReflect() protoreflect.Message {
	mi := &file_main_bootstrap_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BSAccount.ProtoReflect.Descriptor instead.
func (*BSAccount) Descriptor() ([]byte, []int) {
	return file_main_bootstrap_models_proto_rawDescGZIP(), []int{3}
}

func (x *BSAccount) GetInitialSuperuser() *v2.LoginRequest {
	if x != nil {
		return x.InitialSuperuser
	}
	return nil
}

func (x *BSAccount) GetSurveyValue() *v21.NewSurveyUserRequest {
	if x != nil {
		return x.SurveyValue
	}
	return nil
}

type NewBSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// either file_path or file content (in bytes)
	FileExtension string     `protobuf:"bytes,1,opt,name=file_extension,json=fileExtension,proto3" json:"file_extension,omitempty"`
	BsRequest     *BSRequest `protobuf:"bytes,2,opt,name=bs_request,json=bsRequest,proto3" json:"bs_request,omitempty"`
}

func (x *NewBSRequest) Reset() {
	*x = NewBSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_bootstrap_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewBSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBSRequest) ProtoMessage() {}

func (x *NewBSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_bootstrap_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBSRequest.ProtoReflect.Descriptor instead.
func (*NewBSRequest) Descriptor() ([]byte, []int) {
	return file_main_bootstrap_models_proto_rawDescGZIP(), []int{4}
}

func (x *NewBSRequest) GetFileExtension() string {
	if x != nil {
		return x.FileExtension
	}
	return ""
}

func (x *NewBSRequest) GetBsRequest() *BSRequest {
	if x != nil {
		return x.BsRequest
	}
	return nil
}

type NewBSResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (x *NewBSResponse) Reset() {
	*x = NewBSResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_bootstrap_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewBSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBSResponse) ProtoMessage() {}

func (x *NewBSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_bootstrap_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBSResponse.ProtoReflect.Descriptor instead.
func (*NewBSResponse) Descriptor() ([]byte, []int) {
	return file_main_bootstrap_models_proto_rawDescGZIP(), []int{5}
}

func (x *NewBSResponse) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type GetBSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (x *GetBSRequest) Reset() {
	*x = GetBSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_bootstrap_models_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBSRequest) ProtoMessage() {}

func (x *GetBSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_bootstrap_models_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBSRequest.ProtoReflect.Descriptor instead.
func (*GetBSRequest) Descriptor() ([]byte, []int) {
	return file_main_bootstrap_models_proto_rawDescGZIP(), []int{6}
}

func (x *GetBSRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type BS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId    string               `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Content   *BSRequest           `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *BS) Reset() {
	*x = BS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_bootstrap_models_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BS) ProtoMessage() {}

func (x *BS) ProtoReflect() protoreflect.Message {
	mi := &file_main_bootstrap_models_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BS.ProtoReflect.Descriptor instead.
func (*BS) Descriptor() ([]byte, []int) {
	return file_main_bootstrap_models_proto_rawDescGZIP(), []int{7}
}

func (x *BS) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *BS) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *BS) GetContent() *BSRequest {
	if x != nil {
		return x.Content
	}
	return nil
}

type ListBSResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bootstraps []*BS `protobuf:"bytes,1,rep,name=bootstraps,proto3" json:"bootstraps,omitempty"`
}

func (x *ListBSResponse) Reset() {
	*x = ListBSResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_bootstrap_models_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBSResponse) ProtoMessage() {}

func (x *ListBSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_bootstrap_models_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBSResponse.ProtoReflect.Descriptor instead.
func (*ListBSResponse) Descriptor() ([]byte, []int) {
	return file_main_bootstrap_models_proto_rawDescGZIP(), []int{8}
}

func (x *ListBSResponse) GetBootstraps() []*BS {
	if x != nil {
		return x.Bootstraps
	}
	return nil
}

type ListBSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PerPageEntries int64  `protobuf:"varint,1,opt,name=per_page_entries,json=perPageEntries,proto3" json:"per_page_entries,omitempty"`
	OrderBy        string `protobuf:"bytes,2,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	CurrentPageId  string `protobuf:"bytes,3,opt,name=current_page_id,json=currentPageId,proto3" json:"current_page_id,omitempty"`
	IsDescending   bool   `protobuf:"varint,4,opt,name=isDescending,proto3" json:"isDescending,omitempty"`
	Filters        []byte `protobuf:"bytes,5,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListBSRequest) Reset() {
	*x = ListBSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_bootstrap_models_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBSRequest) ProtoMessage() {}

func (x *ListBSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_bootstrap_models_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBSRequest.ProtoReflect.Descriptor instead.
func (*ListBSRequest) Descriptor() ([]byte, []int) {
	return file_main_bootstrap_models_proto_rawDescGZIP(), []int{9}
}

func (x *ListBSRequest) GetPerPageEntries() int64 {
	if x != nil {
		return x.PerPageEntries
	}
	return 0
}

func (x *ListBSRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListBSRequest) GetCurrentPageId() string {
	if x != nil {
		return x.CurrentPageId
	}
	return ""
}

func (x *ListBSRequest) GetIsDescending() bool {
	if x != nil {
		return x.IsDescending
	}
	return false
}

func (x *ListBSRequest) GetFilters() []byte {
	if x != nil {
		return x.Filters
	}
	return nil
}

var File_main_bootstrap_models_proto protoreflect.FileDescriptor

var file_main_bootstrap_models_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x76,
	0x32, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x57, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x6f, 0x77, 0x2f, 0x73, 0x79, 0x73, 0x2d, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x2f, 0x73, 0x79, 0x73, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x79, 0x73, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x4d, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x61, 0x67, 0x65,
	0x6e, 0x6f, 0x77, 0x2f, 0x6d, 0x6f, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x2d, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x6f, 0x64, 0x5f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfb, 0x01, 0x0a, 0x09, 0x42, 0x53, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x41, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x79, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x56, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x76,
	0x32, 0x2e, 0x6d, 0x6f, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x53, 0x0a, 0x0d, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x76, 0x32, 0x2e, 0x6d, 0x6f, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0c, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x22, 0x3e, 0x0a, 0x05, 0x42, 0x53, 0x4f, 0x72, 0x67, 0x12, 0x35, 0x0a, 0x03, 0x6f, 0x72, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x79, 0x73, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x03, 0x6f, 0x72, 0x67,
	0x22, 0xcc, 0x01, 0x0a, 0x09, 0x42, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35,
	0x0a, 0x04, 0x6f, 0x72, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76,
	0x32, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x53, 0x4f, 0x72, 0x67, 0x52,
	0x04, 0x6f, 0x72, 0x67, 0x73, 0x12, 0x41, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x76, 0x32, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x53, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x0a, 0x73, 0x75, 0x70, 0x65,
	0x72, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x76,
	0x32, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x53, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22,
	0xaf, 0x01, 0x0a, 0x09, 0x42, 0x53, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x52, 0x0a,
	0x11, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x79,
	0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x10, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x65, 0x72, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x4e, 0x0a, 0x0c, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x76, 0x32, 0x2e, 0x6d, 0x6f, 0x64,
	0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4e, 0x65, 0x77, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x7b, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x42, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0a, 0x62, 0x73, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x76,
	0x32, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x53, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x09, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28,
	0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x42, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42,
	0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x22, 0x99, 0x01, 0x0a, 0x02, 0x42, 0x53, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3f, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x76, 0x32, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61,
	0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x53, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x50, 0x0a,
	0x0e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x32, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x62, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x42, 0x53, 0x52, 0x0a, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x73, 0x22,
	0xba, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x70, 0x65, 0x72,
	0x50, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x69, 0x73, 0x44, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x44, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x42, 0x3e, 0x5a, 0x3c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x6f, 0x77, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_main_bootstrap_models_proto_rawDescOnce sync.Once
	file_main_bootstrap_models_proto_rawDescData = file_main_bootstrap_models_proto_rawDesc
)

func file_main_bootstrap_models_proto_rawDescGZIP() []byte {
	file_main_bootstrap_models_proto_rawDescOnce.Do(func() {
		file_main_bootstrap_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_bootstrap_models_proto_rawDescData)
	})
	return file_main_bootstrap_models_proto_rawDescData
}

var file_main_bootstrap_models_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_main_bootstrap_models_proto_goTypes = []interface{}{
	(*BSProject)(nil),                   // 0: v2.main_bootstrap.services.BSProject
	(*BSOrg)(nil),                       // 1: v2.main_bootstrap.services.BSOrg
	(*BSRequest)(nil),                   // 2: v2.main_bootstrap.services.BSRequest
	(*BSAccount)(nil),                   // 3: v2.main_bootstrap.services.BSAccount
	(*NewBSRequest)(nil),                // 4: v2.main_bootstrap.services.NewBSRequest
	(*NewBSResponse)(nil),               // 5: v2.main_bootstrap.services.NewBSResponse
	(*GetBSRequest)(nil),                // 6: v2.main_bootstrap.services.GetBSRequest
	(*BS)(nil),                          // 7: v2.main_bootstrap.services.BS
	(*ListBSResponse)(nil),              // 8: v2.main_bootstrap.services.ListBSResponse
	(*ListBSRequest)(nil),               // 9: v2.main_bootstrap.services.ListBSRequest
	(*v2.ProjectRequest)(nil),           // 10: v2.sys_account.services.ProjectRequest
	(*v21.NewDiscoProjectRequest)(nil),  // 11: v2.mod_disco.services.NewDiscoProjectRequest
	(*v21.NewSurveyProjectRequest)(nil), // 12: v2.mod_disco.services.NewSurveyProjectRequest
	(*v2.OrgRequest)(nil),               // 13: v2.sys_account.services.OrgRequest
	(*v2.LoginRequest)(nil),             // 14: v2.sys_account.services.LoginRequest
	(*v21.NewSurveyUserRequest)(nil),    // 15: v2.mod_disco.services.NewSurveyUserRequest
	(*timestamp.Timestamp)(nil),         // 16: google.protobuf.Timestamp
}
var file_main_bootstrap_models_proto_depIdxs = []int32{
	10, // 0: v2.main_bootstrap.services.BSProject.project:type_name -> v2.sys_account.services.ProjectRequest
	11, // 1: v2.main_bootstrap.services.BSProject.project_details:type_name -> v2.mod_disco.services.NewDiscoProjectRequest
	12, // 2: v2.main_bootstrap.services.BSProject.survey_schema:type_name -> v2.mod_disco.services.NewSurveyProjectRequest
	13, // 3: v2.main_bootstrap.services.BSOrg.org:type_name -> v2.sys_account.services.OrgRequest
	1,  // 4: v2.main_bootstrap.services.BSRequest.orgs:type_name -> v2.main_bootstrap.services.BSOrg
	0,  // 5: v2.main_bootstrap.services.BSRequest.projects:type_name -> v2.main_bootstrap.services.BSProject
	3,  // 6: v2.main_bootstrap.services.BSRequest.superusers:type_name -> v2.main_bootstrap.services.BSAccount
	14, // 7: v2.main_bootstrap.services.BSAccount.initial_superuser:type_name -> v2.sys_account.services.LoginRequest
	15, // 8: v2.main_bootstrap.services.BSAccount.survey_value:type_name -> v2.mod_disco.services.NewSurveyUserRequest
	2,  // 9: v2.main_bootstrap.services.NewBSRequest.bs_request:type_name -> v2.main_bootstrap.services.BSRequest
	16, // 10: v2.main_bootstrap.services.BS.created_at:type_name -> google.protobuf.Timestamp
	2,  // 11: v2.main_bootstrap.services.BS.content:type_name -> v2.main_bootstrap.services.BSRequest
	7,  // 12: v2.main_bootstrap.services.ListBSResponse.bootstraps:type_name -> v2.main_bootstrap.services.BS
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_main_bootstrap_models_proto_init() }
func file_main_bootstrap_models_proto_init() {
	if File_main_bootstrap_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_main_bootstrap_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BSProject); i {
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
		file_main_bootstrap_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BSOrg); i {
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
		file_main_bootstrap_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BSRequest); i {
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
		file_main_bootstrap_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BSAccount); i {
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
		file_main_bootstrap_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewBSRequest); i {
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
		file_main_bootstrap_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewBSResponse); i {
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
		file_main_bootstrap_models_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBSRequest); i {
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
		file_main_bootstrap_models_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BS); i {
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
		file_main_bootstrap_models_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBSResponse); i {
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
		file_main_bootstrap_models_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBSRequest); i {
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
			RawDescriptor: file_main_bootstrap_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_main_bootstrap_models_proto_goTypes,
		DependencyIndexes: file_main_bootstrap_models_proto_depIdxs,
		MessageInfos:      file_main_bootstrap_models_proto_msgTypes,
	}.Build()
	File_main_bootstrap_models_proto = out.File
	file_main_bootstrap_models_proto_rawDesc = nil
	file_main_bootstrap_models_proto_goTypes = nil
	file_main_bootstrap_models_proto_depIdxs = nil
}
