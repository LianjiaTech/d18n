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
)

func ExamplePhone() {
	fmt.Println(Phone("13000000123"))
	fmt.Println(Phone(13000000123))
	// Output:
	// 130*****123 <nil>
	// 130*****123 <nil>
}

func ExampleUSCC() {
	fmt.Println(USCC("71797173LM37QP0D4H"))
	// Output:
	// 717971********0D4H <nil>
}

func ExampleAge() {
	fmt.Println(Age(39))
	fmt.Println(Age(23))

	// Output:
	// 30 <nil>
	// 20 <nil>
}

func ExampleBirthday() {
	fmt.Println(Birthday("20200534"))
	fmt.Println(Birthday("2020-05-34"))
	// Output:
	// NNNNNNNN <nil>
	// NNNN-NN-NN <nil>
}

func ExamplePassword() {
	fmt.Println(Password("asfa@12323ssda"))
	fmt.Println(Password("asfa@123"))
	// Output:
	// ********* <nil>
	// ********* <nil>
}

func ExampleCreditCard() {
	fmt.Println(CreditCard("6227612145830440"))
	// Output:
	// 62276121****0440 <nil>
}

func ExampleIP() {
	fmt.Println(IP("192.168.0.1"))
	// Output:
	// 127.0.0.1 <nil>
}

func ExampleSalary() {
	fmt.Println(Salary(1300))
	fmt.Println(Salary(500))
	// Output:
	// 1000 <nil>
	// 0 <nil>
}

func ExampleUsername() {
	fmt.Println(Username("张三"))
	fmt.Println(Username("王二狗"))
	fmt.Println(Username("Dave Li"))
	// Output:
	// 张* <nil>
	// 王** <nil>
	// D****** <nil>
}

func ExampleLicensePlate() {
	fmt.Println(LicensePlate("鄂D71D44"))
	// Output:
	// 鄂*****4 <nil>
}

func ExampleDomain() {
	fmt.Println(Domain("example.com"))
	// Output:
	// e*********m <nil>
}

func ExampleMail() {
	fmt.Println(Mail("zhangsan001@d18n.com"))
	// Output:
	// z**********@*******m <nil>
}

func ExampleOrganizationCode() {
	fmt.Println(OrganizationCode("100000439"))
	// Output:
	// 100***439 <nil>
}

func ExamplePersonalID() {
	fmt.Println(PersonalID("110223700003697"))
	// Output:
	// 110223*****3697 <nil>
}
