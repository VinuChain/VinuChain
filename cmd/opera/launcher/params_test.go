package launcher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultBootnodesUseNetworkNames(t *testing.T) {
	require := require.New(t)

	require.Equal(vinuChainTestnetNetworkName, vinuChainTestnetHeader.NetworkName)
	require.Equal(vinuChainStagingMainnetNetworkName, vinuChainTestMainnetHeader.NetworkName)
	require.Equal(vinuChainMainnetNetworkName, vinuChainMainnetHeader.NetworkName)

	require.Equal(Bootnodes["test"], Bootnodes[vinuChainTestnetHeader.NetworkName])
	require.NotEmpty(Bootnodes[vinuChainTestnetHeader.NetworkName])

	require.Equal(Bootnodes["main"], Bootnodes[vinuChainNetworkName])
	require.NotEmpty(Bootnodes[vinuChainNetworkName])

	require.Equal(Bootnodes["main"], Bootnodes[vinuChainMainnetHeader.NetworkName])
	require.NotEmpty(Bootnodes[vinuChainMainnetHeader.NetworkName])

	_, ok := Bootnodes[vinuChainTestMainnetHeader.NetworkName]
	require.True(ok)
}
