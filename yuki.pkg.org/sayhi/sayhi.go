package sayhi

import "yuki.pkg.org/tools/logger"

func init() {
	logger.InfoLog("sayhi go pkg initial successed!", "in sayhi/sayhi.go")
}

func SayHi() string {
	logger.InfoLog("Sayhi() done!", "in sayhi/sayhi.go")
	return "Hi!"
}
