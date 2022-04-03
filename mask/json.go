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
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// JSON fake json mask function
func JSON(args ...interface{}) (ret string, err error) {
	return
}

func (m *MaskStruct) JSONMask(value interface{}) []byte {
	var in []byte
	switch v := value.(type) {
	case []byte:
		in = v
	case json.RawMessage:
		in, _ = v.MarshalJSON()
	default:
		in = []byte(fmt.Sprint(value))
	}
	if out, err := m.maskMap(in); err == nil {
		return out
	}
	if out, err := m.maskList(in); err == nil {
		return out
	}
	return in
}

func (m *MaskStruct) maskList(in []byte) ([]byte, error) {
	var out []*json.RawMessage
	if err := json.Unmarshal(in, &out); err != nil {
		return in, err
	}

	for i, v := range out {
		if v == nil {
			continue
		}
		r := json.RawMessage(m.JSONMask(*v))
		out[i] = &r
	}
	return json.Marshal(out)
}

func (m *MaskStruct) maskMap(in []byte) ([]byte, error) {
	var out map[string]*json.RawMessage
	if err := json.Unmarshal(in, &out); err != nil {
		return in, err
	}

	for i, v := range out {
		if v == nil {
			continue
		}
		if r, ok := m.Config[strings.ToLower(i)]; ok {
			b, _ := v.MarshalJSON()

			var quoted bool
			if strings.HasPrefix(string(b), `"`) && strings.HasSuffix(string(b), `"`) {
				quoted = true
			}

			// concat mask args
			var args []interface{}
			args = append(args, strings.Trim(string(b), `"`))
			for _, arg := range r.Args {
				args = append(args, arg)
			}

			// run mask function
			var tmp string
			var err error
			if r.MaskFunc == "fake" { // generate fake data no need origin value
				tmp, err = maskFuncs[r.MaskFunc](args[1:]...)
			} else {
				tmp, err = maskFuncs[r.MaskFunc](args...)
			}
			if err != nil {
				return in, err
			}

			if quoted {
				tmp = strconv.Quote(tmp)
			}
			raw := json.RawMessage(tmp)
			out[i] = &raw
			continue
		}
		raw := json.RawMessage(m.JSONMask(*v))
		out[i] = &raw
	}
	return json.Marshal(out)
}
