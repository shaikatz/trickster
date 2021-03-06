/**
* Copyright 2018 Comcast Cable Communications Management, LLC
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
* http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// Package timeconv provides time conversion capabilities to Trickster
package timeconv

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Comcast/trickster/internal/proxy/errors"
)

// ParseDuration returns a duration from a string. Slightly improved over the builtin, since it supports units larger than hour.
func ParseDuration(input string) (time.Duration, error) {
	for i := range input {
		if input[i] > 47 && input[i] < 58 {
			continue
		}
		if input[i] == 46 {
			break
		}
		if i > 0 {
			units, ok := UnitMap[input[i:]]
			if !ok {
				return errors.ParseDuration(input)
			}
			v, err := strconv.ParseInt(input[0:i], 10, 64)
			if err != nil {
				return errors.ParseDuration(input)
			}
			v = v * units
			return time.Duration(v), nil
		}
	}
	return errors.ParseDuration(input)
}

// ParseDurationParts returns a time.Duration from a value and unit
func ParseDurationParts(value int64, units string) (time.Duration, error) {
	if _, ok := UnitMap[units]; !ok {
		return errors.ParseDuration(fmt.Sprintf("%d%s", value, units))
	}
	return time.Duration(value * UnitMap[units]), nil
}

// UnitMap provides a map of common time unit indicators to nanoseconds of duration per unit
var UnitMap = map[string]int64{
	"ns": int64(time.Nanosecond),
	"us": int64(time.Microsecond),
	"µs": int64(time.Microsecond), // U+00B5 = micro symbol
	"μs": int64(time.Microsecond), // U+03BC = Greek letter mu
	"ms": int64(time.Millisecond),
	"s":  int64(time.Second),
	"m":  int64(time.Minute),
	"h":  int64(time.Hour),
	"d":  int64(24 * time.Hour),
	"w":  int64(24 * 7 * time.Hour),
	"y":  int64(24 * 365 * time.Hour),
}
