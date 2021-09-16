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

// PII: Personal Identification Information

import (
	"fmt"
	"strings"

	"d18n/common"
)

// Phone phone default desensitize method
// e.g., 130*****123
func Phone(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	return ReserveMargin(args[0], 3, "*")
}

// Mail mail default desensitize method
// e.g., z****@***m
func Mail(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	s := strings.Split(fmt.Sprint(args[0]), "@")
	if len(s) == 2 {
		left, err := ReserveLeft(s[0], 1, "*")
		if err != nil {
			return ret, err
		}
		right, err := ReserveRight(s[1], 1, "*")
		if err != nil {
			return ret, err
		}
		ret = left + "@" + right
	}
	return ret, err
}

// Username username default desensitize method
// e.g., 王**
func Username(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	return ReserveLeft(args[0], 1, "*") // e.g., 张xx , Axx
}

// Domain domain default desensitize method
// e.g., s*******m
func Domain(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	return ReserveMargin(args[0], 1, "*")
}

// CreditCard credit card default desensitize method
func CreditCard(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	return SmokeInner(args[0], 8, 4, "*")
}

// PersonalID personal ID default desensitize method
func PersonalID(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	return SmokeInner(args[0], 6, 4, "*")
}

// Age age default desensitize method
func Age(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	return NumberFloor(args[0], 1)
}

// Salary salary default desensitize method
func Salary(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	return NumberFloor(args[0], 3)
}

// Birthday birthday default desensitize method
func Birthday(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	return Number2Const(args[0], "N")
}

// IP IP default desensitize method
func IP(args ...interface{}) (ret string, err error) {
	return "127.0.0.1", nil
}

// LicensePlate licensePlate default desensitize method
// 京*****
func LicensePlate(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	return ReserveMargin(args[0], 1, "*") //
}

// Password password default desensitize method
func Password(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	return "*********", nil
}

// USCC uscc default desensitize method
func USCC(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	return SmokeInner(args[0], 6, 4, "*")
}

// OrganizationCode organizationCode default desensitize method
func OrganizationCode(args ...interface{}) (ret string, err error) {
	if len(args) < 1 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	return ReserveMargin(args[0], 3, "*")
}
