// Code generated by protoc-gen-go. DO NOT EDIT.
// source: image_info.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//for uploading image
type UploadImageRequest struct {
	// Types that are valid to be assigned to Data:
	//	*UploadImageRequest_Info
	//	*UploadImageRequest_Chunkdata
	Data                 isUploadImageRequest_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *UploadImageRequest) Reset()         { *m = UploadImageRequest{} }
func (m *UploadImageRequest) String() string { return proto.CompactTextString(m) }
func (*UploadImageRequest) ProtoMessage()    {}
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8085f4b4731c381e, []int{0}
}

func (m *UploadImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageRequest.Unmarshal(m, b)
}
func (m *UploadImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageRequest.Marshal(b, m, deterministic)
}
func (m *UploadImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageRequest.Merge(m, src)
}
func (m *UploadImageRequest) XXX_Size() int {
	return xxx_messageInfo_UploadImageRequest.Size(m)
}
func (m *UploadImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageRequest proto.InternalMessageInfo

type isUploadImageRequest_Data interface {
	isUploadImageRequest_Data()
}

type UploadImageRequest_Info struct {
	Info *ImageInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadImageRequest_Chunkdata struct {
	Chunkdata []byte `protobuf:"bytes,2,opt,name=chunkdata,proto3,oneof"`
}

func (*UploadImageRequest_Info) isUploadImageRequest_Data() {}

func (*UploadImageRequest_Chunkdata) isUploadImageRequest_Data() {}

func (m *UploadImageRequest) GetData() isUploadImageRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UploadImageRequest) GetInfo() *ImageInfo {
	if x, ok := m.GetData().(*UploadImageRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (m *UploadImageRequest) GetChunkdata() []byte {
	if x, ok := m.GetData().(*UploadImageRequest_Chunkdata); ok {
		return x.Chunkdata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UploadImageRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UploadImageRequest_Info)(nil),
		(*UploadImageRequest_Chunkdata)(nil),
	}
}

type ImageInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Created              string   `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Modified             string   `protobuf:"bytes,3,opt,name=modified,proto3" json:"modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageInfo) Reset()         { *m = ImageInfo{} }
func (m *ImageInfo) String() string { return proto.CompactTextString(m) }
func (*ImageInfo) ProtoMessage()    {}
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8085f4b4731c381e, []int{1}
}

func (m *ImageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageInfo.Unmarshal(m, b)
}
func (m *ImageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageInfo.Marshal(b, m, deterministic)
}
func (m *ImageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageInfo.Merge(m, src)
}
func (m *ImageInfo) XXX_Size() int {
	return xxx_messageInfo_ImageInfo.Size(m)
}
func (m *ImageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ImageInfo proto.InternalMessageInfo

func (m *ImageInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ImageInfo) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *ImageInfo) GetModified() string {
	if m != nil {
		return m.Modified
	}
	return ""
}

type UploadImageResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size                 uint32   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadImageResponse) Reset()         { *m = UploadImageResponse{} }
func (m *UploadImageResponse) String() string { return proto.CompactTextString(m) }
func (*UploadImageResponse) ProtoMessage()    {}
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8085f4b4731c381e, []int{2}
}

func (m *UploadImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageResponse.Unmarshal(m, b)
}
func (m *UploadImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageResponse.Marshal(b, m, deterministic)
}
func (m *UploadImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageResponse.Merge(m, src)
}
func (m *UploadImageResponse) XXX_Size() int {
	return xxx_messageInfo_UploadImageResponse.Size(m)
}
func (m *UploadImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageResponse proto.InternalMessageInfo

func (m *UploadImageResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UploadImageResponse) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ImageList struct {
	Images               []*ImageInfo `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ImageList) Reset()         { *m = ImageList{} }
func (m *ImageList) String() string { return proto.CompactTextString(m) }
func (*ImageList) ProtoMessage()    {}
func (*ImageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8085f4b4731c381e, []int{3}
}

func (m *ImageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageList.Unmarshal(m, b)
}
func (m *ImageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageList.Marshal(b, m, deterministic)
}
func (m *ImageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageList.Merge(m, src)
}
func (m *ImageList) XXX_Size() int {
	return xxx_messageInfo_ImageList.Size(m)
}
func (m *ImageList) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageList.DiscardUnknown(m)
}

var xxx_messageInfo_ImageList proto.InternalMessageInfo

func (m *ImageList) GetImages() []*ImageInfo {
	if m != nil {
		return m.Images
	}
	return nil
}

//for downloading image
type DownloadImageRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadImageRequest) Reset()         { *m = DownloadImageRequest{} }
func (m *DownloadImageRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadImageRequest) ProtoMessage()    {}
func (*DownloadImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8085f4b4731c381e, []int{4}
}

func (m *DownloadImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadImageRequest.Unmarshal(m, b)
}
func (m *DownloadImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadImageRequest.Marshal(b, m, deterministic)
}
func (m *DownloadImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadImageRequest.Merge(m, src)
}
func (m *DownloadImageRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadImageRequest.Size(m)
}
func (m *DownloadImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadImageRequest proto.InternalMessageInfo

func (m *DownloadImageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DownloadImageResponse struct {
	// Types that are valid to be assigned to Data:
	//	*DownloadImageResponse_Info
	//	*DownloadImageResponse_Chunkdata
	Data                 isDownloadImageResponse_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *DownloadImageResponse) Reset()         { *m = DownloadImageResponse{} }
func (m *DownloadImageResponse) String() string { return proto.CompactTextString(m) }
func (*DownloadImageResponse) ProtoMessage()    {}
func (*DownloadImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8085f4b4731c381e, []int{5}
}

func (m *DownloadImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadImageResponse.Unmarshal(m, b)
}
func (m *DownloadImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadImageResponse.Marshal(b, m, deterministic)
}
func (m *DownloadImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadImageResponse.Merge(m, src)
}
func (m *DownloadImageResponse) XXX_Size() int {
	return xxx_messageInfo_DownloadImageResponse.Size(m)
}
func (m *DownloadImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadImageResponse proto.InternalMessageInfo

type isDownloadImageResponse_Data interface {
	isDownloadImageResponse_Data()
}

type DownloadImageResponse_Info struct {
	Info *ImageInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type DownloadImageResponse_Chunkdata struct {
	Chunkdata []byte `protobuf:"bytes,2,opt,name=chunkdata,proto3,oneof"`
}

func (*DownloadImageResponse_Info) isDownloadImageResponse_Data() {}

func (*DownloadImageResponse_Chunkdata) isDownloadImageResponse_Data() {}

func (m *DownloadImageResponse) GetData() isDownloadImageResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DownloadImageResponse) GetInfo() *ImageInfo {
	if x, ok := m.GetData().(*DownloadImageResponse_Info); ok {
		return x.Info
	}
	return nil
}

func (m *DownloadImageResponse) GetChunkdata() []byte {
	if x, ok := m.GetData().(*DownloadImageResponse_Chunkdata); ok {
		return x.Chunkdata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DownloadImageResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DownloadImageResponse_Info)(nil),
		(*DownloadImageResponse_Chunkdata)(nil),
	}
}

func init() {
	proto.RegisterType((*UploadImageRequest)(nil), "proto.UploadImageRequest")
	proto.RegisterType((*ImageInfo)(nil), "proto.ImageInfo")
	proto.RegisterType((*UploadImageResponse)(nil), "proto.UploadImageResponse")
	proto.RegisterType((*ImageList)(nil), "proto.ImageList")
	proto.RegisterType((*DownloadImageRequest)(nil), "proto.DownloadImageRequest")
	proto.RegisterType((*DownloadImageResponse)(nil), "proto.DownloadImageResponse")
}

func init() { proto.RegisterFile("image_info.proto", fileDescriptor_8085f4b4731c381e) }

var fileDescriptor_8085f4b4731c381e = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x4d, 0x6c, 0xac, 0x66, 0x6a, 0xa1, 0x8c, 0x0a, 0x31, 0x94, 0x52, 0xf6, 0x20, 0xc1, 0x43,
	0x2a, 0x15, 0x8f, 0x7a, 0x10, 0x0f, 0x2d, 0x08, 0x42, 0x4a, 0xbd, 0xca, 0xb6, 0xd9, 0xc4, 0xc5,
	0x76, 0x37, 0x66, 0x13, 0x0b, 0xfe, 0xb4, 0xbf, 0x20, 0xd9, 0xa4, 0xd5, 0xda, 0x80, 0x17, 0x4f,
	0xbb, 0x93, 0xf7, 0xe6, 0xcd, 0xdb, 0x97, 0x81, 0x0e, 0x5f, 0xd2, 0x98, 0x3d, 0x73, 0x11, 0x49,
	0x3f, 0x49, 0x65, 0x26, 0x71, 0x5f, 0x1f, 0x6e, 0x2f, 0x96, 0x32, 0x5e, 0xb0, 0x81, 0xae, 0x66,
	0x79, 0x34, 0x58, 0xa5, 0x34, 0x49, 0x58, 0xaa, 0x4a, 0x1a, 0x09, 0x01, 0xa7, 0xc9, 0x42, 0xd2,
	0x70, 0x5c, 0x08, 0x04, 0xec, 0x2d, 0x67, 0x2a, 0xc3, 0x73, 0xb0, 0x0a, 0x29, 0xc7, 0xec, 0x9b,
	0x5e, 0x6b, 0xd8, 0x29, 0xb9, 0xbe, 0xa6, 0x8c, 0x45, 0x24, 0x47, 0x46, 0xa0, 0x71, 0xec, 0x81,
	0x3d, 0x7f, 0xc9, 0xc5, 0x6b, 0x48, 0x33, 0xea, 0xec, 0xf5, 0x4d, 0xef, 0x68, 0x64, 0x04, 0xdf,
	0x9f, 0xee, 0x9a, 0x60, 0x15, 0x27, 0x99, 0x82, 0xbd, 0x69, 0x46, 0x04, 0x4b, 0xd0, 0x25, 0xd3,
	0xe2, 0x76, 0xa0, 0xef, 0xe8, 0xc0, 0xc1, 0x3c, 0x65, 0x34, 0x63, 0xa1, 0x96, 0xb1, 0x83, 0x75,
	0x89, 0x2e, 0x1c, 0x2e, 0x65, 0xc8, 0x23, 0xce, 0x42, 0xa7, 0xa1, 0xa1, 0x4d, 0x4d, 0x6e, 0xe0,
	0x78, 0xcb, 0xbc, 0x4a, 0xa4, 0x50, 0xac, 0x76, 0x00, 0x82, 0xa5, 0xf8, 0x07, 0xd3, 0xea, 0xed,
	0x40, 0xdf, 0xc9, 0x75, 0xe5, 0xea, 0x81, 0xab, 0x0c, 0x3d, 0x68, 0xea, 0x0c, 0x95, 0x63, 0xf6,
	0x1b, 0x75, 0x8f, 0x0e, 0x2a, 0x9c, 0x5c, 0xc0, 0xc9, 0xbd, 0x5c, 0x89, 0x9d, 0xd0, 0x6a, 0xc6,
	0x92, 0x18, 0x4e, 0x7f, 0x71, 0x2b, 0x8f, 0xff, 0x9c, 0xf0, 0xf0, 0xd3, 0x04, 0xd4, 0xdd, 0x65,
	0x20, 0x13, 0x96, 0xbe, 0xf3, 0x39, 0xc3, 0x11, 0xb4, 0x7e, 0x24, 0x84, 0x67, 0xd5, 0x9c, 0xdd,
	0x5f, 0xee, 0xba, 0x75, 0x50, 0x69, 0x96, 0x18, 0x9e, 0x89, 0xb7, 0x00, 0x45, 0x4e, 0x1a, 0x50,
	0xd8, 0xf5, 0xcb, 0xbd, 0xf2, 0xd7, 0x7b, 0xe5, 0x4f, 0xb2, 0x94, 0x8b, 0xf8, 0x89, 0x2e, 0x72,
	0xe6, 0x6e, 0x3d, 0xa7, 0xe8, 0x22, 0x06, 0x3e, 0x42, 0x7b, 0x2b, 0x89, 0x3f, 0x24, 0xba, 0x95,
	0x44, 0x6d, 0x7a, 0xc4, 0xb8, 0x34, 0x67, 0x4d, 0x4d, 0xb8, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x9d, 0x78, 0x67, 0xb5, 0xfb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ImageUploadServiceClient is the client API for ImageUploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImageUploadServiceClient interface {
	UploadImage(ctx context.Context, opts ...grpc.CallOption) (ImageUploadService_UploadImageClient, error)
	ListImages(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*ImageList, error)
	DownloadImage(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (ImageUploadService_DownloadImageClient, error)
}

type imageUploadServiceClient struct {
	cc *grpc.ClientConn
}

func NewImageUploadServiceClient(cc *grpc.ClientConn) ImageUploadServiceClient {
	return &imageUploadServiceClient{cc}
}

func (c *imageUploadServiceClient) UploadImage(ctx context.Context, opts ...grpc.CallOption) (ImageUploadService_UploadImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ImageUploadService_serviceDesc.Streams[0], "/proto.ImageUploadService/UploadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageUploadServiceUploadImageClient{stream}
	return x, nil
}

type ImageUploadService_UploadImageClient interface {
	Send(*UploadImageRequest) error
	CloseAndRecv() (*UploadImageResponse, error)
	grpc.ClientStream
}

type imageUploadServiceUploadImageClient struct {
	grpc.ClientStream
}

func (x *imageUploadServiceUploadImageClient) Send(m *UploadImageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imageUploadServiceUploadImageClient) CloseAndRecv() (*UploadImageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadImageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imageUploadServiceClient) ListImages(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*ImageList, error) {
	out := new(ImageList)
	err := c.cc.Invoke(ctx, "/proto.ImageUploadService/ListImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageUploadServiceClient) DownloadImage(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (ImageUploadService_DownloadImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ImageUploadService_serviceDesc.Streams[1], "/proto.ImageUploadService/DownloadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageUploadServiceDownloadImageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ImageUploadService_DownloadImageClient interface {
	Recv() (*DownloadImageResponse, error)
	grpc.ClientStream
}

type imageUploadServiceDownloadImageClient struct {
	grpc.ClientStream
}

func (x *imageUploadServiceDownloadImageClient) Recv() (*DownloadImageResponse, error) {
	m := new(DownloadImageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageUploadServiceServer is the server API for ImageUploadService service.
type ImageUploadServiceServer interface {
	UploadImage(ImageUploadService_UploadImageServer) error
	ListImages(context.Context, *wrappers.StringValue) (*ImageList, error)
	DownloadImage(*wrappers.StringValue, ImageUploadService_DownloadImageServer) error
}

// UnimplementedImageUploadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedImageUploadServiceServer struct {
}

func (*UnimplementedImageUploadServiceServer) UploadImage(srv ImageUploadService_UploadImageServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (*UnimplementedImageUploadServiceServer) ListImages(ctx context.Context, req *wrappers.StringValue) (*ImageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImages not implemented")
}
func (*UnimplementedImageUploadServiceServer) DownloadImage(req *wrappers.StringValue, srv ImageUploadService_DownloadImageServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadImage not implemented")
}

func RegisterImageUploadServiceServer(s *grpc.Server, srv ImageUploadServiceServer) {
	s.RegisterService(&_ImageUploadService_serviceDesc, srv)
}

func _ImageUploadService_UploadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImageUploadServiceServer).UploadImage(&imageUploadServiceUploadImageServer{stream})
}

type ImageUploadService_UploadImageServer interface {
	SendAndClose(*UploadImageResponse) error
	Recv() (*UploadImageRequest, error)
	grpc.ServerStream
}

type imageUploadServiceUploadImageServer struct {
	grpc.ServerStream
}

func (x *imageUploadServiceUploadImageServer) SendAndClose(m *UploadImageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imageUploadServiceUploadImageServer) Recv() (*UploadImageRequest, error) {
	m := new(UploadImageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ImageUploadService_ListImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageUploadServiceServer).ListImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ImageUploadService/ListImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageUploadServiceServer).ListImages(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageUploadService_DownloadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrappers.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImageUploadServiceServer).DownloadImage(m, &imageUploadServiceDownloadImageServer{stream})
}

type ImageUploadService_DownloadImageServer interface {
	Send(*DownloadImageResponse) error
	grpc.ServerStream
}

type imageUploadServiceDownloadImageServer struct {
	grpc.ServerStream
}

func (x *imageUploadServiceDownloadImageServer) Send(m *DownloadImageResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ImageUploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ImageUploadService",
	HandlerType: (*ImageUploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListImages",
			Handler:    _ImageUploadService_ListImages_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadImage",
			Handler:       _ImageUploadService_UploadImage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadImage",
			Handler:       _ImageUploadService_DownloadImage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "image_info.proto",
}
