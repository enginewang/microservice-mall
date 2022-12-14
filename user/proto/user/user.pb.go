// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	UserRegisterRequest
	UserRegisterResponse
	UserLoginRequest
	UserLoginResponse
	UserInfoRequest
	UserInfoResponse
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserRegisterRequest struct {
	UserName  string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	Pwd       string `protobuf:"bytes,3,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *UserRegisterRequest) Reset()                    { *m = UserRegisterRequest{} }
func (m *UserRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRegisterRequest) ProtoMessage()               {}
func (*UserRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserRegisterRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserRegisterRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserRegisterRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type UserRegisterResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *UserRegisterResponse) Reset()                    { *m = UserRegisterResponse{} }
func (m *UserRegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*UserRegisterResponse) ProtoMessage()               {}
func (*UserRegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserRegisterResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UserLoginRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Pwd      string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *UserLoginRequest) Reset()                    { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()               {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserLoginRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserLoginRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type UserLoginResponse struct {
	IsSuccess bool `protobuf:"varint,1,opt,name=is_success,json=isSuccess" json:"is_success,omitempty"`
}

func (m *UserLoginResponse) Reset()                    { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string            { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()               {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserLoginResponse) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

type UserInfoRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
}

func (m *UserInfoRequest) Reset()                    { *m = UserInfoRequest{} }
func (m *UserInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*UserInfoRequest) ProtoMessage()               {}
func (*UserInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserInfoRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type UserInfoResponse struct {
	UserId    int64  `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserName  string `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
}

func (m *UserInfoResponse) Reset()                    { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()               {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserInfoResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfoResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserInfoResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRegisterRequest)(nil), "go.micro.service.user.UserRegisterRequest")
	proto.RegisterType((*UserRegisterResponse)(nil), "go.micro.service.user.UserRegisterResponse")
	proto.RegisterType((*UserLoginRequest)(nil), "go.micro.service.user.UserLoginRequest")
	proto.RegisterType((*UserLoginResponse)(nil), "go.micro.service.user.UserLoginResponse")
	proto.RegisterType((*UserInfoRequest)(nil), "go.micro.service.user.UserInfoRequest")
	proto.RegisterType((*UserInfoResponse)(nil), "go.micro.service.user.UserInfoResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0xb5, 0x89, 0xb6, 0xcd, 0x28, 0x5a, 0x57, 0xc5, 0x52, 0x29, 0xc8, 0x1e, 0x6c, 0x51, 0x58,
	0xa5, 0x1e, 0x3d, 0xe9, 0x45, 0x0a, 0xe2, 0x21, 0xe2, 0x45, 0x84, 0x5a, 0xd3, 0x69, 0xd8, 0x43,
	0xb2, 0x75, 0x27, 0xd1, 0xaf, 0xf2, 0x1f, 0x65, 0x37, 0x89, 0xa6, 0x52, 0xd2, 0x9c, 0xb2, 0x3b,
	0xef, 0xbd, 0x79, 0x93, 0x37, 0x0b, 0x90, 0x12, 0x6a, 0xb1, 0xd0, 0x2a, 0x51, 0xec, 0x28, 0x54,
	0x22, 0x92, 0x81, 0x56, 0x82, 0x50, 0x7f, 0xca, 0x00, 0x85, 0x01, 0x79, 0x00, 0x07, 0xcf, 0x84,
	0xda, 0xc7, 0x50, 0x52, 0x62, 0xbe, 0x1f, 0x29, 0x52, 0xc2, 0x4e, 0xc0, 0x33, 0xf0, 0x24, 0x9e,
	0x46, 0xd8, 0x6d, 0x9c, 0x36, 0x86, 0x9e, 0xdf, 0x36, 0x85, 0xc7, 0x69, 0x84, 0xac, 0x0f, 0x30,
	0x97, 0x9a, 0x92, 0x0c, 0x75, 0x2c, 0xea, 0xd9, 0x8a, 0x85, 0x3b, 0xe0, 0x2e, 0xbe, 0x66, 0x5d,
	0xd7, 0xd6, 0xcd, 0x91, 0x5f, 0xc1, 0xe1, 0xb2, 0x09, 0x2d, 0x54, 0x4c, 0xc8, 0xba, 0xd0, 0x8a,
	0x90, 0x68, 0x1a, 0x16, 0x1e, 0xc5, 0x95, 0xdf, 0x42, 0xc7, 0x28, 0x1e, 0x54, 0x28, 0xe3, 0x5a,
	0x33, 0xe5, 0xa6, 0xce, 0x9f, 0xe9, 0x08, 0xf6, 0x4b, 0x2d, 0x72, 0xc7, 0x3e, 0x80, 0xa4, 0x09,
	0xa5, 0x41, 0x80, 0x44, 0xb6, 0x49, 0xdb, 0xf7, 0x24, 0x3d, 0x65, 0x05, 0x2e, 0x60, 0xcf, 0x68,
	0xc6, 0xf1, 0x5c, 0xd5, 0x71, 0xe5, 0x61, 0x36, 0x66, 0xc6, 0xcf, 0x2d, 0x8e, 0xa1, 0x65, 0x05,
	0x72, 0x66, 0xe9, 0xae, 0xdf, 0x34, 0xd7, 0xf1, 0x6c, 0xb9, 0x93, 0x53, 0x99, 0xa9, 0xfb, 0x2f,
	0xd3, 0xd1, 0xb7, 0x03, 0x9b, 0xc6, 0x89, 0x21, 0xb4, 0x8b, 0x18, 0xd9, 0xb9, 0x58, 0xb9, 0x53,
	0xb1, 0x62, 0xa1, 0xbd, 0x8b, 0x5a, 0xdc, 0xec, 0x17, 0xf8, 0x06, 0x7b, 0x85, 0x2d, 0x1b, 0x1c,
	0x1b, 0x54, 0xe8, 0xca, 0xdb, 0xe9, 0x0d, 0xd7, 0x13, 0x7f, 0xbb, 0xbf, 0xc1, 0xf6, 0x3d, 0x26,
	0x45, 0x72, 0xec, 0xac, 0x42, 0x5a, 0x5a, 0x45, 0x6f, 0xb0, 0x96, 0x57, 0x38, 0xdc, 0xed, 0xbe,
	0xec, 0x88, 0x4b, 0xfb, 0xf0, 0x6f, 0x0c, 0xe5, 0xbd, 0x69, 0xcf, 0xd7, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x8e, 0xf5, 0xa8, 0x02, 0x12, 0x03, 0x00, 0x00,
}
