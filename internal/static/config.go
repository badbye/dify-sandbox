package static

import (
	"os"
	"strconv"

	"github.com/langgenius/dify-sandbox/internal/types"
	"github.com/langgenius/dify-sandbox/internal/utils/log"
	"gopkg.in/yaml.v3"
)

var difySandboxGlobalConfigurations types.DifySandboxGlobalConfigurations

func InitConfig(path string) error {
	difySandboxGlobalConfigurations = types.DifySandboxGlobalConfigurations{}

	// read config file
	configFile, err := os.Open(path)
	if err != nil {
		return err
	}

	defer configFile.Close()

	// parse config file
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&difySandboxGlobalConfigurations)
	if err != nil {
		return err
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err == nil {
		difySandboxGlobalConfigurations.App.Debug = debug
	}

	max_workers := os.Getenv("MAX_WORKERS")
	if max_workers != "" {
		difySandboxGlobalConfigurations.MaxWorkers, _ = strconv.Atoi(max_workers)
	}

	max_requests := os.Getenv("MAX_REQUESTS")
	if max_requests != "" {
		difySandboxGlobalConfigurations.MaxRequests, _ = strconv.Atoi(max_requests)
	}

	port := os.Getenv("SANDBOX_PORT")
	if port != "" {
		difySandboxGlobalConfigurations.App.Port, _ = strconv.Atoi(port)
	}

	timeout := os.Getenv("WORKER_TIMEOUT")
	if timeout != "" {
		difySandboxGlobalConfigurations.WorkerTimeout, _ = strconv.Atoi(timeout)
	}

	api_key := os.Getenv("API_KEY")
	if api_key != "" {
		difySandboxGlobalConfigurations.App.Key = api_key
	}

	python_path := os.Getenv("PYTHON_PATH")
	if python_path != "" {
		difySandboxGlobalConfigurations.PythonPath = python_path
	}

	if difySandboxGlobalConfigurations.PythonPath == "" {
		difySandboxGlobalConfigurations.PythonPath = "/usr/local/bin/python3"
	}

	nodejs_path := os.Getenv("NODEJS_PATH")
	if nodejs_path != "" {
		difySandboxGlobalConfigurations.NodejsPath = nodejs_path
	}

	if difySandboxGlobalConfigurations.NodejsPath == "" {
		difySandboxGlobalConfigurations.NodejsPath = "/usr/local/bin/node"
	}

	enable_network := os.Getenv("ENABLE_NETWORK")
	if enable_network != "" {
		difySandboxGlobalConfigurations.EnableNetwork, _ = strconv.ParseBool(enable_network)
	}

	if difySandboxGlobalConfigurations.EnableNetwork {
		log.Info("network has been enabled")
		socks5_proxy := os.Getenv("SOCKS5_PROXY")
		if socks5_proxy != "" {
			difySandboxGlobalConfigurations.Proxy.Socks5 = socks5_proxy
		}

		if difySandboxGlobalConfigurations.Proxy.Socks5 != "" {
			log.Info("using socks5 proxy: %s", difySandboxGlobalConfigurations.Proxy.Socks5)
		}

		https_proxy := os.Getenv("HTTPS_PROXY")
		if https_proxy != "" {
			difySandboxGlobalConfigurations.Proxy.Https = https_proxy
		}

		if difySandboxGlobalConfigurations.Proxy.Https != "" {
			log.Info("using https proxy: %s", difySandboxGlobalConfigurations.Proxy.Https)
		}

		http_proxy := os.Getenv("HTTP_PROXY")
		if http_proxy != "" {
			difySandboxGlobalConfigurations.Proxy.Http = http_proxy
		}

		if difySandboxGlobalConfigurations.Proxy.Http != "" {
			log.Info("using http proxy: %s", difySandboxGlobalConfigurations.Proxy.Http)
		}
	}
	return nil
}

// avoid global modification, use value copy instead
func GetDifySandboxGlobalConfigurations() types.DifySandboxGlobalConfigurations {
	return difySandboxGlobalConfigurations
}

type RunnerDependencies struct {
	PythonRequirements string
}

var runnerDependencies RunnerDependencies

func GetRunnerDependencies() RunnerDependencies {
	return runnerDependencies
}

func SetupRunnerDependencies() error {
	file, err := os.ReadFile("dependencies/python-requirements.txt")
	if err != nil {
		if err == os.ErrNotExist {
			return nil
		}
		return err
	}

	runnerDependencies.PythonRequirements = string(file)

	return nil
}
