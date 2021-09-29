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

func TestParseSensitiveConfig(t *testing.T) {
	// new test detect struct
	d, err := NewDetectStruct(common.Cfg)
	if err != nil {
		t.Errorf(err.Error())
	}
	d.parseConfig()

	pretty.Println(d.SensitiveConfig["mac"])
}

func TestCheckHeader(t *testing.T) {

	// new test detect struct
	d, err := NewDetectStruct(common.Cfg)
	if err != nil {
		t.Errorf(err.Error())
	}
	d.parseConfig()

	// only detect header, reset all columns info
	d.Status.Columns = make(map[string][]string)

	// cases
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
	d.Status.Header = headers

	d.checkHeader()
	for k, v := range d.Status.Columns {
		if len(v) > 1 || len(v) == 0 {
			t.Error("get wrong types return", k, v)
		}
	}
}
func TestCheckValue(t *testing.T) {
	// new test detect struct
	d, err := NewDetectStruct(common.Cfg)
	if err != nil {
		t.Errorf(err.Error())
	}
	d.parseConfig()

	// cases
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
				types := d.checkValue(v)
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

func TestDetectQuery(t *testing.T) {
	orgCfg := common.Cfg

	common.Cfg.Query = "select * from address limit 10"
	common.Cfg.Database = "sakila"

	d, err := NewDetectStruct(common.Cfg)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = d.DetectQuery()
	if err != nil {
		t.Error(err.Error())
	}

	common.Cfg = orgCfg
}

func TestDetectFile(t *testing.T) {
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
		d, _ := NewDetectStruct(common.Cfg)
		err := d.DetectFile()
		if err != nil {
			t.Error(f, err.Error())
		}
	}
	common.Cfg = orgCfg
}
