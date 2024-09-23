package impl

import (
	"fmt"
	. "github.com/go-tech-lab/framework/logger/src/logger"
	"strings"
)

func NewTroughLogTagsWrapper(logFieldsGetter LogFieldsGetter) *troughLogTagsDecorator {
	return &troughLogTagsDecorator{
		logFieldsGetter: logFieldsGetter,
	}
}

type troughLogTagsDecorator struct {
	logFieldsGetter LogFieldsGetter
}

func (impl *troughLogTagsDecorator) LogfWrapper(format string, params ...interface{}) (string, []interface{}) {
	throughLogFields := impl.logFieldsGetter()
	if len(throughLogFields) <= 0 {
		return format, params
	}
	newFormatBuilder := &strings.Builder{}
	newAddParams := make([]interface{}, 0, len(throughLogFields))
	addFormatBuilder := &strings.Builder{}
	addFormatBuilder.WriteString("\t{ ")
	for k, v := range throughLogFields {
		if len(newAddParams) > 0 {
			addFormatBuilder.WriteString(",")
			addFormatBuilder.WriteString("\"" + k + "\"")
			addFormatBuilder.WriteString(":\"%v\"")
		} else {
			addFormatBuilder.WriteString("\"" + k + "\"")
			addFormatBuilder.WriteString(":\"%v\"")
		}
		newAddParams = append(newAddParams, v)
	}
	addFormatBuilder.WriteString(" }")
	newFormatBuilder.WriteString(format)
	newFormatBuilder.WriteString(addFormatBuilder.String())
	newParams := make([]interface{}, 0, len(params)+len(newAddParams))
	newParams = append(newParams, params...)
	newParams = append(newParams, newAddParams...)
	return newFormatBuilder.String(), newParams
}

func (impl *troughLogTagsDecorator) LogWrapper(v ...interface{}) []interface{} {
	throughLogFields := impl.logFieldsGetter()
	if len(throughLogFields) <= 0 {
		return v
	}
	count := 0
	var builder strings.Builder
	builder.WriteString("\t{ ")
	for k, v := range throughLogFields {
		msg := ""
		if count > 0 {
			msg = fmt.Sprintf(",\"%s\":\"%v\"", k, v)
		} else {
			msg = fmt.Sprintf("\"%s\":\"%v\"", k, v)
		}
		builder.WriteString(msg)
		count++
	}
	builder.WriteString(" }")
	newV := make([]interface{}, 0, len(v)+1)
	newV = append(newV, v...)
	newV = append(newV, builder.String())
	return newV
}

func (impl *troughLogTagsDecorator) LogwWrapper(msg string, logFields LogFields) (string, LogFields) {
	throughLogFields := impl.logFieldsGetter()
	if len(throughLogFields) <= 0 {
		return msg, logFields
	}
	newLogFields := make(LogFields, len(logFields)+len(throughLogFields))
	for k, v := range throughLogFields {
		newLogFields[k] = v
	}
	if len(logFields) > 0 {
		for k, v := range logFields {
			newLogFields[k] = v
		}
	}
	return msg, newLogFields
}
