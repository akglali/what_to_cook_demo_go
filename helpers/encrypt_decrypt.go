package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func Encrypt(userToken string) (string, error) {
	theSecretText := os.Getenv("SECRET-TEXT")
	// Key must be 16, 24, or 32 bytes long
	key := []byte(os.Getenv("ENCRYPT-KEY"))
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	plaintext := []byte(theSecretText + "," + currentTime + "," + userToken)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err)
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	//fmt.Printf("Ciphertext: %x\n", ciphertext)

	return string(ciphertext), nil
}

func Decrypt(userToken string) (string, error) {
	key := []byte(os.Getenv("ENCRYPT-KEY"))
	ciphertext := []byte(userToken)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil

}

func UserAuth(userToken string) (bool, error, string) {
	theSecretTextEnv := os.Getenv("SECRET-TEXT")
	decryptedText, err := Decrypt(userToken)
	if err != nil {
		return false, err, ""
	}
	//if user change the token format then strings.Split might raise an error and shut down the program
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Something went wrong!!\nPlease sign out and sign in again.")
			//NotAcceptableAbort(c, "")
			return
		}
	}()
	splitText := strings.Split(decryptedText, ",")
	theSecretText := splitText[0]
	//check if the secret text is as same as .env value
	if theSecretText != theSecretTextEnv {
		return false, errors.New("secret Key is wrong"), ""
	}
	//check the time format is correct
	createdTime := splitText[1]
	_, err = time.Parse("2006-01-02 15:04:05", createdTime)
	if err != nil {
		return false, errors.New("time format is not valid"), ""
	}
	//if they both correct then user auth is succeed.
	actualToken := splitText[2]

	return true, nil, actualToken

}
