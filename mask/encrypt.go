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
	// "crypto/rand"
	"bytes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/LianjiaTech/d18n/common"

	"github.com/capitalone/fpe/ff1"
	"github.com/capitalone/fpe/ff3"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/tjfoc/gmsm/sm4"

	// "github.com/tjfoc/gmsm/sm4"
	"github.com/wumansgy/goEncrypt"
)

// FF1 format-preserving encryption, ff1 algorithm
// arg 0: value
// arg 1: radix, default 10, max size 62, min size 2. [0-9a-zA-Z]
// arg 2: key
// arg 3: tweak
func FF1(args ...interface{}) (ret string, err error) {
	var ffKey, ffTweak []byte

	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	var radix = 10
	if len(args) > 1 {
		radix, err = strconv.Atoi(fmt.Sprint(args[1]))
		if err != nil {
			return ret, err
		}
	}

	if len(args) < 4 {
		ffKey = defaultCipher.FFKey
		ffTweak = defaultCipher.FFTweak
	} else {
		ffKey, err = hex.DecodeString(fmt.Sprint(args[2]))
		if err != nil {
			return ret, err
		}
		ffTweak, err = hex.DecodeString(fmt.Sprint(args[3]))
		if err != nil {
			return ret, err
		}
	}

	cipher, err := ff1.NewCipher(radix, len(ffTweak), ffKey, ffTweak)
	if err != nil {
		return ret, err
	}
	return cipher.Encrypt(fmt.Sprint(args[0]))
}

// FF3 format-preserving encryption, ff3 algorithm
// arg 0: value
// arg 1: radix, default 10, max size 62, min size 2. [0-9a-zA-Z]
// arg 2: key
// arg 3: tweak
func FF3(args ...interface{}) (ret string, err error) {
	var ffKey, ffTweak []byte

	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	var radix = 10
	if len(args) > 1 {
		radix, err = strconv.Atoi(fmt.Sprint(args[1]))
		if err != nil {
			return ret, err
		}
	}

	if len(args) < 4 {
		ffKey = defaultCipher.FFKey
		ffTweak = defaultCipher.FFTweak
	} else {
		ffKey, err = hex.DecodeString(fmt.Sprint(args[2]))
		if err != nil {
			return ret, err
		}
		ffTweak, err = hex.DecodeString(fmt.Sprint(args[3]))
		if err != nil {
			return ret, err
		}
	}

	cipher, err := ff3.NewCipher(radix, ffKey, ffTweak)
	if err != nil {
		return ret, err
	}
	return cipher.Encrypt(fmt.Sprint(args[0]))

}

// Base64 ...
func Base64(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	ret = fmt.Sprint(args[0])
	ret = base64.StdEncoding.EncodeToString([]byte(ret))
	return ret, err
}

type goEncryptFunc func(plainText, key []byte, ivAes ...byte) ([]byte, error)

// doGoEncrypt ...
// goEncrypt.AesCtrDecrypt 、goEncrypt.AesCbcDecrypt、goEncrypt.TripleDesEncrypt
func doGoEncrypt(f goEncryptFunc, args ...interface{}) (ret string, err error) {
	if len(args) < 3 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	cryptText, err := f(
		[]byte(fmt.Sprint(args[0])), // arg 1: plainText []byte
		args[1].([]byte),            // arg 2: key []byte
		args[2].([]byte)...,         // arg 3: ivDes ...byte
	)
	if err != nil {
		return ret, err
	}
	return base64.StdEncoding.EncodeToString(cryptText), err
}

// DES ...
func DES(args ...interface{}) (ret string, err error) {
	var key, ivAes []byte
	if len(args) > 1 {
		key = []byte(fmt.Sprint(args[1]))
	} else {
		key = defaultCipher.DESKey
	}
	if len(args) > 2 {
		ivAes = []byte(fmt.Sprint(args[2]))
	} else {
		ivAes = defaultCipher.DESIV
	}
	return doGoEncrypt(goEncrypt.DesCbcEncrypt, args[0], key, ivAes)
}

// AES ...
func AES(args ...interface{}) (ret string, err error) {
	var key, ivAes []byte
	if len(args) > 1 {
		key = []byte(fmt.Sprint(args[1]))
	} else {
		key = defaultCipher.AESKey
	}
	if len(args) > 2 {
		ivAes = []byte(fmt.Sprint(args[2]))
	} else {
		ivAes = defaultCipher.AESIV
	}
	return doGoEncrypt(goEncrypt.AesCbcEncrypt, args[0], key, ivAes)
}

// TDEA ...
func TDEA(args ...interface{}) (ret string, err error) {
	var key, ivAes []byte
	if len(args) > 1 {
		key = []byte(fmt.Sprint(args[1]))
	} else {
		key = defaultCipher.TDEAKey
	}
	if len(args) > 2 {
		ivAes = []byte(fmt.Sprint(args[2]))
	} else {
		ivAes = defaultCipher.TDEAIV
	}
	return doGoEncrypt(goEncrypt.TripleDesEncrypt, args[0], key, ivAes)
}

// RSA ...
// arg1: cipherText
// arg2: publicKey
func AESCTR(args ...interface{}) (ret string, err error) {
	var key, ivAes []byte
	if len(args) > 1 {
		key = []byte(fmt.Sprint(args[1]))
	} else {
		key = defaultCipher.AESCTRKey
	}
	if len(args) > 2 {
		ivAes = []byte(fmt.Sprint(args[2]))
	} else {
		ivAes = defaultCipher.AESCTRIV
	}
	return doGoEncrypt(goEncrypt.AesCtrEncrypt, args[0], key, ivAes)
}

// RSA ...
// arg1: cipherText
func RSA(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	crypttext, err := goEncrypt.RsaEncrypt([]byte(fmt.Sprint(args[0])), defaultCipher.PublicKeyRSA)
	if err != nil {
		return ret, err
	}
	return hex.EncodeToString(crypttext), err
}

// ECC ...
// arg1: cipherText
func ECC(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	crypttext, err := goEncrypt.EccEncrypt([]byte(fmt.Sprint(args[0])), defaultCipher.PublicKeyECC)
	if err != nil {
		return ret, err
	}
	return hex.EncodeToString(crypttext), err
}

// SM2 ...
// arg1: cipherText
func SM2(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	pub := &defaultCipher.PrivateKeySM2.PublicKey
	crypttext, err := pub.Encrypt([]byte(fmt.Sprint(args[0])))
	if err != nil {
		return ret, err
	}
	return hex.EncodeToString(crypttext), err
}

// SM3 ...
// arg1: cipherText
func SM3(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	h := sm3.New()
	h.Write([]byte(fmt.Sprint(args[0])))
	crypttext := h.Sum(nil)
	return hex.EncodeToString(crypttext), err
}

// SM4 ...
// arg1: cipherText
func SM4(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	crypttext, err := sm4Encrypt(defaultCipher.SM4Key, defaultCipher.SM4IV, []byte(fmt.Sprint(args[0])))
	if err != nil {
		return ret, err
	}

	return hex.EncodeToString(crypttext), err
}

func sm4Encrypt(key, iv, plainText []byte) ([]byte, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData := pkcs5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil
}

// pkcs5Padding ...
func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// pkcs5UnPadding ...
func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	if length == 0 {
		return nil
	}
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
