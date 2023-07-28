package pokecache

import (
	"testing"
	"time"
)

func TestAdd_Get(t *testing.T) {
	c := NewCache(time.Millisecond * 10)
	cases := []struct {
		inp string
		out []byte
	}{{
		inp: "key",
		out: []byte("helloworld"),
	}, {
		inp: "key",
		out: []byte(""),
	},
		{
			inp: "",
			out: []byte("helloworld"),
		},
	}

	for _, cs := range cases {
		c.Add(cs.inp, cs.out)

		res, ok := c.Get(cs.inp)
		if !ok {
			t.Errorf("Key not found")
			continue
		}

		if string(res) != string(cs.out) {
			t.Errorf("value doesnt match")
			continue
		}
	}
}

func TestReapa(t *testing.T) {
	interval := 10 * time.Millisecond
	c := NewCache(interval)

	key := "key1"
	c.Add(key, []byte("hello"))
	time.Sleep(interval + time.Millisecond)

	_, ok := c.Get(key)
	if ok {
		t.Errorf("The record isnt removed")

	}
}
