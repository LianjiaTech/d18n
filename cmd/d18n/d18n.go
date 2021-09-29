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
	"d18n/common"
	"d18n/detect"
	"d18n/emport"
	"d18n/lint"
	"d18n/mask"
	"d18n/preview"
	"d18n/save"
)

var c common.Config

func main() {
	var err error

	// limit cpu 1 core, memory 2GB
	common.PanicIfError(common.ResourceLimit(1, 2*1024*1024*1024))

	// parse config
	// common.PanicIfError(common.ParseFlag())
	c, err = common.ParseFlags()
	common.PanicIfError(err)

	// parse cipher config
	common.PanicIfError(mask.ParseCipherConfig(c.Cipher))

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
	common.PanicIfError(mask.InitMaskCorpus(c.RandSeed))

	// import file
	if c.Import {
		common.PanicIfError(emportFile())
		return
	}

	common.PanicIfError(saveRows())
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
	common.PanicIfError(s.Save())

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
