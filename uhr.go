package uhr

import (
	"fmt"
	"sort"
	"time"
)

type formatFn func(t time.Time) string

func doFormatInformal(t time.Time, fn formatFn) []string {
	result := []string{fn(t)}
	if t.Hour() > 12 {
		result = append(result, fn(informalHour(t)))
	}
	return result
}

func informalHour(t time.Time) time.Time {
	if t.Hour() > 12 {
		return t.Add(-12 * time.Hour)
	}
	return t
}

func format(format string, args ...int) string {
	aargs := make([]any, 0, len(args))
	for _, arg := range args {
		aargs = append(aargs, number(arg))
	}
	return fmt.Sprintf(format, aargs...)
}

func Uhr(t time.Time) []string {
	result := []string{}
	ti := informalHour(t)

	if t.Minute() != 0 {
		result = append(result,
			doFormatInformal(t, func(t time.Time) string {
				return format("%s Uhr %s", t.Hour(), t.Minute())
			})...)
	}

	switch t.Minute() {
	case 0:
		result = append(result,
			doFormatInformal(t, func(t time.Time) string {
				return format("%s Uhr", t.Hour())
			})...)
		result = append(result,
			doFormatInformal(t, func(t time.Time) string {
				return format("punkt %s", t.Hour())
			})...)
	case 15:
		result = append(result, format("viertel nach %s", ti.Hour()))
	case 30:
		result = append(result, format("halb %s", ti.Hour()+1))
	case 45:
		result = append(result, format("viertel vor %s", ti.Hour()+1))
	}

	if t.Minute() > 0 && t.Minute() <= 20 {
		result = append(result, format("%s nach %s", t.Minute(), ti.Hour()))
		if t.Minute() < 5 {
			result = append(result, format("kurz nach %s", ti.Hour()))
		}
	}
	if t.Minute() >= 40 && t.Minute() <= 59 {
		result = append(result, format("%s vor %s", 60-t.Minute(), ti.Hour()+1))
		if t.Minute() > 55 {
			result = append(result, format("kurz vor %s", ti.Hour()+1))
		}
	}

	if t.Minute() >= 20 && t.Minute() < 30 {
		result = append(result, format("%s vor halb %s", 30-t.Minute(), ti.Hour()+1))
	}

	if t.Minute() > 30 && t.Minute() <= 40 {
		result = append(result, format("%s nach halb %s", t.Minute()-30, ti.Hour()+1))
	}

	sort.Strings(result)
	return result
}

type Hourer interface {
	Hour() int
}

func PartOfDay(h Hourer) string {
	t := h.Hour()
	if t >= 6 && t < 12 {
		return "am Morgen"
	}
	if t >= 12 && t < 14 {
		return "am Mittag"
	}
	if t >= 14 && t < 18 {
		return "am Nachmittag"
	}
	if t >= 18 && t < 22 {
		return "am Abend"
	}
	return "in der Nacht"
}

type Weekdayer interface {
	Weekday() time.Weekday
}

func Weekday(t Weekdayer) string {
	switch t.Weekday() {
	case time.Sunday:
		return "Sonntag"
	case time.Monday:
		return "Montag"
	case time.Tuesday:
		return "Dienstag"
	case time.Wednesday:
		return "Mittwoch"
	case time.Thursday:
		return "Donnerstag"
	case time.Friday:
		return "Freitag"
	case time.Saturday:
		return "Samstag"
	default:
		return ""
	}
}

func number(i int) string {
	switch abs(i) {
	case 0:
		return "null"
	case 1:
		return "eins"
	case 2:
		return "zwei"
	case 3:
		return "drei"
	case 4:
		return "vier"
	case 5:
		return "fünf"
	case 6:
		return "sechs"
	case 7:
		return "sieben"
	case 8:
		return "acht"
	case 9:
		return "neun"
	case 10:
		return "zehn"
	case 11:
		return "elf"
	case 12:
		return "zwölf"
	case 13:
		return "dreizehn"
	case 14:
		return "vierzehn"
	case 15:
		return "fünfzehn"
	case 16:
		return "sechzehn"
	case 17:
		return "siebzehn"
	case 18:
		return "achtzehn"
	case 19:
		return "neunzehn"
	case 20:
		return "zwanzig"
	case 30:
		return "dreißig"
	case 40:
		return "vierzig"
	case 50:
		return "fünfzig"
	}
	return handleEins(number(i%10)) + "und" + number((i/10)*10)
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return i * -1
}

func handleEins(s string) string {
	if s == "eins" {
		return "ein"
	}
	return s
}
