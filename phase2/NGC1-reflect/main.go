package main

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

type Avenger struct {
	Name  string `required:"true" minLen:"2" maxLen:"25"`
	Age   int    `required:"true" min:"17" max:"60"`
	Email string `required:"true" regex:"true"`
}

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

func ValidateAvenger(a Avenger) error{
	t := reflect.TypeOf(a)
	for i := 0;i<t.NumField();i++{
		field := t.Field(i)
		if field.Tag.Get("required") == "true"{
			value := reflect.ValueOf(a).Field(i).String()
			if value == "" {
				return errors.New(field.Name+" is required")
			}
		}
		if field.Tag.Get("minLen") == "2" || field.Tag.Get("maxLen") == "25"{
			value := reflect.ValueOf(a).Field(i).String()
			if len(value) < 2 || len(value) > 25{
				return errors.New(field.Name+" character length should be greater than 2 and less than 25.")
			}
		}
		if field.Tag.Get("min") == "17" || field.Tag.Get("max") == "60"{
			value := reflect.ValueOf(a).Field(i).Int()
			if value < 17 || value > 60{
				return errors.New(field.Name+" should be greater than 17 and less than 60.")
			}
		}
		if field.Tag.Get("regex") == "true"{
			value := reflect.ValueOf(a).Field(i).String()
			if !validateEmail(value) {
				return errors.New("Email is invalid, an example of a correct email format : example@mail.com")
			}
		}
	}
	return nil
}



func main() {
	// correct example validated
	Avenger1 := Avenger{
		Name:  "Captain Amerika",
		Age:   40,
		Email: "captainamerika@mail.com",
	}

	// incorrect example 1 validated
	Avenger2 := Avenger{
		Name:  "I",
		Age:   40,
		Email: "ironman@mail.com",
	}

	// incorrect example 2 validated
	Avenger3 := Avenger{
		Name:  "Hulk",
		Age:   40,
		Email: "hulk@ma-ilcom",
	}

	err := ValidateAvenger(Avenger1)
	if err != nil {
		fmt.Println(err)
	}

	err = ValidateAvenger(Avenger2)
	if err != nil {
		fmt.Println(err)
	}

	err = ValidateAvenger(Avenger3)
	if err != nil {
		fmt.Println(err)
	}
}