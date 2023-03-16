// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// pprof is a tool for collection, manipulation and visualization
// of performance profiles.
package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/et-zone/ppcli/consts"
	"github.com/et-zone/ppcli/driver"
	"github.com/et-zone/ppcli/job"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	router, err := driver.PPInit(&driver.Options{UI: newUI()})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	router.HandleFunc("/target/fresh", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()

		a := request.Form.Get("arg")
		u, err := url.Parse(a)
		if err != nil {
			http.Error(writer, "args err... arg="+a,
				http.StatusOK)
			return
		}
		if a == "" {
			http.Error(writer, "args can not nil,like: ?arg="+consts.ExportAddr+"/debug/pprof/profile?seconds=1",
				http.StatusOK)
			return
		}
		v := u.Query().Get("seconds")
		if v == "" {
			a = a + "?seconds=1"
			driver.UpdateSource([]string{a}, 1)
			return
		} else {
			in, err := strconv.Atoi(v)
			if err != nil {
				http.Error(writer, "?seconds err... "+a, http.StatusOK)
				return
			}
			msg, err := driver.UpdateSource([]string{a}, in)
			if err != nil {
				fmt.Println(err.Error())
				http.Error(writer, "exec err... "+err.Error(), http.StatusOK)
				return
			}
			http.Error(writer, msg, http.StatusOK)
			return
		}

	})

	// can close this code
	job.Run()

	log.Println("please click link to browse Info: ", " http://"+consts.IP+":"+consts.Port+"/target")
	log.Println("please click link to update Source: ", " http://"+consts.IP+":"+consts.Port+"/target/fresh")
	log.Println(http.ListenAndServe(":"+consts.Port, router))
}

// readlineUI implements the driver.UI interface using the
// github.com/chzyer/readline library.
// This is contained in pprof.go to avoid adding the readline
// dependency in the vendored copy of pprof in the Go distribution,
// which does not use this file.
type readlineUI struct {
	rl *readline.Instance
}

func newUI() driver.UI {
	rl, err := readline.New("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "readline: %v", err)
		return nil
	}
	return &readlineUI{
		rl: rl,
	}
}

// ReadLine returns a line of text (a command) read from the user.
// prompt is printed before reading the command.
func (r *readlineUI) ReadLine(prompt string) (string, error) {
	r.rl.SetPrompt(prompt)
	return r.rl.Readline()
}

// Print shows a message to the user.
// It is printed over stderr as stdout is reserved for regular output.
func (r *readlineUI) Print(args ...interface{}) {
	text := fmt.Sprint(args...)
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	fmt.Fprint(r.rl.Stderr(), text)
}

// PrintErr shows a message to the user, colored in red for emphasis.
// It is printed over stderr as stdout is reserved for regular output.
func (r *readlineUI) PrintErr(args ...interface{}) {
	text := fmt.Sprint(args...)
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	if readline.IsTerminal(int(syscall.Stderr)) {
		text = colorize(text)
	}
	fmt.Fprint(r.rl.Stderr(), text)
}

// colorize the msg using ANSI color escapes.
func colorize(msg string) string {
	var red = 31
	var colorEscape = fmt.Sprintf("\033[0;%dm", red)
	var colorResetEscape = "\033[0m"
	return colorEscape + msg + colorResetEscape
}

// IsTerminal returns whether the UI is known to be tied to an
// interactive terminal (as opposed to being redirected to a file).
func (r *readlineUI) IsTerminal() bool {
	return readline.IsTerminal(int(syscall.Stdout))
}

// WantBrowser starts a browser on interactive mode.
func (r *readlineUI) WantBrowser() bool {
	return r.IsTerminal()
}

// SetAutoComplete instructs the UI to call complete(cmd) to obtain
// the auto-completion of cmd, if the UI supports auto-completion at all.
func (r *readlineUI) SetAutoComplete(complete func(string) string) {
	// TODO: Implement auto-completion support.
}
