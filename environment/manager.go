package environment

import (
	"math/rand/v2"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

var RunningFlights chan Flight

func assign_flights() {
	// 用当前时间戳做种子
	now := uint64(time.Now().UnixNano())
	src := rand.NewPCG(now, now^0xDEADBEEF)
	rnd := rand.New(src)

	// 飞行时长服从 10 到 60 分钟的均匀分布
	sleepTimeSampler := distuv.Uniform{
		Min: 5,
		Max: 10,
		Src: rnd,
	}
	for {
		minutes := sleepTimeSampler.Rand()
		// duration := time.Duration(minutes) * time.Minute
		duration := time.Duration(minutes) * time.Second

		time.Sleep(duration)

		flight := SampleFlight()

		RunningFlights <- flight
	}
}

func init() {
	RunningFlights = make(chan Flight, 1000)
	go assign_flights()

}
