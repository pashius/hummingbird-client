package rollup

import (
	"hummingbird/node"
	"hummingbird/node/lightlink/types"
	"math/big"
	"testing"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func TestValidateBundles(t *testing.T) {
	type args struct {
		bundles []*node.Bundle
		head    uint64
	}
	tests := []struct {
		name   string
		args   args
		errStr string
	}{
		{
			name: "happy path should pass",
			args: args{
				bundles: []*node.Bundle{
					{Blocks: []*types.Block{
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(2),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(3),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(4),
						})}},
					{Blocks: []*types.Block{
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(5),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(6),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(7),
						})}},
					{Blocks: []*types.Block{
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(8),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(9),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(10),
						})}},
				},
				head: 1,
			},
			errStr: "",
		},
		{
			name: "test should fail with empty bundles",
			args: args{
				bundles: []*node.Bundle{},
				head:    6,
			},
			errStr: "bundles are empty",
		},
		{
			name: "test should fail as bundle 0 is empty",
			args: args{
				bundles: []*node.Bundle{
					{Blocks: []*types.Block{}},
				},
				head: 6,
			},
			errStr: "bundle 0 is empty",
		},
		{
			name: "test should fail as bundle 1 is empty",
			args: args{
				bundles: []*node.Bundle{
					{Blocks: []*types.Block{
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(2),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(3),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(4),
						})}},
					{Blocks: []*types.Block{}},
				},
				head: 1,
			},
			errStr: "bundle 1 is empty",
		},
		{
			name: "test should fail as the first block in bundle 0 is not the head+1",
			args: args{
				bundles: []*node.Bundle{
					{Blocks: []*types.Block{
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(2),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(3),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(4),
						})}},
					{Blocks: []*types.Block{
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(5),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(6),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(7),
						})}},
					{Blocks: []*types.Block{
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(8),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(9),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(10),
						})}},
				},
				head: 0,
			},
			errStr: "first block in bundle 0 is not the correct height",
		},
		{
			name: "test should fail as bundle 1 has a nil block",
			args: args{
				bundles: []*node.Bundle{
					{Blocks: []*types.Block{
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(2),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(3),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(4),
						})}},
					{Blocks: []*types.Block{nil}},
				},
				head: 1,
			},
			errStr: "block 0 in bundle 1 is nil",
		},
		{
			name: "test should fail as the blocks in bundle 2 are not sequential",
			args: args{
				bundles: []*node.Bundle{
					{Blocks: []*types.Block{
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(2),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(3),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(4),
						})}},
					{Blocks: []*types.Block{
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(5),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(7),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(8),
						})}},
					{Blocks: []*types.Block{
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(9),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(10),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(11),
						})}},
				},
				head: 1,
			},
			errStr: "block 1 in bundle 1 is not sequential",
		},
		{
			name: "test should fail as the first block in bundles 2 & 3 are not parent block + 1",
			args: args{
				bundles: []*node.Bundle{
					{Blocks: []*types.Block{
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(2),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(3),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(4),
						})}},
					{Blocks: []*types.Block{
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(7),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(8),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(9),
						})}},
					{Blocks: []*types.Block{
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(12),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(13),
						}),
						types.NewBlockWithHeader(&ethtypes.Header{
							Number: big.NewInt(14),
						})}},
				},
				head: 1,
			},
			errStr: "first block in bundle 1 is not the correct height",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateBundles(tt.args.bundles, tt.args.head)
			if err != nil && err.Error() != tt.errStr {
				t.Errorf("ValidateBundles() error = %v, wantErrStr %v", err, tt.errStr)
			} else if err == nil && tt.errStr != "" {
				t.Errorf("ValidateBundles() expected error but got none, wantErrStr: %v", tt.errStr)
			}
		})
	}
}
