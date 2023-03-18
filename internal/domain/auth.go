package domain

type AuthEmail struct {
	Email string `json:"email" binding:"required"`
}

type AuthCode struct {
	Email      string `json:"email" binding:"required"`
	SecretCode int32  `json:"secret_code" bson:"secret_code" binding:"required"`
}
