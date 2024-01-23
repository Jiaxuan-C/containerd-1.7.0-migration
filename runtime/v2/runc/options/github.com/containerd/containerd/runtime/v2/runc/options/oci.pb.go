// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.3
// source: oci.proto

package options

import (
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

type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// disable pivot root when creating a container
	NoPivotRoot bool `protobuf:"varint,1,opt,name=no_pivot_root,json=noPivotRoot,proto3" json:"no_pivot_root,omitempty"`
	// create a new keyring for the container
	NoNewKeyring bool `protobuf:"varint,2,opt,name=no_new_keyring,json=noNewKeyring,proto3" json:"no_new_keyring,omitempty"`
	// place the shim in a cgroup
	ShimCgroup string `protobuf:"bytes,3,opt,name=shim_cgroup,json=shimCgroup,proto3" json:"shim_cgroup,omitempty"`
	// set the I/O's pipes uid
	IoUid uint32 `protobuf:"varint,4,opt,name=io_uid,json=ioUid,proto3" json:"io_uid,omitempty"`
	// set the I/O's pipes gid
	IoGid uint32 `protobuf:"varint,5,opt,name=io_gid,json=ioGid,proto3" json:"io_gid,omitempty"`
	// binary name of the runc binary
	BinaryName string `protobuf:"bytes,6,opt,name=binary_name,json=binaryName,proto3" json:"binary_name,omitempty"`
	// runc root directory
	Root string `protobuf:"bytes,7,opt,name=root,proto3" json:"root,omitempty"`
	// criu binary path.
	//
	// Deprecated: runc option --criu is now ignored (with a warning), and the
	// option will be removed entirely in a future release. Users who need a non-
	// standard criu binary should rely on the standard way of looking up binaries
	// in $PATH.
	//
	// Deprecated: Do not use.
	CriuPath string `protobuf:"bytes,8,opt,name=criu_path,json=criuPath,proto3" json:"criu_path,omitempty"`
	// enable systemd cgroups
	SystemdCgroup bool `protobuf:"varint,9,opt,name=systemd_cgroup,json=systemdCgroup,proto3" json:"systemd_cgroup,omitempty"`
	// criu image path
	CriuImagePath string `protobuf:"bytes,10,opt,name=criu_image_path,json=criuImagePath,proto3" json:"criu_image_path,omitempty"`
	// criu work path
	CriuWorkPath string `protobuf:"bytes,11,opt,name=criu_work_path,json=criuWorkPath,proto3" json:"criu_work_path,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oci_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_oci_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_oci_proto_rawDescGZIP(), []int{0}
}

func (x *Options) GetNoPivotRoot() bool {
	if x != nil {
		return x.NoPivotRoot
	}
	return false
}

func (x *Options) GetNoNewKeyring() bool {
	if x != nil {
		return x.NoNewKeyring
	}
	return false
}

func (x *Options) GetShimCgroup() string {
	if x != nil {
		return x.ShimCgroup
	}
	return ""
}

func (x *Options) GetIoUid() uint32 {
	if x != nil {
		return x.IoUid
	}
	return 0
}

func (x *Options) GetIoGid() uint32 {
	if x != nil {
		return x.IoGid
	}
	return 0
}

func (x *Options) GetBinaryName() string {
	if x != nil {
		return x.BinaryName
	}
	return ""
}

func (x *Options) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

// Deprecated: Do not use.
func (x *Options) GetCriuPath() string {
	if x != nil {
		return x.CriuPath
	}
	return ""
}

func (x *Options) GetSystemdCgroup() bool {
	if x != nil {
		return x.SystemdCgroup
	}
	return false
}

func (x *Options) GetCriuImagePath() string {
	if x != nil {
		return x.CriuImagePath
	}
	return ""
}

func (x *Options) GetCriuWorkPath() string {
	if x != nil {
		return x.CriuWorkPath
	}
	return ""
}

type CheckpointOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// exit the container after a checkpoint
	Exit bool `protobuf:"varint,1,opt,name=exit,proto3" json:"exit,omitempty"`
	// checkpoint open tcp connections
	OpenTcp bool `protobuf:"varint,2,opt,name=open_tcp,json=openTcp,proto3" json:"open_tcp,omitempty"`
	// checkpoint external unix sockets
	ExternalUnixSockets bool `protobuf:"varint,3,opt,name=external_unix_sockets,json=externalUnixSockets,proto3" json:"external_unix_sockets,omitempty"`
	// checkpoint terminals (ptys)
	Terminal bool `protobuf:"varint,4,opt,name=terminal,proto3" json:"terminal,omitempty"`
	// allow checkpointing of file locks
	FileLocks bool `protobuf:"varint,5,opt,name=file_locks,json=fileLocks,proto3" json:"file_locks,omitempty"`
	// restore provided namespaces as empty namespaces
	EmptyNamespaces []string `protobuf:"bytes,6,rep,name=empty_namespaces,json=emptyNamespaces,proto3" json:"empty_namespaces,omitempty"`
	// set the cgroups mode, soft, full, strict
	CgroupsMode string `protobuf:"bytes,7,opt,name=cgroups_mode,json=cgroupsMode,proto3" json:"cgroups_mode,omitempty"`
	// checkpoint image path
	ImagePath string `protobuf:"bytes,8,opt,name=image_path,json=imagePath,proto3" json:"image_path,omitempty"`
	// checkpoint work path
	WorkPath string `protobuf:"bytes,9,opt,name=work_path,json=workPath,proto3" json:"work_path,omitempty"`
	// path for pre-dump TODO(cjx.bupt)
	ParentPath string `protobuf:"bytes,10,opt,name=parent_path,json=parentPath,proto3" json:"parent_path,omitempty"`
}

func (x *CheckpointOptions) Reset() {
	*x = CheckpointOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oci_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckpointOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckpointOptions) ProtoMessage() {}

func (x *CheckpointOptions) ProtoReflect() protoreflect.Message {
	mi := &file_oci_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckpointOptions.ProtoReflect.Descriptor instead.
func (*CheckpointOptions) Descriptor() ([]byte, []int) {
	return file_oci_proto_rawDescGZIP(), []int{1}
}

func (x *CheckpointOptions) GetExit() bool {
	if x != nil {
		return x.Exit
	}
	return false
}

func (x *CheckpointOptions) GetOpenTcp() bool {
	if x != nil {
		return x.OpenTcp
	}
	return false
}

func (x *CheckpointOptions) GetExternalUnixSockets() bool {
	if x != nil {
		return x.ExternalUnixSockets
	}
	return false
}

func (x *CheckpointOptions) GetTerminal() bool {
	if x != nil {
		return x.Terminal
	}
	return false
}

func (x *CheckpointOptions) GetFileLocks() bool {
	if x != nil {
		return x.FileLocks
	}
	return false
}

func (x *CheckpointOptions) GetEmptyNamespaces() []string {
	if x != nil {
		return x.EmptyNamespaces
	}
	return nil
}

func (x *CheckpointOptions) GetCgroupsMode() string {
	if x != nil {
		return x.CgroupsMode
	}
	return ""
}

func (x *CheckpointOptions) GetImagePath() string {
	if x != nil {
		return x.ImagePath
	}
	return ""
}

func (x *CheckpointOptions) GetWorkPath() string {
	if x != nil {
		return x.WorkPath
	}
	return ""
}

func (x *CheckpointOptions) GetParentPath() string {
	if x != nil {
		return x.ParentPath
	}
	return ""
}

type ProcessDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// exec process id if the process is managed by a shim
	ExecId string `protobuf:"bytes,1,opt,name=exec_id,json=execId,proto3" json:"exec_id,omitempty"`
}

func (x *ProcessDetails) Reset() {
	*x = ProcessDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oci_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessDetails) ProtoMessage() {}

func (x *ProcessDetails) ProtoReflect() protoreflect.Message {
	mi := &file_oci_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessDetails.ProtoReflect.Descriptor instead.
func (*ProcessDetails) Descriptor() ([]byte, []int) {
	return file_oci_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessDetails) GetExecId() string {
	if x != nil {
		return x.ExecId
	}
	return ""
}

var File_oci_proto protoreflect.FileDescriptor

var file_oci_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6f, 0x63, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x22,
	0xed, 0x02, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6e,
	0x6f, 0x5f, 0x70, 0x69, 0x76, 0x6f, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x6e, 0x6f, 0x50, 0x69, 0x76, 0x6f, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12,
	0x24, 0x0a, 0x0e, 0x6e, 0x6f, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6e, 0x6f, 0x4e, 0x65, 0x77, 0x4b, 0x65,
	0x79, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x69, 0x6d, 0x5f, 0x63, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x69, 0x6d,
	0x43, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x6f, 0x5f, 0x75, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6f, 0x55, 0x69, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x69, 0x6f, 0x5f, 0x67, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69,
	0x6f, 0x47, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x1f, 0x0a, 0x09, 0x63, 0x72, 0x69,
	0x75, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x08, 0x63, 0x72, 0x69, 0x75, 0x50, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x64, 0x5f, 0x63, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x64, 0x43, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72, 0x69, 0x75, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72, 0x69, 0x75,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x72, 0x69,
	0x75, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x72, 0x69, 0x75, 0x57, 0x6f, 0x72, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x22,
	0xdc, 0x02, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x65, 0x78, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x70, 0x65,
	0x6e, 0x5f, 0x74, 0x63, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6f, 0x70, 0x65,
	0x6e, 0x54, 0x63, 0x70, 0x12, 0x32, 0x0a, 0x15, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x55, 0x6e, 0x69,
	0x78, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x4c, 0x6f,
	0x63, 0x6b, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x29,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x78, 0x65, 0x63, 0x49, 0x64, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x75, 0x6e, 0x63, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oci_proto_rawDescOnce sync.Once
	file_oci_proto_rawDescData = file_oci_proto_rawDesc
)

func file_oci_proto_rawDescGZIP() []byte {
	file_oci_proto_rawDescOnce.Do(func() {
		file_oci_proto_rawDescData = protoimpl.X.CompressGZIP(file_oci_proto_rawDescData)
	})
	return file_oci_proto_rawDescData
}

var file_oci_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_oci_proto_goTypes = []interface{}{
	(*Options)(nil),           // 0: containerd.runc.v1.Options
	(*CheckpointOptions)(nil), // 1: containerd.runc.v1.CheckpointOptions
	(*ProcessDetails)(nil),    // 2: containerd.runc.v1.ProcessDetails
}
var file_oci_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oci_proto_init() }
func file_oci_proto_init() {
	if File_oci_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oci_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
		file_oci_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckpointOptions); i {
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
		file_oci_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessDetails); i {
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
			RawDescriptor: file_oci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oci_proto_goTypes,
		DependencyIndexes: file_oci_proto_depIdxs,
		MessageInfos:      file_oci_proto_msgTypes,
	}.Build()
	File_oci_proto = out.File
	file_oci_proto_rawDesc = nil
	file_oci_proto_goTypes = nil
	file_oci_proto_depIdxs = nil
}
