package dto

import (
	"errors"

	"github.com/medivh13/mnc-test/pkg/common/crypto"
	"github.com/medivh13/mnc-test/pkg/common/env"
	"github.com/medivh13/mnc-test/pkg/common/validator"
	util "github.com/medivh13/mnc-test/pkg/utils"
)

type PalindromeReqDTO struct {
	Text      string `json:"text"`
	Signature string `json:"signature" valid:"required" validname="signature"`
}

func (dto *PalindromeReqDTO) Validate() error {
	v := validator.NewValidate(dto)
	v.SetCustomValidation(true, func() error {
		return dto.customValidation()
	})
	return v.Validate()
}
func (dto *PalindromeReqDTO) customValidation() error {

	signature := crypto.EncodeSHA256HMAC(util.GetPrivKeySignature(), dto.Text)
	if signature != dto.Signature {
		if env.IsProduction() {
			return errors.New("invalid signature")
		}
		return errors.New("invalid signature" + " --> " + signature)
	}

	return nil
}

type PalindromeRespDTO struct {
	Result string `json:"result"`
}
