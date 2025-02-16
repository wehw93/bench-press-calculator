package creator

type Request struct{
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Weight float32 `json:"weight"`
	Quantity float32 `json:"quantity"`
}

type Responce struct{
	MaxWeight float32 `json:"max weight: "`
	PercentBetter float32 `json:"better then average on: "`
}