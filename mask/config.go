/*
 * Copyright(c)  2021 Lianjia, Inc.  All Rights Reserved
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mask

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/csv"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"gopkg.in/yaml.v2"
)

type MaskRule struct {
	MaskFunc string   `yaml:"func"`
	Args     []string `yaml:"args"`
}

var defaultMaskConfig map[string]MaskRule

func ParseMaskConfig(file string) error {

	defaultMaskConfig = make(map[string]MaskRule)

	// not config mask
	if file == "" {
		return nil
	}

	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	r := csv.NewReader(fd)
	r.FieldsPerRecord = -1 // fix wrong number of fields
	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(file), "."))
	switch suffix {
	case "csv":
		r.Comma = ','
	case "psv":
		r.Comma = '|'
	case "tsv":
		r.Comma = '\t'
	case "txt":
		r.Comma = ' '
	default:
		err = fmt.Errorf("not support extension: " + suffix)
	}
	for {
		row, err := r.Read()
		if err == io.EOF { // end of file
			break
		} else if err != nil {
			return err
		}

		if len(row) > 1 {
			defaultMaskConfig[strings.ToLower(row[0])] = MaskRule{
				MaskFunc: strings.ToLower(row[1]),
				Args:     row[2:],
			}
		}
	}
	return err
}

type EncryptCipherString struct {
	FFKey   string `yaml:"FFKey"`
	FFTweak string `yaml:"FFTweak"`

	PublicKeyRSA  string `yaml:"PublicKeyRSA"`
	PrivateKeyRSA string `yaml:"PrivateKeyRSA"`

	PublicKeyECC  string `yaml:"PublicKeyECC"`
	PrivateKeyECC string `yaml:"PrivateKeyECC"`

	PrivateKeySM2 string `yaml:"PrivateKeySM2"`

	SM4Key string `yaml:"SM4Key"`
	SM4IV  string `yaml:"SM4IV"`

	AESKey string `yaml:"AESKey"`
	AESIV  string `yaml:"AESIV"`

	DESKey string `yaml:"DESKey"`
	DESIV  string `yaml:"DESIV"`

	TDEAKey string `yaml:"TDEAKey"`
	TDEAIV  string `yaml:"TDEAIV"`

	AESCTRKey string `yaml:"AESCTRKey"`
	AESCTRIV  string `yaml:"AESCTRIV"`
}

var defaultCipherString EncryptCipherString

type EncryptCipher struct {
	FFKey   []byte
	FFTweak []byte

	PublicKeyRSA  []byte
	PrivateKeyRSA []byte

	PublicKeyECC  []byte
	PrivateKeyECC []byte

	PrivateKeySM2 *sm2.PrivateKey

	SM3Hash hash.Hash

	SM4Key []byte
	SM4IV  []byte

	AESKey []byte
	AESIV  []byte

	DESKey []byte
	DESIV  []byte

	TDEAKey []byte
	TDEAIV  []byte

	AESCTRKey []byte
	AESCTRIV  []byte
}

var defaultCipher EncryptCipher

// GenerateEncryptCipher ...
func GenerateEncryptCipher() error {

	// ffKey
	ffKey := make([]byte, 32)
	n, err := rand.Read(ffKey)
	if err != nil {
		return err
	}
	if n != 32 {
		return fmt.Errorf("get wrong random bytes")
	}

	// ffTweak
	ffTweak := make([]byte, 8)
	n, err = rand.Read(ffTweak)
	if err != nil {
		return err
	}
	if n != 8 {
		return fmt.Errorf("get wrong random bytes")
	}

	// defaultEncryptKey
	encryptKey := make([]byte, 128)
	n, err = rand.Read(encryptKey)
	if err != nil {
		return err
	}
	if n != 128 {
		return fmt.Errorf("get wrong random bytes")
	}

	// defaultEncryptKey
	encryptIV := make([]byte, 16)
	n, err = rand.Read(encryptIV)
	if err != nil {
		return err
	}
	if n != 16 {
		return fmt.Errorf("get wrong random bytes")
	}

	// SM2 KEY
	privateKeySM2, err := sm2.GenerateKey()
	if err != nil {
		return err
	}

	// SM3 KEY
	sm3Hash := sm3.New()

	// RSA KEY
	privateKeyRSA, publicKeyRSA, err := genRSAKey()
	if err != nil {
		return err
	}

	// ECC KEY
	privateKeyECC, publicKeyECC, err := genECCKey()

	defaultCipher = EncryptCipher{
		FFKey:         ffKey,
		FFTweak:       ffTweak,
		PublicKeyRSA:  publicKeyRSA,
		PrivateKeyRSA: privateKeyRSA,
		PublicKeyECC:  publicKeyECC,
		PrivateKeyECC: privateKeyECC,
		PrivateKeySM2: privateKeySM2,
		SM3Hash:       sm3Hash,
		SM4Key:        encryptKey[:16],
		SM4IV:         encryptIV[:16],
		AESKey:        encryptKey[:32],
		AESIV:         encryptIV[:16],
		DESKey:        encryptKey[:8],
		DESIV:         encryptIV[:8],
		TDEAKey:       encryptKey[:24],
		TDEAIV:        encryptIV[:8],
		AESCTRKey:     encryptKey[:24],
		AESCTRIV:      encryptIV[:16],
	}
	defaultCipherString = encodeCipher(defaultCipher)

	return err
}

// genRSAKey ...
// generate RSA privatekey 、publickey
func genRSAKey() (privatekey []byte, publickey []byte, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return privatekey, publickey, err
	}
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := pem.Block{
		Bytes: x509PrivateKey,
	}
	privatekey = pem.EncodeToMemory(&privateBlock)

	x509PublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return privatekey, publickey, err
	}
	publicBlock := pem.Block{
		Bytes: x509PublicKey,
	}
	publickey = pem.EncodeToMemory(&publicBlock)
	return privatekey, publickey, err
}

// genECCKey ...
// generate ECC privatekey 、publickey
func genECCKey() (privatekey []byte, publickey []byte, err error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return privatekey, publickey, err
	}
	x509PrivateKey, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return privatekey, publickey, err
	}

	privateBlock := pem.Block{
		Bytes: x509PrivateKey,
	}
	privatekey = pem.EncodeToMemory(&privateBlock)

	x509PublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return privatekey, publickey, err
	}
	publicBlock := pem.Block{
		Bytes: x509PublicKey,
	}
	publickey = pem.EncodeToMemory(&publicBlock)
	return privatekey, publickey, err
}

func PrintCipher() {
	cipher, err := yaml.Marshal(defaultCipherString)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(cipher))
}

func ParseCipherConfig(file string) error {

	// not config mask
	if file == "" {
		return GenerateEncryptCipher()
	}

	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	buf, err := io.ReadAll(fd)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, &defaultCipherString)
	if err != nil {
		return err
	}
	defaultCipher, err = decodeCipher(defaultCipherString)
	return err
}

func encodeCipher(cipher EncryptCipher) EncryptCipherString {

	sm2PrivateKey, err := sm2.WritePrivateKeytoMem(cipher.PrivateKeySM2, nil)
	if err != nil {
		sm2PrivateKey = []byte("WRONG PRIVATE KEY")
	}

	return EncryptCipherString{
		FFKey:         hex.EncodeToString(cipher.FFKey),
		FFTweak:       hex.EncodeToString(cipher.FFTweak),
		PublicKeyRSA:  string(cipher.PublicKeyRSA),
		PrivateKeyRSA: string(cipher.PrivateKeyRSA),
		PublicKeyECC:  string(cipher.PublicKeyECC),
		PrivateKeyECC: string(cipher.PrivateKeyECC),
		PrivateKeySM2: string(sm2PrivateKey),

		SM4Key:    hex.EncodeToString(cipher.SM4Key),
		SM4IV:     hex.EncodeToString(cipher.SM4IV),
		AESKey:    hex.EncodeToString(cipher.AESKey),
		AESIV:     hex.EncodeToString(cipher.AESIV),
		DESKey:    hex.EncodeToString(cipher.DESKey),
		DESIV:     hex.EncodeToString(cipher.DESIV),
		TDEAKey:   hex.EncodeToString(cipher.TDEAKey),
		TDEAIV:    hex.EncodeToString(cipher.TDEAIV),
		AESCTRKey: hex.EncodeToString(cipher.AESCTRKey),
		AESCTRIV:  hex.EncodeToString(cipher.AESCTRIV),
	}
}

func decodeCipher(cipher EncryptCipherString) (EncryptCipher, error) {
	var enc EncryptCipher

	ffKey, err := hex.DecodeString(cipher.FFKey)
	if err != nil {
		return enc, err
	}
	ffTweak, err := hex.DecodeString(cipher.FFTweak)
	if err != nil {
		return enc, err
	}
	sm4Key, err := hex.DecodeString(cipher.SM4Key)
	if err != nil {
		return enc, err
	}
	sm4IV, err := hex.DecodeString(cipher.SM4IV)
	if err != nil {
		return enc, err
	}
	aesKey, err := hex.DecodeString(cipher.AESKey)
	if err != nil {
		return enc, err
	}
	aesIV, err := hex.DecodeString(cipher.AESIV)
	if err != nil {
		return enc, err
	}
	desKey, err := hex.DecodeString(cipher.DESKey)
	if err != nil {
		return enc, err
	}
	desIV, err := hex.DecodeString(cipher.DESIV)
	if err != nil {
		return enc, err
	}
	tdeaKey, err := hex.DecodeString(cipher.TDEAKey)
	if err != nil {
		return enc, err
	}
	tdeaIV, err := hex.DecodeString(cipher.TDEAIV)
	if err != nil {
		return enc, err
	}
	aesctrKey, err := hex.DecodeString(cipher.AESCTRKey)
	if err != nil {
		return enc, err
	}
	aesctrIV, err := hex.DecodeString(cipher.AESCTRIV)
	if err != nil {
		return enc, err
	}
	sm2PrivateKey, err := sm2.ReadPrivateKeyFromMem([]byte(cipher.PrivateKeySM2), nil)
	if err != nil {
		return enc, err
	}

	enc = EncryptCipher{
		FFKey:         ffKey,
		FFTweak:       ffTweak,
		PublicKeyRSA:  []byte(cipher.PublicKeyRSA),
		PrivateKeyRSA: []byte(cipher.PrivateKeyRSA),
		PublicKeyECC:  []byte(cipher.PublicKeyECC),
		PrivateKeyECC: []byte(cipher.PrivateKeyECC),
		PrivateKeySM2: sm2PrivateKey,

		SM4Key: sm4Key,
		SM4IV:  sm4IV,

		AESKey:    aesKey,
		AESIV:     aesIV,
		DESKey:    desKey,
		DESIV:     desIV,
		TDEAKey:   tdeaKey,
		TDEAIV:    tdeaIV,
		AESCTRKey: aesctrKey,
		AESCTRIV:  aesctrIV,
	}
	return enc, err
}
