package user

import (
	"fmt"
	"regexp"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func (u User) ValidateUser() []error {
	var errors []error
	// Ismni tekshir:
	if u.Name == "" {
		errors = append(errors, fmt.Errorf("error: ism bo'sh bo'lishi mumkin emas"))
	}
	if len(u.Name) < 6 {
		errors = append(errors, fmt.Errorf("error: ism kamida 6 belgidan iborat bo'lishi kerak"))
	}
	// YOSHIni tekshir:
	if u.Age < 0 || u.Age > 120 {
		errors = append(errors, fmt.Errorf("error: yosh 0 dan katta va 120 dan kichik bo'lishi kerak"))
	}

	// EMAILini tekshir
	if u.Email == "" {
		errors = append(errors, fmt.Errorf("error: email bo'sh bo'lishi mumkin emas"))
	}
	email := `^[a-zA-Z0-9._%+-]+@gmail\.com$`

	regex := regexp.MustCompile(email)

	b := regex.MatchString(u.Email)
	if !b {
		errors = append(errors, fmt.Errorf("error: Email kiritishda xatolik yuz berdi"))
	}
	return errors
}
