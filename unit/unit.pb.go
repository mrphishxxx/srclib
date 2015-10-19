// Code generated by protoc-gen-gogo.
// source: unit.proto
// DO NOT EDIT!

/*
Package unit is a generated protocol buffer package.

It is generated from these files:
	unit.proto

It has these top-level messages:
	RepoSourceUnit
*/
package unit

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A RepoSourceUnit is the "concrete" form of SourceUnit that includes
// information about which repository (and commit) the source unit
// exists in. In general, type SourceUnit is used during analysis of a
// single source unit and type RepoSourceUnit is used afterwards
// (either in cross-source-unit analysis, such as cross-reference
// resolution, or in after-the-fact DB/API queries).
type RepoSourceUnit struct {
	Repo     string `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	CommitID string `protobuf:"bytes,2,opt,name=commit_id,proto3" json:"commit_id,omitempty"`
	UnitType string `protobuf:"bytes,3,opt,name=unit_type,proto3" json:"unit_type,omitempty"`
	Unit     string `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit,omitempty"`
	// Data is the JSON of the underlying SourceUnit.
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *RepoSourceUnit) Reset()         { *m = RepoSourceUnit{} }
func (m *RepoSourceUnit) String() string { return proto.CompactTextString(m) }
func (*RepoSourceUnit) ProtoMessage()    {}
