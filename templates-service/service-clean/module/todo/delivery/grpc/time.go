package grpc

import "time"

func convertStrToTime(input string) (*time.Time, error) {
	time, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return nil, err
	}
	return &time, nil
}
