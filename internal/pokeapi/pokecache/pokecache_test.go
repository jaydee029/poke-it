package pokecache

import "testing"

func TestAdd_Get(t *testing.T) {
	c := NewCache()
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
