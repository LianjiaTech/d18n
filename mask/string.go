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
	"bytes"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"d18n/common"

	"github.com/bykof/gostradamus"
	"github.com/dnnrly/abbreviate/data"
	"github.com/dnnrly/abbreviate/domain"
)

// Smoke replace every character with mask
// args 0: value
// args 1: replacement
func Smoke(args ...interface{}) (ret string, err error) {
	if len(args) < 2 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	return strings.Repeat(fmt.Sprint(args[1]), len(fmt.Sprint(args[0]))), err
}

// SmokeLeft replace left n characters
// args 0: value
// args 1: left n character
// args 2: replacement
func SmokeLeft(args ...interface{}) (ret string, err error) {
	if len(args) < 3 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	n, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}
	if n < 0 {
		return ret, fmt.Errorf(common.WrongLargeThan0)
	}

	ret = fmt.Sprint(args[0])
	runes := []rune(ret)
	l := len(runes)
	src := fmt.Sprint(args[2])

	if n < l {
		ret = strings.Repeat(src, n) + string(runes[n:l])
	} else {
		ret = strings.Repeat(src, l)
	}
	return ret, err
}

// ReserveLeft reserve left n characters
// args 0: value
// args 1: left n character
// args 2: strings.Repeat(src string)
func ReserveLeft(args ...interface{}) (ret string, err error) {
	if len(args) < 3 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	n, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}

	if n < 0 {
		return ret, fmt.Errorf(common.WrongLargeThan0)
	}

	ret = fmt.Sprint(args[0])
	runes := []rune(ret)
	l := len(runes)
	src := fmt.Sprint(args[2])

	if n < l {
		ret = string(runes[:n]) + strings.Repeat(src, l-n)
	}
	return ret, err
}

// SmokeRight smoke right n characters
// args 0: value
// args 1: right n character
// args 2: strings.Repeat(src string)
func SmokeRight(args ...interface{}) (ret string, err error) {
	if len(args) < 3 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	n, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}

	if n < 0 {
		return ret, fmt.Errorf(common.WrongLargeThan0)
	}

	ret = fmt.Sprint(args[0])
	l := len(ret)
	src := fmt.Sprint(args[2])

	if n < l {
		ret = ret[:l-n] + strings.Repeat(src, n)
	} else {
		ret = strings.Repeat(src, l)
	}
	return ret, err
}

// ReserveRight reserve right n characters
// args 0: value
// args 1: right n character
// args 2: strings.Repeat(src string)
func ReserveRight(args ...interface{}) (ret string, err error) {
	if len(args) < 3 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	n, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}

	if n < 0 {
		return ret, fmt.Errorf(common.WrongLargeThan0)
	}

	ret = fmt.Sprint(args[0])
	runes := []rune(ret)
	l := len(runes)
	src := fmt.Sprint(args[2])

	if n < l {
		ret = strings.Repeat(src, l-n) + string(runes[l-n:])
	}
	return ret, err
}

// ReserveRight smoke margin n characters
// args 0: value
// args 1: margin n character
// args 2: strings.Repeat(src string)
func SmokeMargin(args ...interface{}) (ret string, err error) {
	if len(args) < 3 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	n, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}

	if n < 0 {
		return ret, fmt.Errorf(common.WrongLargeThan0)
	}

	ret = fmt.Sprint(args[0])
	runes := []rune(ret)
	l := len(runes)
	src := fmt.Sprint(args[2])

	if 2*n < l {
		ret = strings.Repeat(src, n) + string(runes[n:l-n]) + strings.Repeat(src, n)
	} else {
		ret = strings.Repeat(src, l)
	}
	return ret, err
}

// ReserveMargin reserve margin n characters
// args 0: value
// args 1: margin n character
// args 2: strings.Repeat(src string)
func ReserveMargin(args ...interface{}) (ret string, err error) {
	if len(args) < 3 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	n, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}

	if n < 0 {
		return ret, fmt.Errorf(common.WrongLargeThan0)
	}

	ret = fmt.Sprint(args[0])
	runes := []rune(ret)
	l := len(runes)
	src := fmt.Sprint(args[2])

	if 2*n < l {
		ret = string(runes[:n]) + strings.Repeat(src, l-2*n) + string(runes[l-n:])
	} else {
		ret = strings.Repeat(src, l)
	}
	return ret, err
}

// SmokeOuter mysql mask_outer
// args 0: value
// args 1, 2: left, right int
// args 3: strings.Repeat(src string)
func SmokeOuter(args ...interface{}) (ret string, err error) {
	if len(args) < 4 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	left, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}
	right, err := strconv.Atoi(fmt.Sprint(args[2]))
	if err != nil {
		return ret, err
	}

	if left < 0 || right < 0 {
		return ret, fmt.Errorf(common.WrongLargeThan0)
	}

	ret = fmt.Sprint(fmt.Sprint(args[0]))
	runes := []rune(ret)
	l := len(runes)
	src := fmt.Sprint(args[3])

	if (left + right) > l {
		return strings.Repeat(src, l), err
	} else {
		ret = strings.Repeat(src, left) + string(runes[left:l-right]) + strings.Repeat(src, right)
	}
	return ret, err
}

// ReserveOuter
func ReserveOuter(args ...interface{}) (ret string, err error) {
	return SmokeInner(args...)
}

// SmokeInner mysql mask_inner
// args 0: value
// args 1, 2: left, right int
// args 3: strings.Repeat(src string)
func SmokeInner(args ...interface{}) (ret string, err error) {
	if len(args) < 4 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	left, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}
	right, err := strconv.Atoi(fmt.Sprint(args[2]))
	if err != nil {
		return ret, err
	}

	if left < 0 || right < 0 {
		return ret, fmt.Errorf(common.WrongLargeThan0)
	}

	ret = fmt.Sprint(args[0])
	runes := []rune(ret)
	l := len(runes)
	src := fmt.Sprint(args[3])

	if (left + right) > l {
		return strings.Repeat(src, l), err
	} else {
		ret = string(runes[:left]) + strings.Repeat(src, l-right-left) + string(runes[l-right:])
	}
	return ret, err
}

// ReserveInner
func ReserveInner(args ...interface{}) (ret string, err error) {
	return SmokeOuter(args...)
}

// Replace strings.Replace
// args 0-3: strings.Replace(s, old, new string, n int)
func Replace(args ...interface{}) (ret string, err error) {
	if len(args) < 3 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	var n int = -1
	if len(args) > 3 {
		n, err = strconv.Atoi(fmt.Sprint(args[3]))
		if err != nil {
			return ret, err
		}
	}
	return strings.Replace(fmt.Sprint(args[0]), fmt.Sprint(args[1]), fmt.Sprint(args[2]), n), err
}

// RegexpReplace regexp.ReplaceAllString
// args 1: regexp.Compile(expr string)
// args 0, 2: regexp.ReplaceAllString(src, repl string）
func RegexpReplace(args ...interface{}) (ret string, err error) {
	if len(args) < 3 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	re, err := regexp.Compile(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}
	ret = re.ReplaceAllString(fmt.Sprint(args[0]), fmt.Sprint(args[2]))

	return ret, err
}

// RegexpRandomReplace through regular random data relpace
// args 0: value
// args 1: regexp.Compile(expr string)
// args 2: max uint  Maximum number of instances to generate for unbounded repeat expressions (e.g., ".*" and "{1,}")
// args 3: min uint  Minimum number of instances to generate for unbounded repeat expressions (e.g., ".*")
func RegexpRandomReplace(args ...interface{}) (ret string, err error) {
	if len(args) < 4 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	re, err := regexp.Compile(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}
	min, err := strconv.ParseUint(fmt.Sprint(args[2]), 10, 32)
	if err != nil {
		return ret, fmt.Errorf(common.WrongArgValue)
	}

	max, err := strconv.ParseUint(fmt.Sprint(args[3]), 10, 32)
	if err != nil {
		return ret, fmt.Errorf(common.WrongArgValue)
	}

	substitute, err := Fake("regexp-rand", fmt.Sprint(args[1]), min, max)
	if err != nil {
		return ret, err
	}
	ret = re.ReplaceAllString(fmt.Sprint(args[0]), substitute)
	return ret, err
}

// Reverse string reverse
func Reverse(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	ret = fmt.Sprint(args[0])
	runes := []rune(ret)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	ret = string(runes)
	return ret, err
}

// ToUpper strings.ToUpper
func ToUpper(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	return strings.ToUpper(fmt.Sprint(args[0])), err
}

// ToLower strings.ToLower
func ToLower(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	return strings.ToLower(fmt.Sprint(args[0])), err
}

// Const replace string with const
// args 0: value
// args 1: const mask string
func Const(args ...interface{}) (ret string, err error) {
	if len(args) < 2 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	return fmt.Sprint(args[1]), err
}

// Number2Const replace all number to 9
// args 0: value
// args 1: const mask string
func Number2Const(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	var c = "9"
	if len(args) > 1 {
		c = fmt.Sprint(args[1])
	}

	re := regexp.MustCompile(`[0-9]`)
	return re.ReplaceAllString(fmt.Sprint(args[0]), c), err
}

// Char2Const replace [a-zA-Z] to N
// args 0: value
// args 1: const mask string
func Char2Const(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	var c = "N"
	if len(args) > 1 {
		c = fmt.Sprint(args[1])
	}

	re := regexp.MustCompile(`[a-zA-Z]`)
	return re.ReplaceAllString(fmt.Sprint(args[0]), c), err
}

// SmokeCharLeft mask left before specify char
// e.g., ***@example.com
// arg0 value
// arg1 char like "@" 、"&"、 "."
// arg2 replace char like '*','#'
func SmokeCharLeft(args ...interface{}) (ret string, err error) {
	if len(args) < 3 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	n := strings.IndexAny(fmt.Sprint(args[0]), fmt.Sprint(args[1]))
	if n > 0 {
		return string(bytes.Repeat([]byte(fmt.Sprint(args[2])), n)) + fmt.Sprint(args[0])[n:], err
	} else {
		return fmt.Sprint(args[0]), err
	}
}

// SmokeCharRight mask right after specify char
// e.g., user@****
// arg0 value
// arg1 char like "@" 、"&"、 "."
// arg2 replace char like '*','#'
func SmokeCharRight(args ...interface{}) (ret string, err error) {
	if len(args) < 3 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	n := strings.IndexAny(fmt.Sprint(args[0]), fmt.Sprint(args[1]))
	if n > -1 && n < len(fmt.Sprint(args[0])) {
		return fmt.Sprint(args[0])[:n+1] + string(bytes.Repeat([]byte(fmt.Sprint(args[2])), len(fmt.Sprint(args[0]))-n-1)), err
	} else {
		return fmt.Sprint(args[0]), err
	}
}

// https://help.aliyun.com/document_detail/150101.html?spm=a2c4g.11186623.6.595.243a5a787EiWOe
// NumberFloor...
// eg -12.78->-12、4856->4000
// arg1: value
// arg2: floor num
func NumberFloor(args ...interface{}) (ret string, err error) {
	if len(args) < 2 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	value, err := strconv.ParseFloat(fmt.Sprint(args[0]), 64)
	if err != nil {
		return "", err
	}
	n, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}
	if value > 0 {
		value = math.Floor(value)
	} else {
		value = math.Ceil(value)
	}
	ret = strconv.FormatFloat(value, 'f', -1, 64)
	if n > len(ret)-1 {
		return "0", err
	} else {
		return ret[:len(ret)-n] + string(bytes.Repeat([]byte("0"), n)), err
	}
}

// DateRound ...
// arg1: date
// arg2: dateFormat,
// accuracy: accuracy second, minute, hour(default), day, month, year
// e.g., 2021-07-23 17:00:00、2021-01-01 00:00:00
// https://docs.oracle.com/javase/6/docs/api/java/text/SimpleDateFormat.html#rfc822timezone
func DateRound(args ...interface{}) (ret string, err error) {
	var accuracy time.Duration
	if len(args) < 2 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	dateFormat := "YYYY-MM-DD HH:mm:ss"
	if len(args) > 2 {
		dateFormat = fmt.Sprint(args[2])
	}

	formatDate, err := gostradamus.Parse(fmt.Sprint(args[0]), dateFormat)
	if err != nil {
		return ret, err
	}
	switch fmt.Sprint(args[1]) {
	case "second":
		accuracy = time.Second
	case "minute":
		accuracy = time.Minute
	case "hour":
		accuracy = time.Hour
	case "day":
		accuracy = time.Hour * 24
	case "month":
		return gostradamus.NewDateTime(formatDate.Year(), formatDate.Month(), 1, 0, 0, 0, 0, formatDate.Timezone()).Format(fmt.Sprint(args[2])), err
	case "year":
		return gostradamus.NewDateTime(formatDate.Year(), 1, 1, 0, 0, 0, 0, formatDate.Timezone()).Format(fmt.Sprint(args[2])), err
	default:
		accuracy = time.Hour
	}
	return gostradamus.DateTimeFromTime(formatDate.Time().Round(accuracy)).Format(dateFormat), err
}

// DateFormat convert date format
// arg1: date
// arg2: old date format,
// arg2: new date format,
func DateFormat(args ...interface{}) (ret string, err error) {
	if len(args) < 2 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	newDateFormat := "YYYY-MM-DD HH:mm:ss"
	if len(args) > 2 {
		newDateFormat = fmt.Sprint(args[2])
	}

	formatDate, err := gostradamus.Parse(fmt.Sprint(args[0]), fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}

	return gostradamus.DateTimeFromTime(formatDate.Time()).Format(newDateFormat), err

}

// LoopMoveLeft ...
// e.g., abcdefg left move 3 defgabc
// arg1: value
// arg2: index
func LoopMoveLeft(args ...interface{}) (ret string, err error) {
	if len(args) < 2 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	n, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}
	n = n % len(fmt.Sprint(args[0]))
	return fmt.Sprint(args[0])[n:] + fmt.Sprint(args[0])[:n], err
}

// LoopMoveRight ...
// e.g., abcdefg left move 3 efgabcd
// arg1: value
// arg2: index
func LoopMoveRight(args ...interface{}) (ret string, err error) {
	if len(args) < 2 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	n, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return ret, err
	}
	valueLength := len(fmt.Sprint(args[0]))
	n = valueLength - n%valueLength
	return fmt.Sprint(args[0])[n:] + fmt.Sprint(args[0])[:n], err
}

// TruncateLeft truncate left n characters
// arg 0: value
// arg 1: index
func TruncateLeft(args ...interface{}) (ret string, err error) {
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
		return string(runes[len(runes)-n:]), nil
	}
}

func TruncateRight(args ...interface{}) (ret string, err error) {
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
		return string(runes[:n]), nil
	}
}

// Abbreviate english words abbreviate
// https://github.com/dnnrly/abbreviate
// strategy-limited => stg-ltd
func Abbreviate(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	value := fmt.Sprint(args[0])
	matcher := data.Abbreviations["en-us"]["common"]
	ret = domain.AsSeparated(matcher, value, "-", 1, true)
	return ret, err
}

// Initialism english words initialism
// hello world => hw
func Initialism(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	var b bytes.Buffer
	for _, s := range strings.Fields(fmt.Sprint(args[0])) {
		b.WriteRune([]rune(s)[0])
	}
	return b.String(), err
}

// Numeronym a number-based word.
// internationalization => i8n
func Numeronym(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	ret = fmt.Sprint(args[0])
	for i, s := range strings.Fields(ret) {
		runes := []rune(s)
		l := len(runes)
		if l > 2 {
			if i == 0 {
				ret = fmt.Sprintf("%c%d%c", runes[0], l-2, runes[l-1])
			} else {
				ret += fmt.Sprintf(" %c%d%c", runes[0], l-2, runes[l-1])
			}
		}
	}

	return ret, err
}
