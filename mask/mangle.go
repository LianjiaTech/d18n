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

package mask

// Copy from https://github.com/grugnog/mangle

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"hash/crc32"
	"path/filepath"
	"strings"
	"unicode"

	"d18n/common"
)

type mangleCorpus [255][]string

var mangleMap = make(map[string]mangleCorpus)

func InitMangle() error {
	files, err := corpusFS.ReadDir("corpus")
	if err != nil {
		return err
	}

	for _, file := range files {
		if !strings.HasPrefix(file.Name(), "mangle") {
			continue
		}
		language := strings.ToLower(strings.TrimLeft(filepath.Ext(file.Name()), "."))

		// load corpus from embed.FS
		en, err := corpusFS.ReadFile("corpus/" + file.Name())
		if err != nil {
			return err
		}
		s := bufio.NewScanner(bytes.NewReader(en))
		c, err := loadCorpus(s)
		if err != nil {
			return err
		}
		mangleMap[language] = c
	}
	return nil
}

// Mangle this func shuffle English article FPE(Format Preserve Encrypt)
// Chinese article should use Shuffle func, because Chinese sentences ard made
// by multi single character, English sentences are made by multi single words.
func Mangle(args ...interface{}) (ret string, err error) {
	if len(args) < 3 {
		return ret, fmt.Errorf(common.WrongArgsCount)
	}

	language := strings.ToLower(fmt.Sprint(args[1]))
	if _, ok := mangleMap[language]; !ok {
		return ret, fmt.Errorf(common.WrongArgValue)
	}

	var m = MangleConfig{
		Corpus: mangleMap[language],
		Secret: fmt.Sprint(args[2]),
	}
	return m.string(fmt.Sprint(args[0])), err
}

// MangleConfig is used to configure an instance prior to mangling.
type MangleConfig struct {
	// Corpus of words to use as replacements. An array of word lengths, each
	// containing an array of words of that length.
	Corpus mangleCorpus
	// A sufficiently long secret, used as a salt so rainbow tables cannot be
	// used to reverse the hashes.
	Secret string
}

// loadCorpus is a helper function that reads a bufio.Scanner of words and
// returns an array of word lengths, each containing an array of words of that
// length.
func loadCorpus(scanner *bufio.Scanner) (mangleCorpus, error) {
	var c mangleCorpus
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordlen := len(scanner.Text())
		c[wordlen] = append(c[wordlen], scanner.Text())
	}
	if len(c[1]) == 0 {
		return c, errors.New("corpus must contain at least one single character word")
	}
	return c, scanner.Err()
}

// string operates on strings, and is preferable if you have many short
// strings to operate on.
func (m MangleConfig) string(s string) string {
	var output string
	var word []rune
	runes := []rune(s)
	strlen := len(runes)
	for i := 0; i < strlen; i++ {
		rune := runes[i]
		if unicode.IsLetter(rune) || unicode.IsNumber(rune) {
			// In word.
			word = append(word, rune)
		} else {
			// Inter-word.
			if len(word) > 0 {
				// Process previous word.
				output += m.words(word)
				// Reset word.
				word = word[0:0]
			}
			output += string(rune)
		}
	}
	// Process last word.
	output += m.words(word)
	return output
}

// Performs the core mangling function on a word. The approach is to hash the
// word and the secret salt, then map the hash value into the available corpus
// words of the appropriate length (or the longest available length, padding
// with whitespace as needed). The resulting word is then adjusted to
// match the original capitalization.
func (m MangleConfig) words(word []rune) string {
	var crc float64
	var pos uint32
	var replacementRunes []rune
	const MaxUint32 = 1<<32 - 1
	replacement := ""

	pad := 0
	wordLen := len(word)
	if wordLen > 0 {
		// SHA256 the string, together with the secret.
		hash := sha256.New()
		hash.Write([]byte(string(word)))
		hash.Write([]byte(m.Secret))
		// Use crc32 to map the hash to a conveniently sized number.
		crc = float64(crc32.ChecksumIEEE([]byte(hash.Sum(nil))))
		// If the word is too long, or we can't find a sufficiently long word in the
		// corpus, look for a shorter one and adjust the padding.
		for wordLen >= 255 || (len(m.Corpus[wordLen]) == 0 && wordLen != 0) {
			wordLen--
			pad++
		}
		if wordLen != 0 {
			// Map the CRC value onto the available corpus words.
			pos = uint32((crc / MaxUint32) * float64(len(m.Corpus[wordLen])))
			// Select the word from the corpus and pad it if it was shorter than the original.
			replacement = m.Corpus[wordLen][pos] + strings.Repeat(" ", pad)
			// Capitalize as per the original string.
			if wordLen > 0 && unicode.IsUpper(word[0]) {
				if wordLen > 1 && unicode.IsUpper(word[1]) {
					replacement = strings.ToUpper(replacement)
				} else {
					replacementRunes = []rune(replacement)
					replacement = strings.ToUpper(string(replacementRunes[0])) + string(replacementRunes[1:])
				}
			}
		}
	}
	return replacement
}
