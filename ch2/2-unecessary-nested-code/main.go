package main

import (
	"errors"
)

func main() {

}

func listing1_join(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

// Remove unecessary nested code
func listing2_join(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 empty")
	}
	if s2 == "" {
		return "", errors.New("s2 empty")
	}
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}
	if len(concat) > max {
		return concat[:max], nil
	}
	return concat, nil

}

func concatenate(s1, s2 string) (string, error) {
	// function to join some strings, error check etc.
	return "", nil
}
