package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type mailProof struct {
	Mail string `json:"mail"`
	jwt.StandardClaims
}

const (
	ProofExpireDuration = time.Hour * 48
)

var MailSecret = []byte("snailMail")

func GenResetProof(mail string) (string, error) {
	resetProof := new(mailProof)
	resetProof.Mail = mail
	resetProof.ExpiresAt = time.Now().Add(ProofExpireDuration).Unix()
	proof := jwt.NewWithClaims(jwt.SigningMethodHS256, resetProof)
	return proof.SignedString(MailSecret)
}

func ParseMailProof(proofString string, mail string) (bool, error) {
	proof, err := jwt.ParseWithClaims(proofString, &mailProof{}, func(token *jwt.Token) (interface{}, error) {
		return MailSecret, nil
	})
	if err != nil {
		return false, err
	}
	if mailInfo, ok := proof.Claims.(*mailProof); ok && proof.Valid {
		return mailInfo.Mail == mail, nil
	}
	return false, errors.New("invalid proof")
}
