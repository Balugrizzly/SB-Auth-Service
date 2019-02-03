package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	// dbusPath is the default path for dbus machine id.
	dbusPath = "/var/lib/dbus/machine-id"
	// dbusPathEtc is the default path for dbus machine id located in /etc.
	// Some systems (like Fedora 20) only know this path.
	// Sometimes it's the other way round.
	dbusPathEtc = "/etc/machine-id"
)

func generateToken() string {
	// Create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"created": time.Now().Unix(),
	})

	id, _ := machineID()
	hashedMachineID := protect("MtU8u9ZwJYnUAuBCktHRDH9cskZsW", id)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte(hashedMachineID))
	return tokenString
}

// machineID returns the uuid specified at `/var/lib/dbus/machine-id` or `/etc/machine-id`.
// If there is an error reading the files an empty string is returned.
// See https://unix.stackexchange.com/questions/144812/generate-consistent-machine-unique-id
func machineID() (string, error) {
	id, err := readFile(dbusPath)
	if err != nil {
		// try fallback path
		id, err = readFile(dbusPathEtc)
	}
	if err != nil {
		return "", err
	}
	return trim(string(id)), nil
}

func readFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func trim(s string) string {
	return strings.TrimSpace(strings.Trim(s, "\n"))
}

// protect calculates HMAC-SHA256 of the application ID, keyed by the machine ID and returns a hex-encoded string.
func protect(appID, id string) string {
	mac := hmac.New(sha256.New, []byte(id))
	mac.Write([]byte(appID))
	return hex.EncodeToString(mac.Sum(nil))
}
