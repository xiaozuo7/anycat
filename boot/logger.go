package boot

import (
	"anycat/global/variable"
	"anycat/util/zaputil"
)

func InitLogger() {
	variable.ZapLog = zaputil.CreateZapUtil()
}
