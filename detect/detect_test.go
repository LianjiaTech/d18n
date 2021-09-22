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

package detect

import (
	"d18n/common"
	"testing"

	"github.com/kr/pretty"
)

func init() {
	ParseSensitiveConfig()
}

func TestParseSensitiveConfig(t *testing.T) {
	err := ParseSensitiveConfig()
	if err != nil {
		t.Error(err.Error())
	}
	pretty.Println(sensitiveConfig["mac"])
}

func TestcheckFileHeader(t *testing.T) {
	var cases = []string{
		"name",
		"birthday",
		"ssn", "security_sequence_number",
		"sex", "gender",
		"address",
		"location",
		"city", "country", "gps",
		"email", "mail",
		"password", "passwd",
		"imei",
		"ip",
		"mac",
		"iban",
		"vin",
		"license_plate", "car_plate",
		"uscc", "uniform_social_credit_code", "business_license",
		"phone", "telephone", "phone_num",
		"passport", "passport_num",
		"postal_code",
		"company",
		"company_name",
	}
	var headers []common.HeaderColumn
	for _, v := range cases {
		headers = append(headers, common.HeaderColumn{Name: v})
	}
	checkFileHeader(headers)
	for k, v := range detectStatus.Columns {
		if len(v) > 1 || len(v) == 0 {
			t.Error("get wrong types return", k, v)
		}
	}
}
func TestCheckValue(t *testing.T) {
	var cases = []map[string][]string{
		{
			"phone":   []string{"13123385678"},
			"email":   []string{"user@example.com"},
			"address": []string{"湖北省荆州市"},
		},
	}
	for _, c := range cases {
		for k := range c {
			for _, v := range c[k] {
				types := checkValue(v)
				if len(types) == 0 {
					t.Error("get wrong types return", v, types)
					return
				}
				if len(types) > 0 && types[0] != k {
					t.Error("wrong type", v, k, types[0])
				}
			}
		}
	}
}

func TestDetectFromFile(t *testing.T) {
	orgCfg := common.Cfg
	files := []string{
		"actor.csv",
		"actor.xlsx",
		"actor.html",
		"actor.json",
		"actor.sql",
		"actor.txt",
		"actor.tsv",
	}

	for _, f := range files {
		common.Cfg.File = common.TestPath + "/test/" + f
		err := DetectFile()
		if err != nil {
			t.Error(f, err.Error())
		}
	}
	common.Cfg = orgCfg
}
