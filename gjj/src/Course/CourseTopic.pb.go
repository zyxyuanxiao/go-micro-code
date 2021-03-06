// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.1
// source: CourseTopic.proto

package Course

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

type CourseTopic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32      `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	CourseId   int32      `protobuf:"varint,2,opt,name=CourseId,proto3" json:"CourseId,omitempty"`
	CourseDid  int32      `protobuf:"varint,3,opt,name=CourseDid,proto3" json:"CourseDid,omitempty"`
	Likes      int32      `protobuf:"varint,4,opt,name=Likes,proto3" json:"Likes,omitempty"`
	Unlikes    int32      `protobuf:"varint,5,opt,name=Unlikes,proto3" json:"Unlikes,omitempty"`
	Title      string     `protobuf:"bytes,6,opt,name=Title,proto3" json:"Title,omitempty"`
	Content    string     `protobuf:"bytes,7,opt,name=Content,proto3" json:"Content,omitempty"`
	UserId     int32      `protobuf:"varint,8,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Addtime    *Timestamp `protobuf:"bytes,9,opt,name=Addtime,proto3" json:"Addtime,omitempty"`
	Updatetime *Timestamp `protobuf:"bytes,10,opt,name=Updatetime,proto3" json:"Updatetime,omitempty"`
}

func (x *CourseTopic) Reset() {
	*x = CourseTopic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CourseTopic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseTopic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseTopic) ProtoMessage() {}

func (x *CourseTopic) ProtoReflect() protoreflect.Message {
	mi := &file_CourseTopic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseTopic.ProtoReflect.Descriptor instead.
func (*CourseTopic) Descriptor() ([]byte, []int) {
	return file_CourseTopic_proto_rawDescGZIP(), []int{0}
}

func (x *CourseTopic) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CourseTopic) GetCourseId() int32 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *CourseTopic) GetCourseDid() int32 {
	if x != nil {
		return x.CourseDid
	}
	return 0
}

func (x *CourseTopic) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *CourseTopic) GetUnlikes() int32 {
	if x != nil {
		return x.Unlikes
	}
	return 0
}

func (x *CourseTopic) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CourseTopic) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CourseTopic) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CourseTopic) GetAddtime() *Timestamp {
	if x != nil {
		return x.Addtime
	}
	return nil
}

func (x *CourseTopic) GetUpdatetime() *Timestamp {
	if x != nil {
		return x.Updatetime
	}
	return nil
}

type CourseTopicReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32      `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	TopicId int32      `protobuf:"varint,2,opt,name=TopicId,proto3" json:"TopicId,omitempty"`
	Content string     `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	UserId  int32      `protobuf:"varint,4,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Likes   int32      `protobuf:"varint,5,opt,name=Likes,proto3" json:"Likes,omitempty"`
	Unlikes int32      `protobuf:"varint,6,opt,name=Unlikes,proto3" json:"Unlikes,omitempty"`
	Addtime *Timestamp `protobuf:"bytes,7,opt,name=Addtime,proto3" json:"Addtime,omitempty"`
}

func (x *CourseTopicReply) Reset() {
	*x = CourseTopicReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CourseTopic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseTopicReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseTopicReply) ProtoMessage() {}

func (x *CourseTopicReply) ProtoReflect() protoreflect.Message {
	mi := &file_CourseTopic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseTopicReply.ProtoReflect.Descriptor instead.
func (*CourseTopicReply) Descriptor() ([]byte, []int) {
	return file_CourseTopic_proto_rawDescGZIP(), []int{1}
}

func (x *CourseTopicReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CourseTopicReply) GetTopicId() int32 {
	if x != nil {
		return x.TopicId
	}
	return 0
}

func (x *CourseTopicReply) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CourseTopicReply) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CourseTopicReply) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *CourseTopicReply) GetUnlikes() int32 {
	if x != nil {
		return x.Unlikes
	}
	return 0
}

func (x *CourseTopicReply) GetAddtime() *Timestamp {
	if x != nil {
		return x.Addtime
	}
	return nil
}

type TopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId int32 `protobuf:"varint,1,opt,name=CourseId,proto3" json:"CourseId,omitempty"`
}

func (x *TopicRequest) Reset() {
	*x = TopicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CourseTopic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicRequest) ProtoMessage() {}

func (x *TopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CourseTopic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicRequest.ProtoReflect.Descriptor instead.
func (*TopicRequest) Descriptor() ([]byte, []int) {
	return file_CourseTopic_proto_rawDescGZIP(), []int{2}
}

func (x *TopicRequest) GetCourseId() int32 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

type TopicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*CourseTopic `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *TopicResponse) Reset() {
	*x = TopicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CourseTopic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicResponse) ProtoMessage() {}

func (x *TopicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_CourseTopic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicResponse.ProtoReflect.Descriptor instead.
func (*TopicResponse) Descriptor() ([]byte, []int) {
	return file_CourseTopic_proto_rawDescGZIP(), []int{3}
}

func (x *TopicResponse) GetResult() []*CourseTopic {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_CourseTopic_proto protoreflect.FileDescriptor

var file_CourseTopic_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x0c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x02, 0x0a, 0x0b, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x44, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x6e, 0x6c,
	0x69, 0x6b, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x55, 0x6e, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x41, 0x64, 0x64, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x10,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6b,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x55, 0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x41, 0x64, 0x64, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x0c, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0d, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x32, 0x4d, 0x0a, 0x12, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x14, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CourseTopic_proto_rawDescOnce sync.Once
	file_CourseTopic_proto_rawDescData = file_CourseTopic_proto_rawDesc
)

func file_CourseTopic_proto_rawDescGZIP() []byte {
	file_CourseTopic_proto_rawDescOnce.Do(func() {
		file_CourseTopic_proto_rawDescData = protoimpl.X.CompressGZIP(file_CourseTopic_proto_rawDescData)
	})
	return file_CourseTopic_proto_rawDescData
}

var file_CourseTopic_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_CourseTopic_proto_goTypes = []interface{}{
	(*CourseTopic)(nil),      // 0: Course.CourseTopic
	(*CourseTopicReply)(nil), // 1: Course.CourseTopicReply
	(*TopicRequest)(nil),     // 2: Course.TopicRequest
	(*TopicResponse)(nil),    // 3: Course.TopicResponse
	(*Timestamp)(nil),        // 4: Course.Timestamp
}
var file_CourseTopic_proto_depIdxs = []int32{
	4, // 0: Course.CourseTopic.Addtime:type_name -> Course.Timestamp
	4, // 1: Course.CourseTopic.Updatetime:type_name -> Course.Timestamp
	4, // 2: Course.CourseTopicReply.Addtime:type_name -> Course.Timestamp
	0, // 3: Course.TopicResponse.result:type_name -> Course.CourseTopic
	2, // 4: Course.CourseTopicService.GetTopic:input_type -> Course.TopicRequest
	3, // 5: Course.CourseTopicService.GetTopic:output_type -> Course.TopicResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_CourseTopic_proto_init() }
func file_CourseTopic_proto_init() {
	if File_CourseTopic_proto != nil {
		return
	}
	file_Common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CourseTopic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseTopic); i {
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
		file_CourseTopic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseTopicReply); i {
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
		file_CourseTopic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicRequest); i {
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
		file_CourseTopic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicResponse); i {
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
			RawDescriptor: file_CourseTopic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_CourseTopic_proto_goTypes,
		DependencyIndexes: file_CourseTopic_proto_depIdxs,
		MessageInfos:      file_CourseTopic_proto_msgTypes,
	}.Build()
	File_CourseTopic_proto = out.File
	file_CourseTopic_proto_rawDesc = nil
	file_CourseTopic_proto_goTypes = nil
	file_CourseTopic_proto_depIdxs = nil
}
