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
	"testing"

	"d18n/common"
)

func init() {
	common.InitTestEnv()
	ParseCipherConfig(common.Cfg.Cipher)
	InitMaskCorpus(common.Cfg.RandSeed)
}

func TestLaplaceDPFloat64(t *testing.T) {
	salary := map[string]float64{
		"张三":  3200.53,
		"李四":  5700.45,
		"王五":  3220.65,
		"张二狗": 8210.73,
		"李蛋":  9250.20,
	}
	salaryNoise := map[string]string{}
	for k, v := range salary {
		ret, err := LaplaceDPFloat64(v, 100, 1, 1, 0)
		if err != nil {
			t.Error(err)
		}
		salaryNoise[k] = ret
	}
	fmt.Println(salary)
	fmt.Println(salaryNoise)
}

func TestLaplaceDPInt64(t *testing.T) {
	salary := map[string]float64{
		"张三":  3200,
		"李四":  5700,
		"王五":  3220,
		"张二狗": 8210,
		"李蛋":  9250,
	}
	salaryNoise := map[string]string{}
	for k, v := range salary {
		ret, err := LaplaceDPInt64(v, 100, 1, 1, 0.5)
		if err != nil {
			t.Error(err)
		}
		salaryNoise[k] = ret
	}
	fmt.Println(salaryNoise)
}
