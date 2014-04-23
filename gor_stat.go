package main

import (
	"time"
	"log"
	"strconv"
)

const (
	rate = 5
)

type GorStat struct {
	statName	string
	latest		int
	mean	 	int
	max			int
	count		int
}

func NewGorStat(statName string) (s *GorStat) {
	s = new(GorStat)
	s.statName = statName
	s.latest = 0
	s.mean = 0
	s.max = 0
	s.count = 0

	if Settings.stats {
		log.Println(s.statName + ":latest,mean,max,count")
		go s.reportStats()
	}
	return
}

func (s *GorStat) Write(latest int) {
	if Settings.stats {
		if latest > s.max {
			s.max = latest
		}
		if latest != 0 {
			s.mean = (s.mean + latest) / 2
		}
		s.latest = latest
		s.count = s.count + 1
	}
}

func (s *GorStat) Reset() {
	s.latest = 0
	s.max = 0
	s.mean = 0
	s.count = 0
}

func (s *GorStat) String() string {
	return s.statName + ":" + strconv.Itoa(s.latest) + "," + strconv.Itoa(s.mean) + "," + strconv.Itoa(s.max) + "," + strconv.Itoa(s.count)
}

func (s *GorStat) reportStats() {
	for {
			log.Println(s)
			s.Reset()
			time.Sleep(rate * time.Second)
	}
}

