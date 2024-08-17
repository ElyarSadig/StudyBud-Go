package utils

import "testing"

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
