// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: camion.proto

package camion

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PaqueteRegistro struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPaquete   string `protobuf:"bytes,1,opt,name=idPaquete,proto3" json:"idPaquete,omitempty"`
	Seguimiento int64  `protobuf:"varint,2,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
	Tipo        string `protobuf:"bytes,3,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Valor       int64  `protobuf:"varint,4,opt,name=valor,proto3" json:"valor,omitempty"`
	Intentos    int64  `protobuf:"varint,5,opt,name=intentos,proto3" json:"intentos,omitempty"`
	Estado      string `protobuf:"bytes,6,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *PaqueteRegistro) Reset() {
	*x = PaqueteRegistro{}
	if protoimpl.UnsafeEnabled {
		mi := &file_camion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaqueteRegistro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaqueteRegistro) ProtoMessage() {}

func (x *PaqueteRegistro) ProtoReflect() protoreflect.Message {
	mi := &file_camion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaqueteRegistro.ProtoReflect.Descriptor instead.
func (*PaqueteRegistro) Descriptor() ([]byte, []int) {
	return file_camion_proto_rawDescGZIP(), []int{0}
}

func (x *PaqueteRegistro) GetIdPaquete() string {
	if x != nil {
		return x.IdPaquete
	}
	return ""
}

func (x *PaqueteRegistro) GetSeguimiento() int64 {
	if x != nil {
		return x.Seguimiento
	}
	return 0
}

func (x *PaqueteRegistro) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *PaqueteRegistro) GetValor() int64 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *PaqueteRegistro) GetIntentos() int64 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

func (x *PaqueteRegistro) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type InformeCamion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPaquete int64  `protobuf:"varint,1,opt,name=idPaquete,proto3" json:"idPaquete,omitempty"`
	Estado    string `protobuf:"bytes,2,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *InformeCamion) Reset() {
	*x = InformeCamion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_camion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InformeCamion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InformeCamion) ProtoMessage() {}

func (x *InformeCamion) ProtoReflect() protoreflect.Message {
	mi := &file_camion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InformeCamion.ProtoReflect.Descriptor instead.
func (*InformeCamion) Descriptor() ([]byte, []int) {
	return file_camion_proto_rawDescGZIP(), []int{1}
}

func (x *InformeCamion) GetIdPaquete() int64 {
	if x != nil {
		return x.IdPaquete
	}
	return 0
}

func (x *InformeCamion) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

var File_camion_proto protoreflect.FileDescriptor

var file_camion_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x22, 0xaf, 0x01, 0x0a, 0x0f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64,
	0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x64, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x67, 0x75,
	0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73,
	0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x45, 0x0a, 0x0d, 0x69, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x65, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x50,
	0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x64,
	0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x32,
	0x53, 0x0a, 0x0d, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x42, 0x0a, 0x0c, 0x6e, 0x75, 0x65, 0x76, 0x6f, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65,
	0x12, 0x17, 0x2e, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x6d, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x6f, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_camion_proto_rawDescOnce sync.Once
	file_camion_proto_rawDescData = file_camion_proto_rawDesc
)

func file_camion_proto_rawDescGZIP() []byte {
	file_camion_proto_rawDescOnce.Do(func() {
		file_camion_proto_rawDescData = protoimpl.X.CompressGZIP(file_camion_proto_rawDescData)
	})
	return file_camion_proto_rawDescData
}

var file_camion_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_camion_proto_goTypes = []interface{}{
	(*PaqueteRegistro)(nil), // 0: camion.paqueteRegistro
	(*InformeCamion)(nil),   // 1: camion.informeCamion
}
var file_camion_proto_depIdxs = []int32{
	0, // 0: camion.CamionService.nuevoPaquete:input_type -> camion.paqueteRegistro
	0, // 1: camion.CamionService.nuevoPaquete:output_type -> camion.paqueteRegistro
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_camion_proto_init() }
func file_camion_proto_init() {
	if File_camion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_camion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaqueteRegistro); i {
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
		file_camion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InformeCamion); i {
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
			RawDescriptor: file_camion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_camion_proto_goTypes,
		DependencyIndexes: file_camion_proto_depIdxs,
		MessageInfos:      file_camion_proto_msgTypes,
	}.Build()
	File_camion_proto = out.File
	file_camion_proto_rawDesc = nil
	file_camion_proto_goTypes = nil
	file_camion_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CamionServiceClient is the client API for CamionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CamionServiceClient interface {
	NuevoPaquete(ctx context.Context, in *PaqueteRegistro, opts ...grpc.CallOption) (*PaqueteRegistro, error)
}

type camionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCamionServiceClient(cc grpc.ClientConnInterface) CamionServiceClient {
	return &camionServiceClient{cc}
}

func (c *camionServiceClient) NuevoPaquete(ctx context.Context, in *PaqueteRegistro, opts ...grpc.CallOption) (*PaqueteRegistro, error) {
	out := new(PaqueteRegistro)
	err := c.cc.Invoke(ctx, "/camion.CamionService/nuevoPaquete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CamionServiceServer is the server API for CamionService service.
type CamionServiceServer interface {
	NuevoPaquete(context.Context, *PaqueteRegistro) (*PaqueteRegistro, error)
}

// UnimplementedCamionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCamionServiceServer struct {
}

func (*UnimplementedCamionServiceServer) NuevoPaquete(context.Context, *PaqueteRegistro) (*PaqueteRegistro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NuevoPaquete not implemented")
}

func RegisterCamionServiceServer(s *grpc.Server, srv CamionServiceServer) {
	s.RegisterService(&_CamionService_serviceDesc, srv)
}

func _CamionService_NuevoPaquete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaqueteRegistro)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CamionServiceServer).NuevoPaquete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/camion.CamionService/NuevoPaquete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CamionServiceServer).NuevoPaquete(ctx, req.(*PaqueteRegistro))
	}
	return interceptor(ctx, in, info, handler)
}

var _CamionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "camion.CamionService",
	HandlerType: (*CamionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "nuevoPaquete",
			Handler:    _CamionService_NuevoPaquete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "camion.proto",
}