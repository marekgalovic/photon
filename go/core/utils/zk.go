package utils

import (
    "fmt";
    "regexp";
    "strconv";
)

var (
    seqRegexp = regexp.MustCompile(`.*-(\d+)`)
)

func ZNodeSeq(name string) (uint64, error) {
    matches := seqRegexp.FindStringSubmatch(name)
    if len(matches) != 2 {
        return 0, fmt.Errorf("Invalid znode name: '%s'", name)
    }

    return strconv.ParseUint(matches[1], 10, 64)
}
