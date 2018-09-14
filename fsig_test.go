package main

import (
	"os"
	"strconv"
	"testing"
	"time"
)

func TestFsig(t *testing.T) {
	tests := []struct {
		name   string
		signal string
	}{
		{
			"numeric signal",
			"1",
		},
		{
			"long string signal",
			"SIGHUP",
		},
		{
			"short string signal",
			"HUP",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			os.Args = []string{"fsig", "-w", "test", test.signal, "--", "sh", "test.sh"}

			go func() {
				time.Sleep(2 * time.Second)

				_, err := os.Create("test/" + strconv.FormatInt(time.Now().Unix(), 10))
				if err != nil {
					t.Error(err)
				}
			}()

			main()
		})
	}
}
