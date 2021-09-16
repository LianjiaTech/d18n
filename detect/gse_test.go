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
	_ "embed"
	"fmt"
)

func ExampleGSE() {
	fmt.Println(GSE("我叫岳云鹏"))
	fmt.Println(GSE("我是张韶涵"))
	fmt.Println(GSE("毛阿敏"))
	fmt.Println(GSE("敲黑板，划重点"))
	fmt.Println(GSE("我住在北京市大兴区庞各庄镇"))
	fmt.Println(GSE("我叫伊藤中郎，我住在三重県員弁町石仏北勢町京ヶ野新田"))
	// Output:
	// name
	// name
	// name
	//
	// address
	// name
}
