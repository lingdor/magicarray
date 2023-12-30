package utils

import "golang.org/x/exp/constraints"

func In[T constraints.Ordered](v T, vals ...T) bool {
	for _, val := range vals {
		if val == v {
			return true
		}
	}
	return false
}
