package campaign

import (
	campaing "emailn/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign campaing.NewCampaignDTO) (string, error) {

	campaing, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	return campaing.ID, nil
}
