package jp_gengo

import (
	"errors"
	"fmt"
	"time"
)

var asiaTokyo *time.Location

var supportedOldestDay = time.Date(1868, 1, 1, 0, 0, 0, 0, Location())

var ErrUnsupportedDate = errors.New("unsupported date passed")

var gengolist = []struct {
	symbol               Symbol
	startDay             *time.Time
	endDay               *time.Time
	duplicatePreviousEra bool
}{
	{symbol: SymbolReiwa, startDay: date(2019, 5, 1), endDay: nil},
	{symbol: SymbolHeisei, startDay: date(1989, 1, 8), endDay: date(2019, 4, 30)},
	{symbol: SymbolSyowa, startDay: date(1926, 7, 30), endDay: date(1989, 1, 7), duplicatePreviousEra: true},
	{symbol: SymbolTaisho, startDay: date(1912, 9, 8), endDay: date(1926, 7, 30), duplicatePreviousEra: true},
	{symbol: SymbolMeiji, startDay: date(1868, 1, 1), endDay: date(1912, 9, 8)},
}

type Symbol string

const (
	SymbolReiwa  Symbol = "R"
	SymbolHeisei Symbol = "H"
	SymbolSyowa  Symbol = "S"
	SymbolTaisho Symbol = "T"
	SymbolMeiji  Symbol = "M"
)

func (s Symbol) String() string {
	return string(s)
}

func (s Symbol) Ja() string {
	switch s {
	case SymbolReiwa:
		return "令和"
	case SymbolHeisei:
		return "平成"
	case SymbolSyowa:
		return "昭和"
	case SymbolTaisho:
		return "大正"
	case SymbolMeiji:
		return "明治"
	}
	panic("unreachable")
}

func Location() *time.Location {
	if asiaTokyo == nil {
		if loc, err := time.LoadLocation("Asia/Tokyo"); err != nil {
			panic(err)
		} else {
			asiaTokyo = loc
		}
	}
	return asiaTokyo
}

func date(y, m, d int) *time.Time {
	t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, Location())
	return &t
}

func ymdEqual(a, b time.Time) bool {
	return a.Year() == b.Year() && a.Month() == b.Month() && a.Day() == b.Day()
}

func NewGengo(t time.Time) (*Gengo, error) {
	if t.Before(supportedOldestDay) {
		return nil, ErrUnsupportedDate
	}

	for _, gengo := range gengolist {
		if !t.After(*gengo.startDay) {
			continue
		}

		if gengo.endDay == nil || t.Before(*gengo.endDay) {
			g := new(Gengo)
			if gengo.duplicatePreviousEra && ymdEqual(*gengo.startDay, t) {
			//	return &Gengo{
			//		symbol:   gengo.symbol,
			//		startDay: *gengo.startDay,
			//		endDay:   gengo.endDay,
			//	}, nil
			}
			g.symbol = gengo.symbol
			g.startDay = *gengo.startDay
			g.endDay = gengo.endDay
			g.year = g.calculateYear(t)
			return g, nil
		}
	}
	return nil, nil
}

type Gengo struct {
	symbol   Symbol
	year     uint16
	startDay time.Time
	endDay   *time.Time
}

func (g *Gengo) calculateYear(t time.Time) uint16 {
	sub := g.startDay.Sub(t)
	if sub == 0 {
		return 1
	}
	fmt.Println("=====>", uint16(sub.Hours() / (24 * 365)))
	return uint16(sub.Hours() / (24 * 365))
}

func (g *Gengo) Ended() bool {
	return g.endDay != nil
}