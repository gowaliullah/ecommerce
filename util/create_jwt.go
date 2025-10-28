package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role" db:"role"`
}

func CreateJWT(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	hearderB64 := base64UrlEncode(byteArrHeader)

	byteArrData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	paylodB64 := base64UrlEncode(byteArrData)

	message := hearderB64 + "." + paylodB64

	byteArrSecret := []byte(secret)
	byteArrMsg := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMsg)

	signature := h.Sum(nil)
	signatureB64 := base64UrlEncode(signature)

	jwt := hearderB64 + "." + paylodB64 + "." + signatureB64

	return jwt, nil
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}

func Verify(jwt, secret string) (bool, error) {
	jwtChunkArr := strings.Split(jwt, ".")
	headerAndPayloadHash := jwtChunkArr[0] + "." + jwtChunkArr[1]

	byteSecret := []byte(secret)
	byteHeaderAndPayload := []byte(headerAndPayloadHash)

	hash := hmac.New(sha256.New, byteSecret)
	hash.Write(byteHeaderAndPayload)

	signature := hash.Sum(nil)

	signatureB64 := base64UrlEncode(signature)

	verifyJwt := headerAndPayloadHash + "." + signatureB64

	if jwt == verifyJwt {
		return true, nil
	}
	return false, errors.New("invalid token")
}

func DecodeToken(token, secret string) (*Payload, error) {

	verified, err := Verify(token, secret)
	if err != nil {
		return nil, err
	}
	if !verified {
		return nil, errors.New("invalid token")
	}
	jwtChunkArr := strings.Split(token, ".")

	payloadHash := jwtChunkArr[1]

	payloadByte, err := base64.RawURLEncoding.DecodeString(payloadHash)
	if err != nil {
		return nil, err
	}

	var payload Payload
	err = json.Unmarshal(payloadByte, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}
