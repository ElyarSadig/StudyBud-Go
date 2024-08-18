package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	result := strings.Builder{}
	result.Grow(length)
	for i := 0; i < length; i++ {
		num := rand.Intn(10)
		result.WriteString(strconv.Itoa(num))
	}
	return result.String()
}

func (u *UserUseCase) setSession(ctx context.Context, sessionValue domain.SessionValue) (string, error) {
	key := u.generateDigitString(6)
	v, err := json.Marshal(sessionValue)
	if err != nil {
		u.logger.Error(err.Error())
		return "", u.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	err = u.redis.Set(ctx, "session", time.Minute*15, key, v)
	if err != nil {
		u.logger.Error(err.Error())
		return "", u.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	u.logger.Info("session set in redis", "key", key, "user", sessionValue)
	return key, nil
}
