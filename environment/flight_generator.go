package environment

import (
    "math/rand/v2"
    "time"

    "gonum.org/v1/gonum/stat/distuv"
)

// SampleFlight 从 DISTRICTS 中随机抽取两个不同区，并随机生成飞行时长（10–60 分钟）
func SampleFlight() ([]District, time.Duration) {
    n := len(DISTRICTS)

    // 用当前时间戳做种子
	now := uint64(time.Now().UnixNano())
    src := rand.NewPCG(now, now^0xDEADBEEF)
    rnd := rand.New(src)

	weights := make([]float64, n)
	for i := 0; i < n; i++ {
		weights[i] = 1.0/float64(n)
	}

    // 等概率抽取两个不同的索引
    cat := distuv.NewCategorical(weights,src)

    i := cat.Rand()
    j := cat.Rand()
    for j == i {
        j = cat.Rand()
    }

    // 飞行时长服从 10 到 60 分钟的均匀分布
    durationDist := distuv.Uniform{
        Min: 10,
        Max: 60,
        Src: rnd,
    }
    minutes := durationDist.Rand()
    duration := time.Duration(minutes) * time.Minute

    return []District{DISTRICTS[int(i)], DISTRICTS[int(j)]}, duration
}