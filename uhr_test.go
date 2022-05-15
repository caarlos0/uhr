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

func TestUhr(t *testing.T) {
	t.Parallel()
	for _, v := range []string{
		"2022-05-14 23:08",
		"2022-05-13 13:00",
		"2022-05-12 13:30",
		"2022-05-11 13:45",
		"2022-05-11 13:15",
	} {
		t.Run(v, func(t *testing.T) {
			t.Parallel()
			requireEqual(t, Uhr(parse(t, v)))
		})
	}
}

func TestNumber(t *testing.T) {
	var numbers []string
	for i := 1; i < 60; i++ {
		numbers = append(numbers, fmt.Sprintf("%d=%s", i, number(i)))
		fmt.Println(number(i))
	}
	requireEqual(t, numbers)
}

const layout = "2006-01-02 15:04"

func parse(tb testing.TB, s string) time.Time {
	tb.Helper()
	t, err := time.Parse(layout, s)
	is.New(tb).NoErr(err)
	return t
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
