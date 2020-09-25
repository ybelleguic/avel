package wind

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetWind(key string) (string, error) {
	return key, nil
}

//Check returns an error if one of the service attrs fail to check themselves
func (s *Service) Check() error {
	return nil
}
