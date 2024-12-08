package commentutil

import (
	"strings"
	"testing"

	"github.com/khulnasoft/codebase"
	"github.com/khulnasoft/codebase/filter"
	"github.com/khulnasoft/codebase/proto/rdf"
)

func TestCommentBody(t *testing.T) {
	tests := []struct {
		in   *codebase.Comment
		want string
	}{
		{
			in: &codebase.Comment{
				ToolName: "tool-name",
				Result: &filter.FilteredDiagnostic{
					Diagnostic: &rdf.Diagnostic{
						Message: "test message 1",
					},
				},
			},
			want: `
**[tool-name]** <sub>reported by [codebase](https://github.com/khulnasoft/codebase) :dog:</sub><br>test message 1
`,
		},
		{
			in: &codebase.Comment{
				Result: &filter.FilteredDiagnostic{
					Diagnostic: &rdf.Diagnostic{
						Message: "test message 2 (no tool)",
					},
				},
			},
			want: `
<sub>reported by [codebase](https://github.com/khulnasoft/codebase) :dog:</sub><br>test message 2 (no tool)
`,
		},
		{
			in: &codebase.Comment{
				ToolName: "global-tool-name",
				Result: &filter.FilteredDiagnostic{
					Diagnostic: &rdf.Diagnostic{
						Message: "test message 3",
						Source:  &rdf.Source{Name: "custom-tool-name"},
					},
				},
			},
			want: `
**[custom-tool-name]** <sub>reported by [codebase](https://github.com/khulnasoft/codebase) :dog:</sub><br>test message 3
`,
		},
		{
			in: &codebase.Comment{
				Result: &filter.FilteredDiagnostic{
					Diagnostic: &rdf.Diagnostic{
						Message:  "test message 4",
						Source:   &rdf.Source{Name: "tool-name"},
						Severity: rdf.Severity_WARNING,
					},
				},
			},
			want: `
⚠️ **[tool-name]** <sub>reported by [codebase](https://github.com/khulnasoft/codebase) :dog:</sub><br>test message 4
`,
		},
		{
			in: &codebase.Comment{
				Result: &filter.FilteredDiagnostic{
					Diagnostic: &rdf.Diagnostic{
						Message: "test message 5 (code)",
						Source:  &rdf.Source{Name: "tool-name"},
						Code: &rdf.Code{
							Value: "CODE14",
						},
					},
				},
			},
			want: `
**[tool-name]** <CODE14> <sub>reported by [codebase](https://github.com/khulnasoft/codebase) :dog:</sub><br>test message 5 (code)
`,
		},
		{
			in: &codebase.Comment{
				Result: &filter.FilteredDiagnostic{
					Diagnostic: &rdf.Diagnostic{
						Message: "test message 6 (code with URL)",
						Source:  &rdf.Source{Name: "tool-name"},
						Code: &rdf.Code{
							Value: "CODE14",
							Url:   "https://example.com/#CODE14",
						},
					},
				},
			},
			want: `
**[tool-name]** <[CODE14](https://example.com/#CODE14)> <sub>reported by [codebase](https://github.com/khulnasoft/codebase) :dog:</sub><br>test message 6 (code with URL)
`,
		},
	}
	for _, tt := range tests {
		want := strings.Trim(tt.want, "\n")
		if got := MarkdownComment(tt.in); got != want {
			t.Errorf("got unexpected comment.\ngot:\n%s\nwant:\n%s", got, want)
		}
	}
}
