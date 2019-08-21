// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cmm.proto

package discussproto

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CMMUser) Validate() error {
	return nil
}
func (this *CMMLink) Validate() error {
	return nil
}
func (this *CMM) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	if this.Link != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Link); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Link", err)
		}
	}
	return nil
}
func (this *CreateCMMRequest) Validate() error {
	if this.Cmm != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Cmm); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Cmm", err)
		}
	}
	return nil
}
func (this *CreateCMMResponse) Validate() error {
	if this.Cmm != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Cmm); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Cmm", err)
		}
	}
	return nil
}
func (this *UpdateCMMRequest) Validate() error {
	if this.Cmm != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Cmm); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Cmm", err)
		}
	}
	return nil
}
func (this *UpdateCMMResponse) Validate() error {
	if this.Cmm != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Cmm); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Cmm", err)
		}
	}
	return nil
}
func (this *DeleteCMMRequest) Validate() error {
	return nil
}
func (this *DeleteCMMResponse) Validate() error {
	return nil
}
func (this *GetCMMByIDRequest) Validate() error {
	return nil
}
func (this *GetCMMByIDResponse) Validate() error {
	if this.Cmm != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Cmm); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Cmm", err)
		}
	}
	return nil
}
func (this *SearchCMMsRequest) Validate() error {
	return nil
}
func (this *SearchCMMsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetCMMsRequest) Validate() error {
	return nil
}
func (this *GetCMMsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetCMMsByMeRequest) Validate() error {
	return nil
}
func (this *GetCMMsByMeResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetCMMsByUserIDRequest) Validate() error {
	return nil
}
func (this *GetCMMsByUserIDResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
