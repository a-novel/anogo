package anogo

import "reflect"

func SliceEqual(a, b interface{}) bool {
	if reflect.TypeOf(a).Kind() != reflect.Slice || reflect.TypeOf(b).Kind() != reflect.Slice {
		return false
	}

	as := reflect.ValueOf(a)
	bs := reflect.ValueOf(b)

	if as.Len() != bs.Len() {
		return false
	}

	for i := 0; i < as.Len(); i++ {
		found := false
		aval := as.Index(i).Interface()

		for j := 0; j < bs.Len(); j++ {
			if bs.Index(j).Interface() == aval {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}
