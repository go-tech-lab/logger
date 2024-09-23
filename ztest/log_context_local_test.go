package test

import (
	. "github.com/go-tech-lab/framework/logger/src/log_context"
	"github.com/go-tech-lab/framework/logger/src/log_context/impl"
	_ "github.com/go-tech-lab/framework/logger/src/logger/impl"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLogContextLocal(t *testing.T) {
	asserts := require.New(t)
	logContextLocal := LogContextLocal()

	//test nil
	logContext := GetLogContext()
	asserts.Nil(logContext)

	//test put and get
	newLogContext := &LogContext{}
	logContextLocal.PutLogContext(newLogContext)
	retLogContext := logContextLocal.GetLogContext()
	asserts.Equal(newLogContext, retLogContext)

	//test remove
	logContextLocal.Clear()
	retLogContext = logContextLocal.GetLogContext()
	asserts.Nil(retLogContext)
}

func TestLogSkipDepth(t *testing.T) {
	asserts := require.New(t)
	logContextLocal := LogContextLocal()
	skipDepth := logContextLocal.SkipDepth()
	asserts.Equal(impl.DefaultSkipDepth, skipDepth)
}

func TestAddLogSkipDepth(t *testing.T) {
	asserts := require.New(t)
	logContextLocal := LogContextLocal()
	addSkipDepth := 2
	logContextLocal.AddSkipDepth(addSkipDepth)
	skipDepth := logContextLocal.SkipDepth()
	asserts.Equal(impl.DefaultSkipDepth+addSkipDepth, skipDepth)
	logContextLocal.AddSkipDepth(-addSkipDepth)
	skipDepth = logContextLocal.SkipDepth()
	asserts.Equal(impl.DefaultSkipDepth, skipDepth)
}

func TestModuleTag(t *testing.T) {
	asserts := require.New(t)
	logContextLocal := LogContextLocal()
	moduleTag := logContextLocal.ModuleTag()
	asserts.Equal(impl.DefaultModuleTag, moduleTag)
	newModuleTag := "MyModuleTag"
	logContextLocal.WithModuleTag(newModuleTag)
	moduleTag = logContextLocal.ModuleTag()
	asserts.Equal(newModuleTag, moduleTag)

	newModuleTag = ""
	logContextLocal.WithModuleTag(newModuleTag)
	moduleTag = logContextLocal.ModuleTag()
	asserts.Equal(newModuleTag, moduleTag)

	logContextLocal.ResetModuleTag()
	moduleTag = logContextLocal.ModuleTag()
	asserts.Equal(impl.DefaultModuleTag, moduleTag)
}
