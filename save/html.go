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
	"strings"

	"github.com/LianjiaTech/d18n/common"

	"golang.org/x/net/html"
)

// saveRows2HTML save rows result into HTML format file
func saveRows2HTML(s *SaveStruct, rows *sql.Rows) error {

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

	w := bufio.NewWriterSize(file, s.Config.MaxBufferSize)
	if s.Config.Watermark != "" {
		_, err = w.WriteString("<!-- " + s.Config.Watermark + " -->\n" + fmt.Sprintf(common.WatermarkPrefix, s.Config.Watermark))
		if err != nil {
			return err
		}
	}

	_, err = w.WriteString("<TABLE>\n")
	if err != nil {
		return err
	}

	// column info
	columnNames, err := rows.Columns()
	if err != nil {
		return err
	}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}

	// set table header with column name
	if !s.Config.NoHeader {
		_, err = w.WriteString("<TR>")
		if err != nil {
			return err
		}
		for _, h := range columnNames {
			_, err = w.WriteString("<TH>" + html.EscapeString(h) + "</TH>")
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

		// write <TR>
		_, err = w.WriteString("<TR>")
		if err != nil {
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

	if s.Config.Watermark != "" {
		_, err = w.WriteString("<!-- " + s.Config.Watermark + " -->" + common.WatermarkSuffix)
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
