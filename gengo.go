package jp_gengo

import (
	"time.Time"
)

type Gengo struct {
	name string
	startDay time.Time
	endDay   *time.Time
}

func NewGengoFromString(str, layout string) (*Gengo, error) {
	t, err := time.Parse(layout, str)
	if err != nil {
		return nil, err
	}

	g := new(Gengo)
	return g, nil
}