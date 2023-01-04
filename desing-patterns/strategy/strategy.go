package main

import "fmt"

type HashAlgorithm interface {
	Hash(p *PasswordProtect)
}

type PasswordProtect struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

func NewPasswordProtect(user string, passwordName string, hash HashAlgorithm) *PasswordProtect {
	return &PasswordProtect{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: hash,
	}
}

func (p *PasswordProtect) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *PasswordProtect) Hash() {
	p.hashAlgorithm.Hash(p)
}

type SHA struct{}

func (SHA) Hash(p *PasswordProtect) {
	fmt.Printf("Hashing using SHA for %s\n", p.passwordName)
}

type MD5 struct{}

func (MD5) Hash(p *PasswordProtect) {
	fmt.Printf("Hashing using MD5 for %s\n", p.passwordName)
}

func main() {
	sha := &SHA{}
	md5 := &MD5{}

	PasswordProtect := NewPasswordProtect("Vic", "gmail password", sha)
	PasswordProtect.Hash()
	PasswordProtect.SetHashAlgorithm(md5)
	PasswordProtect.Hash()
}
