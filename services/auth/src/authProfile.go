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
	regexString := "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

	isMatch, err := regexp.Match(regexString, []byte(email))
	if err != nil {
		fmt.Println("Error with regex: ", err)
		return false
	}
	return isMatch
}

func isPasswordValid(password string) bool {
	if len(password) < 5 {
		return false
	}
	// TODO: Check Password composition (enough letters & numbers)
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
