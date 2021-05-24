package services

import (
	"github.com/medivh13/mnc-test/pkg/dto"
)

type Services interface {
	Palindrome(req *dto.PalindromeReqDTO) (*dto.PalindromeRespDTO, error)
	GetLangaugeById(*dto.LanguageReqDTO) (*dto.LanguageRespDTO, error)
}