// Copyright 2021 MIKAMAI s.r.l
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package predicate

import (
	"github.com/mikamai/karavel/cli/internal/utils"
	"strings"
)

type StringPredicate func(s string) bool

func StringOr(predicates ...StringPredicate) StringPredicate {
	if len(predicates) == 0 {
		panic("predicates list cannot be empty")
	}

	return func(s string) bool {
		res := false
		for _, p := range predicates {
			res = res || p(s)
		}
		return res
	}
}

func StringTrue(_ string) bool {
	return true
}

func StringFalse(_ string) bool {
	return false
}

func IsStringInSlice(haystack []string) StringPredicate {
	return func(s string) bool { return utils.IsStringInSlice(haystack, s) }
}

func StringHasPrefix(prefix string) StringPredicate {
	return func(s string) bool {
		return strings.HasPrefix(s, prefix)
	}
}
