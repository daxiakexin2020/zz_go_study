package main

import (
   "log_manage/util"
)

func main()  {
   myLogger :=util.NewLog("info")
   for  {
      myLogger.Debug("这是debug日志")
      myLogger.Trace("这是trace日志")
      myLogger.Info("这是info日志")
   }
}
