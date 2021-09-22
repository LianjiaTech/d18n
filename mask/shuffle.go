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
	"fmt"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"d18n/common"

	"github.com/andrewarchi/gocipher/gocipher"
)

var shuffleMap map[string]map[string]string

func InitShuffle(seed int64) error {
	rand.Seed(seed)

	files, err := corpusFS.ReadDir("corpus")
	if err != nil {
		return err
	}

	shuffleMap = make(map[string]map[string]string, len(files))
	for _, file := range files {
		if !strings.HasPrefix(file.Name(), "shuffle") {
			continue
		}
		c := strings.ToLower(strings.TrimLeft(filepath.Ext(file.Name()), "."))

		// load corpus from embed.FS
		characters, err := corpusFS.ReadFile("corpus/" + file.Name())
		if err != nil {
			return err
		}
		runes := []rune(string(characters))

		// shuffle []rune
		var shuffle []string
		for _, r := range runes {
			shuffle = append(shuffle, string(r))
		}
		rand.Shuffle(len(runes), func(i, j int) {
			runes[i], runes[j] = runes[j], runes[i]
		})

		// convert []rune to map[string]string
		m := make(map[string]string, len(runes))
		for i, s := range runes {
			m[string(s)] = shuffle[i]
		}

		// set shuffleMap
		shuffleMap[c] = m
	}

	// default shuffle map, keep data type replace
	defaultMap := make(map[string]string)
	for _, c := range []string{"0-9", "lower", "upper", "zh_CN"} {
		for k, v := range shuffleMap[strings.ToLower(c)] {
			defaultMap[k] = v
		}
	}
	shuffleMap["default"] = defaultMap
	return nil
}

// Shuffle shuffle and keep data type
// arg 0: value
// arg 1: corpus name
func Shuffle(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	var corpus = "default"
	if len(args) > 1 {
		corpus = strings.ToLower(fmt.Sprint(args[1]))
		if _, ok := shuffleMap[corpus]; !ok {
			return ret, fmt.Errorf(common.WrongArgValue)
		}
	}

	for _, s := range fmt.Sprint(args[0]) {
		if c, ok := shuffleMap[corpus][fmt.Sprintf("%c", s)]; ok {
			ret += c
		} else {
			ret += fmt.Sprintf("%c", s)
		}
	}
	return
}

// ShuffleRight shuffle right and keep data type
// 保持前n位不变，混淆其余部分。可针对字母和数字字符在同为字母或数字范围内进行混淆，特殊符号将保留。
func ShuffleRight(args ...interface{}) (ret string, err error) {
	if len(args) < 2 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	n, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}

	runes := []rune(fmt.Sprint(args[0]))
	if len(runes) <= n {
		return string(runes), nil
	} else {
		s, err := Shuffle(string(runes[n:]))
		if err != nil {
			return string(runes), err
		}
		ret = string(runes[:n]) + s
	}

	return
}

// ShuffleLeft shuffle left and keep data type
// 保持后n位不变，混淆其余部分。可针对字母和数字字符在同为字母或数字范围内进行混淆，特殊符号将保留。
func ShuffleLeft(args ...interface{}) (ret string, err error) {
	if len(args) < 2 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	n, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}

	runes := []rune(fmt.Sprint(args[0]))
	if len(runes) <= n {
		return string(runes), nil
	} else {
		s, err := Shuffle(string(runes[:len(runes)-n]))
		if err != nil {
			return string(runes), err
		}
		ret = s + string(runes[len(runes)-n:])
	}

	return
}

// Rot letter substitution with the Nth letter after it in the alphabet.
func Rot(args ...interface{}) (ret string, err error) {
	if len(args) < 2 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	n, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}

	switch n {
	case 47:
		ret = gocipher.ROT47.Encipher(fmt.Sprint(args[0]))
	case 18:
		ret = gocipher.ROT18.Encipher(fmt.Sprint(args[0]))
	case 13:
		ret = gocipher.ROT13.Encipher(fmt.Sprint(args[0]))
	case 5:
		ret = gocipher.ROT5.Encipher(fmt.Sprint(args[0]))
	case 32768:
		ret = rot32768(fmt.Sprint(args[0]))
	default:
		err = fmt.Errorf(common.WrongMaskFunc)
	}
	return
}

// rot32768 rotates utf8 string
// https://www.socketloop.com/tutorials/golang-rot32768-rotate-by-0x80-utf-8-strings-example
func rot32768(input string) string {
	var result []string
	rot5map := map[rune]rune{'0': '5', '1': '6', '2': '7', '3': '8', '4': '9', '5': '0', '6': '1', '7': '2', '8': '3', '9': '4'}

	for _, i := range input {
		switch {
		case unicode.IsSpace(i):
			result = append(result, " ")
		case i >= '0' && i <= '9':
			result = append(result, string(rot5map[i]))
		case utf8.ValidRune(i):
			result = append(result, string(rune(i)^utf8.RuneSelf))
		}
	}
	return strings.Join(result, "")
}

func Morse(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	morse := gocipher.NewMorse()
	return morse.Encode(fmt.Sprint(args[0]))
}

func Caesar(args ...interface{}) (ret string, err error) {
	if len(args) < 2 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	n, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}
	caesar := gocipher.NewCaesar(n)
	return caesar.Encipher(fmt.Sprint(args[0])), err
}
