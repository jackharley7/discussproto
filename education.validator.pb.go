// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: education.proto

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

func (this *Education) Validate() error {
	if this.School == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("School", fmt.Errorf(`value '%v' must not be an empty string`, this.School))
	}
	if this.Concentration == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Concentration", fmt.Errorf(`value '%v' must not be an empty string`, this.Concentration))
	}
	if this.DegreeType == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DegreeType", fmt.Errorf(`value '%v' must not be an empty string`, this.DegreeType))
	}
	if !(this.GraduationYear > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("GraduationYear", fmt.Errorf(`value '%v' must be greater than '0'`, this.GraduationYear))
	}
	return nil
}
func (this *CreateEducationRequest) Validate() error {
	if this.Education != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Education); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Education", err)
		}
	}
	return nil
}
func (this *CreateEducationResponse) Validate() error {
	if this.Education != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Education); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Education", err)
		}
	}
	return nil
}
func (this *UpdateEducationRequest) Validate() error {
	if this.Education != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Education); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Education", err)
		}
	}
	return nil
}
func (this *UpdateEducationResponse) Validate() error {
	if this.Education != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Education); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Education", err)
		}
	}
	return nil
}
func (this *DeleteEducationRequest) Validate() error {
	return nil
}
func (this *DeleteEducationResponse) Validate() error {
	return nil
}
