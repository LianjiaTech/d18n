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
)

func ExampleConst() {
	fmt.Println(Const("123", "MASKED"))
	// Output:
	// MASKED <nil>
}

func ExampleSmoke() {
	fmt.Println(Smoke(123, "*"))
	fmt.Println(Smoke("123", "*"))
	fmt.Println(Smoke(123, "x"))
	// Output:
	// *** <nil>
	// *** <nil>
	// xxx <nil>
}

func ExampleSmokeLeft() {
	fmt.Println(SmokeLeft(123, 2, "*"))
	fmt.Println(SmokeLeft("123", 2, "•"))
	fmt.Println(SmokeLeft(123, 2, "x"))
	fmt.Println(SmokeLeft(123, 3, "x"))
	fmt.Println(SmokeLeft(123, 4, "x"))
	fmt.Println(SmokeLeft(123, 0, "x"))
	fmt.Println(SmokeLeft(123, -1, "x"))
	// Output:
	// **3 <nil>
	// ••3 <nil>
	// xx3 <nil>
	// xxx <nil>
	// xxx <nil>
	// 123 <nil>
	//  n should large than 0
}

func ExampleReserveLeft() {
	fmt.Println(ReserveLeft(123, 2, "*"))
	fmt.Println(ReserveLeft("123", 2, "*"))
	fmt.Println(ReserveLeft(123, 2, "x"))
	fmt.Println(ReserveLeft(123, 3, "x"))
	fmt.Println(ReserveLeft(123, 4, "x"))
	fmt.Println(ReserveLeft(123, 0, "x"))
	fmt.Println(ReserveLeft(123, -1, "x"))
	fmt.Println(ReserveLeft("张三", 1, "某"))
	// Output:
	// 12* <nil>
	// 12* <nil>
	// 12x <nil>
	// 123 <nil>
	// 123 <nil>
	// xxx <nil>
	//  n should large than 0
	// 张某 <nil>
}

func ExampleSmokeRight() {
	fmt.Println(SmokeRight(123, 2, "*"))
	fmt.Println(SmokeRight("123", 2, "*"))
	fmt.Println(SmokeRight(123, 2, "x"))
	fmt.Println(SmokeRight(123, 3, "x"))
	fmt.Println(SmokeRight(123, 4, "x"))
	fmt.Println(SmokeRight(123, 0, "x"))
	fmt.Println(SmokeRight(123, -1, "x"))
	// Output:
	// 1** <nil>
	// 1** <nil>
	// 1xx <nil>
	// xxx <nil>
	// xxx <nil>
	// 123 <nil>
	//  n should large than 0
}

func ExampleReserveRight() {
	fmt.Println(ReserveRight(123, 2, "*"))
	fmt.Println(ReserveRight("123", 2, "*"))
	fmt.Println(ReserveRight(123, 2, "x"))
	fmt.Println(ReserveRight(123, 3, "x"))
	fmt.Println(ReserveRight(123, 4, "x"))
	fmt.Println(ReserveRight(123, 0, "x"))
	fmt.Println(ReserveRight(123, -1, "x"))
	// Output:
	// *23 <nil>
	// *23 <nil>
	// x23 <nil>
	// 123 <nil>
	// 123 <nil>
	// xxx <nil>
	//  n should large than 0
}

func ExampleReserveMargin() {
	fmt.Println(ReserveMargin(123456, 2, "*"))
	fmt.Println(ReserveMargin(12, 1, "*"))
	fmt.Println(ReserveMargin("123456", 2, "*"))
	fmt.Println(ReserveMargin(123456, 2, "x"))
	fmt.Println(ReserveMargin(1234, 3, "x"))
	fmt.Println(ReserveMargin(123, 4, "x"))
	fmt.Println(ReserveMargin(123, 0, "x"))
	fmt.Println(ReserveMargin(123, -1, "x"))
	fmt.Println(ReserveMargin("王老五", 1, "某"))
	// Output:
	// 12**56 <nil>
	// ** <nil>
	// 12**56 <nil>
	// 12xx56 <nil>
	// xxxx <nil>
	// xxx <nil>
	// xxx <nil>
	//  n should large than 0
	// 王某五 <nil>
}

func ExampleSmokeMargin() {
	fmt.Println(SmokeMargin(123456, 2, "*"))
	fmt.Println(SmokeMargin("123456", 2, "*"))
	fmt.Println(SmokeMargin(123456, 2, "x"))
	fmt.Println(SmokeMargin(1234, 3, "x"))
	fmt.Println(SmokeMargin(123, 4, "x"))
	fmt.Println(SmokeMargin(123, 0, "x"))
	fmt.Println(SmokeMargin(123, -1, "x"))
	// Output:
	// **34** <nil>
	// **34** <nil>
	// xx34xx <nil>
	// xxxx <nil>
	// xxx <nil>
	// 123 <nil>
	//  n should large than 0
}

func ExampleSmokeOuter() {
	fmt.Println(SmokeOuter(123456, 2, 1, "*"))
	fmt.Println(SmokeOuter(123, 2, 1, "*"))
	fmt.Println(SmokeOuter("123456", 2, 1, "*"))
	fmt.Println(SmokeOuter(123456, 2, 1, "x"))
	fmt.Println(SmokeOuter(1234, 3, 1, "x"))
	fmt.Println(SmokeOuter(123, 1, 4, "x"))
	fmt.Println(SmokeOuter(123, 0, 0, "x"))
	fmt.Println(SmokeOuter(123, -1, 1, "x"))
	// Output:
	// **345* <nil>
	// *** <nil>
	// **345* <nil>
	// xx345x <nil>
	// xxxx <nil>
	// xxx <nil>
	// 123 <nil>
	//  n should large than 0
}

func ExampleReserveOuter() {
	fmt.Println(ReserveOuter(123456, 2, 1, "*"))
	fmt.Println(ReserveOuter(123, 2, 1, "*"))
	fmt.Println(ReserveOuter("123456", 2, 1, "*"))
	fmt.Println(ReserveOuter(123456, 2, 1, "x"))
	fmt.Println(ReserveOuter(1234, 3, 1, "x"))
	fmt.Println(ReserveOuter(123, 1, 4, "x"))
	fmt.Println(ReserveOuter(123, 0, 0, "x"))
	fmt.Println(ReserveOuter(123, -1, 1, "x"))
	// Output:
	// 12***6 <nil>
	// 123 <nil>
	// 12***6 <nil>
	// 12xxx6 <nil>
	// 1234 <nil>
	// xxx <nil>
	// xxx <nil>
	//  n should large than 0
}

func ExampleSmokeInner() {
	fmt.Println(SmokeInner(123456, 2, 1, "*"))
	fmt.Println(SmokeInner(123, 2, 1, "*"))
	fmt.Println(SmokeInner("123456", 2, 1, "*"))
	fmt.Println(SmokeInner(123456, 2, 1, "x"))
	fmt.Println(SmokeInner(1234, 3, 1, "x"))
	fmt.Println(SmokeInner(123, 1, 4, "x"))
	fmt.Println(SmokeInner(123, 0, 0, "x"))
	fmt.Println(SmokeInner(123, -1, 1, "x"))
	// Output:
	// 12***6 <nil>
	// 123 <nil>
	// 12***6 <nil>
	// 12xxx6 <nil>
	// 1234 <nil>
	// xxx <nil>
	// xxx <nil>
	//  n should large than 0
}

func ExampleReserveInner() {
	fmt.Println(ReserveInner(123456, 2, 1, "*"))
	fmt.Println(ReserveInner(123, 2, 1, "*"))
	fmt.Println(ReserveInner("123456", 2, 1, "*"))
	fmt.Println(ReserveInner(123456, 2, 1, "x"))
	fmt.Println(ReserveInner(1234, 3, 1, "x"))
	fmt.Println(ReserveInner(123, 1, 4, "x"))
	fmt.Println(ReserveInner(123, 0, 0, "x"))
	fmt.Println(ReserveInner(123, -1, 1, "x"))
	// Output:
	// **345* <nil>
	// *** <nil>
	// **345* <nil>
	// xx345x <nil>
	// xxxx <nil>
	// xxx <nil>
	// 123 <nil>
	//  n should large than 0
}

func ExampleReplace() {
	fmt.Println(Replace(123, "2", "*", -1))
	fmt.Println(Replace("123", "2", "*", -1))
	fmt.Println(Replace(123, "2", "x", -1))
	fmt.Println(Replace(123, "2", "x", -1))
	fmt.Println(Replace(12223, "2", "x", 2))
	fmt.Println(Replace(123, 1, 4, -1))
	// Output:
	// 1*3 <nil>
	// 1*3 <nil>
	// 1x3 <nil>
	// 1x3 <nil>
	// 1xx23 <nil>
	// 423 <nil>
}

func ExampleRegexpReplace() {
	fmt.Println(RegexpReplace("abcdef", "[bc]", "*"))
	// Output:
	// a**def <nil>
}

func ExampleReverse() {
	fmt.Println(Reverse("abc"))
	fmt.Println(Reverse(123))
	fmt.Println(Reverse("信息隐藏实验"))
	// Output:
	// cba <nil>
	// 321 <nil>
	// 验实藏隐息信 <nil>
}

func ExampleToUpper() {
	fmt.Println(ToUpper("aBc"))
	fmt.Println(ToUpper(123))
	fmt.Println(ToUpper("123"))
	// Output:
	// ABC <nil>
	// 123 <nil>
	// 123 <nil>
}

func ExampleToLower() {
	fmt.Println(ToLower("aBc"))
	fmt.Println(ToUpper(123))
	fmt.Println(ToUpper("123"))
	// Output:
	// abc <nil>
	// 123 <nil>
	// 123 <nil>
}

func ExampleNumber2Const() {
	fmt.Println(Number2Const("(+086)130-1234-123"))
	fmt.Println(Number2Const("(+086)130-1234-123", 0))
	// Output:
	// (+999)999-9999-999 <nil>
	// (+000)000-0000-000 <nil>
}

func ExampleChar2Const() {
	fmt.Println(Char2Const("abc-def"))
	fmt.Println(Char2Const("ABC-DEF", "A"))
	// Output:
	// NNN-NNN <nil>
	// AAA-AAA <nil>
}

func TestRegexpRandomReplace(t *testing.T) {
	orgCfg := common.Cfg
	ret, err := RegexpRandomReplace("13782430405", "^1[3-9][\\d]{9}$", 5, 10)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(ret)
	common.Cfg = orgCfg
}

func ExampleTruncateLeft() {
	fmt.Println(TruncateLeft("abcdef", 2))
	fmt.Println(TruncateLeft("abcdef", 20))
	fmt.Println(TruncateLeft("中文abc", 4))
	// Output:
	// ef <nil>
	// abcdef <nil>
	// 文abc <nil>
}

func ExampleTruncateRight() {
	fmt.Println(TruncateRight("abcdef", 2))
	fmt.Println(TruncateRight("abcdef", 20))
	fmt.Println(TruncateRight("中文abc", 2))
	// Output:
	// ab <nil>
	// abcdef <nil>
	// 中文 <nil>
}

func ExampleSmokeCharLeft() {
	fmt.Println(SmokeCharLeft("zhangsan123@example.com", "@", "*"))
	fmt.Println(SmokeCharLeft("zhangsan123&example.com", "&", "*"))
	fmt.Println(SmokeCharLeft("zhangsan123@example.com", ".", "*"))
	fmt.Println(SmokeCharLeft("@example.com", "@", "*"))
	fmt.Println(SmokeCharLeft("example.com@", "@", "*"))
	fmt.Println(SmokeCharLeft("@", "@", "*"))
	fmt.Println(SmokeCharLeft("", "@", "*"))
	// Output:
	// ***********@example.com <nil>
	// ***********&example.com <nil>
	// *******************.com <nil>
	// @example.com <nil>
	// ***********@ <nil>
	// @ <nil>
	//  <nil>
}

func ExampleSmokeCharRight() {
	fmt.Println(SmokeCharRight("zhangsan123@example.com", "@", "*"))
	fmt.Println(SmokeCharRight("zhangsan123&example.com", "&", "*"))
	fmt.Println(SmokeCharRight("zhangsan123@example.com", ".", "*"))
	fmt.Println(SmokeCharRight("@example.com", "@", "*"))
	fmt.Println(SmokeCharRight("example.com@", "@", "*"))
	fmt.Println(SmokeCharLeft("@", "@", "*"))
	fmt.Println(SmokeCharRight("", "@", "*"))
	// Output:
	// zhangsan123@*********** <nil>
	// zhangsan123&*********** <nil>
	// zhangsan123@example.*** <nil>
	// @*********** <nil>
	// example.com@ <nil>
	// @ <nil>
	//  <nil>
}

func ExampleDateRound() {

	fmt.Println(DateRound("2021-07-23 17:26:45", "hour"))
	fmt.Println(DateRound("2021-07-23 17:26:45", "second", "YYYY-MM-DD HH:mm:ss"))
	fmt.Println(DateRound("2021/07/23 17-26-45", "minute", "YYYY/MM/DD HH-mm-ss"))
	fmt.Println(DateRound("2021-07-23 17:26:45", "hour", "YYYY-MM-DD HH:mm:ss"))
	fmt.Println(DateRound("2021-07-23 17:26:45", "day", "YYYY-MM-DD HH:mm:ss"))
	fmt.Println(DateRound("2021-07-23 17:26:45", "month", "YYYY-MM-DD HH:mm:ss"))
	fmt.Println(DateRound("2021-07-23 17:26:45", "year", "YYYY-MM-DD HH:mm:ss"))
	fmt.Println(DateRound("2021@07@23 17@26@45", "xxxxx", "YYYY@MM@DD HH@mm@ss"))

	// Output:
	// 2021-07-23 17:00:00 <nil>
	// 2021-07-23 17:26:45 <nil>
	// 2021/07/23 17-27-00 <nil>
	// 2021-07-23 17:00:00 <nil>
	// 2021-07-24 00:00:00 <nil>
	// 2021-07-01 00:00:00 <nil>
	// 2021-01-01 00:00:00 <nil>
	// 2021@07@23 17@00@00 <nil>
}

func ExampleDateFormat() {
	fmt.Println(DateFormat("2021/07/23 17-26-45", "YYYY/MM/DD HH-mm-ss"))
	fmt.Println(DateFormat("2021-07-23 17:26:45", "YYYY-MM-DD HH:mm:ss", "YYYY/MM/DD HH-mm-ss"))
	// Output:
	// 2021-07-23 17:26:45 <nil>
	// 2021/07/23 17-26-45 <nil>
}

func ExampleLoopMoveLeft() {
	fmt.Println(LoopMoveLeft("abcdefg", 3))
	// Output:
	// defgabc <nil>
}

func ExampleLoopMoveRight() {
	fmt.Println(LoopMoveRight("abcdefg", 3))
	// Output:
	// efgabcd <nil>
}

func ExampleNumberFloor() {
	fmt.Println(NumberFloor(123456.789, 3))
	fmt.Println(NumberFloor(43, 1))
	fmt.Println(NumberFloor(3654, 3))
	fmt.Println(NumberFloor(56.789, 3))
	// Output:
	// 123000 <nil>
	// 40 <nil>
	// 3000 <nil>
	// 0 <nil>
}

func ExampleAbbreviate() {
	fmt.Println(Abbreviate("strategy-limited"))
	// Output:
	// stg-ltd <nil>
}

func ExampleInitialism() {
	fmt.Println(Initialism("hello world"))
	// Output:
	// hw <nil>
}

func ExampleNumeronym() {
	fmt.Println(Numeronym("internationalization"))
	fmt.Println(Numeronym("data-desensitization"))
	fmt.Println(Numeronym("hello world"))
	// Output:
	// i18n <nil>
	// d18n <nil>
	// h3o w3d <nil>
}
