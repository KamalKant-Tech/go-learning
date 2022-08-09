package opertation

// func main() {

// 	data := []struct{ in, out []string }{
// 		{[]string{}, []string{}},
// 		{[]string{"", "", ""}, []string{""}},
// 		{[]string{"a", "a"}, []string{"a"}},
// 		{[]string{"a", "b", "a"}, []string{"a", "b"}},
// 		{[]string{"a", "b", "a", "b"}, []string{"a", "b"}},
// 		{[]string{"a", "b", "b", "a", "b"}, []string{"a", "b"}},
// 		{[]string{"a", "a", "b", "b", "a", "b"}, []string{"a", "b"}},
// 		{[]string{"a", "b", "c", "a", "b", "c"}, []string{"a", "b", "c"}},
// 	}

// 	for _, exp := range data {
// 		res := UniqStr(exp.in)
// 		if !reflect.DeepEqual(res, exp.out) {
// 			log.Fatalf("%q didn't match %q\n", res, exp.out)
// 		}
// 	}
// }

// UniqStr returns a copy if the passed slice with only unique string results.
func UniqStr(col []string) []string {
	m := map[string]struct{}{}
	for _, v := range col {
		_, ok := m[v]
		if !ok {
			m[v] = struct{}{}
		}
	}
	list := make([]string, len(m))

	i := 0
	for v := range m {
		list[i] = v
		i++
	}
	return list
}
