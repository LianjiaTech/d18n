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

var shuffleTestCases = []interface{}{
	1234567890,
	"1234567890",
	"123abc4567890",
	"123ABC4567890",
	"中文abcABC",
	"10.199.90.105",
	"abc",
}

func ExampleShuffle() {
	for _, c := range shuffleTestCases {
		r, err := Shuffle(c)
		fmt.Println(r, err)
	}
	// Output:
	// 4802731596 <nil>
	// 4802731596 <nil>
	// 480rgf2731596 <nil>
	// 480UMK2731596 <nil>
	// 肥涡rgfUMK <nil>
	// 46.499.96.467 <nil>
	// rgf <nil>
}

func TestShuffleRight(t *testing.T) {
	for _, c := range shuffleTestCases {
		r, err := ShuffleRight(c, 4)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(r)
	}
}

func TestShuffleLeft(t *testing.T) {
	for _, c := range shuffleTestCases {
		r, err := ShuffleLeft(c, 4)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(r)
	}
}

func TestShuffle9(t *testing.T) {
	for _, c := range shuffleTestCases {
		r, err := Shuffle(c, "1-9")
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(r)
	}
}

func TestShuffle10(t *testing.T) {
	for _, c := range shuffleTestCases {
		r, err := Shuffle(c, "0-9")
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(r)
	}
}

func TestShuffleUpper(t *testing.T) {
	for _, c := range shuffleTestCases {
		r, err := Shuffle(c, "upper")
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(r)
	}
}

func TestShuffleLower(t *testing.T) {
	for _, c := range shuffleTestCases {
		r, err := Shuffle(c, "lower")
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(r)
	}
}

func TestShuffleAlphabet(t *testing.T) {
	for _, c := range shuffleTestCases {
		r, err := Shuffle(c, "alphabet")
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(r)
	}
}

func ExampleRot() {
	for _, c := range shuffleTestCases {
		fmt.Println(Rot(c, 5))
		fmt.Println(Rot(c, 13))
		fmt.Println(Rot(c, 18))
		fmt.Println(Rot(c, 47))
		fmt.Println(Rot(c, 32768))
		fmt.Println(Rot(c, 3))
	}
	// Output:
	// 6789012345 <nil>
	// 1234567890 <nil>
	// 6789012345 <nil>
	// `abcdefgh_ <nil>
	// 6789012345 <nil>
	//  wrong mask function
	// 6789012345 <nil>
	// 1234567890 <nil>
	// 6789012345 <nil>
	// `abcdefgh_ <nil>
	// 6789012345 <nil>
	//  wrong mask function
	// 678abc9012345 <nil>
	// 123nop4567890 <nil>
	// 678nop9012345 <nil>
	// `ab234cdefgh_ <nil>
	// 678áâã9012345 <nil>
	//  wrong mask function
	// 678ABC9012345 <nil>
	// 123NOP4567890 <nil>
	// 678NOP9012345 <nil>
	// `abpqrcdefgh_ <nil>
	// 678ÁÂÃ9012345 <nil>
	//  wrong mask function
	// 中文abcABC <nil>
	// 中文nopNOP <nil>
	// 中文nopNOP <nil>
	// 中文234pqr <nil>
	// 亭攇áâãÁÂÃ <nil>
	//  wrong mask function
	// 65.644.45.650 <nil>
	// 10.199.90.105 <nil>
	// 65.644.45.650 <nil>
	// `_]`hh]h_]`_d <nil>
	// 65®644®45®650 <nil>
	//  wrong mask function
	// abc <nil>
	// nop <nil>
	// nop <nil>
	// 234 <nil>
	// áâã <nil>
	//  wrong mask function
}

func ExampleMorse() {
	for _, c := range shuffleTestCases {
		fmt.Println(Morse(c))
	}
	// Output:
	// .---- ..--- ...-- ....- ..... -.... --... ---.. ----. ----- <nil>
	// .---- ..--- ...-- ....- ..... -.... --... ---.. ----. ----- <nil>
	// .---- ..--- ...-- .- -... -.-. ....- ..... -.... --... ---.. ----. ----- <nil>
	// .---- ..--- ...-- .- -... -.-. ....- ..... -.... --... ---.. ----. ----- <nil>
	// # # .- -... -.-. .- -... -.-. error in input: #中##文#ABCABC
	// .---- ----- .-.-.- .---- ----. ----. .-.-.- ----. ----- .-.-.- .---- ----- ..... <nil>
	// .- -... -.-. <nil>
}

func ExampleCaesar() {
	for _, c := range shuffleTestCases {
		fmt.Println(Caesar(c, 3))
	}
	// Output:
	// 1234567890 <nil>
	// 1234567890 <nil>
	// 123def4567890 <nil>
	// 123DEF4567890 <nil>
	// 中文defDEF <nil>
	// 10.199.90.105 <nil>
	// def <nil>
}
