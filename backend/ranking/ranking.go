package ranking

import (
	"fmt"
	"strconv"

	"github.com/Synaxis/nfsc-fesl/backend/model"
)

const (
	PointTypeFloat = iota
)

type Stats map[string]string

type Adder func(p *model.HeroStats, value string, pointType int) error

type Getter func(p *model.HeroStats) (string, error)

type Setter func(p *model.HeroStats, value string, pointType int) error

func GetStats(p *model.HeroStats, keys ...string) (Stats, error) {
	s := Stats{}

	for _, key := range keys {
		val, err := getStatsValue(p, key)
		if err != nil {
			return nil, err
		}
		s[key] = val
	}

	return s, nil
}

func getStatsValue(p *model.HeroStats, key string) (string, error) {
	if getter, ok := getters[key]; ok {
		val, err := getter(p)
		if err != nil {
			return "", fmt.Errorf("ranking: cannot fetch value for key='%s', %v", key, err)
		}

		return val, nil
	}

	return "", fmt.Errorf("ranking: cannot get value for key='%s'", key)
}

const (
	UpdateTypeReplace = iota
	UpdateTypeMax
	_ // UpdateTypeDecrement?
	UpdateTypeAdd
)

func UpdateStatValue(p *model.HeroStats, key, value string, updateType, pointType string) error {
	ut, err := strconv.Atoi(updateType)
	if err != nil {
		return fmt.Errorf(
			"ranking: unknown updateType for key='%s' (%s)",
			key,
			updateType,
		)
	}

	pt, err := strconv.Atoi(pointType)
	if err != nil {
		return err
	}

	switch ut {
	case UpdateTypeReplace:
		return setStatsValue(p, key, value, pt)
	case UpdateTypeMax:
		// return maxStatsValue(p, key, value, pt)
		return nil
	case UpdateTypeAdd:
		return addStatsValue(p, key, value, pt)
	default:
		return fmt.Errorf(
			"ranking: cannot update stats value key='%s' (method=%s, value=%s)",
			key,
			updateType,
			value,
		)
	}
}

func addStatsValue(p *model.HeroStats, key, val string, pt int) error {
	if adder, ok := adders[key]; ok {
		if err := adder(p, val, pt); err != nil {
			return fmt.Errorf("ranking: cannot add value for key='%s'", key)
		}
	}

	return nil
}

func setStatsValue(p *model.HeroStats, key, val string, pt int) error {
	if setter, ok := setters[key]; ok {
		if err := setter(p, val, pt); err != nil {
			return fmt.Errorf("ranking: cannot set value for key='%s'", key)
		}
	}

	return nil
}
