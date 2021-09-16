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
)

func TestFF1(t *testing.T) {
	fmt.Println(FF1(123, 10))
	fmt.Println(FF1("123", 10))
	fmt.Println(FF1("A0", 16))
	fmt.Println(FF1("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01233456789", 62)) // radix max size 62 [0-9A-Za-z]
	ret, err := FF1(123, 10, "048e82ab0ab7c27afd1c76b7d39e4a1a62a21ef663522811790c88fe4e890db5", "ded666507c90c790")
	if err != nil {
		t.Errorf(err.Error())
	}
	if ret != "603" {
		t.Errorf("wrong result: %s", ret)
	}
}

func TestFF3(t *testing.T) {
	fmt.Println(FF3(123, 10))
	fmt.Println(FF3("123", 10))
	fmt.Println(FF3("A0", 16))
	ret, err := FF3(123, 10, "048e82ab0ab7c27afd1c76b7d39e4a1a62a21ef663522811790c88fe4e890db5", "ded666507c90c790")
	if err != nil {
		t.Errorf(err.Error())
	}
	if ret != "103" {
		t.Errorf("wrong result: %s", ret)
	}
}

func ExampleBase64() {
	fmt.Println(Base64("abc"))
	fmt.Println(Base64(123))
	fmt.Println(Base64("123"))
	// Output:
	// YWJj <nil>
	// MTIz <nil>
	// MTIz <nil>
}

func ExampleDES() {
	fmt.Println(DES("hello world"))
	fmt.Println(DES("hello world", "asdfghjk"))
	// Output:
	// jZLr1ir1An0IQc30XbLL3A== <nil>
	// hf5Kqc1nUS++YuxhzQeCIw== <nil>
}

func ExampleAES() {
	fmt.Println(AES("hello world"))
	fmt.Println(AES("hello world", "asdfghjk12345678d18nd18n"))
	// Output:
	// F5QMumnZOlCchKi2nu99rA== <nil>
	// FfX2LiivVQbq+w9Kat0Z3w== <nil>
}

func ExampleAESCTR() {
	fmt.Println(AESCTR("hello world"))
	fmt.Println(AESCTR("hello world", "asdfghjk12345678d18nd18n"))
	// Output:
	// 0URdmXIvWjcZS3U= <nil>
	// 6JXTR6DHyv9iRRw= <nil>

}

func ExampleTDEA() {
	fmt.Println(TDEA("hello world"))
	fmt.Println(TDEA("hello world", "asdfghjk12345678d18nd18n"))
	// Output:
	// F0HMxhk+uKKBSlR1IAyt+Q== <nil>
	// lrorDbcC2s92Pn6TYeOl5A== <nil>
}

func TestRSA(t *testing.T) {
	if ret, err := RSA("hello d18n"); err != nil {
		t.Error(err)
	} else {
		fmt.Println(ret)
	}
}

func TestECC(t *testing.T) {
	if ret, err := ECC("hello d18n"); err != nil {
		t.Error(err)
	} else {
		fmt.Println(ret)
	}
}

func TestSM2(t *testing.T) {
	crypttext, err := SM2("hello world")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(crypttext)
}

func ExampleSM3() {
	fmt.Println(SM3("hello world"))
	// Output:
	// 44f0061e69fa6fdfc290c494654a05dc0c053da7e5c52b84ef93a9d67d3fff88 <nil>
}

func ExampleSM4() {
	fmt.Println(SM4("hello world"))
	// Output:
	// b9b1742de155fe5720c0b8b1b95e3134 <nil>
}
