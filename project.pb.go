// Code generated by protoc-gen-go.
// source: project.proto
// DO NOT EDIT!

/*
Package crypto_pb is a generated protocol buffer package.

It is generated from these files:
	project.proto

It has these top-level messages:
	CredentialOperation
	ProjectOperation
	Operation
	Credential
	Project
	ProjectOperationResponse
	CredentialOperationResponse
	Response
*/
package crypto_pb

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

type CredentialOperation_Command int32

const (
	CredentialOperation_GET    CredentialOperation_Command = 0
	CredentialOperation_SET    CredentialOperation_Command = 1
	CredentialOperation_DELETE CredentialOperation_Command = 2
)

var CredentialOperation_Command_name = map[int32]string{
	0: "GET",
	1: "SET",
	2: "DELETE",
}
var CredentialOperation_Command_value = map[string]int32{
	"GET":    0,
	"SET":    1,
	"DELETE": 2,
}

func (x CredentialOperation_Command) String() string {
	return proto.EnumName(CredentialOperation_Command_name, int32(x))
}
func (CredentialOperation_Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type ProjectOperation_Command int32

const (
	ProjectOperation_LIST             ProjectOperation_Command = 0
	ProjectOperation_CREATE           ProjectOperation_Command = 1
	ProjectOperation_UPDATE           ProjectOperation_Command = 2
	ProjectOperation_DELETE           ProjectOperation_Command = 3
	ProjectOperation_LIST_CREDENTIALS ProjectOperation_Command = 4
	ProjectOperation_ADD_MEMBER       ProjectOperation_Command = 5
	ProjectOperation_DELETE_MEMBER    ProjectOperation_Command = 6
)

var ProjectOperation_Command_name = map[int32]string{
	0: "LIST",
	1: "CREATE",
	2: "UPDATE",
	3: "DELETE",
	4: "LIST_CREDENTIALS",
	5: "ADD_MEMBER",
	6: "DELETE_MEMBER",
}
var ProjectOperation_Command_value = map[string]int32{
	"LIST":             0,
	"CREATE":           1,
	"UPDATE":           2,
	"DELETE":           3,
	"LIST_CREDENTIALS": 4,
	"ADD_MEMBER":       5,
	"DELETE_MEMBER":    6,
}

func (x ProjectOperation_Command) String() string {
	return proto.EnumName(ProjectOperation_Command_name, int32(x))
}
func (ProjectOperation_Command) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Response_Status int32

const (
	Response_ERROR   Response_Status = 0
	Response_SUCCESS Response_Status = 1
)

var Response_Status_name = map[int32]string{
	0: "ERROR",
	1: "SUCCESS",
}
var Response_Status_value = map[string]int32{
	"ERROR":   0,
	"SUCCESS": 1,
}

func (x Response_Status) String() string {
	return proto.EnumName(Response_Status_name, int32(x))
}
func (Response_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

type CredentialOperation struct {
	Command CredentialOperation_Command `protobuf:"varint,1,opt,name=command,enum=crypto_pb.CredentialOperation_Command" json:"command,omitempty"`
	Project int32                       `protobuf:"varint,2,opt,name=project" json:"project,omitempty"`
	Key     string                      `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Value   string                      `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (m *CredentialOperation) Reset()                    { *m = CredentialOperation{} }
func (m *CredentialOperation) String() string            { return proto.CompactTextString(m) }
func (*CredentialOperation) ProtoMessage()               {}
func (*CredentialOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ProjectOperation struct {
	Command     ProjectOperation_Command `protobuf:"varint,1,opt,name=command,enum=crypto_pb.ProjectOperation_Command" json:"command,omitempty"`
	Name        string                   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Environment string                   `protobuf:"bytes,3,opt,name=environment" json:"environment,omitempty"`
	ProjectId   int32                    `protobuf:"varint,4,opt,name=projectId" json:"projectId,omitempty"`
	MemberId    int32                    `protobuf:"varint,5,opt,name=memberId" json:"memberId,omitempty"`
	UserId      int32                    `protobuf:"varint,6,opt,name=userId" json:"userId,omitempty"`
	AccessLevel string                   `protobuf:"bytes,7,opt,name=accessLevel" json:"accessLevel,omitempty"`
}

func (m *ProjectOperation) Reset()                    { *m = ProjectOperation{} }
func (m *ProjectOperation) String() string            { return proto.CompactTextString(m) }
func (*ProjectOperation) ProtoMessage()               {}
func (*ProjectOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Operation struct {
	// Types that are valid to be assigned to ProjectOrCredentialOp:
	//	*Operation_ProjectOp
	//	*Operation_CredentialOp
	ProjectOrCredentialOp isOperation_ProjectOrCredentialOp `protobuf_oneof:"ProjectOrCredentialOp"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isOperation_ProjectOrCredentialOp interface {
	isOperation_ProjectOrCredentialOp()
}

type Operation_ProjectOp struct {
	ProjectOp *ProjectOperation `protobuf:"bytes,1,opt,name=projectOp,oneof"`
}
type Operation_CredentialOp struct {
	CredentialOp *CredentialOperation `protobuf:"bytes,2,opt,name=credentialOp,oneof"`
}

func (*Operation_ProjectOp) isOperation_ProjectOrCredentialOp()    {}
func (*Operation_CredentialOp) isOperation_ProjectOrCredentialOp() {}

func (m *Operation) GetProjectOrCredentialOp() isOperation_ProjectOrCredentialOp {
	if m != nil {
		return m.ProjectOrCredentialOp
	}
	return nil
}

func (m *Operation) GetProjectOp() *ProjectOperation {
	if x, ok := m.GetProjectOrCredentialOp().(*Operation_ProjectOp); ok {
		return x.ProjectOp
	}
	return nil
}

func (m *Operation) GetCredentialOp() *CredentialOperation {
	if x, ok := m.GetProjectOrCredentialOp().(*Operation_CredentialOp); ok {
		return x.CredentialOp
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operation_OneofMarshaler, _Operation_OneofUnmarshaler, _Operation_OneofSizer, []interface{}{
		(*Operation_ProjectOp)(nil),
		(*Operation_CredentialOp)(nil),
	}
}

func _Operation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operation)
	// ProjectOrCredentialOp
	switch x := m.ProjectOrCredentialOp.(type) {
	case *Operation_ProjectOp:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProjectOp); err != nil {
			return err
		}
	case *Operation_CredentialOp:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CredentialOp); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Operation.ProjectOrCredentialOp has unexpected type %T", x)
	}
	return nil
}

func _Operation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Operation)
	switch tag {
	case 1: // ProjectOrCredentialOp.projectOp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProjectOperation)
		err := b.DecodeMessage(msg)
		m.ProjectOrCredentialOp = &Operation_ProjectOp{msg}
		return true, err
	case 2: // ProjectOrCredentialOp.credentialOp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CredentialOperation)
		err := b.DecodeMessage(msg)
		m.ProjectOrCredentialOp = &Operation_CredentialOp{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Operation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Operation)
	// ProjectOrCredentialOp
	switch x := m.ProjectOrCredentialOp.(type) {
	case *Operation_ProjectOp:
		s := proto.Size(x.ProjectOp)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_CredentialOp:
		s := proto.Size(x.CredentialOp)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Credential struct {
	Id     int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Cipher string `protobuf:"bytes,3,opt,name=cipher" json:"cipher,omitempty"`
}

func (m *Credential) Reset()                    { *m = Credential{} }
func (m *Credential) String() string            { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()               {}
func (*Credential) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Project struct {
	Id          int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Environment string `protobuf:"bytes,3,opt,name=environment" json:"environment,omitempty"`
}

func (m *Project) Reset()                    { *m = Project{} }
func (m *Project) String() string            { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()               {}
func (*Project) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ProjectOperationResponse struct {
	Project     *Project      `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	MemberId    int32         `protobuf:"varint,2,opt,name=memberId" json:"memberId,omitempty"`
	Credentials []*Credential `protobuf:"bytes,3,rep,name=credentials" json:"credentials,omitempty"`
}

func (m *ProjectOperationResponse) Reset()                    { *m = ProjectOperationResponse{} }
func (m *ProjectOperationResponse) String() string            { return proto.CompactTextString(m) }
func (*ProjectOperationResponse) ProtoMessage()               {}
func (*ProjectOperationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ProjectOperationResponse) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *ProjectOperationResponse) GetCredentials() []*Credential {
	if m != nil {
		return m.Credentials
	}
	return nil
}

type CredentialOperationResponse struct {
	Credential *Credential `protobuf:"bytes,1,opt,name=credential" json:"credential,omitempty"`
}

func (m *CredentialOperationResponse) Reset()                    { *m = CredentialOperationResponse{} }
func (m *CredentialOperationResponse) String() string            { return proto.CompactTextString(m) }
func (*CredentialOperationResponse) ProtoMessage()               {}
func (*CredentialOperationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CredentialOperationResponse) GetCredential() *Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

type Response struct {
	Status Response_Status `protobuf:"varint,1,opt,name=status,enum=crypto_pb.Response_Status" json:"status,omitempty"`
	Error  string          `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Info   string          `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
	// Types that are valid to be assigned to ProjectOrCredentialResponse:
	//	*Response_ProjectOpResponse
	//	*Response_CredentialOpResponse
	ProjectOrCredentialResponse isResponse_ProjectOrCredentialResponse `protobuf_oneof:"ProjectOrCredentialResponse"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isResponse_ProjectOrCredentialResponse interface {
	isResponse_ProjectOrCredentialResponse()
}

type Response_ProjectOpResponse struct {
	ProjectOpResponse *ProjectOperationResponse `protobuf:"bytes,4,opt,name=projectOpResponse,oneof"`
}
type Response_CredentialOpResponse struct {
	CredentialOpResponse *CredentialOperationResponse `protobuf:"bytes,5,opt,name=credentialOpResponse,oneof"`
}

func (*Response_ProjectOpResponse) isResponse_ProjectOrCredentialResponse()    {}
func (*Response_CredentialOpResponse) isResponse_ProjectOrCredentialResponse() {}

func (m *Response) GetProjectOrCredentialResponse() isResponse_ProjectOrCredentialResponse {
	if m != nil {
		return m.ProjectOrCredentialResponse
	}
	return nil
}

func (m *Response) GetProjectOpResponse() *ProjectOperationResponse {
	if x, ok := m.GetProjectOrCredentialResponse().(*Response_ProjectOpResponse); ok {
		return x.ProjectOpResponse
	}
	return nil
}

func (m *Response) GetCredentialOpResponse() *CredentialOperationResponse {
	if x, ok := m.GetProjectOrCredentialResponse().(*Response_CredentialOpResponse); ok {
		return x.CredentialOpResponse
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Response) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Response_OneofMarshaler, _Response_OneofUnmarshaler, _Response_OneofSizer, []interface{}{
		(*Response_ProjectOpResponse)(nil),
		(*Response_CredentialOpResponse)(nil),
	}
}

func _Response_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Response)
	// ProjectOrCredentialResponse
	switch x := m.ProjectOrCredentialResponse.(type) {
	case *Response_ProjectOpResponse:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProjectOpResponse); err != nil {
			return err
		}
	case *Response_CredentialOpResponse:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CredentialOpResponse); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Response.ProjectOrCredentialResponse has unexpected type %T", x)
	}
	return nil
}

func _Response_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Response)
	switch tag {
	case 4: // ProjectOrCredentialResponse.projectOpResponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProjectOperationResponse)
		err := b.DecodeMessage(msg)
		m.ProjectOrCredentialResponse = &Response_ProjectOpResponse{msg}
		return true, err
	case 5: // ProjectOrCredentialResponse.credentialOpResponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CredentialOperationResponse)
		err := b.DecodeMessage(msg)
		m.ProjectOrCredentialResponse = &Response_CredentialOpResponse{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Response_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Response)
	// ProjectOrCredentialResponse
	switch x := m.ProjectOrCredentialResponse.(type) {
	case *Response_ProjectOpResponse:
		s := proto.Size(x.ProjectOpResponse)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_CredentialOpResponse:
		s := proto.Size(x.CredentialOpResponse)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*CredentialOperation)(nil), "crypto_pb.CredentialOperation")
	proto.RegisterType((*ProjectOperation)(nil), "crypto_pb.ProjectOperation")
	proto.RegisterType((*Operation)(nil), "crypto_pb.Operation")
	proto.RegisterType((*Credential)(nil), "crypto_pb.Credential")
	proto.RegisterType((*Project)(nil), "crypto_pb.Project")
	proto.RegisterType((*ProjectOperationResponse)(nil), "crypto_pb.ProjectOperationResponse")
	proto.RegisterType((*CredentialOperationResponse)(nil), "crypto_pb.CredentialOperationResponse")
	proto.RegisterType((*Response)(nil), "crypto_pb.Response")
	proto.RegisterEnum("crypto_pb.CredentialOperation_Command", CredentialOperation_Command_name, CredentialOperation_Command_value)
	proto.RegisterEnum("crypto_pb.ProjectOperation_Command", ProjectOperation_Command_name, ProjectOperation_Command_value)
	proto.RegisterEnum("crypto_pb.Response_Status", Response_Status_name, Response_Status_value)
}

func init() { proto.RegisterFile("project.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xad, 0x9d, 0xd8, 0xae, 0xaf, 0xbf, 0x56, 0xee, 0x7c, 0x2d, 0x58, 0x2d, 0xa0, 0xca, 0x48,
	0xc0, 0x02, 0x65, 0x11, 0x84, 0x58, 0x20, 0x24, 0xda, 0xd8, 0x40, 0xa4, 0x94, 0x54, 0xe3, 0x74,
	0x87, 0x14, 0x39, 0xce, 0x20, 0x0c, 0xf1, 0x8f, 0x6c, 0x27, 0x52, 0x1f, 0x86, 0x35, 0xef, 0xc1,
	0x73, 0xf0, 0x0a, 0xbc, 0x03, 0xe3, 0xf1, 0xd8, 0x1e, 0x92, 0x10, 0x24, 0x56, 0xbe, 0xf7, 0xce,
	0xb9, 0x67, 0xe6, 0x1e, 0x9f, 0x19, 0x38, 0x48, 0xb3, 0xe4, 0x33, 0x09, 0x8a, 0x1e, 0xfd, 0x16,
	0x09, 0xd2, 0x83, 0xec, 0x36, 0x2d, 0x92, 0x69, 0x3a, 0xb3, 0xbf, 0x4b, 0xf0, 0xff, 0x20, 0x23,
	0x73, 0x12, 0x17, 0xa1, 0xbf, 0x18, 0xa7, 0x24, 0xf3, 0x8b, 0x30, 0x89, 0xd1, 0x6b, 0xd0, 0x82,
	0x24, 0x8a, 0xfc, 0x78, 0x6e, 0x49, 0xe7, 0xd2, 0x93, 0xc3, 0xfe, 0xa3, 0x5e, 0xd3, 0xd4, 0xdb,
	0xd2, 0xd0, 0x1b, 0x54, 0x68, 0x5c, 0xb7, 0x21, 0x0b, 0x34, 0xbe, 0xab, 0x25, 0x53, 0x06, 0x05,
	0xd7, 0x29, 0x32, 0xa1, 0xf3, 0x85, 0xdc, 0x5a, 0x1d, 0x5a, 0xd5, 0x71, 0x19, 0xa2, 0x63, 0x50,
	0x56, 0xfe, 0x62, 0x49, 0xac, 0x2e, 0xab, 0x55, 0x89, 0xfd, 0x18, 0x34, 0xce, 0x8a, 0x34, 0xe8,
	0xbc, 0x75, 0x27, 0xe6, 0x5e, 0x19, 0x78, 0x34, 0x90, 0x10, 0x80, 0xea, 0xb8, 0x23, 0x77, 0xe2,
	0x9a, 0xb2, 0xfd, 0x53, 0x06, 0xf3, 0xba, 0x22, 0x6f, 0x27, 0x78, 0xb5, 0x3e, 0xc1, 0x43, 0x61,
	0x82, 0x75, 0xf4, 0xe6, 0xf1, 0x11, 0x74, 0x63, 0x3f, 0x22, 0xec, 0xec, 0x3a, 0x66, 0x31, 0x3a,
	0x07, 0x83, 0xc4, 0xab, 0x30, 0x4b, 0xe2, 0x88, 0x8e, 0xcf, 0x07, 0x10, 0x4b, 0xe8, 0x1e, 0xe8,
	0x7c, 0xca, 0xe1, 0x9c, 0x0d, 0xa3, 0xe0, 0xb6, 0x80, 0x4e, 0x61, 0x3f, 0x22, 0xd1, 0x8c, 0x64,
	0x74, 0x51, 0x61, 0x8b, 0x4d, 0x8e, 0xee, 0x80, 0xba, 0xcc, 0xd9, 0x8a, 0xca, 0x56, 0x78, 0x56,
	0xee, 0xe9, 0x07, 0x01, 0xc9, 0xf3, 0x11, 0x59, 0x91, 0x85, 0xa5, 0x55, 0x7b, 0x0a, 0x25, 0x3b,
	0x6d, 0x65, 0xda, 0x87, 0xee, 0x68, 0xe8, 0x95, 0x3a, 0x51, 0x79, 0x06, 0xd8, 0xbd, 0xa0, 0xf2,
	0x30, 0xa9, 0x6e, 0xae, 0x9d, 0x32, 0x96, 0x05, 0xd9, 0x3a, 0x54, 0x75, 0xb3, 0x44, 0x4f, 0x29,
	0xd0, 0x71, 0xdf, 0x4f, 0x86, 0x17, 0x23, 0xcf, 0xec, 0xa2, 0x43, 0x80, 0x0b, 0xc7, 0x99, 0x5e,
	0xb9, 0x57, 0x97, 0x2e, 0x36, 0x15, 0x74, 0x04, 0x07, 0x55, 0x47, 0x5d, 0x52, 0xed, 0x6f, 0x12,
	0xe8, 0xad, 0xd0, 0x2f, 0x9b, 0x99, 0xc7, 0x29, 0x93, 0xda, 0xe8, 0x9f, 0xed, 0x90, 0xfa, 0xdd,
	0x1e, 0x6e, 0xf1, 0xc8, 0x81, 0xff, 0x02, 0xc1, 0x4d, 0x4c, 0x6e, 0xa3, 0xff, 0x60, 0xb7, 0xd9,
	0x28, 0xc5, 0x6f, 0x5d, 0x97, 0x77, 0xe1, 0xa4, 0xde, 0x26, 0x13, 0xf1, 0xf6, 0x1b, 0x80, 0x36,
	0xa7, 0xa3, 0xc9, 0x61, 0xe5, 0x06, 0x05, 0xd3, 0xa8, 0x36, 0xa2, 0xdc, 0x1a, 0x91, 0xfe, 0x85,
	0x20, 0x4c, 0x3f, 0x91, 0x8c, 0xff, 0x5c, 0x9e, 0xd9, 0x63, 0xd0, 0xf8, 0x06, 0x1b, 0x24, 0xff,
	0x64, 0x14, 0xfb, 0xab, 0x04, 0xd6, 0xba, 0x32, 0x98, 0xe4, 0x69, 0x12, 0xe7, 0x04, 0x3d, 0x6d,
	0xaf, 0x4e, 0xa5, 0x27, 0xda, 0xd4, 0xb3, 0xbd, 0x4e, 0xa2, 0xab, 0xe4, 0x35, 0x57, 0xbd, 0x00,
	0xa3, 0x15, 0x2a, 0xa7, 0x07, 0xe9, 0x50, 0xb6, 0x93, 0xad, 0xea, 0x62, 0x11, 0x69, 0x4f, 0xe0,
	0x6c, 0x8b, 0xf0, 0xcd, 0x09, 0x9f, 0x03, 0xb4, 0x68, 0x7e, 0xc8, 0x3f, 0xd0, 0x0a, 0x40, 0xfb,
	0x87, 0x0c, 0xfb, 0x0d, 0x47, 0x1f, 0xd4, 0xbc, 0xf0, 0x8b, 0x65, 0xce, 0xef, 0xe7, 0xa9, 0xd0,
	0x5f, 0x83, 0x7a, 0x1e, 0x43, 0x60, 0x8e, 0x2c, 0x1f, 0x0a, 0x92, 0x65, 0x49, 0xc6, 0xd5, 0xae,
	0x92, 0xf2, 0x17, 0x84, 0xf1, 0xc7, 0x84, 0xeb, 0xcc, 0x62, 0xe4, 0xc1, 0x51, 0xe3, 0xb2, 0x9a,
	0x8d, 0xdd, 0x48, 0x63, 0xe7, 0x43, 0x50, 0x43, 0xa9, 0xc5, 0x36, 0xfb, 0xd1, 0x07, 0x38, 0x16,
	0x7d, 0xd7, 0xf0, 0x2a, 0x8c, 0xf7, 0x2f, 0x4f, 0xa4, 0x40, 0xbd, 0x95, 0xc5, 0x3e, 0x07, 0xb5,
	0x1a, 0x17, 0xe9, 0xa0, 0xb8, 0x18, 0x8f, 0x31, 0xbd, 0xc8, 0x06, 0x68, 0xde, 0xcd, 0x60, 0xe0,
	0x7a, 0x9e, 0x29, 0x5d, 0xde, 0x87, 0xb3, 0x2d, 0x3e, 0xaf, 0x09, 0x66, 0x2a, 0x7b, 0xde, 0x9f,
	0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x12, 0x75, 0x58, 0xef, 0x05, 0x00, 0x00,
}
