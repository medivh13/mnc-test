package dto

import (
	"errors"

	"github.com/medivh13/mnc-test/pkg/common/crypto"
	"github.com/medivh13/mnc-test/pkg/common/env"
	"github.com/medivh13/mnc-test/pkg/common/validator"
	util "github.com/medivh13/mnc-test/pkg/utils"
)

type LanguageRespDTO struct {
	Language       string       `json:"language"`
	Appeared       string       `json:"appeared"`
	CreatedName    []string     `json:"created"`
	Functional     bool         `json:"functional"`
	ObjectOriented bool         `json:"object-oriented"`
	Relation       *RelationDTO `json:"relation"`
}

type RelationDTO struct {
	InfluenceBy []string `json:"influence-by"`
	Influences  []string `json:"influences"`
}

type LanguageReqDTO struct {
	Id        string `json:"id" valid:"required" validname="id"`
	Signature string `json:"signature" valid:"required" validname="signature"`
}

func (dto *LanguageReqDTO) Validate() error {
	v := validator.NewValidate(dto)
	v.SetCustomValidation(true, func() error {
		return dto.customValidation()
	})
	return v.Validate()
}
func (dto *LanguageReqDTO) customValidation() error {

	signature := crypto.EncodeSHA256HMAC(util.GetPrivKeySignature(), dto.Id)
	if signature != dto.Signature {
		if env.IsProduction() {
			return errors.New("invalid signature")
		}
		return errors.New("invalid signature" + " --> " + signature)
	}

	return nil
}
