package main

import (
	"github.com/AlexanderChiuluvB/bfs/pitchfork/conf"
	myzk "github.com/AlexanderChiuluvB/bfs/store/zk"
	"fmt"
	"testing"
)

func TestPitchfork(t *testing.T) {

	var (
		config        *conf.Config
		zk            *myzk.Zookeeper
		p             *Pitchfork
		storelist     StoreList
		store         *conf.Store
		pitchforklist PitchforkList
		err           error
	)

	if config, err = conf.NewConfig(configFile); err != nil {
		t.Errorf("NewConfig(\"%s\") error(%v)", configFile, err)
		return
	}

	if zk, err = myzk.NewZookeeper(); err != nil {
		t.Errorf("NewZookeeper() error(%v)", err)
		t.FailNow()
	}

	p, err = NewPitchfork(config)
	if err = p.Register(); err != nil {
		t.Errorf("pitchfork Register() failed, Quit now")
		t.FailNow()
	}

	storelist, _, err = p.WatchGetStores()
	if err != nil {
		t.Errorf("pitchfork WatchGetStores() failed, Quit now")
		t.FailNow()
	}
	for _, store = range storelist {
		fmt.Println(store.rack, store.ID, store.host, store.status)
	}

	pitchforklist, _, err = p.WatchGetPitchforks()
	if err != nil {
		t.Errorf("pitchfork WatchGetPitchforks() failed, Quit now")
		t.FailNow()
	}
	for _, p = range pitchforklist {
		fmt.Println(p.ID)
	}

	for _, store = range storelist {
		if err = p.getStore(store); err != nil {
			t.Errorf("probeStore() called error(%v)", err)
			t.FailNow()
		}
	}

}
