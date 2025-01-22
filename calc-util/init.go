package calc_util

/*
Intersection computes the intersection of two slices.

Example:
a := []int{1, 2, 3, 4, 5}
b := []int{2, 4, 5, 6}
result := Intersection(a, b)
// result: {2, 4, 5}

Notes:
- Slice `a` is not deduplicated.
*/
func Intersection[T comparable](a, b []T) []T {
	k := []T{}
	m := map[T]bool{}
	for _, v := range b {
		m[v] = true
	}
	for _, v := range a {
		if m[v] {
			k = append(k, v)
		}
	}
	return k
}

/*
Sub computes the difference of two slices.

Example:
a := []int{1, 2, 3, 4, 5}
b := []int{2, 4, 5, 6}
aMinusB := Sub(a, b) // {1, 3}
bMinusA := Sub(b, a) // {6}
*/
func Sub[T comparable](a, b []T) []T {
	k := []T{}
	m := map[T]bool{}
	for _, v := range b {
		m[v] = true
	}
	for _, v := range a {
		if !m[v] {
			k = append(k, v)
		}
	}
	return k
}

/*
Com computes the union of two slices.

Example:
a := []int{1, 2, 3, 4, 5}
b := []int{2, 4, 5, 6}
result := Com(a, b)
// result: {1, 2, 3, 4, 5, 6}
*/
func Com[T comparable](a, b []T) []T {
	m := map[T]bool{}
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if !m[v] {
			a = append(a, v)
		}
	}
	return a
}

/*
UpdateListWithList updates elements in slice `a` with values from slice `b`
based on provided key functions and an update function.

Example:

	a := []struct{id, age int; c interface{}}{
	    {id: 1, age: 1, c: nil},
	    {id: 2, age: 2, c: nil},
	}

	b := []struct{uid int; c interface{}}{
	    {uid: 1, c: map[string]string{"key": "value1"}},
	    {uid: 2, c: map[string]string{"key": "value2"}},
	}

keyFuncA := func(a struct{id, age int; c interface{}}) int { return a.id }
keyFuncB := func(b struct{uid int; c interface{}}) int { return b.uid }

	updateFunc := func(a struct{id, age int; c interface{}}, b struct{uid int; c interface{}}) struct{id, age int; c interface{}} {
	    a.c = b.c
	    return a
	}

result := UpdateListWithList(a, b, keyFuncA, keyFuncB, updateFunc)
// result: [{id: 1, age: 1, c: {"key": "value1"}}, {id: 2, age: 2, c: {"key": "value2"}}]
*/
func UpdateListWithList[A any, B any, K comparable](
	a []A,
	b []B,
	keyFuncA func(A) K,
	keyFuncB func(B) K,
	updateFunc func(A, B) A,
) []A {
	bMap := make(map[K]B)
	for _, itemB := range b {
		key := keyFuncB(itemB)
		bMap[key] = itemB
	}
	var res []A
	for i := range a {
		key := keyFuncA(a[i])
		if itemB, exists := bMap[key]; exists {
			a[i] = updateFunc(a[i], itemB)
		}
		res = append(res, a[i])
	}
	return res
}

/*
FilterListByList filters elements in slice `a` based on the presence of corresponding elements in slice `b`.

Example:

	a := []struct{id, age int}{
	    {id: 1, age: 1},
	    {id: 2, age: 2},
	    {id: 3, age: 3},
	}

	b := []struct{uid int}{
	    {uid: 1},
	    {uid: 3},
	}

keyFuncA := func(a struct{id, age int}) int { return a.id }
keyFuncB := func(b struct{uid int}) int { return b.uid }
result := FilterListByList(a, b, keyFuncA, keyFuncB)
// result: [{id: 1, age: 1}, {id: 3, age: 3}]
*/
func FilterListByList[A any, B any, K comparable](
	a []A,
	b []B,
	keyFuncA func(A) K,
	keyFuncB func(B) K,
) []A {
	var m = make(map[K]bool)
	for _, cur := range b {
		m[keyFuncB(cur)] = true
	}
	var res []A
	for _, cur := range a {
		if m[keyFuncA(cur)] {
			res = append(res, cur)
		}
	}
	return res
}
