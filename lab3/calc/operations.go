package calc

import (
	"errors"
	
)

func init() {
}

func Sum(nums ...float64) float64 {
	var sum float64
	for _, n := range nums {
		sum += n
	}
	return sum
}

func Max(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	maxVal := nums[0]
	for _, n := range nums {
		if n > maxVal {
			maxVal = n
		}
	}
	return maxVal
}

func Min(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	minVal := nums[0]
	for _, n := range nums {
		if n < minVal {
			minVal = n
		}
	}
	return minVal
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("помилка: ділення на нуль заборонено")
	}
	return a / b, nil
}

type Calculator interface {
	Sum(nums ...float64) float64
	Max(nums ...float64) float64
	Min(nums ...float64) float64
	Divide(a, b float64) (float64, error)
}

type Calc struct{}


func (c Calc) Sum(nums ...float64) float64 {
	return Sum(nums...)
}

func (c Calc) Max(nums ...float64) float64 {
	return Max(nums...)
}

func (c Calc) Min(nums ...float64) float64 {
	return Min(nums...)
}

func (c Calc) Divide(a, b float64) (float64, error) {
	return Divide(a, b)
}
