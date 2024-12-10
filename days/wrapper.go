package days

import "strconv"

func NumericResult(result int, err error) (string, error) {
	if err != nil {
		return "", err
	}

	return strconv.Itoa(result), nil
}
