package main

import (
	"sync"
    "math/rand"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
    block := 64
    length := len(src)
    result := make([]int64, length)
    if length < block {
        Sort(src, 0, length-1, result, 10)
    } else {
        InsertSort(src, block)
        //SplitQuickSort(src, block)
        var wait sync.WaitGroup
        for lens := block; lens < length; lens*=2 {
            total := (length + lens*2 -1) / (lens *2)
            wait.Add(total)
            for current := 0; current < total; current++ {
                go Merges(src, current, lens, &wait, result)
            }
           wait.Wait()
        }
    }
}

func Merges(src []int64, current,lens int, wait *sync.WaitGroup, result []int64) {
    defer wait.Done()
    length := len(src)
    lowLeft := 2*current*lens
    lowRight := minInt(lowLeft+lens, length)
    highLeft := lowRight
    highRight := minInt(highLeft+lens, length)

    leftPtr, rightPtr, resultPtr := lowLeft, lowRight, lowLeft

    for leftPtr < highLeft && rightPtr < highRight {
        if src[leftPtr] < src[rightPtr] {
            result[resultPtr] = src[leftPtr]
            leftPtr++
        }else {
            result[resultPtr] = src[rightPtr]
            rightPtr++
        }
        resultPtr++
    }

    for i := leftPtr; i < highLeft; i++ {
        result[resultPtr] = src[i]
        resultPtr++
    }

    for i:= rightPtr; i < highRight; i++ {
        result[resultPtr] = src[i]
        resultPtr++
    }

    for i := lowLeft; i < highRight; i++ {
        src[i] = result[i]
    }
}

func InsertSort(src []int64, block int) {
    length := len(src)
    for current := 0; current < length; current+=block {
        end := minInt(current+block, length)
        for i := current + 1; i < end; i++ {
            for j := i; j > current && src[j-1] > src[j]; j-- {
                src[j-1], src[j] = src[j], src[j-1]
            }
        }
    }
}

func SplitQuickSort(src []int64, block int) {
    length := len(src)
    for current := 0; current < length; current+=block {
        end :=  minInt(current+block, length)
        QuickSort(src[current:end])
    }
}

func QuickSort(src []int64) {
    length := len(src)
    if length == 0 {
        return
    }
    left, right := 0, length - 1
    rand := rand.Int() % length
    src[right], src[rand] = src[rand], src[right]
    for i := 0; i < length; i++ {
        if src[i] < src[right] {
            src[left], src[i] = src[i], src[left]
            left++
        }
    }
    src[left], src[right] = src[right], src[left]
    QuickSort(src[:left])
    QuickSort(src[left+1:])
}

func minInt(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func Sort(src []int64, left int, right int, tmp []int64, thread int) {
	if left < right {
		mid := (left + right) / 2
		if thread < len(src) {
			Sort(src, left, mid, tmp, thread)
			Sort(src, mid+1, right, tmp, thread)
		} else {
			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				defer func() { wg.Done() }()
				Sort(src, left, mid, tmp, thread)
			}()

			go func() {
				defer func() { wg.Done() }()
				Sort(src, mid+1, right, tmp, thread)
			}()
            wg.Wait()
		}

        Merge(left, right, mid, src, tmp)
    }
}

func Merge(left int, right int, mid int, src []int64, tmp []int64) {
	i := left
	j := mid + 1
	t := 0
	for i <= mid && j <= right {
		if src[i] <= src[j] {
			tmp[t], t, i = src[i], t+1, i+1
		} else {
			tmp[t], t, j = src[j], t+1, j+1
		}
	}

	for i <= mid {
		tmp[t], t, i = src[i], t+1, i+1
	}

	for j <= right {
		tmp[t], t, j = src[j], t+1, j+1
	}

	t = 0
	for left <= right {
		src[left], left, t = tmp[t], left+1, t+1
	}
}

//func main() {
//    arr := []int64{3,6,8,0,1,7,34,67,90,12,54}
//    MergeSort(arr)
//    fmt.Println(arr)
//}
