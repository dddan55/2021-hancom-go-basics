package main

import "fmt"

// Slice
func main() {
	// Slice
	rainbow := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	rainbow = append(rainbow, "sky", "pink", "purple")
	fmt.Println(rainbow)
}

// array 와 slice
/*
func changeArr(arr [3]int) {
	arr[0] = 100 // 함수에 인자로 들어간 arr이라는 파라미터가 함수 스코프(중괄호) 내에서 지역변수로 생성. 즉, 원본의 메모리 주소 자체와 다름.

	fmt.Println("This is Array ==>", arr)
}

func changeSlice(slice []int) {
	slice[0] = 100 // 여기에서의 slice는 참조값이므로 원본 값 자체를 변경.
	fmt.Println("This is Slice Func ==>", slice)

}

func main() {
	arr := [3]int{1, 2, 3} // {1,2,3}
	slice := []int{1, 2, 3, 4}

	changeArr(arr)
	changeSlice(slice)

	fmt.Println("Main func array ==>", arr)   // 원본은 바뀌지 않음.
	fmt.Println("Main func slice ==>", slice) // 참조값이므로 원본이 바뀜.
}
*/
