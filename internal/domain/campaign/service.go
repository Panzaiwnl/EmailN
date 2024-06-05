package campaign

import (
	campaing "emailn/internal/contract"
	"emailn/internal/internalErrors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign campaing.NewCampaignDTO) (string, error) {

	campaing, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaing)
	if err != nil {
		return "", internalErrors.ErrInternal
	}

	return campaing.ID, nil
}
