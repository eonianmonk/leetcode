package main

import "testing"

func Test_disemvoveling(t *testing.T) {
	t.Run("remove vovels from known strings", func(t *testing.T) {
		entry := "This website is for losers LOL!"
		wanted := "Ths wbst s fr lsrs LL!"

		reply := Disemvowel(entry)
		if reply != wanted {
			t.Fatalf("failed to process string: got \"%s\", wanted \"%s\"", wanted, reply)
		}
	})
}
