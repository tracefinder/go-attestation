// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tpm.proto

package go_attestation_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type TpmVersion int32

const (
	TpmVersion_TPM_VERSION_UNSPECIFIED TpmVersion = 0
	TpmVersion_TPM_12                  TpmVersion = 1
	TpmVersion_TPM_20                  TpmVersion = 2
)

var TpmVersion_name = map[int32]string{
	0: "TPM_VERSION_UNSPECIFIED",
	1: "TPM_12",
	2: "TPM_20",
}

var TpmVersion_value = map[string]int32{
	"TPM_VERSION_UNSPECIFIED": 0,
	"TPM_12":                  1,
	"TPM_20":                  2,
}

func (x TpmVersion) String() string {
	return proto.EnumName(TpmVersion_name, int32(x))
}

func (TpmVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_63ac7bc02f9d1279, []int{0}
}

type TpmInterface int32

const (
	TpmInterface_TPM_INTERFACE_UNSPECIFIED TpmInterface = 0
	TpmInterface_DIRECT                    TpmInterface = 1
	TpmInterface_KERNEL_MANAGED            TpmInterface = 2
	TpmInterface_DAEMON_MANAGED            TpmInterface = 3
)

var TpmInterface_name = map[int32]string{
	0: "TPM_INTERFACE_UNSPECIFIED",
	1: "DIRECT",
	2: "KERNEL_MANAGED",
	3: "DAEMON_MANAGED",
}

var TpmInterface_value = map[string]int32{
	"TPM_INTERFACE_UNSPECIFIED": 0,
	"DIRECT":                    1,
	"KERNEL_MANAGED":            2,
	"DAEMON_MANAGED":            3,
}

func (x TpmInterface) String() string {
	return proto.EnumName(TpmInterface_name, int32(x))
}

func (TpmInterface) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_63ac7bc02f9d1279, []int{1}
}

type EndorsementKey_DataType int32

const (
	EndorsementKey_DATA_TYPE_UNSPECIFIED EndorsementKey_DataType = 0
	EndorsementKey_PUBLIC_BLOB           EndorsementKey_DataType = 1
	EndorsementKey_X509_CERT_BLOB        EndorsementKey_DataType = 2
)

var EndorsementKey_DataType_name = map[int32]string{
	0: "DATA_TYPE_UNSPECIFIED",
	1: "PUBLIC_BLOB",
	2: "X509_CERT_BLOB",
}

var EndorsementKey_DataType_value = map[string]int32{
	"DATA_TYPE_UNSPECIFIED": 0,
	"PUBLIC_BLOB":           1,
	"X509_CERT_BLOB":        2,
}

func (x EndorsementKey_DataType) String() string {
	return proto.EnumName(EndorsementKey_DataType_name, int32(x))
}

func (EndorsementKey_DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_63ac7bc02f9d1279, []int{1, 0}
}

type ChallengeInfo_ChallengeType int32

const (
	ChallengeInfo_CHALLENGE_UNSPECIFIED ChallengeInfo_ChallengeType = 0
	ChallengeInfo_CHALLENGE_CA          ChallengeInfo_ChallengeType = 1
)

var ChallengeInfo_ChallengeType_name = map[int32]string{
	0: "CHALLENGE_UNSPECIFIED",
	1: "CHALLENGE_CA",
}

var ChallengeInfo_ChallengeType_value = map[string]int32{
	"CHALLENGE_UNSPECIFIED": 0,
	"CHALLENGE_CA":          1,
}

func (x ChallengeInfo_ChallengeType) String() string {
	return proto.EnumName(ChallengeInfo_ChallengeType_name, int32(x))
}

func (ChallengeInfo_ChallengeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_63ac7bc02f9d1279, []int{5, 0}
}

type StatusReport_ReportType int32

const (
	StatusReport_REPORT_UNSPECIFIED           StatusReport_ReportType = 0
	StatusReport_REPORT_TPM_UNSUITABLE        StatusReport_ReportType = 1
	StatusReport_REPORT_TPM_OPERATION_FAILURE StatusReport_ReportType = 2
	StatusReport_REPORT_LOG_UNAVAILABLE       StatusReport_ReportType = 3
)

var StatusReport_ReportType_name = map[int32]string{
	0: "REPORT_UNSPECIFIED",
	1: "REPORT_TPM_UNSUITABLE",
	2: "REPORT_TPM_OPERATION_FAILURE",
	3: "REPORT_LOG_UNAVAILABLE",
}

var StatusReport_ReportType_value = map[string]int32{
	"REPORT_UNSPECIFIED":           0,
	"REPORT_TPM_UNSUITABLE":        1,
	"REPORT_TPM_OPERATION_FAILURE": 2,
	"REPORT_LOG_UNAVAILABLE":       3,
}

func (x StatusReport_ReportType) String() string {
	return proto.EnumName(StatusReport_ReportType_name, int32(x))
}

func (StatusReport_ReportType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_63ac7bc02f9d1279, []int{7, 0}
}

// TpmInfo encapsulates version / device information
// about the TPM, and how the attestation client interfaces
// with it.
type TpmInfo struct {
	TpmVersion   TpmVersion   `protobuf:"varint,1,opt,name=tpm_version,json=tpmVersion,proto3,enum=go_attestation.proto.TpmVersion" json:"tpm_version,omitempty"`
	Manufacturer string       `protobuf:"bytes,2,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	TpmInterface TpmInterface `protobuf:"varint,3,opt,name=tpm_interface,json=tpmInterface,proto3,enum=go_attestation.proto.TpmInterface" json:"tpm_interface,omitempty"`
	// This number represents the version of the support code which
	// interfaces with the TPM.
	TpmInterfaceVersion uint32 `protobuf:"varint,4,opt,name=tpm_interface_version,json=tpmInterfaceVersion,proto3" json:"tpm_interface_version,omitempty"` // Deprecated: Do not use.
	// This is the string provided by the TPM.
	TpmOpaqueInfo string `protobuf:"bytes,5,opt,name=tpm_opaque_info,json=tpmOpaqueInfo,proto3" json:"tpm_opaque_info,omitempty"`
	// This is set if challenges must be generated
	// in TrouSerS format for TPM 1.2 devices.
	TrousersFormat       bool     `protobuf:"varint,6,opt,name=trousers_format,json=trousersFormat,proto3" json:"trousers_format,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TpmInfo) Reset()         { *m = TpmInfo{} }
func (m *TpmInfo) String() string { return proto.CompactTextString(m) }
func (*TpmInfo) ProtoMessage()    {}
func (*TpmInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_63ac7bc02f9d1279, []int{0}
}

func (m *TpmInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpmInfo.Unmarshal(m, b)
}
func (m *TpmInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpmInfo.Marshal(b, m, deterministic)
}
func (m *TpmInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpmInfo.Merge(m, src)
}
func (m *TpmInfo) XXX_Size() int {
	return xxx_messageInfo_TpmInfo.Size(m)
}
func (m *TpmInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TpmInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TpmInfo proto.InternalMessageInfo

func (m *TpmInfo) GetTpmVersion() TpmVersion {
	if m != nil {
		return m.TpmVersion
	}
	return TpmVersion_TPM_VERSION_UNSPECIFIED
}

func (m *TpmInfo) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *TpmInfo) GetTpmInterface() TpmInterface {
	if m != nil {
		return m.TpmInterface
	}
	return TpmInterface_TPM_INTERFACE_UNSPECIFIED
}

// Deprecated: Do not use.
func (m *TpmInfo) GetTpmInterfaceVersion() uint32 {
	if m != nil {
		return m.TpmInterfaceVersion
	}
	return 0
}

func (m *TpmInfo) GetTpmOpaqueInfo() string {
	if m != nil {
		return m.TpmOpaqueInfo
	}
	return ""
}

func (m *TpmInfo) GetTrousersFormat() bool {
	if m != nil {
		return m.TrousersFormat
	}
	return false
}

type EndorsementKey struct {
	Datatype             EndorsementKey_DataType `protobuf:"varint,1,opt,name=datatype,proto3,enum=go_attestation.proto.EndorsementKey_DataType" json:"datatype,omitempty"`
	Data                 []byte                  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *EndorsementKey) Reset()         { *m = EndorsementKey{} }
func (m *EndorsementKey) String() string { return proto.CompactTextString(m) }
func (*EndorsementKey) ProtoMessage()    {}
func (*EndorsementKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_63ac7bc02f9d1279, []int{1}
}

func (m *EndorsementKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndorsementKey.Unmarshal(m, b)
}
func (m *EndorsementKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndorsementKey.Marshal(b, m, deterministic)
}
func (m *EndorsementKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndorsementKey.Merge(m, src)
}
func (m *EndorsementKey) XXX_Size() int {
	return xxx_messageInfo_EndorsementKey.Size(m)
}
func (m *EndorsementKey) XXX_DiscardUnknown() {
	xxx_messageInfo_EndorsementKey.DiscardUnknown(m)
}

var xxx_messageInfo_EndorsementKey proto.InternalMessageInfo

func (m *EndorsementKey) GetDatatype() EndorsementKey_DataType {
	if m != nil {
		return m.Datatype
	}
	return EndorsementKey_DATA_TYPE_UNSPECIFIED
}

func (m *EndorsementKey) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Tpm20AikInfo describes an AIK using TPM 2.0 structures.
type Tpm20AikInfo struct {
	// This is a TPMT_PUBLIC structure.
	PublicBlob []byte `protobuf:"bytes,1,opt,name=public_blob,json=publicBlob,proto3" json:"public_blob,omitempty"`
	// This is a TPMS_CREATION_DATA structure.
	CreationData []byte `protobuf:"bytes,2,opt,name=creation_data,json=creationData,proto3" json:"creation_data,omitempty"`
	// This is a TPMU_ATTEST structure, with the dynamic section
	// containing a CREATION_INFO structure.
	AttestationData []byte `protobuf:"bytes,3,opt,name=attestation_data,json=attestationData,proto3" json:"attestation_data,omitempty"`
	// This is a TPMT_SIGNATURE structure.
	SignatureData        []byte   `protobuf:"bytes,4,opt,name=signature_data,json=signatureData,proto3" json:"signature_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tpm20AikInfo) Reset()         { *m = Tpm20AikInfo{} }
func (m *Tpm20AikInfo) String() string { return proto.CompactTextString(m) }
func (*Tpm20AikInfo) ProtoMessage()    {}
func (*Tpm20AikInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_63ac7bc02f9d1279, []int{2}
}

func (m *Tpm20AikInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tpm20AikInfo.Unmarshal(m, b)
}
func (m *Tpm20AikInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tpm20AikInfo.Marshal(b, m, deterministic)
}
func (m *Tpm20AikInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tpm20AikInfo.Merge(m, src)
}
func (m *Tpm20AikInfo) XXX_Size() int {
	return xxx_messageInfo_Tpm20AikInfo.Size(m)
}
func (m *Tpm20AikInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Tpm20AikInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Tpm20AikInfo proto.InternalMessageInfo

func (m *Tpm20AikInfo) GetPublicBlob() []byte {
	if m != nil {
		return m.PublicBlob
	}
	return nil
}

func (m *Tpm20AikInfo) GetCreationData() []byte {
	if m != nil {
		return m.CreationData
	}
	return nil
}

func (m *Tpm20AikInfo) GetAttestationData() []byte {
	if m != nil {
		return m.AttestationData
	}
	return nil
}

func (m *Tpm20AikInfo) GetSignatureData() []byte {
	if m != nil {
		return m.SignatureData
	}
	return nil
}

// Tpm12AikInfo describes an AIK using TPM 1.2 structures.
type Tpm12AikInfo struct {
	// This is a TPM_PUBKEY structure.
	PublicBlob []byte `protobuf:"bytes,1,opt,name=public_blob,json=publicBlob,proto3" json:"public_blob,omitempty"`
	// This is auxillary data, provided for the purpose of debugging.
	// on Windows devices, this represents the contents of PCP_ID_BINDING.
	Aux                  []byte   `protobuf:"bytes,2,opt,name=aux,proto3" json:"aux,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tpm12AikInfo) Reset()         { *m = Tpm12AikInfo{} }
func (m *Tpm12AikInfo) String() string { return proto.CompactTextString(m) }
func (*Tpm12AikInfo) ProtoMessage()    {}
func (*Tpm12AikInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_63ac7bc02f9d1279, []int{3}
}

func (m *Tpm12AikInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tpm12AikInfo.Unmarshal(m, b)
}
func (m *Tpm12AikInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tpm12AikInfo.Marshal(b, m, deterministic)
}
func (m *Tpm12AikInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tpm12AikInfo.Merge(m, src)
}
func (m *Tpm12AikInfo) XXX_Size() int {
	return xxx_messageInfo_Tpm12AikInfo.Size(m)
}
func (m *Tpm12AikInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Tpm12AikInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Tpm12AikInfo proto.InternalMessageInfo

func (m *Tpm12AikInfo) GetPublicBlob() []byte {
	if m != nil {
		return m.PublicBlob
	}
	return nil
}

func (m *Tpm12AikInfo) GetAux() []byte {
	if m != nil {
		return m.Aux
	}
	return nil
}

// AikInfo describes the public key, parameters, and creation information
// of an attestation identity key.
type AikInfo struct {
	// Types that are valid to be assigned to TpmAikInfo:
	//	*AikInfo_Tpm20
	//	*AikInfo_Tpm12
	TpmAikInfo           isAikInfo_TpmAikInfo `protobuf_oneof:"tpm_aik_info"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AikInfo) Reset()         { *m = AikInfo{} }
func (m *AikInfo) String() string { return proto.CompactTextString(m) }
func (*AikInfo) ProtoMessage()    {}
func (*AikInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_63ac7bc02f9d1279, []int{4}
}

func (m *AikInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AikInfo.Unmarshal(m, b)
}
func (m *AikInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AikInfo.Marshal(b, m, deterministic)
}
func (m *AikInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AikInfo.Merge(m, src)
}
func (m *AikInfo) XXX_Size() int {
	return xxx_messageInfo_AikInfo.Size(m)
}
func (m *AikInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AikInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AikInfo proto.InternalMessageInfo

type isAikInfo_TpmAikInfo interface {
	isAikInfo_TpmAikInfo()
}

type AikInfo_Tpm20 struct {
	Tpm20 *Tpm20AikInfo `protobuf:"bytes,1,opt,name=tpm20,proto3,oneof"`
}

type AikInfo_Tpm12 struct {
	Tpm12 *Tpm12AikInfo `protobuf:"bytes,2,opt,name=tpm12,proto3,oneof"`
}

func (*AikInfo_Tpm20) isAikInfo_TpmAikInfo() {}

func (*AikInfo_Tpm12) isAikInfo_TpmAikInfo() {}

func (m *AikInfo) GetTpmAikInfo() isAikInfo_TpmAikInfo {
	if m != nil {
		return m.TpmAikInfo
	}
	return nil
}

func (m *AikInfo) GetTpm20() *Tpm20AikInfo {
	if x, ok := m.GetTpmAikInfo().(*AikInfo_Tpm20); ok {
		return x.Tpm20
	}
	return nil
}

func (m *AikInfo) GetTpm12() *Tpm12AikInfo {
	if x, ok := m.GetTpmAikInfo().(*AikInfo_Tpm12); ok {
		return x.Tpm12
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AikInfo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AikInfo_Tpm20)(nil),
		(*AikInfo_Tpm12)(nil),
	}
}

// ChallengeInfo describes which challenge a nonce corresponds to.
type ChallengeInfo struct {
	Type                 ChallengeInfo_ChallengeType `protobuf:"varint,1,opt,name=type,proto3,enum=go_attestation.proto.ChallengeInfo_ChallengeType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ChallengeInfo) Reset()         { *m = ChallengeInfo{} }
func (m *ChallengeInfo) String() string { return proto.CompactTextString(m) }
func (*ChallengeInfo) ProtoMessage()    {}
func (*ChallengeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_63ac7bc02f9d1279, []int{5}
}

func (m *ChallengeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChallengeInfo.Unmarshal(m, b)
}
func (m *ChallengeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChallengeInfo.Marshal(b, m, deterministic)
}
func (m *ChallengeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengeInfo.Merge(m, src)
}
func (m *ChallengeInfo) XXX_Size() int {
	return xxx_messageInfo_ChallengeInfo.Size(m)
}
func (m *ChallengeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengeInfo proto.InternalMessageInfo

func (m *ChallengeInfo) GetType() ChallengeInfo_ChallengeType {
	if m != nil {
		return m.Type
	}
	return ChallengeInfo_CHALLENGE_UNSPECIFIED
}

// ClientInfo is optional data sent from the client to identify what version
// of racc-client it is running.
type ClientInfo struct {
	MachineTrack         string   `protobuf:"bytes,1,opt,name=machine_track,json=machineTrack,proto3" json:"machine_track,omitempty"`
	ClRollup             string   `protobuf:"bytes,2,opt,name=cl_rollup,json=clRollup,proto3" json:"cl_rollup,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientInfo) Reset()         { *m = ClientInfo{} }
func (m *ClientInfo) String() string { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()    {}
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_63ac7bc02f9d1279, []int{6}
}

func (m *ClientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInfo.Unmarshal(m, b)
}
func (m *ClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInfo.Marshal(b, m, deterministic)
}
func (m *ClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInfo.Merge(m, src)
}
func (m *ClientInfo) XXX_Size() int {
	return xxx_messageInfo_ClientInfo.Size(m)
}
func (m *ClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInfo proto.InternalMessageInfo

func (m *ClientInfo) GetMachineTrack() string {
	if m != nil {
		return m.MachineTrack
	}
	return ""
}

func (m *ClientInfo) GetClRollup() string {
	if m != nil {
		return m.ClRollup
	}
	return ""
}

func (m *ClientInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// StatusReport describes information from a client which is distinct to any
// attestation operation.
type StatusReport struct {
	Type                 StatusReport_ReportType `protobuf:"varint,1,opt,name=type,proto3,enum=go_attestation.proto.StatusReport_ReportType" json:"type,omitempty"`
	Code                 int64                   `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string                  `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Operation            string                  `protobuf:"bytes,4,opt,name=operation,proto3" json:"operation,omitempty"`
	ClientInfo           *ClientInfo             `protobuf:"bytes,5,opt,name=client_info,json=clientInfo,proto3" json:"client_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StatusReport) Reset()         { *m = StatusReport{} }
func (m *StatusReport) String() string { return proto.CompactTextString(m) }
func (*StatusReport) ProtoMessage()    {}
func (*StatusReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_63ac7bc02f9d1279, []int{7}
}

func (m *StatusReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusReport.Unmarshal(m, b)
}
func (m *StatusReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusReport.Marshal(b, m, deterministic)
}
func (m *StatusReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusReport.Merge(m, src)
}
func (m *StatusReport) XXX_Size() int {
	return xxx_messageInfo_StatusReport.Size(m)
}
func (m *StatusReport) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusReport.DiscardUnknown(m)
}

var xxx_messageInfo_StatusReport proto.InternalMessageInfo

func (m *StatusReport) GetType() StatusReport_ReportType {
	if m != nil {
		return m.Type
	}
	return StatusReport_REPORT_UNSPECIFIED
}

func (m *StatusReport) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *StatusReport) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StatusReport) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *StatusReport) GetClientInfo() *ClientInfo {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

func init() {
	proto.RegisterEnum("go_attestation.proto.TpmVersion", TpmVersion_name, TpmVersion_value)
	proto.RegisterEnum("go_attestation.proto.TpmInterface", TpmInterface_name, TpmInterface_value)
	proto.RegisterEnum("go_attestation.proto.EndorsementKey_DataType", EndorsementKey_DataType_name, EndorsementKey_DataType_value)
	proto.RegisterEnum("go_attestation.proto.ChallengeInfo_ChallengeType", ChallengeInfo_ChallengeType_name, ChallengeInfo_ChallengeType_value)
	proto.RegisterEnum("go_attestation.proto.StatusReport_ReportType", StatusReport_ReportType_name, StatusReport_ReportType_value)
	proto.RegisterType((*TpmInfo)(nil), "go_attestation.proto.TpmInfo")
	proto.RegisterType((*EndorsementKey)(nil), "go_attestation.proto.EndorsementKey")
	proto.RegisterType((*Tpm20AikInfo)(nil), "go_attestation.proto.Tpm20AikInfo")
	proto.RegisterType((*Tpm12AikInfo)(nil), "go_attestation.proto.Tpm12AikInfo")
	proto.RegisterType((*AikInfo)(nil), "go_attestation.proto.AikInfo")
	proto.RegisterType((*ChallengeInfo)(nil), "go_attestation.proto.ChallengeInfo")
	proto.RegisterType((*ClientInfo)(nil), "go_attestation.proto.ClientInfo")
	proto.RegisterType((*StatusReport)(nil), "go_attestation.proto.StatusReport")
}

func init() { proto.RegisterFile("tpm.proto", fileDescriptor_63ac7bc02f9d1279) }

var fileDescriptor_63ac7bc02f9d1279 = []byte{
	// 848 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xef, 0x6e, 0xe3, 0x44,
	0x10, 0x3f, 0x27, 0xbd, 0xb6, 0x99, 0xfc, 0xa9, 0xb5, 0xd0, 0x23, 0xe5, 0x0e, 0x11, 0xf9, 0x04,
	0x94, 0x93, 0x88, 0x9a, 0x20, 0x90, 0x40, 0x7c, 0xd9, 0x38, 0xdb, 0x9e, 0x39, 0x37, 0x89, 0xb6,
	0x4e, 0x05, 0x9f, 0xac, 0x8d, 0xbb, 0x49, 0xad, 0xfa, 0x1f, 0xf6, 0x06, 0xd1, 0x0f, 0x3c, 0x04,
	0x12, 0xcf, 0xc0, 0x4b, 0xf0, 0x8d, 0x17, 0xe1, 0x55, 0xd0, 0xae, 0xe3, 0xd8, 0x39, 0xda, 0xd3,
	0x7d, 0xca, 0xec, 0x6f, 0xe6, 0x37, 0xfe, 0xcd, 0x64, 0x66, 0xa0, 0x21, 0x92, 0xb0, 0x9f, 0xa4,
	0xb1, 0x88, 0xd1, 0x87, 0xab, 0xd8, 0x65, 0x42, 0xf0, 0x4c, 0x30, 0xe1, 0xc7, 0x51, 0x8e, 0x1a,
	0xff, 0xd4, 0xe0, 0xc0, 0x49, 0x42, 0x2b, 0x5a, 0xc6, 0x08, 0x43, 0x53, 0x24, 0xa1, 0xfb, 0x2b,
	0x4f, 0x33, 0x3f, 0x8e, 0xba, 0x5a, 0x4f, 0x3b, 0xed, 0x0c, 0x7b, 0xfd, 0x87, 0x78, 0x7d, 0x27,
	0x09, 0xaf, 0xf3, 0x38, 0x0a, 0x62, 0x6b, 0x23, 0x03, 0x5a, 0x21, 0x8b, 0xd6, 0x4b, 0xe6, 0x89,
	0x75, 0xca, 0xd3, 0x6e, 0xad, 0xa7, 0x9d, 0x36, 0xe8, 0x0e, 0x86, 0x2e, 0xa0, 0x2d, 0x3f, 0xe3,
	0x47, 0x82, 0xa7, 0x4b, 0xe6, 0xf1, 0x6e, 0x5d, 0x7d, 0xc8, 0x78, 0xf4, 0x43, 0x56, 0x11, 0x49,
	0x5b, 0xa2, 0xf2, 0x42, 0xdf, 0xc2, 0xf1, 0x4e, 0xa2, 0xad, 0xf2, 0xbd, 0x9e, 0x76, 0xda, 0x1e,
	0xd5, 0xba, 0x1a, 0xfd, 0xa0, 0x4a, 0x28, 0x44, 0x7e, 0x0e, 0x47, 0x92, 0x17, 0x27, 0xec, 0x97,
	0x35, 0x77, 0xfd, 0x68, 0x19, 0x77, 0x9f, 0x2a, 0x9d, 0x52, 0xd7, 0x54, 0xa1, 0xaa, 0x1f, 0x5f,
	0xc0, 0x91, 0x48, 0xe3, 0x75, 0xc6, 0xd3, 0xcc, 0x5d, 0xc6, 0x69, 0xc8, 0x44, 0x77, 0xbf, 0xa7,
	0x9d, 0x1e, 0xd2, 0x4e, 0x01, 0x9f, 0x2b, 0xd4, 0xf8, 0x5b, 0x83, 0x0e, 0x89, 0x6e, 0xe2, 0x34,
	0xe3, 0x21, 0x8f, 0xc4, 0x1b, 0x7e, 0x8f, 0x2c, 0x38, 0xbc, 0x61, 0x82, 0x89, 0xfb, 0x84, 0x6f,
	0x1a, 0xf9, 0xd5, 0xc3, 0xf5, 0xed, 0xf2, 0xfa, 0x63, 0x26, 0x98, 0x73, 0x9f, 0x70, 0xba, 0xa5,
	0x23, 0x04, 0x7b, 0xd2, 0x56, 0xbd, 0x6c, 0x51, 0x65, 0x1b, 0x3f, 0xc2, 0x61, 0x11, 0x89, 0x4e,
	0xe0, 0x78, 0x8c, 0x1d, 0xec, 0x3a, 0x3f, 0xcf, 0x88, 0x3b, 0x9f, 0x5c, 0xcd, 0x88, 0x69, 0x9d,
	0x5b, 0x64, 0xac, 0x3f, 0x41, 0x47, 0xd0, 0x9c, 0xcd, 0x47, 0xb6, 0x65, 0xba, 0x23, 0x7b, 0x3a,
	0xd2, 0x35, 0x84, 0xa0, 0xf3, 0xd3, 0x37, 0x67, 0xdf, 0xb9, 0x26, 0xa1, 0x4e, 0x8e, 0xd5, 0x8c,
	0xbf, 0x34, 0x68, 0x39, 0x49, 0x38, 0x3c, 0xc3, 0xfe, 0x9d, 0xaa, 0xfb, 0x53, 0x68, 0x26, 0xeb,
	0x45, 0xe0, 0x7b, 0xee, 0x22, 0x88, 0x17, 0x4a, 0x7e, 0x8b, 0x42, 0x0e, 0x8d, 0x82, 0x78, 0x81,
	0x5e, 0x42, 0xdb, 0x4b, 0xb9, 0xaa, 0xc2, 0xad, 0x48, 0x6b, 0x15, 0xa0, 0x94, 0x86, 0xbe, 0x04,
	0xbd, 0x52, 0x6d, 0x1e, 0x57, 0x57, 0x71, 0x47, 0x15, 0x5c, 0x85, 0x7e, 0x06, 0x9d, 0xcc, 0x5f,
	0x45, 0x4c, 0xce, 0x47, 0x1e, 0xb8, 0xa7, 0x02, 0xdb, 0x5b, 0x54, 0x86, 0x19, 0x58, 0xe9, 0x1c,
	0x0c, 0xdf, 0x5b, 0xa7, 0x0e, 0x75, 0xb6, 0xfe, 0x6d, 0xa3, 0x4e, 0x9a, 0xc6, 0x1f, 0x1a, 0x1c,
	0x14, 0xf4, 0xef, 0xe1, 0xa9, 0x90, 0x65, 0x2b, 0x62, 0xf3, 0x1d, 0xf3, 0xb7, 0xed, 0xcc, 0xeb,
	0x27, 0x34, 0xa7, 0x6c, 0xb8, 0x83, 0xa1, 0xca, 0xfd, 0x2e, 0xee, 0x56, 0xed, 0x86, 0x3b, 0x18,
	0x8e, 0x3a, 0x20, 0xc7, 0xd8, 0x65, 0xfe, 0x9d, 0x9a, 0x3d, 0xe3, 0x4f, 0x0d, 0xda, 0xe6, 0x2d,
	0x0b, 0x02, 0x1e, 0xad, 0xf2, 0xc1, 0x23, 0xb0, 0x57, 0x19, 0x9c, 0xc1, 0xc3, 0xc9, 0x77, 0x28,
	0xe5, 0x4b, 0x0d, 0x8f, 0xa2, 0x1b, 0x3f, 0x54, 0xf2, 0x16, 0x93, 0x62, 0xbe, 0xc6, 0xb6, 0x4d,
	0x26, 0x17, 0x6f, 0x4f, 0x8a, 0x0e, 0xad, 0xd2, 0x65, 0x62, 0x5d, 0x33, 0x6e, 0x01, 0xcc, 0xc0,
	0xe7, 0x91, 0x50, 0x92, 0x5e, 0x42, 0x3b, 0x64, 0xde, 0xad, 0x1f, 0x71, 0x57, 0xa4, 0xcc, 0xbb,
	0x53, 0xda, 0xd4, 0x66, 0x2b, 0xd0, 0x91, 0x18, 0x7a, 0x0e, 0x0d, 0x2f, 0x70, 0xd3, 0x38, 0x08,
	0xd6, 0xc9, 0x66, 0xf5, 0x0f, 0xbd, 0x80, 0xaa, 0x37, 0xea, 0xc2, 0x41, 0xb1, 0x9f, 0x75, 0xe5,
	0x2a, 0x9e, 0xc6, 0xbf, 0x35, 0x68, 0x5d, 0x09, 0x26, 0xd6, 0x19, 0xe5, 0x49, 0x9c, 0x0a, 0x84,
	0x77, 0xea, 0x7f, 0x64, 0x71, 0xaa, 0x8c, 0x7e, 0xfe, 0x53, 0xd6, 0x2e, 0x97, 0xc6, 0x8b, 0x6f,
	0xb8, 0x52, 0x51, 0xa7, 0xca, 0x96, 0x0a, 0x42, 0x9e, 0x65, 0x6c, 0xc5, 0x0b, 0x05, 0x9b, 0x27,
	0x7a, 0x01, 0x8d, 0x38, 0xe1, 0xa9, 0x4a, 0xaf, 0x66, 0xaf, 0x41, 0x4b, 0x40, 0xde, 0x45, 0x4f,
	0x75, 0xa2, 0xbc, 0x15, 0xcd, 0xc7, 0xee, 0x62, 0xd9, 0x32, 0x0a, 0xde, 0xd6, 0x36, 0x7e, 0x07,
	0x28, 0x25, 0xa2, 0x67, 0x80, 0x28, 0x99, 0x4d, 0xa9, 0xf3, 0xd6, 0x9f, 0x70, 0x02, 0xc7, 0x1b,
	0xdc, 0x99, 0x5d, 0x4a, 0xdf, 0xdc, 0x72, 0xf0, 0xc8, 0x26, 0xba, 0x86, 0x7a, 0xf0, 0xa2, 0xe2,
	0x9a, 0xce, 0x08, 0xc5, 0x8e, 0x35, 0x9d, 0xb8, 0xe7, 0xd8, 0xb2, 0xe7, 0x94, 0xe8, 0x35, 0xf4,
	0x31, 0x3c, 0xdb, 0x44, 0xd8, 0xd3, 0x0b, 0x77, 0x3e, 0xc1, 0xd7, 0xd8, 0xb2, 0x15, 0xbb, 0xfe,
	0x0a, 0x03, 0x94, 0x07, 0x1b, 0x3d, 0x87, 0x8f, 0x64, 0x92, 0x6b, 0x42, 0xaf, 0x64, 0x8a, 0x5d,
	0x0d, 0x00, 0xfb, 0xd2, 0x39, 0x18, 0xea, 0x5a, 0x61, 0x0f, 0xcf, 0xf4, 0xda, 0x2b, 0xa6, 0x96,
	0xaf, 0x3c, 0xbe, 0x9f, 0xc0, 0x89, 0xf4, 0x59, 0x13, 0x87, 0xd0, 0x73, 0x6c, 0x92, 0xff, 0xa7,
	0x19, 0x5b, 0x94, 0x98, 0x4e, 0x7e, 0x74, 0xde, 0x10, 0x3a, 0x21, 0xb6, 0x7b, 0x89, 0x27, 0xf8,
	0x82, 0x8c, 0xf5, 0x9a, 0xc4, 0xc6, 0x98, 0x5c, 0x4e, 0x27, 0x5b, 0xac, 0xbe, 0xd8, 0x57, 0x2d,
	0xfc, 0xfa, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0xe2, 0xe4, 0xeb, 0xb5, 0x06, 0x00, 0x00,
}