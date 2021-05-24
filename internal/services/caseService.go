package services

import (
	"strconv"

	"github.com/medivh13/mnc-test/pkg/dto"
)

var Language = []dto.LanguageRespDTO{
	dto.LanguageRespDTO{
		Language:       "C",
		Appeared:       "1972",
		CreatedName:    []string{"Dennis Ritchie"},
		Functional:     true,
		ObjectOriented: false,
		Relation: &dto.RelationDTO{
			InfluenceBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
			Influences:  []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"},
		},
	},
}

type service struct {
}

func NewService() Services {
	return &service{}
}

func (s *service) Palindrome(req *dto.PalindromeReqDTO) (*dto.PalindromeRespDTO, error) {
	response := dto.PalindromeRespDTO{}
	response.Result = "Palindrome"
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(req.Text)/2; i++ {
		if req.Text[i] != req.Text[len(req.Text)-i-1] {
			response.Result = "Not Palindrome"
			return &response, nil
		}
	}

	return &response, nil

}

func (s *service) GetLangaugeById(req *dto.LanguageReqDTO) (*dto.LanguageRespDTO, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, _ := strconv.Atoi(req.Id)
	return &Language[id], nil

}
