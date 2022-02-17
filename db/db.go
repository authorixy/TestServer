package db

import (
	"fmt"
	"sort"
	"sync"
)

type Score struct {
	Success bool
	LeftChance int8
	UsedTime int64
}

func getSuccessInfo(b bool) string {
	if b {
		return "猜出了"
	} else {
		return "没猜出"
	}
}

func (s Score) ToString(name string) string {
	return fmt.Sprintf("%s%s, 还剩%d次机会,耗时%d", name, getSuccessInfo(s.Success), s.LeftChance, s.UsedTime)
}

type DB struct {
	Scores map[string]Score
	m      sync.RWMutex
}

var db *DB
var once sync.Once

func GetDB() *DB {
	if db == nil {
		once.Do(func() {
			sm := make(map[string]Score, 30)
			m := sync.RWMutex{}
			db = &DB{sm, m}
		})
	}
	return db
}

func (d *DB) lock() {
	d.m.Lock()
}
func (d *DB) unlock() {
	d.m.Unlock()
}

type UScore struct {
	Name string
	Score
}

func (d *DB) GetRankList() string {
	d.m.RLock()
	defer d.m.Unlock()
	scores := make([]UScore, 0, len(d.Scores))
	for k, v := range d.Scores {
		scores = append(scores, UScore{
			Name: k,
			Score:v,
		})
	}
	sort.Slice(scores, func(i, j int) bool {
		if scores[i].Success && scores[j].Success {
			if scores[i].LeftChance == scores[j].LeftChance {
				return scores[i].UsedTime <= scores[j].UsedTime
			} else {
				return scores[i].LeftChance < scores[j].LeftChance
			}
		} else {
			if scores[i].Success {
				return false
			} else {
				return true
			}
		}
	})
	var list string
	for _, us := range scores {
		list += us.Score.ToString(us.Name) + "\n"
	}
	return list
}

func (d *DB) UPDATE(name string, s Score) error {
	d.lock()
	defer d.unlock()
	if _, ok := d.Scores[name]; !ok {
		d.Scores[name] = s
	} else {
		return fmt.Errorf("already upload")
	}
	return nil
}