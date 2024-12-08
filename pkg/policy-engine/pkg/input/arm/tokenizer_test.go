// © 2022-2023 Khulnasoft Limited All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package arm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenizesExpressions(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    string
		expected []token
	}{
		{
			name:  "tokenizes expression containing commas, brackets, and parentheses",
			input: "someFn(arg1, ARG_2)",
			expected: []token{
				identifier("someFn"),
				openParen{},
				identifier("arg1"),
				comma{},
				identifier("ARG_2"),
				closeParen{},
			},
		},
		{
			name:  "tokenizes expression containing variable whitespace",
			input: "someFn (\t\targ1,      ARG_2)",
			expected: []token{
				identifier("someFn"),
				openParen{},
				identifier("arg1"),
				comma{},
				identifier("ARG_2"),
				closeParen{},
			},
		},
		{
			name:  "tokenizes expression containing string literals",
			input: "someFn('a-string', 'another string')",
			expected: []token{
				identifier("someFn"),
				openParen{},
				stringLiteral("a-string"),
				comma{},
				stringLiteral("another string"),
				closeParen{},
			},
		},
		{
			name:  "tokenizes expression containing string literals containing escaped single quotes",
			input: "'some ''escaped'' ''quotes'''",
			expected: []token{
				stringLiteral("some 'escaped' 'quotes'"),
			},
		},
		{
			name:  "tokenizes expression containing string literals with special characters",
			input: "'[some]' 'more(special, characters)'",
			expected: []token{
				stringLiteral("[some]"),
				stringLiteral("more(special, characters)"),
			},
		},
		{
			name:  "tokenizes expression containing property dereferences",
			input: "resourceGroup().location",
			expected: []token{
				identifier("resourceGroup"),
				openParen{},
				closeParen{},
				dot{},
				identifier("location"),
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			output, err := tokenize(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, output)
		})
	}
}
