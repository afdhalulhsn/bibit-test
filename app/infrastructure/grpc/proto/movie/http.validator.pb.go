// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: http.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Http) Validate() error {
	for _, item := range this.Rules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
			}
		}
	}
	return nil
}
func (this *HttpRule) Validate() error {
	if oneOfNester, ok := this.GetPattern().(*HttpRule_Custom); ok {
		if oneOfNester.Custom != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Custom); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Custom", err)
			}
		}
	}
	for _, item := range this.AdditionalBindings {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AdditionalBindings", err)
			}
		}
	}
	return nil
}
func (this *CustomHttpPattern) Validate() error {
	return nil
}