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
	"regexp"

	"github.com/LianjiaTech/d18n/common"

	"gopkg.in/yaml.v2"
)

//go:embed sensitive.yaml
var defaultSensitiveConfig []byte

// sensitiveConfig ...
type sensitiveConfig map[string]BasicDetect

func (d *DetectStruct) parseConfig() error {
	// load sensitive config
	c, err := common.ReadFileString(d.Config.Sensitive)
	if err == nil {
		defaultSensitiveConfig = []byte(c)
	}

	err = yaml.Unmarshal(defaultSensitiveConfig, &d.Sensitive)
	if err != nil {
		return err
	}

	// check config regexp valid
	for _, v := range d.Sensitive {
		for _, r := range v.Key {
			_, err = regexp.Compile(r)
			if err != nil {
				return err
			}
		}

		for _, r := range v.Value {
			_, err = regexp.Compile(r)
			if err != nil {
				return err
			}
		}
	}
	return err
}
