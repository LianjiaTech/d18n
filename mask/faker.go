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
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"path/filepath"
	"regexp/syntax"
	"strconv"
	"strings"
	"time"

	"github.com/LianjiaTech/d18n/common"

	"github.com/brianvoe/gofakeit/v6"
	regen "github.com/zach-klippenstein/goregen"
)

var faker *gofakeit.Faker

type district map[string][]string
type city map[string]district
type province map[string]city

var fakeAddressCorpus map[string]province

type personalName struct {
	FamilyNames []string `json:"family_names"`
	MiddleNames []string `json:"middle_names"`
	FirstNames  []string `json:"first_names"`
}

var fakeNameCorpus map[string]personalName

func InitFaker(seed int64) error {
	faker = gofakeit.New(seed) // or NewCrypto() to use crypto/rand
	return initFakeCorpus()
}

func initFakeCorpus() error {
	files, err := corpusFS.ReadDir("corpus")
	if err != nil {
		return err
	}

	fakeAddressCorpus = make(map[string]province)
	fakeNameCorpus = make(map[string]personalName)

	for _, file := range files {
		// load address corpus
		if strings.HasPrefix(file.Name(), "address.") {
			buf, err := corpusFS.ReadFile("corpus/" + file.Name())
			country := strings.ToLower(strings.TrimLeft(filepath.Ext(file.Name()), "."))
			var addr province
			if err != nil {
				return err
			}
			err = json.Unmarshal(buf, &addr)
			if err != nil {
				return err
			}
			fakeAddressCorpus[country] = addr
		}

		// load name corpus
		if strings.HasPrefix(file.Name(), "name.") {
			buf, err := corpusFS.ReadFile("corpus/" + file.Name())
			country := strings.ToLower(strings.TrimLeft(filepath.Ext(file.Name()), "."))
			var name personalName
			if err != nil {
				return err
			}
			err = json.Unmarshal(buf, &name)
			if err != nil {
				return err
			}
			fakeNameCorpus[country] = name
		}
	}
	return err
}

// Fake generate different type of fake data
func Fake(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	switch strings.ToLower(fmt.Sprint(args[0])) {
	case "address":
		var country, level string
		if len(args) > 1 {
			country = fmt.Sprint(args[1])
		} else {
			country = "zh_CN"
		}

		if len(args) > 2 {
			level = fmt.Sprint(args[2])
		}
		return fakeAddress(country, level)
	case "license-plate":
		var country string
		if len(args) > 1 {
			country = strings.ToLower(fmt.Sprint(args[1]))
		}
		switch country {
		case "zh_CN":
			return fakeRegexRandomData("^(([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z](([0-9]{5}[DF])|([DF]([A-HJ-NP-Z0-9])[0-9]{4})))|([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z][A-HJ-NP-Z0-9]{4}[A-HJ-NP-Z0-9挂学警港澳使领]))$", 1, 1)
		default:
			return fakeRegexRandomData("^(([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z](([0-9]{5}[DF])|([DF]([A-HJ-NP-Z0-9])[0-9]{4})))|([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z][A-HJ-NP-Z0-9]{4}[A-HJ-NP-Z0-9挂学警港澳使领]))$", 1, 1)
		}
	case "name":
		if len(args) > 1 {
			return fakeNameByCountry(fmt.Sprint(args[1]))
		} else {
			return faker.Name(), nil
		}
	case "password":
		// args[1] fake type "password"
		// args[2] password example like 'Ab1 ,' stand for upper, lower, space, special
		// args[3] password length
		var lower, upper, numeric, special, space bool
		// default length
		pwdLen := 16
		switch len(args) {
		case 1:
			// default policy and length
			return faker.Password(lower, upper, numeric, special, space, pwdLen), nil
		case 3:
			// user specified policy and length
			pwdLen, err = strconv.Atoi(fmt.Sprint(args[2]))
			if err != nil {
				return ret, fmt.Errorf(common.WrongArgValue)
			}
		}

		// password example, could be empty string
		example := fmt.Sprint(args[1])

		for i := 0; i < len(example); i++ {
			switch {
			case 64 < example[i] && example[i] < 91:
				upper = true
			case 96 < example[i] && example[i] < 123:
				lower = true
			case 47 < example[i] && example[i] < 58:
				numeric = true
			case example[i] == 32:
				space = true
			default:
				special = true
			}
		}
		return faker.Password(lower, upper, numeric, special, space, pwdLen), nil
	case "email": // gen_rnd_email()
		return faker.Email(), nil
	case "ssn": // American Social Security number
		return faker.SSN(), nil
	case "birthday":
		daysAgo := faker.Number(0, 365*100) // last 100 years random date
		return time.Now().Add(-time.Duration(daysAgo) * 24 * time.Hour).Format("2006-01-02"), nil
	case "cc", "creditcard":
		return faker.CreditCard().Number, nil
	case "url":
		return faker.URL(), nil
	case "phone":
		var country string
		if len(args) > 1 {
			country = strings.ToLower(fmt.Sprint(args[1]))
		}
		switch country {
		case "zh_CN":
			return fakeRegexRandomData("1(?:3\\d{3}|5[^4\\D]\\d{2}|8\\d{3}|7(?:[0-35-9]\\d{2}|4(?:0\\d|1[0-2]|9\\d))|9[0-35-9]\\d{2}|6[2567]\\d{2}|4[579]\\d{2})\\d{6}$\n", 5, 10)
		default:
			return faker.Phone(), nil
		}
	case "uuid":
		return faker.UUID(), nil
	case "ip", "ipv4":
		return faker.IPv4Address(), nil
	case "ipv6":
		return faker.IPv6Address(), nil
	case "number":
		var min = math.MinInt32
		var max = math.MaxInt32
		if len(args) > 2 {
			min, err = strconv.Atoi(fmt.Sprint(args[1]))
			if err != nil {
				return ret, err
			}
			max, err = strconv.Atoi(fmt.Sprint(args[2]))
			if err != nil {
				return ret, err
			}
			if min >= max {
				return ret, fmt.Errorf("min should larger than max")
			}
		}
		return fmt.Sprint(faker.Number(min, max)), nil
	case "uscc": // China unified social credit code
		return fakerUSCC()
	case "regexp-rand":
		min, err := strconv.ParseUint(fmt.Sprint(args[2]), 10, 32)
		if err != nil {
			return ret, err
		}
		max, err := strconv.ParseUint(fmt.Sprint(args[3]), 10, 32)
		if err != nil {
			return ret, err
		}
		return fakeRegexRandomData(fmt.Sprint(args[1]), min, max)
	}
	return ret, nil
}

// fakeAddress generate fake address
func fakeAddress(country, level string) (ret string, err error) {
	// check country name
	country = strings.ToLower(country)
	if _, ok := fakeAddressCorpus[country]; !ok {
		return ret, fmt.Errorf(common.WrongArgValue)
	}

	var street string
	// map is disordered
	for province, cities := range fakeAddressCorpus[country] {
		if level == "province" {
			return province, err
		}
		for city, districts := range cities {
			if level == "city" {
				return province + city, err
			}
			for district, streets := range districts {
				if level == "district" {
					return province + city + district, err
				}
				streetsLength := len(streets)
				if streetsLength == 0 {
					street = ""
				} else {
					street = streets[rand.Intn(streetsLength)]
				}
				return province + city + district + street, err
			}
		}
	}
	return ret, err
}

// fakeRegexRandomData generate fake data based on regular expression
func fakeRegexRandomData(pattern string, min, max uint64) (ret string, err error) {
	if min > max {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	if max > 4000 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	if generator, err := regen.NewGenerator(pattern, &regen.GeneratorArgs{
		RngSource:               rand.NewSource(time.Now().UnixNano()),
		MaxUnboundedRepeatCount: uint(max),
		MinUnboundedRepeatCount: uint(min),
		Flags:                   syntax.Perl, // regen use Posix by default, but regexp.MustCompile use Perl by default
	}); err != nil {
		return "", err
	} else {
		return generator.Generate(), nil
	}
}

// fakeNameByCountry generate fake personal name
func fakeNameByCountry(country string) (ret string, err error) {
	// check country name
	country = strings.ToLower(country)
	if _, ok := fakeNameCorpus[country]; !ok {
		return ret, fmt.Errorf(common.WrongArgValue)
	}

	var familyName, middleName, firstName string
	names := fakeNameCorpus[country]
	familyNameLen := len(names.FamilyNames)
	middleNameLen := len(names.MiddleNames)
	firstNameLen := len(names.FirstNames)

	if familyNameLen == 0 {
		familyName = ""
	} else {
		familyName = names.FamilyNames[rand.Intn(familyNameLen)]
	}
	if middleNameLen == 0 {
		middleName = ""
	} else {
		middleName = names.MiddleNames[rand.Intn(middleNameLen)]
	}
	if firstNameLen == 0 {
		firstName = ""
	} else {
		firstName = names.FirstNames[rand.Intn(firstNameLen)]
	}

	return familyName + middleName + firstName, nil
}

// USCC code table
var usccCode = []struct {
	Department string `json:"Department"`
	Category   []int  `json:"Category"`
}{
	{Department: "1", Category: []int{1, 2, 3, 9}},
	{Department: "2", Category: []int{1, 9}},
	{Department: "3", Category: []int{1, 2, 3, 4, 5, 9}},
	{Department: "4", Category: []int{1, 9}},
	{Department: "5", Category: []int{1, 2, 3, 9}},
	{Department: "6", Category: []int{1, 2, 9}},
	{Department: "7", Category: []int{1, 2, 9}},
	{Department: "8", Category: []int{1, 9}},
	{Department: "9", Category: []int{1, 2, 3}},
	{Department: "A", Category: []int{1, 9}},
	{Department: "N", Category: []int{1, 2, 3, 9}},
	{Department: "Y", Category: []int{1}},
}

// fakerUSCC Chinese unified social credit code
func fakerUSCC() (ret string, err error) {
	uscc, err := fakeRegexRandomData("^\\d{6}[0-9A-HJ-NPQRTUWXY]{10}$", 5, 10)
	if err != nil {
		return ret, err
	}
	code := usccCode[rand.Intn(len(usccCode))]
	return fmt.Sprintf("%s%d", code.Department, code.Category[rand.Intn(len(code.Category))]) + uscc, err
}
