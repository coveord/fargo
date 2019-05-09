package fargo

import (
	"reflect"
	"testing"
)

func TestRoundRobin(t *testing.T) {
	tests := []struct {
		urls []string
		want []string
	}{
		{
			urls: []string{
				"192.168.33.10",
				"192.168.33.11",
				"192.168.33.12",
			},
			want: []string{
				"192.168.33.10",
				"192.168.33.11",
				"192.168.33.12",
				"192.168.33.10",
			},
		},
	}

	for i, test := range tests {
		rr := newRoundRobin(test.urls)

		gots := make([]string, 0, len(test.want))
		for j := 0; j < len(test.want); j++ {
			gots = append(gots, rr.Next())
		}

		if got, want := gots, test.want; !reflect.DeepEqual(got, want) {
			t.Errorf("tests[%d] - RoundRobin is wrong. want: %v, got: %v", i, want, got)
		}
	}
}
