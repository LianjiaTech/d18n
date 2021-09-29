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

package detect

import (
	"embed"
	"strings"

	"github.com/LianjiaTech/d18n/common"

	"github.com/go-ego/gse"
)

var seg gse.Segmenter
var loadCorpus bool

//go:embed corpus
var corpus embed.FS

func initGSE() (err error) {
	files, err := corpus.ReadDir("corpus")
	if err != nil {
		return err
	}

	// load gse corpus
	for _, i := range files {
		if strings.HasPrefix(i.Name(), "gse.") {
			s, err := corpus.ReadFile("corpus/" + i.Name())
			if err != nil {
				return err
			}
			err = seg.LoadDictStr(string(s))
			if err != nil {
				return err
			}
		}
	}
	loadCorpus = true
	return err
}

// GSE use gse to analyze the most likely data types
func GSE(txt string) string {
	if !loadCorpus {
		common.PanicIfError(initGSE())
	}

	pos := seg.PosTrim(txt, false, "x")
	return maxFreqPos(pos)
}

// maxFreqPos compute maximum frequency for pos
func maxFreqPos(pos []gse.SegPos) string {
	var rank = make(map[string]float64)
	var maxPos string
	var maxFreq float64
	for _, v := range pos {
		rank[v.Pos] = rank[v.Pos] + seg.SuggestFreq(v.Text)
	}
	for k, v := range rank {
		if v > maxFreq {
			maxFreq = v
			maxPos = k
		}
	}
	return maxPos
}
