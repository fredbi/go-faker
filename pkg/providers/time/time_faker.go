package time

import (
	"time"

	"github.com/fredbi/go-faker/pkg/generic"
)

type Faker struct {
	generic.Faker
}

func (f *Faker) Century() int            { return 0 }
func (f *Faker) CenturyString() string   { return "" }
func (f *Faker) Date() string            { return "" }
func (f *Faker) DayOfMonth() string      { return "" }
func (f *Faker) DayOfWeek() string       { return "" }
func (f *Faker) MonthName() string       { return "" }
func (f *Faker) Year() int               { return 0 }
func (f *Faker) YearString() string      { return "" }
func (f *Faker) Time() time.Time         { return time.Time{} }
func (f *Faker) TimeString() string      { return f.Time().String() }
func (f *Faker) Duration() time.Duration { return time.Duration(0) }
func (f *Faker) TimePeriod() string      { return "" }
func (f *Faker) Timestamp() string       { return "" }
func (f *Faker) Timezone() string        { return "" }
func (f *Faker) UNIXTime() string        { return "" }
