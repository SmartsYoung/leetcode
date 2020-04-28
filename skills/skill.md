
```
func digui(nums []int, num []bool, arr []int, res *[][]int) {
	if len(arr) == len(nums) {
		//*res = append(*res, append([]int{}, arr...))  //二维切片插入一维切片的两种方式
	
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*res = append(*res, tmp)

        //*res = append(*res, arr)      //这种方式有时结果不对，原因是切片arr本身可能因扩容导致传值仍为原来的底层数组
        return
     }
}
```


```
//for循环平行赋值
for i, j := 1, 10; i < j; i,j=i+1,j+1 {  //死循环
    fmt.Println(i)
}
//不能写成
for i, j := 1, 10; i < j; i++,j++ {
    fmt.Println(i)
}
//第二个表达式 i<j 不能有分号，如 i>0, j>0 可以用 && ||等连接


```