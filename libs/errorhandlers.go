package libs

import (
	"funtastix/backend/dto"
	"strings"
)

func RegisterErrHandler(err error) string {
	var errMessage string
	if strings.Contains(err.Error(), "Field validation for 'Email' failed on the 'min' tag") {
		errMessage = "Email at least have 12 character or more"
	}

	if strings.Contains(err.Error(), "Field validation for 'Email' failed on the 'email' tag") {
		errMessage = "Invalid email format"
	}

	if strings.Contains(err.Error(), "Field validation for 'Password' failed on the 'min' tag") {
		errMessage = "Password at least have 6 character or more"
	}

	return errMessage
}

func StrongPasswordHandler(form dto.RegisterDTO) string {
	var errMess string

	upperCased := strings.ContainsAny(form.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerCased := strings.ContainsAny(form.Password, "abcdefghijklmnopqrstuvwxyz")
	includeNum := strings.ContainsAny(form.Password, "0123456789")
	includeSC := strings.ContainsAny(form.Password, "!@#$%^&*()-_=+[]{}|;:',.<>?/`~")

	if !upperCased {
		errMess = "Password must contains one uppercased, one lowercased, one number, and one special character eg #@$%?!... etc"
	}
	if !lowerCased {
		errMess = "Password must contains one uppercased, one lowercased, one number, and one special character eg #@$%?!... etc"
	}
	if !includeNum {
		errMess = "Password must contains one uppercased, one lowercased, one number, and one special character eg #@$%?!... etc"
	}
	if !includeSC {
		errMess = "Password must contains one uppercased, one lowercased, one number, and one special character eg #@$%?!... etc"
	}

	return errMess
}
