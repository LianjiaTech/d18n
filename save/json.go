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

package save

import (
	"database/sql"
	"fmt"

	json "github.com/json-iterator/go"

	//"encoding/json"
	"os"

	"d18n/common"
	"d18n/mask"
)

// JSON limit
// JSON parser limits - IBM Documentation
// https://www.ibm.com/docs/en/datapower-gateways/7.6?topic=20-json-parser-limits

// saveRows2JSON save rows result into JSON format file
func saveRows2JSON(s *SaveStruct, rows *sql.Rows) error {
	file, err := os.Create(s.CommonConfig.File)
	if err != nil {
		return err
	}
	defer file.Close()

	// new json stream
	stream := json.NewStream(json.Config{IndentionStep: 2}.Froze(), file, 512)
	stream.WriteArrayStart() // [

	// key names as json list first element
	if !s.CommonConfig.NoHeader {
		buf, err := json.Marshal(common.DBParserColumnNames(s.Status.Header))
		if err != nil {
			return err
		}
		_, err = stream.Write(buf)
		if err != nil {
			return err
		}
		err = stream.Flush()
		if err != nil {
			return err
		}
	}

	// init columns
	columns := make([]interface{}, len(s.Status.Header))
	cols := make([]interface{}, len(s.Status.Header))
	for j := range columns {
		cols[j] = &columns[j]
	}

	for rows.Next() {
		s.Status.Lines++
		// limit return rows
		if s.CommonConfig.Limit != 0 && s.Status.Lines > s.CommonConfig.Limit {
			break
		}

		// scan columns
		if err := rows.Scan(cols...); err != nil {
			return err
		}

		values := make([]string, len(columns))
		for j, col := range columns {
			if col == nil {
				values[j] = s.CommonConfig.NULLString
			} else {
				switch col.(type) {
				case []byte:
					values[j] = string(col.([]byte))
				case []string:
					values[j] = common.ParseArray(col.([]string))
				default:
					values[j] = fmt.Sprint(col)
				}

				// data mask
				values[j], err = mask.Mask(s.Status.Header[j].Name(), values[j])
				if err != nil {
					return err
				}

				// hex-blob
				values[j], _ = common.HexBLOB(s.Status.Header[j].Name(), values[j])
			}
		}

		// write each record into JSON file one by one
		buf, err := json.Marshal(values)
		if err != nil {
			return err
		}

		// add comma
		if s.CommonConfig.NoHeader && s.Status.Lines == 1 {
			// -no-header and first line don't add comma
		} else {
			stream.WriteMore() // ,
		}

		// write one row
		_, err = stream.Write(buf)
		if err != nil {
			return err
		}
		err = stream.Flush()
		if err != nil {
			return err
		}
	}
	stream.WriteArrayEnd() // ]
	err = stream.Flush()
	if err != nil {
		return err
	}
	err = rows.Err()
	return err
}
