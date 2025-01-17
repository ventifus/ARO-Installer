package installer

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	"github.com/openshift/ARO-Installer/pkg/cluster/graph"
	"github.com/openshift/ARO-Installer/pkg/util/stringutils"
)

func (m *manager) persistGraph(ctx context.Context, g graph.Graph) error {
	resourceGroup := stringutils.LastTokenByte(m.oc.Properties.ClusterProfile.ResourceGroupID, '/')
	clusterStorageAccountName := "cluster" + m.oc.Properties.StorageSuffix

	exists, err := m.graph.Exists(ctx, resourceGroup, clusterStorageAccountName)
	if err != nil || exists {
		return err
	}

	// the graph is quite big, so we store it in a storage account instead of in cosmosdb
	return m.graph.Save(ctx, resourceGroup, clusterStorageAccountName, g)
}
