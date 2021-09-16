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
	"d18n/mask"

	"golang.org/x/net/html"
)

// saveRows2HTML save rows result into HTML format file
func saveRows2HTML(rows *sql.Rows) error {
	file, err := os.Create(common.Cfg.File)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriterSize(file, common.Cfg.MaxBufferSize)
	if common.Cfg.Watermark != "" {
		_, err = w.WriteString("<!-- " + common.Cfg.Watermark + " -->\n" + fmt.Sprintf(common.WatermarkPrefix, common.Cfg.Watermark))
		if err != nil {
			return err
		}
	}

	_, err = w.WriteString("<TABLE>\n")
	if err != nil {
		return err
	}
	// set table header with column name
	if !common.Cfg.NoHeader {
		_, err = w.WriteString("<TR>")
		if err != nil {
			return err
		}
		for _, h := range saveStatus.Header {
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
	columns := make([]interface{}, len(saveStatus.Header))
	cols := make([]interface{}, len(saveStatus.Header))
	for j := range columns {
		cols[j] = &columns[j]
	}

	for rows.Next() {
		saveStatus.Lines++
		// limit return rows
		if common.Cfg.Limit != 0 && saveStatus.Lines > common.Cfg.Limit {
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
				values[j] = common.Cfg.NULLString
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
				values[j], err = mask.Mask(saveStatus.Header[j].Name(), values[j])
				if err != nil {
					return err
				}

				// hex-blob
				values[j], _ = common.HexBLOB(saveStatus.Header[j].Name(), values[j])
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

	if common.Cfg.Watermark != "" {
		_, err = w.WriteString("<!-- " + common.Cfg.Watermark + " -->" + common.WatermarkSuffix)
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
