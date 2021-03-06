// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package control

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/mattermost/mattermost-load-test-ng/loadtest/store/memstore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	seed := memstore.SetRandomSeed()
	fmt.Printf("Seed value is: %d\n", seed)
	os.Exit(m.Run())
}

func TestRandomizeUserName(t *testing.T) {
	name := RandomizeUserName("test-agent-1-user-4")
	assert.Regexp(t, regexp.MustCompile(`user[[:alpha:]]+-4`), name)

	name = RandomizeUserName("lt1-user4")
	assert.True(t, strings.HasPrefix(name, "lt1-user"))

	name = RandomizeUserName("testuser")
	assert.Equal(t, name, "testuser")
}

func TestRandomizeTeamDisplayName(t *testing.T) {
	name := RandomizeTeamDisplayName("badname")
	assert.Equal(t, "badname", name)

	name = RandomizeTeamDisplayName("team9")
	assert.True(t, strings.HasPrefix(name, "team9-"))

	name = RandomizeTeamDisplayName("team9-k")
	assert.True(t, strings.HasPrefix(name, "team9-"))
}

func TestGetErrOrigin(t *testing.T) {
	var origin string
	test := func() {
		origin = getErrOrigin()
	}
	test()
	fmt.Println(origin)
	require.True(t, strings.HasPrefix(origin, "control.TestGetErrOrigin"))
}

func TestEmulateUserTyping(t *testing.T) {
	search := "this is long enough"
	res := EmulateUserTyping(search, func(term string) UserActionResponse {
		return UserActionResponse{Info: term}
	})
	require.Nil(t, res.Err)
	require.Equal(t, search, res.Info)
	text := ""
	i := 0
	res = EmulateUserTyping(search, func(term string) UserActionResponse {
		text = term
		if i == 2 {
			return UserActionResponse{Err: errors.New("an error")}
		}
		i++
		return UserActionResponse{Info: text}
	})
	require.NotNil(t, res.Err)
	require.Equal(t, "an error", res.Err.Error())
}

func TestGenerateRandomSentences(t *testing.T) {
	randomize := GenerateRandomSentences(8)
	s := strings.Split(randomize, " ")
	require.Len(t, s, 8)

	randomize = GenerateRandomSentences(0)
	s = strings.Split(randomize, " ")
	require.Len(t, s, 1)
	require.Equal(t, s[0], "🙂")
}

func TestSelectWeighted(t *testing.T) {
	t.Run("empty weights", func(t *testing.T) {
		idx, err := SelectWeighted([]int{})
		require.Error(t, err)
		require.Equal(t, -1, idx)
	})

	t.Run("zero sum", func(t *testing.T) {
		weights := []int{
			0,
			0,
			0,
		}
		idx, err := SelectWeighted(weights)
		require.Error(t, err)
		require.Equal(t, -1, idx)
	})

	t.Run("weighted selection", func(t *testing.T) {
		weights := []int{
			1000,
			100,
			10,
		}

		distribution := make(map[int]int, len(weights))

		n := 10000
		for i := 0; i < n; i++ {
			idx, err := SelectWeighted(weights)
			require.NoError(t, err)
			distribution[idx]++
		}

		require.Greater(t, distribution[0], distribution[1])
		require.Greater(t, distribution[1], distribution[2])
	})
}
