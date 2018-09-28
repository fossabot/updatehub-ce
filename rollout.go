package main

import (
	"time"

	"github.com/asdine/storm"
)

type Rollout struct {
	ID         int       `storm:"id,increment" json:"id"`
	Package    string    `storm:"index" json:"package"`
	Devices    []string  `json:"devices"`
	Running    bool      `json:"running"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
}

func (r *Rollout) IsFinished(db *storm.DB) (bool, error) {
	for _, uid := range r.Devices {
		var d Device
		if err := db.One("UID", uid, &d); err != nil {
			return false, err
		}

		if d.Status != "finished" {
			return false, nil
		}
	}

	return true, nil
}
