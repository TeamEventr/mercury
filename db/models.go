// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type EnumAccountStatus string

const (
	EnumAccountStatusActive   EnumAccountStatus = "active"
	EnumAccountStatusDisabled EnumAccountStatus = "disabled"
	EnumAccountStatusBanned   EnumAccountStatus = "banned"
	EnumAccountStatusDeleted  EnumAccountStatus = "deleted"
)

func (e *EnumAccountStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnumAccountStatus(s)
	case string:
		*e = EnumAccountStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for EnumAccountStatus: %T", src)
	}
	return nil
}

type NullEnumAccountStatus struct {
	EnumAccountStatus EnumAccountStatus
	Valid             bool // Valid is true if EnumAccountStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEnumAccountStatus) Scan(value interface{}) error {
	if value == nil {
		ns.EnumAccountStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.EnumAccountStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEnumAccountStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.EnumAccountStatus), nil
}

type EnumBookingStatus string

const (
	EnumBookingStatusOpen   EnumBookingStatus = "open"
	EnumBookingStatusClosed EnumBookingStatus = "closed"
)

func (e *EnumBookingStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnumBookingStatus(s)
	case string:
		*e = EnumBookingStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for EnumBookingStatus: %T", src)
	}
	return nil
}

type NullEnumBookingStatus struct {
	EnumBookingStatus EnumBookingStatus
	Valid             bool // Valid is true if EnumBookingStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEnumBookingStatus) Scan(value interface{}) error {
	if value == nil {
		ns.EnumBookingStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.EnumBookingStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEnumBookingStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.EnumBookingStatus), nil
}

type EnumEventSchedule string

const (
	EnumEventScheduleOntime    EnumEventSchedule = "ontime"
	EnumEventSchedulePrepond   EnumEventSchedule = "prepond"
	EnumEventSchedulePostponed EnumEventSchedule = "postponed"
	EnumEventScheduleCancelled EnumEventSchedule = "cancelled"
)

func (e *EnumEventSchedule) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnumEventSchedule(s)
	case string:
		*e = EnumEventSchedule(s)
	default:
		return fmt.Errorf("unsupported scan type for EnumEventSchedule: %T", src)
	}
	return nil
}

type NullEnumEventSchedule struct {
	EnumEventSchedule EnumEventSchedule
	Valid             bool // Valid is true if EnumEventSchedule is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEnumEventSchedule) Scan(value interface{}) error {
	if value == nil {
		ns.EnumEventSchedule, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.EnumEventSchedule.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEnumEventSchedule) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.EnumEventSchedule), nil
}

type EnumEventVisibility string

const (
	EnumEventVisibilityDraft     EnumEventVisibility = "draft"
	EnumEventVisibilityPublished EnumEventVisibility = "published"
)

func (e *EnumEventVisibility) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnumEventVisibility(s)
	case string:
		*e = EnumEventVisibility(s)
	default:
		return fmt.Errorf("unsupported scan type for EnumEventVisibility: %T", src)
	}
	return nil
}

type NullEnumEventVisibility struct {
	EnumEventVisibility EnumEventVisibility
	Valid               bool // Valid is true if EnumEventVisibility is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEnumEventVisibility) Scan(value interface{}) error {
	if value == nil {
		ns.EnumEventVisibility, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.EnumEventVisibility.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEnumEventVisibility) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.EnumEventVisibility), nil
}

type EnumGenderOptions string

const (
	EnumGenderOptionsMale   EnumGenderOptions = "male"
	EnumGenderOptionsFemale EnumGenderOptions = "female"
	EnumGenderOptionsOthers EnumGenderOptions = "others"
)

func (e *EnumGenderOptions) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnumGenderOptions(s)
	case string:
		*e = EnumGenderOptions(s)
	default:
		return fmt.Errorf("unsupported scan type for EnumGenderOptions: %T", src)
	}
	return nil
}

type NullEnumGenderOptions struct {
	EnumGenderOptions EnumGenderOptions
	Valid             bool // Valid is true if EnumGenderOptions is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEnumGenderOptions) Scan(value interface{}) error {
	if value == nil {
		ns.EnumGenderOptions, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.EnumGenderOptions.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEnumGenderOptions) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.EnumGenderOptions), nil
}

type EnumHostedStatus string

const (
	EnumHostedStatusUnderFive    EnumHostedStatus = "under_five"
	EnumHostedStatusUnderTwenty  EnumHostedStatus = "under_twenty"
	EnumHostedStatusUnderFifty   EnumHostedStatus = "under_fifty"
	EnumHostedStatusUnderHundred EnumHostedStatus = "under_hundred"
	EnumHostedStatusMore         EnumHostedStatus = "more"
)

func (e *EnumHostedStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnumHostedStatus(s)
	case string:
		*e = EnumHostedStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for EnumHostedStatus: %T", src)
	}
	return nil
}

type NullEnumHostedStatus struct {
	EnumHostedStatus EnumHostedStatus
	Valid            bool // Valid is true if EnumHostedStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEnumHostedStatus) Scan(value interface{}) error {
	if value == nil {
		ns.EnumHostedStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.EnumHostedStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEnumHostedStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.EnumHostedStatus), nil
}

type Bookmark struct {
	ID       int32
	Username string
	EventID  uuid.UUID
}

type Event struct {
	ID              uuid.UUID
	Title           string
	HostID          uuid.UUID
	Blurb           pgtype.Text
	Description     pgtype.Text
	CoverPictureUrl pgtype.Text
	BannerUrl       pgtype.Text
	ThumbnailUrl    pgtype.Text
	Visibility      EnumEventVisibility
	Tags            []string
	Venue           pgtype.Text
	Schedule        EnumEventSchedule
	StartTime       interface{}
	EndTime         interface{}
	AgeLimit        pgtype.Int4
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       interface{}
}

type EventImage struct {
	ID        int32
	EventID   uuid.UUID
	ImageType string
	Url       string
}

type Host struct {
	ID                 uuid.UUID
	Username           string
	PanNumber          pgtype.Text
	AccountNumber      pgtype.Text
	FirstName          string
	MiddleName         pgtype.Text
	LastName           string
	PhoneNumber        string
	Dob                string
	CompanyName        string
	CompanyEmail       string
	BackupEmail        string
	Registered         pgtype.Bool
	RegistrationNumber pgtype.Text
	Address            pgtype.Text
	Pincode            pgtype.Int4
	EventCount         pgtype.Int4
	HostedStatus       EnumHostedStatus
	AccountStatus      EnumAccountStatus
	CreatedAt          interface{}
	UpdatedAt          interface{}
}

type HostOnboarding struct {
	ID                 int32
	Username           string
	FirstName          string
	MiddleName         pgtype.Text
	LastName           string
	PhoneNumber        string
	Dob                string
	CompanyName        string
	CompanyEmail       string
	BackupEmail        pgtype.Text
	Registered         bool
	RegistrationNumber pgtype.Text
	HostedStatus       NullEnumHostedStatus
	Otp                string
	CreatedAt          interface{}
	ExpiryAt           time.Time
}

type PriceTier struct {
	ID               int32
	EventID          uuid.UUID
	Name             string
	ValidityStart    time.Time
	ValidityEnd      time.Time
	Price            int32
	SeatAvailable    int32
	TotalSeat        int32
	BookingOpenTime  time.Time
	BookingCloseTime time.Time
	BookingStatus    NullEnumBookingStatus
}

type Ticket struct {
	ID          uuid.UUID
	EventID     uuid.UUID
	Username    string
	TierID      int32
	FirstName   string
	MiddleName  pgtype.Text
	LastName    string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Cancelled   pgtype.Bool
	CheckIn     bool
}

type UserAccount struct {
	ID            uuid.UUID
	Username      string
	PasswordLogin bool
	ProviderID    pgtype.Text
	Password      pgtype.Text
	FirstName     pgtype.Text
	MiddleName    pgtype.Text
	LastName      pgtype.Text
	Gender        NullEnumGenderOptions
	Email         string
	Avatar        pgtype.Text
	City          pgtype.Text
	Status        EnumAccountStatus
	LoggedinAt    interface{}
	RefreshToken  pgtype.Text
	CreatedAt     interface{}
	UpdatedAt     interface{}
	DeletedAt     interface{}
}

type UserOnboarding struct {
	ID        int32
	Username  string
	Password  string
	Email     string
	Otp       string
	CreatedAt interface{}
	ExpiryAt  time.Time
}

type VerificationAction struct {
	ID       int32
	Username string
	Purpose  string
	Otp      string
	ExpiryAt time.Time
}

type Volunteer struct {
	ID         string
	Username   string
	EventID    uuid.UUID
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ShiftStart time.Time
	ShiftEnd   time.Time
	Removed    pgtype.Bool
}
