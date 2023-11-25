package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/hjfitz/gitlab-pipeline-profiling/types"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if err == nil {
		return !info.IsDir()
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func filter(arr []string, criteria string) []string {
	filtered := []string{}
	for _, item := range arr {
		if item != criteria {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func getEnvFiles() []string {
	goEnv := os.Getenv("GO_ENV")
	envFiles := []string{}

	if goEnv != "prod" {
		if goEnv != "" {
			envFiles = append(envFiles, fmt.Sprintf(".env.%s", goEnv))
			envFiles = append(envFiles, fmt.Sprintf(".env.%s.local", goEnv))
		}
		envFiles = append(envFiles, ".env.local")
	} else {
		envFiles = append(envFiles, ".env.prod")
	}

	for _, envFile := range envFiles {
		if !fileExists(envFile) {
			fmt.Printf("File %s does not exist\n", envFile)
			envFiles = filter(envFiles, envFile)
		} else {
			fmt.Printf("Loading %s\n", envFile)
		}
	}

	return envFiles

}

func defaultToEnv(envMap map[string]string, key string) string {
	envVar := envMap[key]

	if envVar == "" {
		return os.Getenv(key)
	}

	return envVar

}

func GetConfig() types.Config {
	envFiles := getEnvFiles()

	rawConfig, err := godotenv.Read(envFiles...)

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	gitlabToken := defaultToEnv(rawConfig, "GITLAB_TOKEN")
	groupIdsRaw := defaultToEnv(rawConfig, "GROUP_IDS")

	groupIds := strings.Split(groupIdsRaw, ",")

	obfuscatedToken := gitlabToken[0:3] + strings.Repeat("*", len(gitlabToken)-3)

	log.Info().Str("gitlabToken", obfuscatedToken).Str("groupIds", groupIdsRaw).Msg("Loaded config")

	return types.Config{
		GitlabToken: gitlabToken,
		GroupIds:    groupIds,
	}

}
