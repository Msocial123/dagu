package dag

import (
	"os"
	"path"
	"testing"

	"github.com/dagu-dev/dagu/internal/config"
	"github.com/dagu-dev/dagu/internal/util"
	"github.com/stretchr/testify/require"
)

var (
	testdataDir = path.Join(util.MustGetwd(), "testdata")
	testHomeDir = path.Join(util.MustGetwd(), "testdata/home")
)

func TestMain(m *testing.M) {
	err := os.Setenv("HOME", testHomeDir)
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestDAG_String(t *testing.T) {
	t.Run("String representation of default.yaml", func(t *testing.T) {
		cfg, err := config.Load()
		require.NoError(t, err)

		loader := NewLoader(cfg)
		dg, err := loader.Load("", path.Join(testdataDir, "default.yaml"), "")
		require.NoError(t, err)

		ret := dg.String()
		require.Contains(t, ret, "Name: default")
	})
}

func TestDAG_SockAddr(t *testing.T) {
	t.Run("Unix Socket", func(t *testing.T) {
		d := &DAG{Location: "testdata/testDag.yml"}
		require.Regexp(t, `^/tmp/@dagu-testDag-[0-9a-f]+\.sock$`, d.SockAddr())
	})
}
