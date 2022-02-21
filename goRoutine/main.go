package main

import(
	"fmt"
	"time"
	"runtime"
	"sync"
)

func main(){
	case_origin()
	case_mutex()
}

func case_origin(){

	// 데이터 접근시 제어 없음

	var data = []int{} // int형 슬라이스 생성
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용
	//runtime.GOMAXPROCS(1) // 모든 CPU 사용

	fmt.Println(runtime.NumCPU()) 	// 본인 PC가 사용하고 있는 코어 및 HT(intel) or SMT(AMD) 갯수

	go func() {
		for i := 0; i < 10000; i++ {
			data = append(data, 1)

			runtime.Gosched()	// 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			data = append(data, 1)
			runtime.Gosched()
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println(len(data))
}	

func case_mutex(){

	// 데이터 접근시 뮤텍스 사용

	var data = []int{} // int형 슬라이스 생성
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용
	//runtime.GOMAXPROCS(1) // 모든 CPU 사용
	fmt.Println(runtime.NumCPU()) 	// 본인 PC가 사용하고 있는 코어 및 HT(intel) or SMT(AMD) 갯수

	var mutex = new(sync.Mutex)

	go func() {	// go rutine
		for i := 0; i < 10000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()

			runtime.Gosched()	// 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()

			runtime.Gosched()
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println(len(data))

}