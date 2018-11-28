// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/apis/prediction_service.proto

package tensorflow_serving

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

func init() {
	proto.RegisterFile("tensorflow_serving/apis/prediction_service.proto", fileDescriptor_6f2588d3ed9ea15a)
}

var fileDescriptor_6f2588d3ed9ea15a = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x29, 0x85, 0x6f, 0xbf, 0xec, 0x41, 0x74, 0x8f, 0x39, 0x2a, 0xf5, 0x67, 0x49, 0x45,
	0xff, 0x03, 0x7b, 0x10, 0x0f, 0x01, 0x89, 0x17, 0x6f, 0x61, 0x4d, 0x26, 0x61, 0x21, 0xd9, 0x8d,
	0x3b, 0x53, 0xc5, 0xff, 0xdc, 0x83, 0x07, 0xa9, 0x3b, 0x1b, 0x69, 0x4d, 0xda, 0x5e, 0x93, 0xcf,
	0x7b, 0x33, 0xef, 0xed, 0x88, 0x6b, 0x02, 0x83, 0xd6, 0x95, 0xb5, 0x7d, 0xcf, 0x10, 0xdc, 0x9b,
	0x36, 0xd5, 0x5c, 0xb5, 0x1a, 0xe7, 0xad, 0x83, 0x42, 0xe7, 0xa4, 0xad, 0xf1, 0xdf, 0x73, 0x88,
	0x5b, 0x67, 0xc9, 0x4a, 0xf9, 0xab, 0x88, 0x59, 0x11, 0xcd, 0x86, 0x5c, 0xf2, 0x5a, 0x21, 0xea,
	0x52, 0xe7, 0x6a, 0xe5, 0xe4, 0x1d, 0xa2, 0xc1, 0x99, 0x15, 0x50, 0xd6, 0xd8, 0x02, 0xea, 0xac,
	0x01, 0x52, 0x85, 0x22, 0xc5, 0x8a, 0xb3, 0x21, 0x85, 0x36, 0x25, 0x38, 0x30, 0x61, 0xb9, 0x68,
	0xba, 0x23, 0x0e, 0x63, 0xe7, 0x43, 0x98, 0x83, 0xca, 0x01, 0x62, 0xb7, 0xeb, 0xcd, 0xd7, 0x58,
	0x1c, 0x3d, 0x76, 0x55, 0x3c, 0xf9, 0x26, 0xa4, 0x12, 0xff, 0x17, 0x3e, 0xd9, 0x87, 0xbc, 0x88,
	0xff, 0x16, 0x12, 0x2f, 0xd6, 0x72, 0xa7, 0xf0, 0xba, 0x04, 0xa4, 0xe8, 0x72, 0x1f, 0x14, 0x5b,
	0x6b, 0x10, 0xe4, 0xb3, 0x98, 0xa4, 0x7e, 0x19, 0x39, 0xed, 0x93, 0xa5, 0xdd, 0xa6, 0xc1, 0xfd,
	0x74, 0x17, 0xc6, 0xce, 0xa9, 0x98, 0x70, 0x22, 0x79, 0xdc, 0x27, 0xe1, 0x9f, 0xc1, 0xf6, 0x64,
	0x2b, 0xc3, 0x9e, 0x95, 0x38, 0x48, 0x96, 0x35, 0xe9, 0x87, 0xf0, 0x1e, 0xfd, 0xb5, 0xac, 0x33,
	0x5b, 0x6b, 0xd9, 0x44, 0x79, 0x50, 0x23, 0x0e, 0xef, 0x81, 0x92, 0xd5, 0x91, 0x24, 0x7c, 0x23,
	0xf2, 0xaa, 0x4f, 0xbf, 0x49, 0x85, 0x61, 0xb3, 0xfd, 0x60, 0x3f, 0xee, 0x6e, 0xfc, 0x39, 0x1a,
	0xbd, 0xfc, 0xfb, 0x39, 0x85, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0x88, 0x87, 0xa4,
	0x2c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	// Classify.
	Classify(ctx context.Context, in *ClassificationRequest, opts ...grpc.CallOption) (*ClassificationResponse, error)
	// Regress.
	Regress(ctx context.Context, in *RegressionRequest, opts ...grpc.CallOption) (*RegressionResponse, error)
	// Predict -- provides access to loaded TensorFlow model.
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error)
	// MultiInference API for multi-headed models.
	MultiInference(ctx context.Context, in *MultiInferenceRequest, opts ...grpc.CallOption) (*MultiInferenceResponse, error)
	// GetModelMetadata - provides access to metadata for loaded models.
	GetModelMetadata(ctx context.Context, in *GetModelMetadataRequest, opts ...grpc.CallOption) (*GetModelMetadataResponse, error)
}

type predictionServiceClient struct {
	cc *grpc.ClientConn
}

func NewPredictionServiceClient(cc *grpc.ClientConn) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) Classify(ctx context.Context, in *ClassificationRequest, opts ...grpc.CallOption) (*ClassificationResponse, error) {
	out := new(ClassificationResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/Classify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) Regress(ctx context.Context, in *RegressionRequest, opts ...grpc.CallOption) (*RegressionResponse, error) {
	out := new(RegressionResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/Regress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) MultiInference(ctx context.Context, in *MultiInferenceRequest, opts ...grpc.CallOption) (*MultiInferenceResponse, error) {
	out := new(MultiInferenceResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/MultiInference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) GetModelMetadata(ctx context.Context, in *GetModelMetadataRequest, opts ...grpc.CallOption) (*GetModelMetadataResponse, error) {
	out := new(GetModelMetadataResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/GetModelMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	// Classify.
	Classify(context.Context, *ClassificationRequest) (*ClassificationResponse, error)
	// Regress.
	Regress(context.Context, *RegressionRequest) (*RegressionResponse, error)
	// Predict -- provides access to loaded TensorFlow model.
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
	// MultiInference API for multi-headed models.
	MultiInference(context.Context, *MultiInferenceRequest) (*MultiInferenceResponse, error)
	// GetModelMetadata - provides access to metadata for loaded models.
	GetModelMetadata(context.Context, *GetModelMetadataRequest) (*GetModelMetadataResponse, error)
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_Classify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClassificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Classify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/Classify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Classify(ctx, req.(*ClassificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_Regress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Regress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/Regress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Regress(ctx, req.(*RegressionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_MultiInference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiInferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).MultiInference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/MultiInference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).MultiInference(ctx, req.(*MultiInferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_GetModelMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).GetModelMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/GetModelMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).GetModelMetadata(ctx, req.(*GetModelMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.serving.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Classify",
			Handler:    _PredictionService_Classify_Handler,
		},
		{
			MethodName: "Regress",
			Handler:    _PredictionService_Regress_Handler,
		},
		{
			MethodName: "Predict",
			Handler:    _PredictionService_Predict_Handler,
		},
		{
			MethodName: "MultiInference",
			Handler:    _PredictionService_MultiInference_Handler,
		},
		{
			MethodName: "GetModelMetadata",
			Handler:    _PredictionService_GetModelMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tensorflow_serving/apis/prediction_service.proto",
}
