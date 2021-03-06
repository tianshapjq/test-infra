/*
Copyright 2016 The Kubernetes Authors.

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

package plugins

import (
	"fmt"
	"strings"

	"k8s.io/test-infra/prow/github"
)

const AboutThisBot = "Instructions for interacting with me using PR comments are available [here](https://github.com/kubernetes/community/blob/master/contributors/devel/pull-request-commands.md).  If you have questions or suggestions related to my behavior, please file an issue against the [kubernetes/test-infra](https://github.com/kubernetes/test-infra/issues/new?title=Prow%20issue:) repository. I understand the commands that are listed [here](https://github.com/kubernetes/test-infra/blob/master/commands.md)."

// FormatResponse nicely formats a response to an issue comment.
func FormatResponse(ic github.IssueComment, s string) string {
	format := `@%s: %s.

<details>

In response to [this comment](%s):

%s

%s
</details>
`
	// Quote the user's comment by prepending ">" to each line.
	var quoted []string
	for _, l := range strings.Split(ic.Body, "\n") {
		quoted = append(quoted, ">"+l)
	}
	return fmt.Sprintf(format, ic.User.Login, s, ic.HTMLURL, strings.Join(quoted, "\n"), AboutThisBot)
}
