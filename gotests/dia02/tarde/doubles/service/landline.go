package service

import "github.com/bgw7/doubles/repository"

type LandlineService struct {
	repository repository.LandlineRepository
}

func NewLandlineService(repository repository.LandlineRepository) *LandlineService {
	return &LandlineService{repository}
}

func (l *LandlineService) GetVersion() string {
	return "1.0.0"
}

func (l *LandlineService) AddEntry(name string, landline string) {
	l.repository.AddEntry(name, landline)
}

// SearchNameByPhone implements LandlineRepository.
func (l *LandlineService) SearchNameByPhone(landline string) string {
	if len(landline) < 3 {
		return ""
	}

	return l.repository.SearchNameByPhone(landline)
}

// SearchPhoneByName implements LandlineRepository.
func (l *LandlineService) SearchPhoneByName(name string) string {
	return l.repository.SearchPhoneByName(name)
}
