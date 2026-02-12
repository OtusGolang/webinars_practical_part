package main

// UserService is the business logic layer
type UserService struct {
	repo    *UserRepository
	printer *PrinterService
	logger  Logger
}

// NewUserService creates a new UserService with all dependencies
func NewUserService(repo *UserRepository, printer *PrinterService, logger Logger) *UserService {
	logger.Info("Initializing UserService")
	return &UserService{
		repo:    repo,
		printer: printer,
		logger:  logger,
	}
}

// GetUserInfo retrieves user info and prints it
func (s *UserService) GetUserInfo(userID string) {
	s.logger.Info("Getting user info for ID: " + userID)
	userInfo := s.repo.GetUser(userID)
	s.printer.Print(userInfo)
}
