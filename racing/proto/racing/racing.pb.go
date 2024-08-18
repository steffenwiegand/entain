// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.21.12
// source: racing/racing.proto

package racing

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

type ListRacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter  *ListRacesRequestFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	OrderBy string                  `protobuf:"bytes,2,opt,name=orderBy,proto3" json:"orderBy,omitempty"`
}

func (x *ListRacesRequest) Reset() {
	*x = ListRacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRacesRequest) ProtoMessage() {}

func (x *ListRacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRacesRequest.ProtoReflect.Descriptor instead.
func (*ListRacesRequest) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{0}
}

func (x *ListRacesRequest) GetFilter() *ListRacesRequestFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ListRacesRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

// Response to ListRaces call.
type ListRacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Races []*Race `protobuf:"bytes,1,rep,name=races,proto3" json:"races,omitempty"`
}

func (x *ListRacesResponse) Reset() {
	*x = ListRacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRacesResponse) ProtoMessage() {}

func (x *ListRacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRacesResponse.ProtoReflect.Descriptor instead.
func (*ListRacesResponse) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{1}
}

func (x *ListRacesResponse) GetRaces() []*Race {
	if x != nil {
		return x.Races
	}
	return nil
}

// Filter for listing races.
type ListRacesRequestFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MeetingIds      []int64 `protobuf:"varint,1,rep,packed,name=meeting_ids,json=meetingIds,proto3" json:"meeting_ids,omitempty"`
	ShowVisibleOnly bool    `protobuf:"varint,2,opt,name=show_visible_only,json=showVisibleOnly,proto3" json:"show_visible_only,omitempty"`
}

func (x *ListRacesRequestFilter) Reset() {
	*x = ListRacesRequestFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRacesRequestFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRacesRequestFilter) ProtoMessage() {}

func (x *ListRacesRequestFilter) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRacesRequestFilter.ProtoReflect.Descriptor instead.
func (*ListRacesRequestFilter) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{2}
}

func (x *ListRacesRequestFilter) GetMeetingIds() []int64 {
	if x != nil {
		return x.MeetingIds
	}
	return nil
}

func (x *ListRacesRequestFilter) GetShowVisibleOnly() bool {
	if x != nil {
		return x.ShowVisibleOnly
	}
	return false
}

// Request for GetRace call.
type GetRaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRaceRequest) Reset() {
	*x = GetRaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRaceRequest) ProtoMessage() {}

func (x *GetRaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRaceRequest.ProtoReflect.Descriptor instead.
func (*GetRaceRequest) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{3}
}

func (x *GetRaceRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Response to GetRace call.
type GetRaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Race *Race `protobuf:"bytes,1,opt,name=race,proto3" json:"race,omitempty"`
}

func (x *GetRaceResponse) Reset() {
	*x = GetRaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRaceResponse) ProtoMessage() {}

func (x *GetRaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRaceResponse.ProtoReflect.Descriptor instead.
func (*GetRaceResponse) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{4}
}

func (x *GetRaceResponse) GetRace() *Race {
	if x != nil {
		return x.Race
	}
	return nil
}

// A race resource.
type Race struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID represents a unique identifier for the race.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// MeetingID represents a unique identifier for the races meeting.
	MeetingId int64 `protobuf:"varint,2,opt,name=meeting_id,json=meetingId,proto3" json:"meeting_id,omitempty"`
	// Name is the official name given to the race.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Number represents the number of the race.
	Number int64 `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	// Visible represents whether or not the race is visible.
	Visible bool `protobuf:"varint,5,opt,name=visible,proto3" json:"visible,omitempty"`
	// AdvertisedStartTime is the time the race is advertised to run.
	AdvertisedStartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=advertised_start_time,json=advertisedStartTime,proto3" json:"advertised_start_time,omitempty"`
	// Status reflects the current status of a race. Values are`OPEN` or `CLOSED`.
	Status string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Race) Reset() {
	*x = Race{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Race) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Race) ProtoMessage() {}

func (x *Race) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Race.ProtoReflect.Descriptor instead.
func (*Race) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{5}
}

func (x *Race) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Race) GetMeetingId() int64 {
	if x != nil {
		return x.MeetingId
	}
	return 0
}

func (x *Race) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Race) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Race) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Race) GetAdvertisedStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AdvertisedStartTime
	}
	return nil
}

func (x *Race) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_racing_racing_proto protoreflect.FileDescriptor

var file_racing_racing_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x22, 0x37, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x72, 0x61, 0x63,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x2e, 0x52, 0x61, 0x63, 0x65, 0x52, 0x05, 0x72, 0x61, 0x63, 0x65, 0x73, 0x22, 0x65, 0x0a,
	0x16, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x77,
	0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x68, 0x6f, 0x77, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x72, 0x61, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67,
	0x2e, 0x52, 0x61, 0x63, 0x65, 0x52, 0x04, 0x72, 0x61, 0x63, 0x65, 0x22, 0xe3, 0x01, 0x0a, 0x04,
	0x52, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x4e, 0x0a, 0x15, 0x61, 0x64, 0x76,
	0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0x8a, 0x01, 0x0a, 0x06, 0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x42, 0x0a, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x61, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09,
	0x5a, 0x07, 0x2f, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_racing_racing_proto_rawDescOnce sync.Once
	file_racing_racing_proto_rawDescData = file_racing_racing_proto_rawDesc
)

func file_racing_racing_proto_rawDescGZIP() []byte {
	file_racing_racing_proto_rawDescOnce.Do(func() {
		file_racing_racing_proto_rawDescData = protoimpl.X.CompressGZIP(file_racing_racing_proto_rawDescData)
	})
	return file_racing_racing_proto_rawDescData
}

var file_racing_racing_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_racing_racing_proto_goTypes = []interface{}{
	(*ListRacesRequest)(nil),       // 0: racing.ListRacesRequest
	(*ListRacesResponse)(nil),      // 1: racing.ListRacesResponse
	(*ListRacesRequestFilter)(nil), // 2: racing.ListRacesRequestFilter
	(*GetRaceRequest)(nil),         // 3: racing.GetRaceRequest
	(*GetRaceResponse)(nil),        // 4: racing.GetRaceResponse
	(*Race)(nil),                   // 5: racing.Race
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_racing_racing_proto_depIdxs = []int32{
	2, // 0: racing.ListRacesRequest.filter:type_name -> racing.ListRacesRequestFilter
	5, // 1: racing.ListRacesResponse.races:type_name -> racing.Race
	5, // 2: racing.GetRaceResponse.race:type_name -> racing.Race
	6, // 3: racing.Race.advertised_start_time:type_name -> google.protobuf.Timestamp
	0, // 4: racing.Racing.ListRaces:input_type -> racing.ListRacesRequest
	3, // 5: racing.Racing.GetRace:input_type -> racing.GetRaceRequest
	1, // 6: racing.Racing.ListRaces:output_type -> racing.ListRacesResponse
	4, // 7: racing.Racing.GetRace:output_type -> racing.GetRaceResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_racing_racing_proto_init() }
func file_racing_racing_proto_init() {
	if File_racing_racing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_racing_racing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRacesRequest); i {
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
		file_racing_racing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRacesResponse); i {
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
		file_racing_racing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRacesRequestFilter); i {
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
		file_racing_racing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRaceRequest); i {
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
		file_racing_racing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRaceResponse); i {
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
		file_racing_racing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Race); i {
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
			RawDescriptor: file_racing_racing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_racing_racing_proto_goTypes,
		DependencyIndexes: file_racing_racing_proto_depIdxs,
		MessageInfos:      file_racing_racing_proto_msgTypes,
	}.Build()
	File_racing_racing_proto = out.File
	file_racing_racing_proto_rawDesc = nil
	file_racing_racing_proto_goTypes = nil
	file_racing_racing_proto_depIdxs = nil
}
