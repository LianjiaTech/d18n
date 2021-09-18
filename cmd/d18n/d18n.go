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

func main() {
	// limit cpu 1 core, memory 2GB
	common.PanicIfError(common.ResourceLimit(1, 2*1024*1024*1024))

	// parse config
	// common.PanicIfError(common.ParseFlag())
	common.PanicIfError(common.ParseFlags())

	// parse cipher config
	common.PanicIfError(mask.ParseCipherConfig(common.Cfg.Cipher))

	// print cipher
	if common.Cfg.PrintCipher {
		mask.PrintCipher()
		return
	}

	// print config
	if common.Cfg.PrintConfig {
		common.PrintConfig()
		return
	}

	// preview file
	if common.Cfg.Preview > 0 {
		common.PanicIfError(preview.Preview())
		return
	}

	// lint file
	if common.Cfg.Lint {
		common.PanicIfError(lintFile())
		return
	}

	// detect sensitive info
	if common.Cfg.Detect {
		common.PanicIfError(detectRows())
		return
	}

	// init mask corpus
	common.PanicIfError(mask.InitMaskCorpus(common.Cfg.RandSeed))

	// import file
	if common.Cfg.Import {
		common.PanicIfError(emportFile())
		return
	}

	common.PanicIfError(saveRows())
}

func saveRows() error {
	// new save struct
	s, err := save.NewSaveStruct(common.Cfg)
	if err != nil {
		return err
	}

	// query and save result
	common.PanicIfError(s.Save())

	// check save status
	return s.CheckStatus()
}

func lintFile() error {
	// check file format
	common.PanicIfError(lint.Lint())

	// check lint status
	return lint.CheckStatus()
}

func emportFile() error {
	e, err := emport.NewEmportStruct(common.Cfg)
	if err != nil {
		return err
	}
	// import file into database
	common.PanicIfError(e.Emport())

	// check emport status
	return e.CheckStatus()
}

func detectRows() error {
	err := detect.ParseSensitiveConfig()
	if err != nil {
		return err
	}

	common.PanicIfError(detect.Detect())

	// check detect status
	return detect.CheckStatus()
}
