package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type WeatherEntry struct {
	gorm.Model
	Timestamp time.Time
	N         int
	Ci        bool
	Nh        int
	Cm        bool
	RRR       int
}

func main() {
	f, err := os.Open("weather.csv")
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(sqlite.Open("file.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&WeatherEntry{})
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(f)
	reader.Comma = ';'
	_, _ = reader.Read()
	var strs []string
	for i := 1; ; i++ {
		if strs, err = reader.Read(); err != nil {
			log.Fatal(err)
		}

		ts, _ := time.Parse("02.01.2006 15:04", strs[0])

		e := WeatherEntry{
			Timestamp: ts,
			N:         n(strs[1]),
			Ci:        ci(strs[2]),
			Nh:        n(strs[3]),
			Cm:        cm(strs[5]),
		}

		if ts.Hour() < 6 || ts.Hour() > 18 {
			continue
		}
		db.Create(&e)
	}

}

func n(in string) int {
	switch {
	case strings.HasPrefix(in, "Небо"):
		return 100
	case strings.HasPrefix(in, "100"):
		return 100
	case strings.HasPrefix(in, "90"):
		return 95
	case strings.HasPrefix(in, "60"):
		return 60
	case strings.HasPrefix(in, "10"):
		return 5
	case strings.HasPrefix(in, "40"):
		return 40
	case strings.HasPrefix(in, "70"):
		return 75
	case strings.HasPrefix(in, "50"):
		return 50
	case strings.HasPrefix(in, "20"):
		return 25
	default:
		return 0
	}
}

func ci(in string) bool {
	switch {
	case strings.HasPrefix(in, "Купчасто-дощові"):
		return true
	case strings.HasPrefix(in, "Шаруваті розірвані або купчасті розірвані хмари поганої погоди"):
		return true

	default:
		return false
	}
}

func cm(in string) bool {
	switch {
	case strings.HasPrefix(in, "Високо-купчасті прозорі або щільні у двох або більше шарах чи високо-купчасті щільні в одному шарі, що не поширюються по небу, або високо-купчасті з високо-шаруватими або шарувато-дощовими."):
		return true
	case strings.HasPrefix(in, "Високо-шаруваті непрозорі або шарувато-дощові."):
		return true

	default:
		return false
	}
}
