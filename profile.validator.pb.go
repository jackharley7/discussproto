// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: profile.proto

package discussproto

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Profile) Validate() error {
	return nil
}
func (this *UpdateProfileRequest) Validate() error {
	if this.Profile != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Profile); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Profile", err)
		}
	}
	return nil
}
func (this *UpdateProfileResponse) Validate() error {
	if this.Profile != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Profile); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Profile", err)
		}
	}
	return nil
}
func (this *GetProfileProgressRequest) Validate() error {
	return nil
}
func (this *GetProfileProgressResponse) Validate() error {
	return nil
}
