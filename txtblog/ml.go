/*
   Copyright 2012 gtalent2@gmail.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package txtblog

import (
	"html"
	"strings"
)

func ml(str string) string {
	str = html.EscapeString(str)
	bs := 20
	ea := 21
	bold := 22
	ital := 23
	str = strings.Replace(str, "\\\\", string(bs), -1)
	str = strings.Replace(str, "\\*", string(ea), -1)
	str = strings.Replace(str, "**", string(bold), -1)
	str = strings.Replace(str, "*", string(ital), -1)
	depth := 0
	for i := 0; i < len(str); i++ {
		v := int(str[i])
		if v == ital {
			if depth == 1 {
				str = strings.Replace(str, string(ital), "</i>", 1)
				depth -= 1
			} else {
				str = strings.Replace(str, string(ital), "<i>", 1)
				depth += 1
			}
		} else if v == bold {
			if depth == 1 {
				str = strings.Replace(str, string(bold), "</b>", 1)
				depth -= 1
			} else {
				str = strings.Replace(str, string(bold), "<b>", 1)
				depth += 1
			}
		}
	}
	str = strings.Replace(str, string(bs), "\\", -1)
	str = strings.Replace(str, string(ea), "*", -1)
	return str
}
