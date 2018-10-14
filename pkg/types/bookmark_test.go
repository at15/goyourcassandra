package types_test

import (
	"testing"

	"github.com/dyweb/gommon/util/testutil"
	"github.com/stretchr/testify/assert"

	"github.com/at15/goyourcassandra/pkg/types"
)

func TestBookmark_Decode(t *testing.T) {
	var b types.Bookmark

	testutil.ReadYAMLToStrict(t, "testdata/bookmark.yml", &b)
	assert.Equal(t, "localhost", b.Host)
}
