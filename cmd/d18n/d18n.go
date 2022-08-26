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

package main

import (
	"github.com/LianjiaTech/d18n/common"
	"github.com/LianjiaTech/d18n/detect"
	"github.com/LianjiaTech/d18n/emport"
	"github.com/LianjiaTech/d18n/lint"
	"github.com/LianjiaTech/d18n/mask"
	"github.com/LianjiaTech/d18n/preview"
	"github.com/LianjiaTech/d18n/save"
)

var c common.Config

func main() {
	var err error

	// limit cpu 1 core, memory 2GB
	common.PanicIfError(common.ResourceLimit(1, 2*1024*1024*1024))

	// parse config
	c, err = common.ParseFlags()
	common.PanicIfError(err)

	// parse cipher config
	if c.Mask != "" { // generate cipher will cost about 200ms
		common.PanicIfError(mask.ParseCipherConfig(c.Cipher))
	}

	// print cipher
	if c.PrintCipher {
		mask.PrintCipher()
		return
	}

	// print config
	if c.PrintConfig {
		common.PrintConfig(c)
		return
	}

	// preview file
	if c.Preview > 0 {
		common.PanicIfError(previewFile())
		return
	}

	// lint file
	if c.Lint {
		common.PanicIfError(lintFile())
		return
	}

	// detect sensitive info
	if c.Detect {
		common.PanicIfError(detectRows())
		return
	}

	// init mask corpus
	if c.Mask != "" {
		common.PanicIfError(mask.InitMaskCorpus(c.RandSeed))
	}

	// import file
	if c.Import {
		common.PanicIfError(emportFile())
		return
	}

	if c.Interactive {
		err = saveRows()
		if err != nil {
			println(err.Error())
		}
		main()
	} else {
		common.PanicIfError(saveRows())
	}
}

func previewFile() error {
	p, err := preview.NewPreviewStruct(c)
	if err != nil {
		return err
	}
	return p.Preview()
}

func saveRows() error {
	// new save struct
	s, err := save.NewSaveStruct(c)
	if err != nil {
		return err
	}

	// query and save result
	err = s.Save()
	if err != nil {
		return err
	}

	// show save status
	return s.ShowStatus()
}

func lintFile() error {
	l, err := lint.NewLintStruct(c)
	if err != nil {
		return err
	}

	// check file format
	common.PanicIfError(l.Lint())

	// show lint status
	return l.ShowStatus()
}

func emportFile() error {
	e, err := emport.NewEmportStruct(c)
	if err != nil {
		return err
	}
	// import file into database
	common.PanicIfError(e.Emport())

	// show emport status
	return e.ShowStatus()
}

func detectRows() error {
	d, err := detect.NewDetectStruct(c)
	if err != nil {
		return err
	}

	// detect sensitive data
	common.PanicIfError(d.Detect())

	// show detect status
	return d.ShowStatus()
}
