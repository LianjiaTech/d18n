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
	"d18n/common"
	"fmt"
	"testing"
	"time"
)

func TestFakeName(t *testing.T) {
	cases := []string{
		"name",
		"name",
		"address",
		"address",
		"license-plate",
		"license-plate",
		"email",
		"email",
		"ssn",
		"ssn",
		"creditcard",
		"creditcard",
		"url",
		"url",
		"number",
		"number",
		"uuid",
		"uuid",
		"ip",
		"ip",
		"ipv6",
		"ipv6",
	}
	var last string
	for _, c := range cases {
		current, _ := Fake(c)
		if last == current {
			t.Error("generate same val", c)
		}
		fmt.Println("type:", c, "current:", current, "last:", last)
		last = current
	}
}

func TestFakeChinaAddress(t *testing.T) {
	levels := []string{
		"province",
		"city",
		"district",
		"street",
		"",
	}
	for _, l := range levels {
		address, err := fakeAddress("zh_CN", l)
		if err != nil {
			t.Error(err.Error())
		}
		fmt.Println(l, ":", address)
	}
}

func TestFakeRegexRandomData(t *testing.T) {
	patterns := []string{
		"^1[3-9][\\d]{9}$", //
		"[a-zA-Z]{6}[a-zA-Z0-9]{3}@[a-zA-Z.]+\\.[a-zA-Z]+",
		"^([1-9]{1})(\\d{14}|\\d{18})$",
		"^[1-9]\\d{7}((0\\d)|(1[0-2]))(([0-2]\\d)|3[0-1])\\d{3}$|^[1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([0-2]\\d)|3[0-1])\\d{3}([0-9]|X)$",
		// https://cloud.tencent.com/developer/article/1361209
		"^(([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z](([0-9]{5}[DF])|([DF]([A-HJ-NP-Z0-9])[0-9]{4})))|([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z][A-HJ-NP-Z0-9]{4}[A-HJ-NP-Z0-9挂学警港澳使领]))$",
		"^C[0-9A-HJ-NP-Z]\\d{7}$", // 大陆港澳台
		"^\\d{6}[0-9A-HJ-NPQRTUWXY]{10}$",
	}
	common.Cfg.RandSeed = time.Now().Unix()
	for _, pattern := range patterns {
		context, err := fakeRegexRandomData(pattern, 5, 10)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(context)
	}
}

func TestFakeChineseName(t *testing.T) {
	t.Log(fakeNameByCountry("zh_CN"))
	t.Log(fakeNameByCountry("zh_CN"))
}

func TestFake(t *testing.T) {
	fakeData, err := Fake("name", "zh_CN")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(fakeData)
	fakeData, err = Fake("uscc")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(fakeData)
	fakeData, err = Fake("lpn", "zh_CN")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(fakeData)
	fakeData, err = Fake("phone", "zh_CN")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(fakeData)

	// `^([A-Za-z0-9_\-\.])+\@[A-Za-z0-9_\-\.]+\.[A-Za-z]{2,4}$`
	fakeData, err = Fake("regexp-rand", `^([A-Za-z0-9_\-\.])+\@[A-Za-z0-9_\-\.]+\.[A-Za-z]{2,4}$`, 1, 1)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(fakeData)
}

func Example_fakePassword() {
	orgCfg := common.Cfg
	common.Cfg.RandSeed = 1989
	var fakeData string
	var err error
	// default policy and length
	fakeData, err = Fake("password")
	fmt.Println(fakeData, err)

	// default length
	fakeData, _ = Fake("password", "1")
	fmt.Println(fakeData, err)

	// all specified
	fakeData, _ = Fake("password", "aA", 12)
	fmt.Println(fakeData, err)

	common.Cfg = orgCfg
	// Output:
	// cq8d1uhfy4l4gk4p <nil>
	// 0023369360722229 <nil>
	// nxCVAJmQUNws <nil>
}

func TestFakerUSCC(t *testing.T) {
	fakeDate, err := fakerUSCC()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(fakeDate)
}
