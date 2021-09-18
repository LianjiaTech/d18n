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
	"bufio"
	"database/sql"
	"fmt"
	"os"

	"d18n/common"

	"golang.org/x/net/html"
)

// saveRows2HTML save rows result into HTML format file
func saveRows2HTML(s *SaveStruct, rows *sql.Rows) error {
	file, err := os.Create(s.CommonConfig.File)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriterSize(file, s.CommonConfig.MaxBufferSize)
	if s.CommonConfig.Watermark != "" {
		_, err = w.WriteString("<!-- " + s.CommonConfig.Watermark + " -->\n" + fmt.Sprintf(common.WatermarkPrefix, s.CommonConfig.Watermark))
		if err != nil {
			return err
		}
	}

	_, err = w.WriteString("<TABLE>\n")
	if err != nil {
		return err
	}
	// set table header with column name
	if !s.CommonConfig.NoHeader {
		_, err = w.WriteString("<TR>")
		if err != nil {
			return err
		}
		for _, h := range s.Status.Header {
			_, err = w.WriteString("<TH>" + html.EscapeString(h.Name()) + "</TH>")
			if err != nil {
				return err
			}
		}
		_, err = w.WriteString("</TR>\n")
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

		// write <TR>
		_, err = w.WriteString("<TR>")
		if err != nil {
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
				values[j], err = s.Masker.Mask(s.Status.Header[j].Name(), values[j])
				if err != nil {
					return err
				}

				// hex-blob
				values[j], _ = common.HexBLOB(s.Status.Header[j].Name(), values[j])
			}

			_, err = w.WriteString("<TD>" + html.EscapeString(values[j]) + "</TD>")
			if err != nil {
				return err
			}
		}

		// write </TR>
		_, err = w.WriteString("</TR>\n")
		if err != nil {
			return err
		}
	}
	err = rows.Err()
	if err != nil {
		return err
	}

	_, err = w.WriteString("</TABLE>")
	if err != nil {
		return err
	}

	if s.CommonConfig.Watermark != "" {
		_, err = w.WriteString("<!-- " + s.CommonConfig.Watermark + " -->" + common.WatermarkSuffix)
		if err != nil {
			return err
		}
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	return err
}
