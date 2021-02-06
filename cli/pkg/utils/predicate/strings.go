package predicate

import (
	"github.com/mikamai/karavel/cli/pkg/utils"
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
