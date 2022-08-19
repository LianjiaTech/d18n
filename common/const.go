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

package common

const (
	// UTF8BOM utf8 file BOM header
	UTF8BOM = "\xEF\xBB\xBF"

	// DefaultExcelMaxFileSize excel file max size 10MB
	DefaultExcelMaxFileSize = 10 * 1024 * 1024
)

const (
	DATETIME_FORMAT = "2006-01-02 15:04:05.999999999"
	DATE_FORMAT     = "2006-01-02"
	YEAR_FORMAT     = "2006"
)

// ERROR message
const (
	WrongEmptySet    = `Empty set`
	WrongJSONFormat  = `JSON format error, only support string list. e.g., [ ["header", "columns" ], [ "col1", "col2" ] ]`
	WrongSQLFormat   = `SQL format error, only support basic INSERT/REPLACE syntax`
	WrongArgsCount   = `arguments count mismatch`
	WrongMaskFunc    = `wrong mask function`
	WrongArgValue    = `wrong arguments`
	WrongQuotesValue = `ANSI_QUOTES mode values not support double quotes`
	WrongColumnsCnt  = `columns count mismatch`
	WrongLargeThan0  = `n should large than 0`
)

const (
	WatermarkPrefix = `
<style type="text/css">
	.watermarked {
	  position: fixed;
	  overflow: hidden;
	}

	.watermarked::before {
	  position: absolute;
	  top: -75%%;
	  left: -75%%;

	  display: block;
	  width: 150%%;
	  height: 150%%;
	  z-index: -1;

	  transform: rotate(-45deg);
	  content: attr(data-watermark);

	  opacity: 0.5;
	  line-height: 3em;
	  letter-spacing: 2px;
	  color: #ccc;
	}
</style>

<div class="watermarked" data-watermark="%s">
`
	WatermarkSuffix = `
</div>

<script type="text/javascript">
	Array.from(document.querySelectorAll('.watermarked')).forEach(function(el) {
	  el.dataset.watermark = (el.dataset.watermark + ' ').repeat(300);
	});
</script>`
)
