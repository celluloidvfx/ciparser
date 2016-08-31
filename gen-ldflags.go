/*
 * Johannes Amorosa, (C) 2016
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	b64 "encoding/base64"
)

func GenLDFlags(cfg *CiConfig) string {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	c := cfg.GetValueName("cvars")

	var ldflagsStr string
	ldflagsStr = "-X main.appVersion=" + timestamp + " "
	ldflagsStr = ldflagsStr + "-X main.appReleaseTag=" + releaseTag(timestamp) + " "
	ldflagsStr = ldflagsStr + "-X main.appCommitID=" + commitID() + " "
	ldflagsStr = ldflagsStr + "-X main.appShortCommitID=" + commitID()[:12] + " "
	ldflagsStr = ldflagsStr + "-X main.appBranch=" + branch() + " "

	for _, v := range c.([]Customvars) {

		if os.Getenv(strings.ToUpper(v.Name)) != "" {
			v.Path = os.Getenv(strings.ToUpper(v.Name))
		}

		if v.Value != "" && v.Path == "" {
			ldflagsStr = ldflagsStr + "-X main.app" + v.Name + "=" + b64.StdEncoding.EncodeToString([]byte(v.Value)) + " "
		}
		if v.Value == "" && v.Path != "" {
			ldflagsStr = ldflagsStr + "-X main.app" + v.Name + "=" + readFileContent(v) + " "
		}
		if v.Value != "" && v.Path != "" {
			fmt.Println("Error reading customvars path and value are mutual exclusive")
			os.Exit(1)
		}

	}

	ldflagsStr = ldflagsStr + "-linkmode external -extldflags \"-static\" -s -w"
	return ldflagsStr
}
func readFileContent(v Customvars) string {
	var (
		out []byte
		e   error
	)
	if out, e = ioutil.ReadFile(v.Path); e != nil {
		fmt.Fprintln(os.Stderr, "Error reading content "+string(out)+":", e)
		os.Exit(1)
	}
	return b64.StdEncoding.EncodeToString(out)
}

func branch() string {
	// git rev-parse --abbrev-ref HEAD
	var (
		branch []byte
		e      error
	)
	cmdName := "git"
	cmdArgs := []string{"rev-parse", "--abbrev-ref", "HEAD"}
	if branch, e = exec.Command(cmdName, cmdArgs...).Output(); e != nil {
		fmt.Fprintln(os.Stderr, "Error generating git branch: ", e)
		os.Exit(1)
	}
	return strings.TrimSpace(string(branch))
}

// genReleaseTag prints release tag to the console for easy git tagging.
func releaseTag(version string) string {
	relPrefix := "DEVELOPMENT"
	if prefix := os.Getenv("APP_RELEASE"); prefix != "" {
		relPrefix = prefix
	}

	relTag := strings.Replace(version, " ", "-", -1)
	relTag = strings.Replace(relTag, ":", "-", -1)
	relTag = strings.Replace(relTag, ",", "", -1)
	return relPrefix + "." + relTag
}

// commitID returns the abbreviated commit-id hash of the last commit.
func commitID() string {
	// git log --format="%h" -n1
	var (
		commit []byte
		e      error
	)
	cmdName := "git"
	cmdArgs := []string{"log", "--format=%H", "-n1"}
	if commit, e = exec.Command(cmdName, cmdArgs...).Output(); e != nil {
		fmt.Fprintln(os.Stderr, "Error generating git commit-id: ", e)
		os.Exit(1)
	}
	return strings.TrimSpace(string(commit))
}
