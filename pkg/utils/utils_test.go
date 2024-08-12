package utils

import "testing"

func TestValidateName(t *testing.T) {
	testCases := []struct {
		name                 string
		expectedValid        bool
		expectedErrorMessage string
		desc                 string
	}{
		{
			name:                 "e",
			expectedValid:        false,
			expectedErrorMessage: "name must be at least 2 characters long",
			desc:                 "Short Name",
		},
		{
			name:                 "Elyar",
			expectedValid:        true,
			expectedErrorMessage: "",
			desc:                 "Valid Name",
		},
		{
			name:                 "الیار",
			expectedValid:        false,
			expectedErrorMessage: "name can only contain letters and spaces",
			desc:                 "Invalid Charachters",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			valid, message := ValidateName(tC.name)
			if valid != tC.expectedValid {
				t.Errorf("Expected isValid for name %s to be %t, but got %t", tC.name, tC.expectedValid, valid)
			}
			if message != tC.expectedErrorMessage {
				t.Errorf("Expected error message to be %s, but got %s", tC.expectedErrorMessage, message)
			}
		})
	}
}

func TestValidateUserName(t *testing.T) {
	testCases := []struct {
		username             string
		expectedValid        bool
		expectedErrorMessage string
		desc                 string
	}{
		{
			username:             "eL",
			expectedValid:        false,
			expectedErrorMessage: "username must be at least 3 characters long",
			desc:                 "Short Username",
		},
		{
			username:             "Elyar",
			expectedValid:        true,
			expectedErrorMessage: "",
			desc:                 "Valid Username",
		},
		{
			username:             "الیار",
			expectedValid:        false,
			expectedErrorMessage: "username can only contain letters, digits, underscores, and dots, and must start with a letter",
			desc:                 "Invalid Charachters",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			valid, message := ValidateUsername(tC.username)
			if valid != tC.expectedValid {
				t.Errorf("Expected isValid for name %s to be %t, but got %t", tC.username, tC.expectedValid, valid)
			}
			if message != tC.expectedErrorMessage {
				t.Errorf("Expected error message to be %s, but got %s", tC.expectedErrorMessage, message)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	testCases := []struct {
		password             string
		expectedValid        bool
		expectedErrorMessage string
		desc                 string
	}{
		{
			password:             "e",
			expectedValid:        false,
			expectedErrorMessage: "password must be at least 8 characters long",
			desc:                 "Short password",
		},
		{
			password:             "Elyar",
			expectedValid:        false,
			expectedErrorMessage: "password must be at least 8 characters long",
			desc:                 "Valid Name but Short password",
		},
		{
			password:             "Password",
			expectedValid:        false,
			expectedErrorMessage: "password must contain at least one digit",
			desc:                 "Missing digit",
		},
		{
			password:             "Password1",
			expectedValid:        false,
			expectedErrorMessage: "password must contain at least one special character",
			desc:                 "Missing special character",
		},
		{
			password:             "password1!",
			expectedValid:        false,
			expectedErrorMessage: "password must contain at least one uppercase letter",
			desc:                 "Missing uppercase letter",
		},
		{
			password:             "PASSWORD1!",
			expectedValid:        false,
			expectedErrorMessage: "password must contain at least one lowercase letter",
			desc:                 "Missing lowercase letter",
		},
		{
			password:             "P@ssw0rd",
			expectedValid:        true,
			expectedErrorMessage: "",
			desc:                 "Valid password",
		},
		{
			password:             "P@ssword123!",
			expectedValid:        true,
			expectedErrorMessage: "",
			desc:                 "Valid password with more characters",
		},
		{
			password:             "P@ssw0",
			expectedValid:        false,
			expectedErrorMessage: "password must be at least 8 characters long",
			desc:                 "Valid components but too short",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			valid, message := ValidatePassword(tC.password)
			if valid != tC.expectedValid {
				t.Errorf("Expected isValid for password %s to be %t, but got %t", tC.password, tC.expectedValid, valid)
			}
			if message != tC.expectedErrorMessage {
				t.Errorf("Expected error message to be %s, but got %s", tC.expectedErrorMessage, message)
			}
		})
	}
}
