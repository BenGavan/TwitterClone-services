package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEmailValidation(t *testing.T) {
	table := []struct {
		email    string
		expected bool
	}{
		{"ben.gavan@student.manchester.ac.uk", true},
		{"", false},
		{"@", false},
		{"someone@somewhere.com", true},
		{"someone@somewhere", false},
	}

	for _, row := range table {
		got := isEmailValid(row.email)
		require.Equalf(t, row.expected, got, "email: %q, expected: %v, got: %v", row.email, row.expected, got)
	}
}

func TestPasswordVaildation(t *testing.T) {
	table := []struct {
		password string
		expected bool
	}{
		{"no", false},
		{"thisisalongpassword", true},
	}

	for _, row := range table {
		require.Equal(t, row.expected, isPasswordValid(row.password))
	}
}

func TestPasswordHashing(t *testing.T) {
	table := []struct {
		passwordRaw  string
		passwordHash string
		expected     bool
	}{
		{"password", hashPassword("password"), true},
		{"password", hashPassword("notpassword"), false},
		{"password", hashPassword("password1"), false},
		{"", hashPassword(""), true},
		{" ", hashPassword(" "), true},
		{"  ", hashPassword(" "), false},
	}

	for _, row := range table {
		require.Equal(t, row.expected, doPasswordsMatch(row.passwordRaw, row.passwordHash))
	}
}

func TestUUIDGeneration(t *testing.T) {
	uuids := make([]string, 1000)
	for i:=0; i<1000; i++ {
		uuids[i] = newUUID()
	}
	testUUID := newUUID()
	for _, uuid := range uuids {
		require.NotEqual(t, uuid, testUUID)
	}
}
