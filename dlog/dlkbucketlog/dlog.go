package dlkbucketlog

import (
	"github.com/libp2p/go-libp2p-kbucket/util"
	"go.uber.org/zap"
)

var L *zap.Logger

func init() {
	L = util.GetXDebugLog("lkbucket")
}
