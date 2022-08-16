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
	"os"
	"strings"

	json "github.com/json-iterator/go"
)

// JSON limit
// JSON parser limits - IBM Documentation
// https://www.ibm.com/docs/en/datapower-gateways/7.6?topic=20-json-parser-limits

// saveRows2JSON save rows result into JSON format file
func saveRows2JSON(s *SaveStruct, rows *sql.Rows) error {

	var err error
	var file *os.File
	if strings.EqualFold(s.Config.File, "stdout") {
		file = os.Stdout
	} else {
		file, err = os.Create(s.Config.File)
		if err != nil {
			return err
		}
	}
	defer file.Close()

	// new json stream
	stream := json.NewStream(json.Config{IndentionStep: 2}.Froze(), file, 512)
	stream.WriteArrayStart() // [

	// column info
	columnNames, err := rows.Columns()
	if err != nil {
		return err
	}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}

	// key names as json list first element
	if !s.Config.NoHeader {
		buf, err := json.Marshal(columnNames)
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
	columnValues := make([]interface{}, len(columnNames))
	cols := make([]interface{}, len(columnNames))
	for j := range columnValues {
		cols[j] = &columnValues[j]
	}

	for rows.Next() {
		s.Status.Lines++
		// limit return rows
		if s.Config.Limit != 0 && s.Status.Lines > s.Config.Limit {
			s.Status.Lines = s.Config.Limit
			break
		}

		// scan columns
		if err := rows.Scan(cols...); err != nil {
			return err
		}

		values := make([]string, len(columnNames))
		for j, col := range columnValues {
			if col == nil {
				values[j] = s.Config.NULLString
			} else {
				values[j] = s.String(col, columnTypes[j])

				// data mask
				values[j], err = s.Masker.Mask(s.FieldName(j), values[j])
				if err != nil {
					return err
				}

				// hex-blob
				values[j], _ = s.Config.Hex(s.FieldName(j), values[j])
			}
		}

		// write each record into JSON file one by one
		buf, err := json.Marshal(values)
		if err != nil {
			return err
		}

		// add comma
		if s.Config.NoHeader && s.Status.Lines == 1 {
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
