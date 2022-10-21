package realTime

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func GetTime() string {
	res, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	return res.String()
}
