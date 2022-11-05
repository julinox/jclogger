/* Package */
package jclogger

/* Imports */

/* Glocals */
const (
  JCLOG_ENV_LEVEL  = "LOG_LEVEL"
  JCLOG_ENV_FORMAT = "LOG_FORMAT"
)

const (
  JCLOG_FILE    = 0
  JCLOG_CONSOLE = 1
)

const (
  JCLOG_FLAG_LOCATION = 1 << 0
)

const (
  LEVEL_TRACE   = 0
  LEVEL_DEBUG   = 1
  LEVEL_INFO    = 2
  LEVEL_WARNING = 3
  LEVEL_ERROR   = 4
  LEVEL_FATAL   = 5
)

var LEVEL_STR = map[int]string {
  LEVEL_TRACE   : "TRACE",
  LEVEL_DEBUG   : "DEBUG",
  LEVEL_INFO    : "INFO",
  LEVEL_WARNING : "WARNING",
  LEVEL_ERROR   : "ERROR",
  LEVEL_FATAL   : "FATAL",
}

/* Types */

/* Interface 'JCLogger' */
type JCLogger interface {
  Close() ()
  SetLevel(int) ()
  SetTimeStamp(string) ()
  FilterLevel(int) ()
  UnFilterLevel(int) ()
  FuncName(int) (string)
  Trace(...interface{}) ()
  TraceSp(string, ...interface{}) ()
  Debug(...interface{}) ()
  DebugSp(string, ...interface{}) ()
  Info(...interface{}) ()
  InfoSp(string, ...interface{}) ()
  Warning(...interface{}) ()
  WarningSp(string, ...interface{}) ()
  Error(...interface{}) ()
  ErrorSp(string, ...interface{}) ()
  Fatal(...interface{}) ()
  FatalSp(string, ...interface{}) ()
}

/* Functions */
