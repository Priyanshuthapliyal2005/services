// Code generated by protoc-gen-go-json. DO NOT EDIT.
// source: state.proto

package proto

import (
	"google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON implements json.Marshaler
func (msg *ServiceState) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ServiceState) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}
