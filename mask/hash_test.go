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

import "fmt"

func ExampleCRC32() {
	fmt.Println(CRC32("abc"))
	fmt.Println(CRC32(123))
	fmt.Println(CRC32("123"))
	// Output:
	// 352441c2 <nil>
	// 884863d2 <nil>
	// 884863d2 <nil>
}

func ExampleMD5() {
	fmt.Println(MD5("abc"))
	fmt.Println(MD5(123))
	fmt.Println(MD5("123"))
	// Output:
	// 900150983cd24fb0d6963f7d28e17f72 <nil>
	// 202cb962ac59075b964b07152d234b70 <nil>
	// 202cb962ac59075b964b07152d234b70 <nil>
}

func ExampleSHA1() {
	fmt.Println(SHA1("abc"))
	fmt.Println(SHA1(123))
	fmt.Println(SHA1("123"))
	// Output:
	// a9993e364706816aba3e25717850c26c9cd0d89d <nil>
	// 40bd001563085fc35165329ea1ff5c5ecbdbbeef <nil>
	// 40bd001563085fc35165329ea1ff5c5ecbdbbeef <nil>
}

func ExampleSHA2() {
	fmt.Println(SHA2("abc"))
	fmt.Println(SHA2(123))
	fmt.Println(SHA2("123"))
	// Output:
	// ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad <nil>
	// a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3 <nil>
	// a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3 <nil>
}

func ExampleHMAC() {
	fmt.Println(HMAC("helle world", "md5", "pass"))
	fmt.Println(HMAC("helle world", "sha1", "pass"))
	fmt.Println(HMAC("helle world", "sha2", "pass"))
	// Output:
	// 37c4d226765f06daa3ad91a6c33a5d3e <nil>
	// 316278d8dab3b11d98501b6f000980a1203d7e4a <nil>
	// 3daa940668aa37073bc91dd8d71c3f42c44e8b5a91a5100a8ad7efb9c2ad6224 <nil>
}
