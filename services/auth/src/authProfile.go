package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"regexp"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randomString(length int) string {
	b := make([]byte, length)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := length-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func newUUID() string {
	return randomString(15)
}

func isEmailValid(email string) bool {
	regexStr := "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])"
	isMatch, err := regexp.Match(regexStr, []byte(email))
	if err != nil {
		fmt.Println("Error with regex: ", err)
		return false
	}
	return isMatch
}

func isPasswordValid(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasUpperCase, err := regexp.Match("[A-Z]", []byte(password))
	haslowerCase, err := regexp.Match("[a-z]", []byte(password))
	hasNumber, err := regexp.Match("\\d", []byte(password))
	hasNonAlpha, err := regexp.Match("\\W", []byte(password))
	if err != nil {
		fmt.Printf("Error from regex: %q\n", err.Error())
		return false
	}
	b2i := map[bool]int8{false: 0, true: 1}
	if b2i[hasUpperCase] + b2i[haslowerCase] + b2i[hasNumber] + b2i[hasNonAlpha] < 3 {
		return false
	}
	return true
}

func hashPassword(word string) string {
	cost := bcrypt.DefaultCost + 2
	result, err := bcrypt.GenerateFromPassword([]byte(word), cost)
	if err != nil {
		return "" // TODO: Handle Hash error
	}
	return string(result)
}

func doPasswordsMatch(given, stored string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(stored), []byte(given))
	if err != nil {
		return false
	}
	return true
}
