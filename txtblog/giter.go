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
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"appengine"
	"appengine/urlfetch"
)

type Index struct {
	List []PostHeader `json:"list"`
}

type PostHeader struct {
	Path    string `json:"path"`
	Note    bool   `json:"note"`
	LastMod int    `json:"last_modified"`
}


var index *Index
var posts []string

func get(path string, r *http.Request) string {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get(path)
	if err != nil {
		return ""
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(out)
}

func getPost(i int, r *http.Request) string {
	in := index
	if in == nil {
		indexFile := get(root+"index.json", r)
		in = new(Index)
		json.Unmarshal([]byte(indexFile), in)
		posts = make([]string, len(in.List))
		index = in
	}
	if i >= len(in.List) {
		return "Invalid post ID: " + strconv.Itoa(i)
	}
	posts := posts
	if posts[i] == "" {
		note := ""
		a := in.List[i]
		if a.Note {
			no := strconv.Itoa(i)
			if len(no) == 1 {
				no = "0" + no
			}
			note = get(root+"notes/"+no+".txt", r)
		}
		content := get(root+a.Path, r)
		title := a.Path[3:strings.LastIndex(a.Path, ".")]
		title = strings.Replace(title, "_", " ", -1)
		posts[i] = temp(title, content, note, i)
	}
	return posts[i]
}
