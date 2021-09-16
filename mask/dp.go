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
	"strconv"

	"d18n/common"

	"github.com/google/differential-privacy/go/noise"
)

// DifferentialPrivacy Laplace Float64
// DifferentialPrivacy Laplace Int64

// LaplaceDPFloat64 differential privacy masking based on laplace
// arg 0: value
// arg 1: l0sensitivity
// arg 2: lInfSensitivity
// arg 3: epsilon
// arg 4: delta
func LaplaceDPFloat64(args ...interface{}) (ret string, err error) {
	if len(args) < 5 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	value, err := strconv.ParseFloat(fmt.Sprint(args[0]), 64)
	if err != nil {
		return "", err
	}
	l0sensitivity, err := strconv.ParseInt(fmt.Sprint(args[1]), 10, 64)
	if err != nil {
		return "", err
	}
	lInfSensitivity, err := strconv.ParseFloat(fmt.Sprint(args[2]), 64)
	if err != nil {
		return "", err
	}
	epsilon, err := strconv.ParseFloat(fmt.Sprint(args[3]), 64)
	if err != nil {
		return "", err
	}

	delta, err := strconv.ParseFloat(fmt.Sprint(args[4]), 64)
	if err != nil {
		return "", err
	}

	ret = strconv.FormatFloat(noise.Laplace().AddNoiseFloat64(value, l0sensitivity, lInfSensitivity, epsilon, delta), 'f', -1, 64)
	return ret, err
}

// LaplaceDPInt64 differential privacy masking based on laplace
// arg 0: value
// arg 1: l0sensitivity
// arg 2: lInfSensitivity
// arg 3: epsilon
// arg 4: delta
func LaplaceDPInt64(args ...interface{}) (ret string, err error) {
	if len(args) < 5 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}
	value, err := strconv.ParseInt(fmt.Sprint(args[0]), 10, 64)
	if err != nil {
		return "", err
	}
	l0sensitivity, err := strconv.ParseInt(fmt.Sprint(args[1]), 10, 64)
	if err != nil {
		return "", err
	}
	lInfSensitivity, err := strconv.ParseInt(fmt.Sprint(args[2]), 10, 64)
	if err != nil {
		return "", err
	}
	epsilon, err := strconv.ParseFloat(fmt.Sprint(args[3]), 64)
	if err != nil {
		return "", err
	}

	delta, err := strconv.ParseFloat(fmt.Sprint(args[4]), 64)
	if err != nil {
		return "", err
	}
	a := noise.Gaussian().AddNoiseInt64(value, l0sensitivity, lInfSensitivity, epsilon, delta)
	ret = strconv.FormatInt(a, 10)
	return ret, err
}
