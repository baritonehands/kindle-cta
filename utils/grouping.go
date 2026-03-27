package utils

type GroupByFn[K comparable, T any] func(item T) K

func GroupBy[K comparable, T any](items []T, groupBy GroupByFn[K, T]) [][]T {
	groups := map[K][]T{}
	for _, busArrival := range items {
		groupKey := groupBy(busArrival)
		_, found := groups[groupKey]
		if !found {
			groups[groupKey] = []T{}
		}
		groups[groupKey] = append(groups[groupKey], busArrival)
	}

	ret := make([][]T, 0, len(groups))
	for _, group := range groups {
		if len(group) > 0 {
			ret = append(ret, group)
		}
	}
	return ret
}
