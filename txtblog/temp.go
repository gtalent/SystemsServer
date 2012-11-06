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
	"strconv"
	"strings"
)

func content(txt string) string {
	out := `
			<div>
		`
	txt = ml(txt)
	p := strings.Split(txt, "\n")
	for _, a := range p {
		s := `
					<p style="margin:auto; width:100%; text-ident:50pt;">
						&nbsp; &nbsp; &nbsp;` + a + `
					</p>`
		out += s
	}
	out += `
			</div>`
	return out
}

func note(txt string) string {
	out := `
			<div>
				<br><br>
				<strong>Note:</strong>
		`
	p := strings.Split(txt, "\n")
	for _, a := range p {
		s := `
					<p style="margin:auto; width:100%; text-ident:50pt;">
						&nbsp; &nbsp; &nbsp;` + a + `
					</p>`
		out += s
	}
	out += `
			</div>`
	return out
}

func temp(title, cont, n string, postNo int) string {
	p := strings.Replace(post, "{{ title }}", title, -1)
	prev := postNo - 1
	if prev < 0 {
		prev = 0
	}
	next := postNo + 1
	if next >= len(posts) {
		next = len(posts) - 1
	}
	p = strings.Replace(p, "{{ prev }}", strconv.Itoa(prev), -1)
	p = strings.Replace(p, "{{ nex }}", strconv.Itoa(next), -1)
	p += content(cont)
	if n != "" {
		p += note(n)
	}
	p += `</div></body></html>`
	return p
}

const post = `
<html>
	<head>
		<style type="text/css">
			.NavButton {
				display: inline;
				background: #cccc99;
			}
			.NavBar {
				display: inline;
				width: 50%;
			}

			a:link {
			  	text-decoration:none;
				color:black;
			}
			a:visited {
			  	text-decoration:none;
				color:black;
			}
			a:hover {
			  	text-decoration:none;
				color:white;
			}
			a:active {
			  	text-decoration:none;
				color:white;
			}
		</style>
		<title>` + blogTitle + `: {{ title }}</title>
	</head>
	<body>
		<h1 style="text-align:center">` + blogTitle + `</h1>
		<h3 style="text-align:center">{{ title }}</h3>

		<center>
			<div id="NavBar">
				<a href="/posts/{{ prev }}">
					<div class="NavButton">&nbsp;&nbsp;&nbsp;&lt&nbsp;&nbsp;&nbsp;</div>
				</a>
				<a href="/posts/{{ nex }}">
					<div class="NavButton">&nbsp;&nbsp;&nbsp;&gt&nbsp;&nbsp;&nbsp;</div>
				</a>
			</div>
		</center>

		<div style="margin: auto; width:50%; text-ident:50pt;">
`
