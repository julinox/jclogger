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
  FilterLevels([]int) ()
}

/* Functions */
