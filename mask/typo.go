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

type MaskFunc func(args ...interface{}) (ret string, err error)

// maskFuncs support functions list, case insensitive
var maskFuncs = map[string]MaskFunc{
	"aes":                 AES,
	"aesctr":              AESCTR,
	"abbreviate":          Abbreviate,
	"age":                 Age,
	"base64":              Base64,
	"birthday":            Birthday,
	"crc32":               CRC32,
	"caesar":              Caesar,
	"char2const":          Char2Const,
	"const":               Const,
	"creditcard":          CreditCard,
	"des":                 DES,
	"dateformat":          DateFormat,
	"dateround":           DateRound,
	"domain":              Domain,
	"ecc":                 ECC,
	"ff1":                 FF1,
	"ff3":                 FF3,
	"fake":                Fake,
	"hmac":                HMAC,
	"ip":                  IP,
	"initialism":          Initialism,
	"json":                JSON,
	"laplacedpfloat64":    LaplaceDPFloat64,
	"laplacedpint64":      LaplaceDPInt64,
	"licenseplate":        LicensePlate,
	"loopmoveleft":        LoopMoveLeft,
	"loopmoveright":       LoopMoveRight,
	"md5":                 MD5,
	"mail":                Mail,
	"mangle":              Mangle,
	"morse":               Morse,
	"number2const":        Number2Const,
	"numberfloor":         NumberFloor,
	"numeronym":           Numeronym,
	"organizationcode":    OrganizationCode,
	"password":            Password,
	"personalid":          PersonalID,
	"phone":               Phone,
	"rsa":                 RSA,
	"regexprandomreplace": RegexpRandomReplace,
	"regexpreplace":       RegexpReplace,
	"replace":             Replace,
	"reserveinner":        ReserveInner,
	"reserveleft":         ReserveLeft,
	"reservemargin":       ReserveMargin,
	"reserveouter":        ReserveOuter,
	"reserveright":        ReserveRight,
	"reverse":             Reverse,
	"rot":                 Rot,
	"sha1":                SHA1,
	"sha2":                SHA2,
	"sm2":                 SM2,
	"sm3":                 SM3,
	"sm4":                 SM4,
	"salary":              Salary,
	"shuffle":             Shuffle,
	"shuffleleft":         ShuffleLeft,
	"shuffleright":        ShuffleRight,
	"smoke":               Smoke,
	"smokecharleft":       SmokeCharLeft,
	"smokecharright":      SmokeCharRight,
	"smokeinner":          SmokeInner,
	"smokeleft":           SmokeLeft,
	"smokemargin":         SmokeMargin,
	"smokeouter":          SmokeOuter,
	"smokeright":          SmokeRight,
	"tdea":                TDEA,
	"tolower":             ToLower,
	"toupper":             ToUpper,
	"truncateleft":        TruncateLeft,
	"truncateright":       TruncateRight,
	"uscc":                USCC,
	"username":            Username,
}
