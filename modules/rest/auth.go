package rest

import "time"

// ClientModel ...
type ClientModel struct {
	ClientPlatformID string `json:"client_platform_id"`
	ClientTypeID     string `json:"client_type_id"`
}

// StandardLoginModel ...
type StandardLoginModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// SessionModel ...
type SessionModel struct {
	ClientPlatformID string    `json:"client_platform_id"`
	ClientTypeID     string    `json:"client_type_id"`
	UserID           string    `json:"user_id"`
	ID               string    `json:"id"`
	RoleID           string    `json:"role_id"`
	IP               string    `json:"ip"`
	Data             string    `json:"data"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	ExpiresAt        time.Time `json:"expires_at"`
}

// TokenModel ...
type TokenModel struct {
	AccessToken      string    `json:"access_token"`
	RefreshToken     string    `json:"refresh_token"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	ExpiresAt        time.Time `json:"expires_at"`
	RefreshInSeconds int       `json:"refresh_in_seconds"`
}

// ClientTypeModel ...
type ClientTypeModel struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	RegisterStrategy string `json:"register_strategy"`
	LoginStrategy    string `json:"login_strategy"`
	RecoveryStrategy string `json:"recovery_strategy"`
}

// AccessModel ...
type AccessModel struct {
	Token  string `json:"token"`
	Method string `json:"method"`
	URL    string `json:"url"`
}

// HasAccessModel ...
type HasAccessModel struct {
	ClientPlatformID string `json:"client_platform_id"`
	ClientTypeID     string `json:"client_type_id"`
	UserID           string `json:"user_id"`
	ID               string `json:"id"`
	RoleID           string `json:"role_id"`
}

// RefreshTokenModel ...
type RefreshTokenModel struct {
	Token string `json:"token" validate:"required"`
}
