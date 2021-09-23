package commit

import "time"

type Commit struct {
	Hash    string
	Author  string
	Date    time.Time
	Message string
}
