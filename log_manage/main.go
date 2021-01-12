package main

import (
   "log_manage/util"
)

func main()  {
   myLogger :=util.NewLog()
   myLogger.Debug("这是debug日志")
}
