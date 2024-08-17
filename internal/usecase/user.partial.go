package usecase

import (
	"errors"

	"math/rand"

	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/utils"
)

func (u *UserUseCase) validateUserRegisterForm(form *domain.UserRegisterForm) error {
	err := utils.ValidateName(form.Name)
	if err != nil {
		return err
	}
	err = utils.ValidateUsername(form.Username)
	if err != nil {
		return err
	}
	err = utils.ValidateEmail(form.Email)
	if err != nil {
		return err
	}
	if form.Password1 != form.Password2 {
		return errors.New("passwords do not match!")
	}
	err = utils.ValidatePassword(form.Password1)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) generateDigitString(length int) string {
	digits := make([]byte, length)
	var num byte
	for i := 0; i < length; i++ {
		num = byte(rand.Intn(10))
		digits[i] = num
	}
	return string(digits)
}
