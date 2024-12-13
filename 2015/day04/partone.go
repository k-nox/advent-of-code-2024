package day04

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartOne(useSample bool) int {
	key := strings.TrimSpace(string(helper.ReadInput(2015, 4, useSample)))
	i := 0
	for {
		test := fmt.Sprintf("%s%d", key, i)
		h := md5.New()
		io.WriteString(h, test)
		if strings.HasPrefix(fmt.Sprintf("%x", h.Sum(nil)), "00000") {
			return i
		}
		i++
	}
}
