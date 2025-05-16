package doreader

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommentDeletor(t *testing.T) {
	code := `
	a = 2
	a = b // однострочный комментарий
	// а теперь один на строке
	
	/* многострочный комментарий на одной строке */
	/*
		многострочный комментарий
		на нескольких строках
	*/

	a = b /* многострочный комментарий на одной строке после кода */
	a = /* многострочный комментарий на одной строке внутри кода */ b 
/`

	target := `
	a = 2
	a = b 
	
	
	
	

	a = b 
	a =  b 
/`

	reader := strings.NewReader(code)
	deletor := NewCommentDeletor(reader)
	res, err := io.ReadAll(deletor)

	require.NoError(t, err)
	require.Equal(t, target, string(res))
}
