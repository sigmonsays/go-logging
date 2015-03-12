
leveled logging libary in go

documentation - http://godoc.org/github.com/sigmonsays/go-logging


quickstart
-----------------------

in log.go add the following 

      package whatever
      import (
          gologging "git.llnw.com/lama/go-logging.git"
      )

      var log gologging.Logger
      func init() {
         log = gologging.Register("whatever", func(newlog gologging.Logger) { log = newlog })
      }


to change the log level of all loggers

      gologging.SetLogLevel("ERROR")
