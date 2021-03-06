package logg

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

type Message struct {
	text string
	datetime time.Time
	dateFormated string
	path string
	levelLabel string
	stack string
	app string
}

//NewMessage create the message that will send to writers
func NewMessage(level int, text, stack, app string) *Message {
	m := &Message{text: text, stack: stack, app: app}
	m.setLevelLabel(level)
	m.setPath()
	m.setDatetime(time.Now())

	return m
}

func (s *Message) setDatetime(datetime time.Time) {
	s.datetime = datetime
	s.dateFormated = datetime.Format("2006/01/02 15:04:05")
}

func (s *Message) setPath() {
	_, fn, line, _ := runtime.Caller(3)
	file := strings.Split(fn, "/")

	s.path = fmt.Sprintf("%v:%v", file[len(file)-1], line)
}

func (s *Message) setLevelLabel(level int) {
	switch level {
	case PanicType:
		s.levelLabel = "PANIC"
	case InfoType:
		s.levelLabel = "INFO"
	case ErrorType:
		s.levelLabel = "ERROR"
	case WarningType:
		s.levelLabel = "WARNING"
	}
}

//ToJSON export the message to JSON
func (s *Message) ToJSON() interface{} {
	return struct {
		Message string `json:"message,omitempty"`
		Path string `json:"path,omitempty"`
		Type string `json:"type,omitempty"`
		Stack string `json:"stack,omitempty"`
		App string `json:"app,omitempty"`
		Date time.Time `json:"date,omitempty"`
	}{
		s.text,
		s.path,
		s.levelLabel,
		s.stack,
		s.app,
		s.datetime,
	}
}

//String export the message to string
func (s *Message) String() string {
	return fmt.Sprintf("%v %s %s - %s", s.dateFormated, s.path, s.levelLabel, s.text)
}