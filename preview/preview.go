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

package preview

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/LianjiaTech/d18n/common"
)

type PreviewStruct struct {
	Config common.Config
}

func NewPreviewStruct(c common.Config) (*PreviewStruct, error) {
	var p *PreviewStruct
	p = &PreviewStruct{
		Config: c,
	}
	return p, nil
}

// Preview preview export file
func (p *PreviewStruct) Preview() error {
	var err error

	switch p.Config.File {
	case "", "stdout":
		return fmt.Errorf("expect -file arg")
	}

	if _, err := os.Stat(p.Config.File); err != nil {
		return err
	}

	suffix := strings.ToLower(strings.TrimLeft(filepath.Ext(p.Config.File), "."))
	switch suffix {
	case "stdout", "":
	case "csv", "psv", "tsv", "txt", "sql":
		err = previewCSV(p)
	case "html":
		err = previewHTML(p)
	case "xlsx":
		err = previewXlsx(p)
	case "json":
		err = previewJSON(p)
	default:
		err = fmt.Errorf("not support extension: " + suffix)
	}

	return err
}
