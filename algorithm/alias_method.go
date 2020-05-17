package algorithm

import (
	"errors"
	"fmt"
	. "github.com/wenlincheng/go-common/collection"
	"math/rand"
	"time"
)

// 抽奖算法实现
type AliasMethod struct {
	probability []float64
	alias       []int
	length      int
}

// 初始化抽奖池
func (a *AliasMethod) Initialize(prob []float64) error {
	// 数据校验
	if len(prob) == 0 {
		fmt.Print("概率数组必须为非空")
		return errors.New("概率数组必须为非空")
	}

	// 概率列表大小
	a.length = len(prob)
	a.probability = make([]float64, a.length)
	a.alias = make([]int, a.length)
	probtemp := make([]float64, a.length)

	// 双向队列
	small := LinkedList{}
	large := LinkedList{}

	for i := 0; i < a.length; i++ {
		// 初始化 probtemp
		probtemp[i] = prob[i] * float64(a.length)
		if probtemp[i] < 1.0 {
			small.Add(i)
		} else {
			large.Add(i)
		}
	}

	for !small.IsEmpty() && !large.IsEmpty() {
		err, less := small.Pop()
		if err != nil {
			return err
		}
		err, more := large.Pop()
		if err != nil {
			return err
		}
		a.probability[less.(int)] = probtemp[less.(int)]
		a.alias[less.(int)] = more.(int)
		probtemp[more.(int)] = probtemp[more.(int)] - (1.0 - a.probability[less.(int)])
		if probtemp[more.(int)] < 1.0 {
			small.Add(more)
		} else {
			large.Add(more)
		}
	}

	for !small.IsEmpty() {
		err, less := small.Pop()
		if err != nil {
			return err
		}
		a.probability[less.(int)] = 1.0
	}
	for !large.IsEmpty() {
		err, more := large.Pop()
		if err != nil {
			return err
		}
		a.probability[more.(int)] = 1.0
	}

	return nil
}

// 抽奖
func (a *AliasMethod) Next() int {
	rand.Seed(time.Now().UnixNano())
	column := rand.Intn(a.length)
	coinToss := rand.Float64() < a.probability[column]
	if coinToss {
		return column
	} else {
		return a.alias[column]
	}
}
