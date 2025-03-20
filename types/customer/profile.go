package types

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

/*
 * Types for handling fetch requests on profile requests
 */

type GetMyUserProfileResponse struct {
	Message           string `json:"message"`
	Username          string `json:"username"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	ProfilePictureUrl string `json:"profile_picture_url"`
}

/*
 * Types for handling CSRF request as well as handling the editing of the
 * user profile's FirstName, LastName and PhoneNumber
 */

type EditMyProfileCsrfResponse struct {
	Message  string `json:"message"`
	Token    string `json:"token"`
	ExpiryAt string `json:"expiry_at"`
}

type EditMyProfileRequest struct {
	Message     string `json:"message"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}

func (r *EditMyProfileRequest) Validate() (err error) {
	err = v.ValidateStruct(&r,
		v.Field(&r.FirstName, v.Required, v.Length(1, 50)),
		v.Field(&r.LastName, v.Required, v.Length(1, 50)),
		v.Field(&r.PhoneNumber, v.Required, is.E164),
	)
	return
}

type EditMyProfileResponse struct {
	Message     string `json:"message"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}

/*
 * Types for requesting preSignedUrl to upload an image for profile picture
 * to an s3-compatible file-storage. Currently, it is cloudflare-r2.
 */

type AddPfpRequest struct {
	FileName string `json:"file_name"`
	FileSize string `json:"file_size"`
	FileType string `json:"file_type"`
}

func (r AddPfpRequest) Validate() (err error) {
	err = v.ValidateStruct(&r,
		v.Field(&r.FileName,
			v.Required,
			v.Length(1, 255)),
		v.Field(&r.FileSize,
			v.Required,
			is.Int),
		v.Field(&r.FileType,
			v.Required,
			v.In("image/jpeg", "image/png", "image/jpg")),
	)
	return
}

type AddPfpResponse struct {
	Message      string `json:"message"`
	PreSignedUrl string `json:"pre_signed_url"`
	FileName     string `json:"file_name"`
	FileType     string `json:"file_type"`
	ExpiryAt     string `json:"expiry_at"`
}

/*
 * Types for deleting a user's image for profile picture
 */

type DeleteUserPfpResponse struct {
	Message string `json:"message"`
}
