package uhr

import (
	"fmt"
	"time"
)

func Uhr(t time.Time) []string {
	result := []string{
		t.Format("15:04"),
	}

	if t.Minute() != 0 {
		result = append(
			result,
			t.Format("15 Ühr 04"),
			number(t.Hour())+" Ühr "+number(t.Minute()),
		)
		if t.Hour() > 12 {
			result = append(
				result,
				t.Format("3 Ühr 04"),
				number(t.Hour()-12)+" Ühr "+number(t.Minute()),
			)
		}
	}

	switch t.Minute() {
	case 0:
		result = append(
			result,
			t.Format("15 Ühr"),
			number(t.Hour())+" Ühr",
			t.Format("Punkt 15"),
			"Punkt "+number(t.Hour()),
		)
		if t.Hour() > 12 {
			result = append(
				result,
				t.Format("3 Ühr"),
				number(t.Hour()-12)+" Ühr",
				t.Format("Punkt 3"),
				"Punkt "+number(t.Hour()-12),
			)
		}
	case 15:
		result = append(
			result,
			t.Format("Viertel nach 15"),
			"Viertel nach "+number(t.Hour()),
		)
		if t.Hour() > 12 {
			result = append(
				result,
				t.Format("Viertel nach 3"),
				"Viertel nach "+number(t.Hour()-12),
			)
		}
	case 30:
		result = append(
			result,
			t.Add(time.Hour).Format("Halb 15"),
			"Halb "+number(t.Hour()+1),
		)
		if t.Hour() > 12 {
			result = append(
				result,
				t.Add(time.Hour).Format("Halb 3"),
				"Halb "+number(t.Hour()+1-12),
			)
		}
	case 45:
		result = append(
			result,
			t.Add(time.Hour).Format("Viertel vor 15"),
			"Viertel vor "+number(t.Hour()+1),
		)
		if t.Hour() > 12 {
			result = append(
				result,
				t.Add(time.Hour).Format("Viertel vor 3"),
				"Viertel vor "+number(t.Hour()+1-12),
			)
		}
	}

	if t.Minute() > 0 && t.Minute() <= 20 {
		result = append(
			result,
			fmt.Sprintf("%d nach %d", t.Minute(), t.Hour()),
			fmt.Sprintf("%s nach %s", number(t.Minute()), number(t.Hour())),
		)
		if t.Minute() < 5 {
			result = append(
				result,
				fmt.Sprintf("kurz nach %s", number(t.Hour())),
			)
		}
		if t.Hour() > 12 {
			result = append(
				result,
				fmt.Sprintf("%d nach %d", t.Minute(), t.Hour()-12),
				fmt.Sprintf("%s nach %s", number(t.Minute()), number(t.Hour()-12)),
			)
			if t.Minute() < 5 {
				result = append(
					result,
					fmt.Sprintf("kurz nach %s", number(t.Hour()-12)),
				)
			}
		}
	}
	if t.Minute() >= 40 && t.Minute() <= 59 {
		result = append(
			result,
			fmt.Sprintf("%d vor %d", 60-t.Minute(), t.Hour()+1),
			fmt.Sprintf("%s vor %s", number(60-t.Minute()), number(t.Hour()+1)),
		)
		if t.Minute() > 55 {
			result = append(
				result,
				fmt.Sprintf("kurz vor %s", number(t.Hour()+1)),
			)
		}
		if t.Hour() > 12 {
			result = append(
				result,
				fmt.Sprintf("%d vor %d", 60-t.Minute(), t.Hour()+1-12),
				fmt.Sprintf("%s vor %s", number(60-t.Minute()), number(t.Hour()+1-12)),
			)
			if t.Minute() > 55 {
				result = append(
					result,
					fmt.Sprintf("kurz vor %s", number(t.Hour()+1-12)),
				)
			}
		}
	}

	if t.Minute() >= 20 && t.Minute() < 30 {
		result = append(
			result,
			fmt.Sprintf("%d vor halb %d", 30-t.Minute(), t.Hour()),
			fmt.Sprintf("%s vor halb %s", number(30-t.Minute()), number(t.Hour())),
		)
		if t.Hour() > 12 {
			result = append(
				result,
				fmt.Sprintf("%d vor halb %d", 30-t.Minute(), t.Hour()-12),
				fmt.Sprintf("%s vor halb %s", number(30-t.Minute()), number(t.Hour()-12)),
			)
		}
	}

	if t.Minute() > 30 && t.Minute() <= 40 {
		result = append(
			result,
			fmt.Sprintf("%d nach halb %d", t.Minute()-30, t.Hour()),
			fmt.Sprintf("%s nach halb %s", number(t.Minute()-30), number(t.Hour())),
		)
		if t.Hour() > 12 {
			result = append(
				result,
				fmt.Sprintf("%d nach halb %d", t.Minute()-30, t.Hour()-12),
				fmt.Sprintf("%s nach halb %s", number(t.Minute()-30), number(t.Hour()-12)),
			)
		}
	}

	return result
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
