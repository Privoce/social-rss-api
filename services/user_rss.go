package services

import (
	"api.privoce.com/rss/models"
	"api.privoce.com/rss/repositorys"
)

// SocialService interface
type SocialService struct{}

func (s *SocialService) GetUserLookIns(name string) []*models.LookIn {
	if len(name) == 0 {
		return nil
	}
	return repositorys.FindLookInsByUserName(name)
}
