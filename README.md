
# calc_util

`calc_util` 提供用于处理切片的交集、差集、并集，以及复杂的列表更新和过滤操作的函数。

## 函数列表

### Intersection
**计算交集**
```go
func Intersection[T comparable](a, b []T) []T
```  
**示例**:
```go
a := []int{1, 2, 3, 4, 5}
b := []int{2, 4, 5, 6}
result := Intersection(a, b)
// result: {2, 4, 5}
```  

---  

### Sub
**计算差集**
```go
func Sub[T comparable](a, b []T) []T
```  
**示例**:
```go
a := []int{1, 2, 3, 4, 5}
b := []int{2, 4, 5, 6}
aMinusB := Sub(a, b) // {1, 3}
bMinusA := Sub(b, a) // {6}
```  

---  

### Com
**计算并集**
```go
func Com[T comparable](a, b []T) []T
```  
**示例**:
```go
a := []int{1, 2, 3, 4, 5}
b := []int{2, 4, 5, 6}
result := Com(a, b)
// result: {1, 2, 3, 4, 5, 6}
```  

---  

### UpdateListWithList
**用另一个切片的数据更新目标切片**
```go
func UpdateListWithList[A any, B any, K comparable](
    a []A,
    b []B,
    keyFuncA func(A) K,
    keyFuncB func(B) K,
    updateFunc func(A, B) A,
) []A
```  
**示例**:
```go
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
```  

---  

### FilterListByList
**根据另一个切片过滤目标切片**
```go
func FilterListByList[A any, B any, K comparable](
    a []A,
    b []B,
    keyFuncA func(A) K,
    keyFuncB func(B) K,
) []A
```  
**示例**:
```go
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
```  

---  

## 注意事项
- 支持所有 `comparable` 类型。
- 根据需求提供正确的 `keyFunc` 和 `updateFunc`。  
