package usecase

// Service is the usecase service
type Service struct {
	repo Repository
}

// NewService creates a new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}
