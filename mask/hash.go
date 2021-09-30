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
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"hash/crc32"
	"strings"

	"github.com/LianjiaTech/d18n/common"
)

// this file if for decryptable function

func CRC32(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	ret = fmt.Sprint(args[0])
	ret = fmt.Sprintf("%08x", crc32.ChecksumIEEE([]byte(ret)))
	return ret, err
}

func MD5(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	ret = fmt.Sprint(args[0])
	h := md5.New()
	h.Write([]byte(ret))
	ret = hex.EncodeToString(h.Sum(nil))
	return ret, err
}

func SHA1(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	ret = fmt.Sprint(args[0])
	h := sha1.New()
	h.Write([]byte(ret))
	ret = hex.EncodeToString(h.Sum(nil))
	return ret, err
}

func SHA2(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	ret = fmt.Sprint(args[0])
	h := sha256.New()
	h.Write([]byte(ret))
	ret = hex.EncodeToString(h.Sum(nil))
	return ret, err
}

func HMAC(args ...interface{}) (ret string, err error) {
	if len(args) < 3 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	type cfg map[string]func() hash.Hash
	var functions = cfg{
		"md5":  md5.New,
		"sha1": sha1.New,
		"sha2": sha256.New,
	}

	h := hmac.New(functions[strings.ToLower(fmt.Sprint(args[1]))], []byte(fmt.Sprint(args[2])))

	h.Write([]byte(fmt.Sprint(args[0])))

	return hex.EncodeToString(h.Sum(nil)), nil
}
