// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user.proto

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

func (this *User) Validate() error {
	if this.Profile != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Profile); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Profile", err)
		}
	}
	for _, item := range this.Education {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Education", err)
			}
		}
	}
	for _, item := range this.WorkExperience {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("WorkExperience", err)
			}
		}
	}
	return nil
}
func (this *LoginRequest) Validate() error {
	if this.Username == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Username", fmt.Errorf(`value '%v' must not be an empty string`, this.Username))
	}
	if this.Password == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must not be an empty string`, this.Password))
	}
	return nil
}
func (this *LoginResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *RefreshUserTokenRequest) Validate() error {
	return nil
}
func (this *RefreshUserTokenResponse) Validate() error {
	return nil
}
func (this *CheckForUsernameRequest) Validate() error {
	if this.Username == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Username", fmt.Errorf(`value '%v' must not be an empty string`, this.Username))
	}
	return nil
}
func (this *CheckForUsernameResponse) Validate() error {
	return nil
}
func (this *CreateUserRequest) Validate() error {
	if this.Username == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Username", fmt.Errorf(`value '%v' must not be an empty string`, this.Username))
	}
	if this.Password == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must not be an empty string`, this.Password))
	}
	if this.RepeatedPassword == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("RepeatedPassword", fmt.Errorf(`value '%v' must not be an empty string`, this.RepeatedPassword))
	}
	return nil
}
func (this *CreateUserResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *ChangePasswordRequest) Validate() error {
	if this.OldPassword == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OldPassword", fmt.Errorf(`value '%v' must not be an empty string`, this.OldPassword))
	}
	if this.NewPassword == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NewPassword", fmt.Errorf(`value '%v' must not be an empty string`, this.NewPassword))
	}
	if this.NewPasswordRepeat == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NewPasswordRepeat", fmt.Errorf(`value '%v' must not be an empty string`, this.NewPasswordRepeat))
	}
	return nil
}
func (this *ChangePasswordResponse) Validate() error {
	return nil
}
func (this *ChangePasswordFromTempRequest) Validate() error {
	if this.TempPassword == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TempPassword", fmt.Errorf(`value '%v' must not be an empty string`, this.TempPassword))
	}
	if this.NewPassword == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NewPassword", fmt.Errorf(`value '%v' must not be an empty string`, this.NewPassword))
	}
	if this.NewPasswordRepeat == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NewPasswordRepeat", fmt.Errorf(`value '%v' must not be an empty string`, this.NewPasswordRepeat))
	}
	return nil
}
func (this *ChangePasswordFromTempResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *RequestPasswordResetRequest) Validate() error {
	if this.Email == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must not be an empty string`, this.Email))
	}
	return nil
}
func (this *RequestPasswordResetResponse) Validate() error {
	return nil
}
func (this *GetUsersRequest) Validate() error {
	return nil
}
func (this *GetUsersResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetByIdsRequest) Validate() error {
	return nil
}
func (this *GetByIdsResponse) Validate() error {
	for _, item := range this.Users {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
			}
		}
	}
	return nil
}
func (this *GetUserByIdRequest) Validate() error {
	return nil
}
func (this *GetUserByIdResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *GetUserMeRequest) Validate() error {
	return nil
}
func (this *GetUserMeResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *UpdateUserRequest) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *UpdateUserResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *DeleteUserRequest) Validate() error {
	return nil
}
func (this *DeleteUserResponse) Validate() error {
	return nil
}
func (this *GetUserByTwitterScreenNameRequest) Validate() error {
	return nil
}
func (this *GetUserByTwitterScreenNameResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *LinkTwitterRequest) Validate() error {
	return nil
}
func (this *LinkTwitterResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *UnLinkTwitterRequest) Validate() error {
	return nil
}
func (this *UnLinkTwitterResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *GetTwitterUserByScreenNameRequest) Validate() error {
	return nil
}
func (this *GetTwitterUserByScreenNameResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *GetProfileImageUploadURLRequest) Validate() error {
	return nil
}
func (this *GetProfileImageUploadURLResponse) Validate() error {
	return nil
}
