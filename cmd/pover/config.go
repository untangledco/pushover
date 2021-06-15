package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

type config struct {
	user string
	token string
}

func configFromFile(name string) (config, error) {
	f, err := os.Open(name)
	if err != nil {
		f.Close()
		return config{}, err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	i := 0
	c := config{}
	for sc.Scan() {
		i++
		s := sc.Text()
		// skip comments
		if strings.HasPrefix(s, "#") {
			continue
		}
		slice := strings.Split(s, " ")
		if len(slice) > 2 {
			return config{}, fmt.Errorf("%s:%d: too many values", f.Name(), i)
		}
		switch slice[0] {
		case "user":
			c.user = slice[1]
		case "token":
			c.token = slice[1]
		default:
			return config{}, fmt.Errorf("%s:%d: unknown key %s", f.Name(), i, slice[0])
		}
	}
	if c.user == "" {
		return config{}, fmt.Errorf("no user")
	} else if c.token == "" {
		return config{}, fmt.Errorf("no token")
	}
	return c, nil
}
