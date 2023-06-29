package main

import (
	"fmt"
	"github.com/kaspanet/kaspad/blockdag"
	"github.com/kaspanet/kaspad/domain/consensus/model/externalapi"
	"github.com/kaspanet/kaspad/infrastructure/logger"
	"github.com/kaspanet/kaspad/stability-tests/common"
	"os"
)

func main() {
	finalityInterval := uint64(5) // Set the desired finality interval

	// Create a new instance of the block DAG
	blockDAG, teardown, err := common.SetupBlockDAG("testdata", false)
	if err != nil {
		fmt.Printf("Error setting up block DAG: %s\n", err)
		os.Exit(1)
	}
	defer teardown(false)

	// Get the current tip of the block DAG
	tipHash, err := blockDAG.VirtualBlockDAG().GetVirtualTopHash()
	if err != nil {
		fmt.Printf("Error getting virtual top hash: %s\n", err)
		os.Exit(1)
	}

	// Generate `finalityInterval - 1` blocks on top of the current tip
	for i := uint64(0); i < finalityInterval-1; i++ {
		tip, err := blockDAG.BlockByHash(tipHash)
		if err != nil {
			fmt.Printf("Error getting block by hash: %s\n", err)
			os.Exit(1)
		}

		// Create a new block template based on the current tip
		blockTemplate, err := blockDAG.BuildBlockWithParents([]*externalapi.DomainHash{tip.Hash()})
		if err != nil {
			fmt.Printf("Error building block template: %s\n", err)
			os.Exit(1)
		}

		// Mine the new block
		newBlock, err := common.MineBlock(blockDAG, blockTemplate, logger.New(logger.LogOff))
		if err != nil {
			fmt.Printf("Error mining block: %s\n", err)
			os.Exit(1)
		}

		// Update the tip to the newly generated block
		tipHash = newBlock.Header().Hash()
	}

	fmt.Printf("Chain of %d blocks successfully generated.\n", finalityInterval-1)
}
