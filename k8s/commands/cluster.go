package commands

import (
	"context"

	"golang.org/x/xerrors"

	tunnel_checks "go.khulnasoft.com/tunnel-checks"
	k8sArtifacts "go.khulnasoft.com/tunnel-kubernetes/pkg/artifacts"
	"go.khulnasoft.com/tunnel-kubernetes/pkg/k8s"
	"go.khulnasoft.com/tunnel-kubernetes/pkg/tunnelk8s"
	"go.khulnasoft.com/pkg/commands/operation"
	"go.khulnasoft.com/pkg/flag"
	"go.khulnasoft.com/pkg/log"
	"go.khulnasoft.com/pkg/types"
)

// clusterRun runs scan on kubernetes cluster
func clusterRun(ctx context.Context, opts flag.Options, cluster k8s.Cluster) error {
	if err := validateReportArguments(opts); err != nil {
		return err
	}
	var artifacts []*k8sArtifacts.Artifact
	var err error
	switch opts.Format {
	case types.FormatCycloneDX:
		artifacts, err = tunnelk8s.New(cluster).ListClusterBomInfo(ctx)
		if err != nil {
			return xerrors.Errorf("get k8s artifacts with node info error: %w", err)
		}
	case types.FormatJSON, types.FormatTable:
		k8sOpts := []tunnelk8s.K8sOption{
			tunnelk8s.WithExcludeNamespaces(opts.ExcludeNamespaces),
			tunnelk8s.WithIncludeNamespaces(opts.IncludeNamespaces),
			tunnelk8s.WithExcludeKinds(opts.ExcludeKinds),
			tunnelk8s.WithIncludeKinds(opts.IncludeKinds),
			tunnelk8s.WithExcludeOwned(opts.ExcludeOwned),
		}
		if opts.Scanners.AnyEnabled(types.MisconfigScanner) && !opts.DisableNodeCollector {
			artifacts, err = tunnelk8s.New(cluster, k8sOpts...).ListArtifactAndNodeInfo(ctx, nodeCollectorOptions(ctx, opts)...)
			if err != nil {
				return xerrors.Errorf("get k8s artifacts with node info error: %w", err)
			}
		} else {
			artifacts, err = tunnelk8s.New(cluster, k8sOpts...).ListArtifacts(ctx)
			if err != nil {
				return xerrors.Errorf("get k8s artifacts error: %w", err)
			}
		}
	default:
		return xerrors.Errorf(`unknown format %q. Use "json" or "table" or "cyclonedx"`, opts.Format)
	}

	if !opts.DisableNodeCollector && !opts.Quiet {
		log.InfoContext(ctx, "Node scanning is enabled")
		log.InfoContext(ctx, "If you want to disable Node scanning via an in-cluster Job, please try '--disable-node-collector' to disable the Node-Collector job.")
	}
	runner := newRunner(opts, cluster.GetCurrentContext())
	return runner.run(ctx, artifacts)
}

func nodeCollectorOptions(ctx context.Context, opts flag.Options) []tunnelk8s.NodeCollectorOption {
	nodeCollectorOptions := []tunnelk8s.NodeCollectorOption{
		tunnelk8s.WithScanJobNamespace(opts.NodeCollectorNamespace),
		tunnelk8s.WithIgnoreLabels(opts.ExcludeNodes),
		tunnelk8s.WithScanJobImageRef(opts.NodeCollectorImageRef),
		tunnelk8s.WithTolerations(opts.Tolerations),
	}

	ctx = log.WithContextPrefix(ctx, log.PrefixMisconfiguration)
	contentPath, err := operation.InitBuiltinChecks(ctx, opts.CacheDir, opts.Quiet, opts.SkipCheckUpdate,
		opts.MisconfOptions.ChecksBundleRepository, opts.RegistryOpts())
	if err != nil {
		log.Error("Falling back to embedded checks", log.Err(err))
		nodeCollectorOptions = append(nodeCollectorOptions,
			[]tunnelk8s.NodeCollectorOption{
				tunnelk8s.WithEmbeddedCommandFileSystem(tunnel_checks.EmbeddedK8sCommandsFileSystem),
				tunnelk8s.WithEmbeddedNodeConfigFilesystem(tunnel_checks.EmbeddedConfigCommandsFileSystem),
			}...)
	}

	complianceCommandsIDs := getComplianceCommands(opts)
	nodeCollectorOptions = append(nodeCollectorOptions, []tunnelk8s.NodeCollectorOption{
		tunnelk8s.WithCommandPaths(contentPath),
		tunnelk8s.WithSpecCommandIds(complianceCommandsIDs),
	}...)
	return nodeCollectorOptions
}

func getComplianceCommands(opts flag.Options) []string {
	var commands []string
	if opts.Compliance.Spec.ID != "" {
		for _, control := range opts.Compliance.Spec.Controls {
			for _, command := range control.Commands {
				if command.ID != "" {
					commands = append(commands, command.ID)
				}
			}
		}
	}
	return commands
}
