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
	"os"
	"testing"

	"d18n/common"
)

func init() {
	common.InitTestEnv()
}

func TestMainSave(t *testing.T) {
	orgCfg := common.Cfg
	orgArgs := os.Args
	args := []string{
		"--defaults-extra-file", common.TestPath + "/test/my.cnf",
		"--query", common.Cfg.Query,
	}
	os.Args = append(os.Args[:1], args...)
	main()
	os.Args = orgArgs
	common.Cfg = orgCfg
}

func TestMainEmport(t *testing.T) {
	orgCfg := common.Cfg
	orgArgs := os.Args
	args := []string{
		"--file", common.TestPath + "/test/actor.csv", "--import",
		"--table", "actor",
		"--limit", "10",
		"--extended-insert", "3",
		"--schema", common.TestPath + "/test/schema.txt",
		"--database", "sakila",
		"--replace", "--disable-foreign-key-checks",
	}
	os.Args = append(os.Args[:1], args...)
	main()
	os.Args = orgArgs
	common.Cfg = orgCfg
}

func Example_lint() {
	orgCfg := common.Cfg
	orgArgs := os.Args
	args := []string{
		"--file", common.TestPath + "/test/actor.csv", "--lint",
	}
	os.Args = append(os.Args[:1], args...)
	main()
	os.Args = orgArgs
	common.Cfg = orgCfg
	// Output:
	// ok
}

func TestMainDetect(t *testing.T) {
	orgCfg := common.Cfg
	orgArgs := os.Args
	args := []string{
		"--query", `select * from sakila.actor limit 10`, "--detect",
	}
	os.Args = append(os.Args[:1], args...)
	main()
	os.Args = orgArgs
	common.Cfg = orgCfg
}

func TestMainPrintConfig(t *testing.T) {
	orgCfg := common.Cfg
	orgArgs := os.Args
	args := []string{
		"--print-config",
	}
	os.Args = append(os.Args[:1], args...)
	main()
	os.Args = orgArgs
	common.Cfg = orgCfg
}

func TestMainPrintCipher(t *testing.T) {
	orgCfg := common.Cfg
	orgArgs := os.Args
	args := []string{
		"--print-cipher",
	}
	os.Args = append(os.Args[:1], args...)
	main()
	os.Args = orgArgs
	common.Cfg = orgCfg
}

func Example_preview() {
	orgCfg := common.Cfg
	orgArgs := os.Args
	args := []string{
		"--preview", "2",
		"--file", common.TestPath + "/test/actor.csv",
	}
	os.Args = append(os.Args[:1], args...)
	main()
	os.Args = orgArgs
	common.Cfg = orgCfg
	// Output:
	// actor_id,first_name,last_name,last_update
	// 1,PENELOPE,GUINESS,2006-02-15 04:34:33
}
