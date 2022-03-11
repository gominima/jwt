package jwt 

import (
	"context"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

type AuthJwt struct {
	algo jwa.SignatureAlgorithm
	signkey interface{}
	verifykey interface{}
	verif jwt.ParseOption
}

func New(alg string, skey interface{}, vkey interface{}) *AuthJwt {
	ja := &AuthJwt{
		algo: jwa.SignatureAlgorithm(alg),
		signkey: skey,
		verifykey: vkey,
	}
	if ja.verifykey != nil {
		ja.verif = jwt.WithVerify(ja.algo, ja.verifykey)
		return ja
	}
	ja.verif = jwt.WithVerify(ja.algo, ja.signkey)
	return ja
}


func Verify()