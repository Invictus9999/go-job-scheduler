// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: scheduler/v1/job.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SimpleJobPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SimpleJobPayload) Reset() {
	*x = SimpleJobPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_v1_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleJobPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleJobPayload) ProtoMessage() {}

func (x *SimpleJobPayload) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_v1_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleJobPayload.ProtoReflect.Descriptor instead.
func (*SimpleJobPayload) Descriptor() ([]byte, []int) {
	return file_scheduler_v1_job_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleJobPayload) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type JobDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EmailId string                 `protobuf:"bytes,2,opt,name=email_id,json=emailId,proto3" json:"email_id,omitempty"`
	NextRun *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=next_run,json=nextRun,proto3" json:"next_run,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*JobDetails_SimpleJobPayload
	Payload isJobDetails_Payload `protobuf_oneof:"payload"`
}

func (x *JobDetails) Reset() {
	*x = JobDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_v1_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobDetails) ProtoMessage() {}

func (x *JobDetails) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_v1_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobDetails.ProtoReflect.Descriptor instead.
func (*JobDetails) Descriptor() ([]byte, []int) {
	return file_scheduler_v1_job_proto_rawDescGZIP(), []int{1}
}

func (x *JobDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobDetails) GetEmailId() string {
	if x != nil {
		return x.EmailId
	}
	return ""
}

func (x *JobDetails) GetNextRun() *timestamppb.Timestamp {
	if x != nil {
		return x.NextRun
	}
	return nil
}

func (m *JobDetails) GetPayload() isJobDetails_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *JobDetails) GetSimpleJobPayload() *SimpleJobPayload {
	if x, ok := x.GetPayload().(*JobDetails_SimpleJobPayload); ok {
		return x.SimpleJobPayload
	}
	return nil
}

type isJobDetails_Payload interface {
	isJobDetails_Payload()
}

type JobDetails_SimpleJobPayload struct {
	SimpleJobPayload *SimpleJobPayload `protobuf:"bytes,21,opt,name=simple_job_payload,json=simpleJobPayload,proto3,oneof"`
}

func (*JobDetails_SimpleJobPayload) isJobDetails_Payload() {}

var File_scheduler_v1_job_proto protoreflect.FileDescriptor

var file_scheduler_v1_job_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6a,
	0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc9, 0x01, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12,
	0x35, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x6e,
	0x65, 0x78, 0x74, 0x52, 0x75, 0x6e, 0x12, 0x4e, 0x0a, 0x12, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x48, 0x00, 0x52, 0x10, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4a, 0x6f, 0x62, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x49, 0x6e, 0x76, 0x69, 0x63, 0x74, 0x75, 0x73, 0x39, 0x39, 0x39, 0x39, 0x2f, 0x67, 0x6f, 0x2d,
	0x6a, 0x6f, 0x62, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scheduler_v1_job_proto_rawDescOnce sync.Once
	file_scheduler_v1_job_proto_rawDescData = file_scheduler_v1_job_proto_rawDesc
)

func file_scheduler_v1_job_proto_rawDescGZIP() []byte {
	file_scheduler_v1_job_proto_rawDescOnce.Do(func() {
		file_scheduler_v1_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_scheduler_v1_job_proto_rawDescData)
	})
	return file_scheduler_v1_job_proto_rawDescData
}

var file_scheduler_v1_job_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_scheduler_v1_job_proto_goTypes = []interface{}{
	(*SimpleJobPayload)(nil),      // 0: scheduler.v1.SimpleJobPayload
	(*JobDetails)(nil),            // 1: scheduler.v1.JobDetails
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_scheduler_v1_job_proto_depIdxs = []int32{
	2, // 0: scheduler.v1.JobDetails.next_run:type_name -> google.protobuf.Timestamp
	0, // 1: scheduler.v1.JobDetails.simple_job_payload:type_name -> scheduler.v1.SimpleJobPayload
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_scheduler_v1_job_proto_init() }
func file_scheduler_v1_job_proto_init() {
	if File_scheduler_v1_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scheduler_v1_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleJobPayload); i {
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
		file_scheduler_v1_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobDetails); i {
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
	file_scheduler_v1_job_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*JobDetails_SimpleJobPayload)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scheduler_v1_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scheduler_v1_job_proto_goTypes,
		DependencyIndexes: file_scheduler_v1_job_proto_depIdxs,
		MessageInfos:      file_scheduler_v1_job_proto_msgTypes,
	}.Build()
	File_scheduler_v1_job_proto = out.File
	file_scheduler_v1_job_proto_rawDesc = nil
	file_scheduler_v1_job_proto_goTypes = nil
	file_scheduler_v1_job_proto_depIdxs = nil
}
