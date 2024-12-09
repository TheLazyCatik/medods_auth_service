package signup

type SignupRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type SignupResponse struct {
	Result SignupResult `json:"result" validate:"required"`
}

type SignupResult struct {
	UserID string `json:"user_id" validate:"required"`
}
