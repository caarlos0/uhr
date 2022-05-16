package uhr

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/matryer/is"
)

const layout = "15:04"

func TestUhr(t *testing.T) {
	t.Parallel()
	for h := 0; h <= 23; h++ {
		for m := 0; m <= 59; m++ {
			s := fmt.Sprintf("%02d:%02d", h, m)
			t.Run(s, func(t *testing.T) {
				t.Parallel()
				is := is.New(t)
				now, err := time.Parse(layout, s)
				is.NoErr(err)
				requireEqual(t, Uhr(now))
			})
		}
	}
}

func TestPartOfDay(t *testing.T) {
	for i := 0; i <= 23; i++ {
		i := i
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			now := mockHourer{i}
			requireEqual(t, []string{PartOfDay(now)})
		})
	}
}

type mockHourer struct {
	h int
}

func (h mockHourer) Hour() int {
	return h.h
}

func TestWeekday(t *testing.T) {
	for _, w := range []time.Weekday{
		time.Sunday,
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
	} {
		w := w
		t.Run(w.String(), func(t *testing.T) {
			t.Parallel()
			now := mockWeekdayer{w}
			requireEqual(t, []string{Weekday(now)})
		})
	}
}

type mockWeekdayer struct {
	w time.Weekday
}

func (w mockWeekdayer) Weekday() time.Weekday { return w.w }

func TestNumber(t *testing.T) {
	var numbers []string
	for i := 1; i < 60; i++ {
		numbers = append(numbers, fmt.Sprintf("%d=%s", i, number(i)))
		fmt.Println(number(i))
	}
	requireEqual(t, numbers)
}

var update = flag.Bool("update", false, "update .golden files")

func requireEqual(tb testing.TB, out []string) {
	tb.Helper()
	is := is.New(tb)

	golden := "testdata/" + tb.Name() + ".golden"
	if *update {
		is.NoErr(os.MkdirAll(filepath.Dir(golden), 0o755))
		is.NoErr(os.WriteFile(golden, []byte(strings.Join(out, "\n")), 0o655))
	}

	gbts, err := os.ReadFile(golden)
	is.NoErr(err)

	is.Equal(string(gbts), strings.Join(out, "\n"))
}
