/*
  'JCLogrus' implements 'JCLogger' interface.
  This implementation uses 'logrus'
*/

/* Package */
package jclogger

/* Imports */
import (
  "os"
  "fmt"
  "errors"
  "strings"
  "github.com/sirupsen/logrus"
)

/* Glocals */
const (
  JCLOGRUS_DEFAULT_LEVEL = LEVEL_INFO
  JCLOGRUS_DEFAULT_FORMAT = "2006/01/02 15:04:05"
)

/* Types */
// Ancillary struct to handle 'hooks' related info
type logAncillary struct {
  filterLevels map[int]int
}

type JCLogrus struct {
  Level int
  TimeStamp string
  OutputType int
  OutputPath string
  Output interface {}
  logrus_ *logrus.Logger
  ancillary *logAncillary
}

/* Interface 'JCLogger' */
func (bl *JCLogrus) SetLevel(level int) () {

  /*
    Set logger level
  */

  if (bl == nil) {
    return
  }

  bl.logrus_.SetLevel(jcloggerLevelToLogrusLevel(level))
}

func (bl *JCLogrus) SetTimeStamp(ts string) () {

  /*
    Set logger time stamp
  */

  var tf logrus.TextFormatter

  if (bl == nil) {
    return
  }

  tf.FullTimestamp = true
  tf.TimestampFormat = ts
  bl.logrus_.SetFormatter(&tf)
}

func (bl *JCLogrus) Close() () {

  /*
    Close the logger
  */

  if (bl == nil || bl.OutputType == JCLOG_CONSOLE) {
    return
  }

  if (bl.OutputType == JCLOG_FILE) {
    ff, tt := bl.Output.(*os.File)
    if (tt) {
      ff.Close()
    }

    bl.logrus_ = nil
  }
}

func (bl *JCLogrus) FilterLevels(levels []int) () {
  fmt.Println("filter levels")
}

// Log levels
func (bl *JCLogrus) Trace(s string) () {

  if (bl == nil || bl.logrus_ == nil) {
    return
  }

  bl.logrus_.Trace(s)
}

func (bl *JCLogrus) Debug(s string) () {

  if (bl == nil || bl.logrus_ == nil) {
    return
  }

  bl.logrus_.Debug(s)
}

func (bl *JCLogrus) Info(s string) () {

  if (bl == nil || bl.logrus_ == nil) {
    return
  }

  bl.logrus_.Info(s)
}

func (bl *JCLogrus) Warning(s string) () {

  if (bl == nil || bl.logrus_ == nil) {
    return
  }

  bl.logrus_.Warning(s)
}

func (bl *JCLogrus) Error(s string) () {

  if (bl == nil || bl.logrus_ == nil) {
    return
  }

  bl.logrus_.Error(s)
}

func (bl *JCLogrus) Fatal(s string) () {

  if (bl == nil || bl.logrus_ == nil) {
    return
  }

  bl.logrus_.Fatal(s)
}

/* Interface 'stringer' */
func (bl *JCLogrus) String() (string) {

  if (bl.OutputType == JCLOG_CONSOLE) {
    return fmt.Sprintf("%v(%v) | %v",
      LEVEL_STR[bl.Level], bl.Level, bl.TimeStamp)
  }

  return fmt.Sprintf("%v(%v) | %v\nPath: %v",
    LEVEL_STR[bl.Level],bl.Level, bl.TimeStamp, bl.OutputPath)
}

/* Functions */
func CreateLogger(outputType int, outputPath string) (*JCLogrus, error) {

  /*
    Creates a new logger. Level and timestamp will be set to
    defaults (JCLOGRUS_DEFAULT_FORMAT, JCLOGRUS_DEFAULT_LEVEL) unless
    environment variables are set
  */

  var jcgrus JCLogrus

  // Setting defaults
  jcgrus.Level =  readLevel()
  jcgrus.TimeStamp = readTimeStamp()
  jcgrus.OutputType = outputType
  jcgrus.logrus_ = logrus.New()
  if (outputType == JCLOG_CONSOLE) {
    jcgrus.OutputPath = ""
    jcgrus.logrus_.SetOutput(os.Stdout)
    jcgrus.Output = nil

  } else {
    if (outputPath == "") {
      return nil, errors.New("'outputPath' cannot be empty [outputType = LOG_FILE]")
    }

    jcgrus.OutputPath = outputPath
    auxMode := os.O_APPEND | os.O_CREATE | os.O_RDWR
    ff, err := os.OpenFile(jcgrus.OutputPath, auxMode, 0666)
    if (err != nil) {
      return nil, err
    }

    jcgrus.logrus_.SetLevel(jcloggerLevelToLogrusLevel(jcgrus.Level))
    jcgrus.logrus_.SetOutput(ff)
    jcgrus.Output = ff
    jcgrus.SetTimeStamp(jcgrus.TimeStamp)
    jcgrus.ancillary = nil
  }

  return &jcgrus, nil
}

/* Internals */
func readLevel() (int) {

  /*
    Level can be set via enviroment variable.
    Check if proper enviroment variable was set
  */

  var retLevel int

  retLevel = JCLOGRUS_DEFAULT_LEVEL
  level := os.Getenv(JCLOG_ENV_LEVEL)
  if (level == "") {
    return JCLOGRUS_DEFAULT_LEVEL
  }

  switch (strings.ToUpper(level)) {
    case LEVEL_STR[LEVEL_TRACE]:
      retLevel = LEVEL_TRACE

    case LEVEL_STR[LEVEL_DEBUG]:
      retLevel = LEVEL_DEBUG

    case LEVEL_STR[LEVEL_INFO]:
      retLevel = LEVEL_INFO

    case LEVEL_STR[LEVEL_WARNING]:
      retLevel = LEVEL_WARNING

    case LEVEL_STR[LEVEL_ERROR]:
      retLevel = LEVEL_ERROR

    case LEVEL_STR[LEVEL_FATAL]:
      retLevel = LEVEL_FATAL
  }

  return retLevel
}

func readTimeStamp() (string) {

  /*
    Format can be set via enviroment variable.
    Check if proper enviroment variable was set
  */

  format := os.Getenv(JCLOG_ENV_FORMAT)
  // No format validation
  if (format != "") {
    return format
  }

  return JCLOGRUS_DEFAULT_FORMAT
}

func jcloggerLevelToLogrusLevel(level int) (logrus.Level) {

  /*
    Conversion 'JCLogger' level to 'logrus' level
  */

  switch(level) {
    case LEVEL_TRACE:
      return logrus.TraceLevel

    case LEVEL_DEBUG:
      return logrus.DebugLevel

    case LEVEL_INFO:
      return logrus.InfoLevel

    case LEVEL_WARNING:
      return logrus.WarnLevel

    case LEVEL_ERROR:
      return logrus.ErrorLevel

    case LEVEL_FATAL:
      return logrus.FatalLevel

    default:
      return logrus.InfoLevel
  }
}
