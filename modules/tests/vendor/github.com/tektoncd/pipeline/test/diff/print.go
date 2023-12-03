/*
Copyright 2020 The Tekton Authors

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

package diff

import "fmt"

// PrintWantGot takes a diff string generated by cmp.Diff and returns it
// in a consistent format for reuse across all of our tests. This
// func assumes that the order of arguments passed to cmp.Diff was
// (want, got) or, in other words, the expectedResult then the actualResult.
func PrintWantGot(diff string) string {
	return fmt.Sprintf("(-want, +got): %s", diff)
}
