package main

import (
	"fmt"
	"log/slog"
	"net/url"
	"os"

	"github.com/ifbyol/redeploy-applications/deployer/api"
)

func main() {
	token := os.Getenv("OKTETO_TOKEN")
	oktetoURL := os.Getenv("OKTETO_URL")

	logLevel := &slog.LevelVar{} // INFO
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))

	u, err := url.Parse(oktetoURL)
	if err != nil {
		logger.Error(fmt.Sprintf("Invalid OKTETO_URL %s", err))
		os.Exit(1)
	}

	nsList, err := api.GetNamespaces(u.Host, token, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("There was an error requesting the namespaces: %s", err))
		os.Exit(1)
	}

	for _, ns := range nsList {
		logger.Info(fmt.Sprintf("Processing namespace '%s'", ns.Name))

		applications, err := api.GetApplicationsWithinNamespace(u.Host, token, ns.Name, logger)
		if err != nil {
			logger.Error(fmt.Sprintf("There was an error requesting the applications within namespace '%s': %s", ns.Name, err))
			logger.Info("-----------------------------------------------")
			continue
		}

		for _, app := range applications {
			logger.Info(fmt.Sprintf("Applications within namespace '%s': %s", ns.Name, app.Name))
		}
		logger.Info("-----------------------------------------------")
	}
}
