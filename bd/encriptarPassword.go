package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la rutinoa que me permite encriptar la password*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err

}
