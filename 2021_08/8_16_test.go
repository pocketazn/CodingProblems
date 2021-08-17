package _2021_08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialize(t *testing.T) {
	type TestCase struct {
		Name string
		Tree Node
		ExpectedSerialization string
	}

	Cases := []TestCase{
		{
			Name: "Single Node",
			Tree: Node{
				Data: "root",
			},
			ExpectedSerialization:"[root..]",
		},
		{
			Name: "Single Node with Left Node",
			Tree: Node{
				Data: "root",
				Left: &Node{
					Data: "Left",
				},
			},
			ExpectedSerialization:"[root.[Left..].]",
		},
		{
			Name: "Single Node with Right Node",
			Tree: Node{
				Data: "root",
				Right: &Node{
					Data: "Right",
				},
			},
			ExpectedSerialization:"[root..[Right..]]",
		},
		{
			Name: "Single Node with Left and Right Node",
			Tree: Node{
				Data: "root",
				Right: &Node{
					Data: "Right",
				},
				Left: &Node{
					Data: "Left",
				},
			},
			ExpectedSerialization:"[root.[Left..].[Right..]]",
		},
		{
			Name: "Deep Left Path",
			Tree: Node{
				Data: "root",
				Left: &Node{
					Data: "Left",
					Left: &Node{
						Data: "Left",
						Left: &Node{
							Data: "Left",
						},
					},
				},
			},
			ExpectedSerialization:"[root.[Left.[Left.[Left..].].].]",
		},
	}

	for _, tt := range Cases {
		t.Run(tt.Name, func(t *testing.T) {
			result := Serialize(&tt.Tree)
			assert.Equal(t, tt.ExpectedSerialization, result)
		})
	}
}