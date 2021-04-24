package poa_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/lifefile/LFT/poa"
	"github.com/lifefile/LFT/thor"
)

var (
	p1 = thor.BytesToAddress([]byte("p1"))
	p2 = thor.BytesToAddress([]byte("p2"))
	p3 = thor.BytesToAddress([]byte("p3"))
	p4 = thor.BytesToAddress([]byte("p4"))
	p5 = thor.BytesToAddress([]byte("p5"))

	proposers = []poa.Proposer{
		{p1, false},
		{p2, true},
		{p3, false},
		{p4, false},
		{p5, false},
	}

	parentTime = uint64(1001)
)

func TestSchedule(t *testing.T) {

	_, err := poa.NewScheduler(thor.BytesToAddress([]byte("px")), proposers, 1, parentTime)
	assert.NotNil(t, err)

	sched, _ := poa.NewScheduler(p1, proposers, 1, parentTime)

	for i := uint64(0); i < 100; i++ {
		now := parentTime + i*thor.BlockInterval/2
		nbt := sched.Schedule(now)
		assert.True(t, nbt >= now)
		assert.True(t, sched.IsTheTime(nbt))
	}
}

func TestIsTheTime(t *testing.T) {
	sched, _ := poa.NewScheduler(p2, proposers, 1, parentTime)

	tests := []struct {
		now  uint64
		want bool
	}{
		{parentTime - 1, false},
		{parentTime + thor.BlockInterval/2, false},
		{parentTime + thor.BlockInterval, true},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, sched.IsTheTime(tt.now))
	}
}

func TestUpdates(t *testing.T) {

	sched, _ := poa.NewScheduler(p1, proposers, 1, parentTime)

	tests := []struct {
		newBlockTime uint64
		want         uint64
	}{
		{parentTime + thor.BlockInterval, 2},
		{parentTime + thor.BlockInterval*30, 1},
	}

	for _, tt := range tests {
		_, score := sched.Updates(tt.newBlockTime)
		assert.Equal(t, tt.want, score)
	}
}