package script

import "github.com/bitfield/script"

func exec() {
	script.Exec(`find . -name "*.go"`)
}
