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
type jcLogrusAnc struct {
  filtered byte
}

type JCLogrus struct {
  Level int
  TimeStamp string
  OutputType int
  OutputPath string
  Output interface {}
  logrus_ *logrus.Logger
  ancillary *jcLogrusAnc
}

/* Interface 'JCLogger' */
func (jl *JCLogrus) SetLevel(level int) () {

  /*
    Set logger level
  */

  if (jl == nil) {
    return
  }

  jl.logrus_.SetLevel(levelConversion(level))
}

func (jl *JCLogrus) SetTimeStamp(ts string) () {

  /*
    Set logger time stamp
  */

  var tf logrus.TextFormatter

  if (jl == nil) {
    return
  }

  tf.FullTimestamp = true
  tf.TimestampFormat = ts
  jl.logrus_.SetFormatter(&tf)
}

func (jl *JCLogrus) Close() () {

  /*
    Close the logger
  */

  if (jl == nil || jl.OutputType == JCLOG_CONSOLE) {
    return
  }

  if (jl.OutputType == JCLOG_FILE) {
    ff, tt := jl.Output.(*os.File)
    if (tt) {
      ff.Close()
    }

    jl.logrus_ = nil
  }
}

func (jl *JCLogrus) FilterLevel(level int) () {

  /*
    Log only this level (you may filter multiple levels)
  */

  if (jl == nil || LEVEL_STR[level] == "") {
    return
  }

  if (jl.ancillary == nil) {
    jl.ancillary = &jcLogrusAnc{}
  }

  jl.ancillary.filtered |= 1 << level
}

func (jl *JCLogrus) UnFilterLevel(level int) () {

  /*
    Disable this filter
  */

  if (jl == nil || LEVEL_STR[level] == "") {
    return
  }

  if (jl.ancillary == nil) {
    return
  }

  jl.ancillary.filtered &= ^(1 << level)
}

// Log levels
func (jl *JCLogrus) Trace(s string) () {

  if (jl == nil || jl.logrus_ == nil) {
    return
  }

  if (jl.ancillary != nil) {
    filterMe := jl.ancillary.filtered & (1 << LEVEL_TRACE)
    if (jl.ancillary.filtered != 0 && filterMe == 0) {
        return
    }
  }

  jl.logrus_.Trace(s)
}

func (jl *JCLogrus) Debug(s string) () {

  if (jl == nil || jl.logrus_ == nil) {
    return
  }

  if (jl.ancillary != nil) {
    filterMe := jl.ancillary.filtered & (1 << LEVEL_DEBUG)
    if (jl.ancillary.filtered != 0 && filterMe == 0) {
        return
    }
  }

  jl.logrus_.Debug(s)
}

func (jl *JCLogrus) Info(s string) () {

  if (jl == nil || jl.logrus_ == nil) {
    return
  }

  if (jl.ancillary != nil) {
    filterMe := jl.ancillary.filtered & (1 << LEVEL_INFO)
    if (jl.ancillary.filtered != 0 && filterMe == 0) {
        return
    }
  }

  jl.logrus_.Info(s)
}

func (jl *JCLogrus) Warning(s string) () {

  if (jl == nil || jl.logrus_ == nil) {
    return
  }

  if (jl.ancillary != nil) {
    filterMe := jl.ancillary.filtered & (1 << LEVEL_WARNING)
    if (jl.ancillary.filtered != 0 && filterMe == 0) {
        return
    }
  }

  jl.logrus_.Warning(s)
}

func (jl *JCLogrus) Error(s string) () {

  if (jl == nil || jl.logrus_ == nil) {
    return
  }

  if (jl.ancillary != nil) {
    filterMe := jl.ancillary.filtered & (1 << LEVEL_ERROR)
    if (jl.ancillary.filtered != 0 && filterMe == 0) {
        return
    }
  }

  jl.logrus_.Error(s)
}

func (jl *JCLogrus) Fatal(s string) () {

  if (jl == nil || jl.logrus_ == nil) {
    return
  }

  if (jl.ancillary != nil) {
    filterMe := jl.ancillary.filtered & (1 << LEVEL_FATAL)
    if (jl.ancillary.filtered != 0 && filterMe == 0) {
        return
    }
  }

  jl.logrus_.Fatal(s)
}

/* Interface 'stringer' */
func (jl *JCLogrus) String() (string) {

  if (jl.OutputType == JCLOG_CONSOLE) {
    return fmt.Sprintf("%v(%v) | %v",
      LEVEL_STR[jl.Level], jl.Level, jl.TimeStamp)
  }

  return fmt.Sprintf("%v(%v) | %v\nPath: %v",
    LEVEL_STR[jl.Level], jl.Level, jl.TimeStamp, jl.OutputPath)
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
  jcgrus.ancillary = nil
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

    jcgrus.logrus_.SetOutput(ff)
    jcgrus.Output = ff
  }

  jcgrus.SetTimeStamp(jcgrus.TimeStamp)
  jcgrus.logrus_.SetLevel(levelConversion(jcgrus.Level))
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

func levelConversion(level int) (logrus.Level) {

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
