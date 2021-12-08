package main

import (
	"fmt"
	"os"
)

func HelloWorld() {
	file, err := os.Open("hello.txt")
	defer file.Close() // defer 용법. 파일을 열었기 때문에 작업이 끝나면 파일을 닫는다.
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 1024)
	fmt.Println("buffer =< ", buffer)
	if _, err = file.Read(buffer); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buffer))
}

// defer
// defer는 함수 내에서 제일 나중에, 즉 끝나기 직전에 실행하는 용법.
func main() {

	HelloWorld()
	fmt.Println("Finished!!")

}

/* 1번
// defer fmt.Println("World")
   fmt.Println("Hello")
*/

/* 2번
defer func() {
	fmt.Println("defer is called at the last")
}()

fmt.Println("This is a main call")
*/

// 3번
/*
	var num1, num2 int = 50, 0

	defer fmt.Println("Finished!!!")

	result := num1 / num2
	fmt.Println(result)
}
*/

/* make 함수를 통해서 slice 생성
func main() {
	x := []int{2, 4, 6, 8, 10, 12, 14}
	fmt.Println(len(x)) // 길이 측정
	fmt.Println(cap(x)) // 용량 측정

	// []int 타입의 슬라이스는 len이 4개, cap이 6개
	slice := make([]int, 4, 6) // make(슬라이스 타입, 슬라이스 길이, 슬라이스 용량(생략 가능))
	fmt.Println(slice)

	slice = append(slice, 1)
	length := len(slice)
	cap := cap(slice)
	fmt.Println(slice, length, cap)
}
*/

/* slice~
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
