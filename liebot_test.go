package liebot

import "testing"

func TestLiebot(t *testing.T) {
	s := Search("onion offense")
	if s != "http://achewood.com/comic.php?date=02102005" {
		t.Errorf("Couldn't find onion offense, found %s", s)
	}
}

func TestTitle(t *testing.T) {
	s := Search("Philippe and the flower")
	if s != "http://achewood.com/comic.php?date=10072002" {
		t.Errorf("Couldn't find comic for title 'Philippe and the flower', found %s", s)
	}
}
