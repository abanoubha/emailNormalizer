package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	emails := []string{
		"abanoub@gmail.com",
		"Abanoub@GMail.com",
		"Abanoub+1@GMail.com",
		"Abanoub+h@GMail.com",
		"A.banoub@GMail.com",
		"Ab.anoub@GMail.com",
		"Aba.noub@GMail.com",
		"Aban.oub@GMail.com",
		"Abano.ub@GMail.com",
		"Abanou.b@GMail.com",
		"A.banou.b@GMail.com",
		"A.b.anou.b@GMail.com",
		"A.b.a.nou.b@GMail.com",
		"A.b.a.n.ou.b@GMail.com",
		"A.b.a.n.o.u.b@GMail.com",
		"A.b.a.n.o.u.b+2@GMail.com",
	}

	for _, e := range emails {
		normalized, err := normalize(e)
		if err != nil {
			fmt.Printf("Error normalizing this email: %v, because: %v\n", e, err.Error())
		}
		fmt.Printf("email: %v, normalized: %v, %v\n", e, normalized, normalized == "abanoub@gmail.com")

	}
}

func normalize(email string) (string, error) {
	// remove leading and trailing whitespaces
	email = strings.TrimSpace(email)

	// convert to lowercase
	email = strings.ToLower(email)

	// remove dots from email's username
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		// the email has more than one @ symbol
		return "", errors.New("E-mail is not valid.")
	}
	usernameWithoutDots := strings.ReplaceAll(parts[0], ".", "")

	// remove the part after '+' symbol of the email's username
	plusParts := strings.Split(usernameWithoutDots, "+")

	return plusParts[0] + "@" + parts[1], nil
}
