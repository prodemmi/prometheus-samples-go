package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/go-faker/faker/v4"
)

func randRange(min int64, max int64) int64 {
	return rand.Int64N(max) + min
}

func randate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int64N(delta) + min
	return time.Unix(sec, 0)
}

func ranbool() bool {
	return rand.IntN(2) == 1
}

func SetupUserSeeder(m *UserMetrics) {
	var lastMaxOrder int64
	lastMaxOrder = 0

	for i := 0; i < 10_000; i++ {
		loginAt, _ := time.Parse("2006-01-02", faker.Date())
		coutries := []string{"Iran", "America", "China"}
		newUser := User{
			ID:          i,
			Username:    fmt.Sprintf("@%v", faker.Username()),
			OrderCount:  rand.Int64N(20),
			LastOrderAt: randate(),
			MaxPayment:  rand.Int64N(5_000_000),
			IsOnline:    ranbool(),
			IP:          faker.IPv4(),
			Country:     coutries[rand.IntN(len(coutries))],
			LoginAt:     loginAt,
			CreatedAt:   time.Now(),
		}

		m.CountMaxOrder.Set(float64(newUser.MaxPayment))

		if newUser.IsOnline {
			m.CountOnline.Inc()
		} else {
			m.CountOnline.Dec()
		}

		if newUser.MaxPayment >= lastMaxOrder {
			m.MaxPaymentByIP.WithLabelValues(newUser.Country).Set(float64(newUser.MaxPayment))
			lastMaxOrder = newUser.MaxPayment
		}

		fmt.Printf("User %d Seted.\n", i)
		time.Sleep(time.Millisecond * 500)
	}
}
