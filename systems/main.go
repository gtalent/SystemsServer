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
package systems

import (
	"fmt"
	"net/http"
	"strconv"
)

func handle(path string, f func(http.ResponseWriter, *http.Request) string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		if url[len(url)-1] != '/' {
			w.Header().Set("Location", url+"/")
			w.WriteHeader(302)
		} else {
			fmt.Fprintf(w, f(w, r))
		}
	})
}

func init() {
	go handle("/", home)
}

func thirdSlash(str string) int {
	c := 0
	for i, a := range str {
		if a == '/' {
			c++ //lol
			if c == 3 {
				return i
			}
		}
	}
	return -1
}

func home(w http.ResponseWriter, r *http.Request) string {
	url := r.URL.String()
	if len(url) > 3 && url[:4] == "http" {
		url = url[thirdSlash(url):]
	}
	if url == "/" {
		w.Header().Set("Location", "/posts/0/")
		w.WriteHeader(302)
		return ""
	} else if len(url) < 5 || url[1:6] != "posts" {
		return "Invalid URL: " + url
	}
	url = url[7 : len(url)-1]
	postID, err := strconv.Atoi(url)
	if err != nil || postID < 0 {
		return "Invalid post ID: " + url
	}
	return getPost(postID, r)
}
