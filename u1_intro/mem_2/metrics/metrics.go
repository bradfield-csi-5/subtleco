package metrics

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
)

type UserData struct {
	ages     []uint8
	payments []uint32
}

func AverageAge(users UserData) float64 {
	total := uint64(0)
	user_count := len(users.ages)
	i := 0
	for ; i < user_count-3; i += 4 {
		a_1, a_2, a_3, a_4 := users.ages[i], users.ages[i+1], users.ages[i+2], users.ages[i+3]
		// total += uint64(a_1 + a_2 + a_3 + a_4) THIS WILL OVERFLOW!
		total += uint64(a_1) + uint64(a_2) + uint64(a_3) + uint64(a_4)
	}
	for ; i < user_count; i++ {
		age := users.ages[i]
		total += uint64(age)
	}

	return float64(total) / float64(user_count)
}

func AveragePaymentAmount(users UserData) float64 {
	amount := uint64(0)
	totalPayments := len(users.payments)
	for _, p := range users.payments {
		amount += uint64(p)
	}
	return float64(amount/uint64(totalPayments)) * 0.01
}

// Compute the standard deviation of payment amounts
// Variance[x] = E[x^2] - E[x]^2
func StdDevPaymentAmount(users UserData) float64 {
	sumSquare, sum := float64(0), float64(0)
	count := float64(len(users.payments))
	for _, p := range users.payments {
		amount := float64(p) * 0.01
		sumSquare += amount * amount
		sum += amount
	}
	avgSquare := sumSquare / float64(count)
	avg := sum / float64(count)
	return math.Sqrt(avgSquare - avg*avg)
}

func LoadData() UserData {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	reader := csv.NewReader(f)
	userLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse users.csv as csv", err)
	}

	ages := make([]uint8, len(userLines))
	for i, line := range userLines {
		age, _ := strconv.Atoi(line[2])
		ages[i] = uint8(age)
	}

	f, err = os.Open("payments.csv")
	if err != nil {
		log.Fatalln("Unable to read payments.csv", err)
	}
	reader = csv.NewReader(f)
	paymentLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse payments.csv as csv", err)
	}

	payments := make([]uint32, len(paymentLines))
	for i, line := range paymentLines {
		paymentCents, _ := strconv.ParseUint(line[0], 10, 32)
		payments[i] = uint32(paymentCents)
	}

	return UserData{ages, payments}
}
