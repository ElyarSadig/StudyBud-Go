package utils

import (
	"testing"
	"time"
)

func TestValidateName(t *testing.T) {
	testCases := []struct {
		name                 string
		expectedErrorMessage string
		desc                 string
	}{
		{
			name:                 "e",
			expectedErrorMessage: "name must be at least 2 characters long",
			desc:                 "Short Name",
		},
		{
			name:                 "Elyar",
			expectedErrorMessage: "",
			desc:                 "Valid Name",
		},
		{
			name:                 "الیار",
			expectedErrorMessage: "name can only contain letters and spaces",
			desc:                 "Invalid Charachters",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := ValidateName(tC.name)
			if err != nil && err.Error() != tC.expectedErrorMessage {
				t.Errorf("expected error to be %s, but got %s", tC.expectedErrorMessage, err.Error())
			}
		})
	}
}

func TestValidateUserName(t *testing.T) {
	testCases := []struct {
		username             string
		expectedErrorMessage string
		desc                 string
	}{
		{
			username:             "eL",
			expectedErrorMessage: "username must be at least 3 characters long",
			desc:                 "Short Username",
		},
		{
			username:             "Elyar",
			expectedErrorMessage: "",
			desc:                 "Valid Username",
		},
		{
			username:             "الیار",
			expectedErrorMessage: "username can only contain letters, digits, underscores, and dots, and must start with a letter",
			desc:                 "Invalid Charachters",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := ValidateUsername(tC.username)
			if err != nil && err.Error() != tC.expectedErrorMessage {
				t.Errorf("expected error to be %s, but got %s", tC.expectedErrorMessage, err.Error())
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	testCases := []struct {
		password             string
		expectedErrorMessage string
		desc                 string
	}{
		{
			password:             "e",
			expectedErrorMessage: "password must be at least 8 characters long",
			desc:                 "Short password",
		},
		{
			password:             "Elyar",
			expectedErrorMessage: "password must be at least 8 characters long",
			desc:                 "Valid Name but Short password",
		},
		{
			password:             "Password",
			expectedErrorMessage: "password must contain at least one digit",
			desc:                 "Missing digit",
		},
		{
			password:             "Password1",
			expectedErrorMessage: "password must contain at least one special character",
			desc:                 "Missing special character",
		},
		{
			password: "password1!",

			expectedErrorMessage: "password must contain at least one uppercase letter",
			desc:                 "Missing uppercase letter",
		},
		{
			password: "PASSWORD1!",

			expectedErrorMessage: "password must contain at least one lowercase letter",
			desc:                 "Missing lowercase letter",
		},
		{
			password: "P@ssw0rd",

			expectedErrorMessage: "",
			desc:                 "Valid password",
		},
		{
			password:             "P@ssword123!",
			expectedErrorMessage: "",
			desc:                 "Valid password with more characters",
		},
		{
			password:             "P@ssw0",
			expectedErrorMessage: "password must be at least 8 characters long",
			desc:                 "Valid components but too short",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := ValidatePassword(tC.password)
			if err != nil && err.Error() != tC.expectedErrorMessage {
				t.Errorf("expected error to be %s, but got %s", tC.expectedErrorMessage, err.Error())
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		email                string
		expectedErrorMessage string
		desc                 string
	}{
		{
			email:                "eL@cg",
			expectedErrorMessage: "invalid email format",
			desc:                 "invalid email",
		},
		{
			email:                "Elyar@",
			expectedErrorMessage: "invalid email format",
			desc:                 "invalid email",
		},
		{
			email:                "elyar@email.com",
			expectedErrorMessage: "",
			desc:                 "valid email",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := ValidateEmail(tC.email)
			if err != nil && err.Error() != tC.expectedErrorMessage {
				t.Errorf("expected error to be %s, but got %s", tC.expectedErrorMessage, err.Error())
			}
		})
	}
}

func TestFormatDuration(t *testing.T) {
	testCases := []struct {
		duration time.Duration
		expected string
		desc     string
	}{
		{
			duration: time.Second * 15,
			expected: "15 seconds",
			desc:     "15 seconds",
		},
		{
			duration: time.Minute * 1,
			expected: "1 minute",
			desc:     "1 minute",
		},
		{
			duration: time.Minute * 34,
			expected: "34 minutes",
			desc:     "34 minutes",
		},
		{
			duration: time.Hour * 1,
			expected: "1 hour",
			desc:     "1 hour",
		},
		{
			duration: time.Hour * 2,
			expected: "2 hours",
			desc:     "2 hours",
		},
		{
			duration: time.Hour * 24,
			expected: "1 day",
			desc:     "1 day",
		},
		{
			duration: time.Hour * 24 * 5,
			expected: "5 days",
			desc:     "5 days",
		},
		{
			duration: time.Hour * 24 * 30,
			expected: "1 month",
			desc:     "1 month",
		},
		{
			duration: time.Hour * 24 * 60,
			expected: "2 months",
			desc:     "2 months",
		},
		{
			duration: time.Hour * 24 * 500,
			expected: "1 year",
			desc:     "500 days",
		},
		{
			duration: time.Hour * 24 * 1000,
			expected: "2 years",
			desc:     "2 years",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := FormatDuration(tC.duration)
			if output != tC.expected {
				t.Errorf("expected %s, but got %s", tC.expected, output)
			}
		})
	}
}
