package apply

type (
	LoginOneApply struct {
		Account  string `json:"account" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)
