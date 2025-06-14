package docker

import (
	"context"
	"os"
	"os/exec"
	"strings"
	"sync"
)

var containerCreationLock sync.Mutex

func pullImage(image string) error {
	command := exec.Command("docker", "pull", image)
	// Stream the command output to the console
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command.Run()
}

func IsContainerRunning(ctx context.Context, containerName string) (bool, error) {
	command := exec.CommandContext(ctx,
		"docker",
		"ps",
		"-q",
		"--filter",
		"name="+containerName,
	)

	output, err := command.Output()
	if err != nil {
		return false, err
	}

	return len(output) > 0, nil
}

func CreateContainer(ctx context.Context, containerName string, image string) error {
	if err := pullImage(image); err != nil {
		return err
	}
	createCmd := exec.CommandContext(ctx,
		"docker",
		"run",
		"--rm",
		"--network", "none",
		"-d",
		"--name", containerName,
		"-v", "/tmp/submissions:/tmp/submissions",
		image,
		"tail", "-f", "/dev/null",
	)

	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr

	return createCmd.Run()
}

func EnsureContainerIsRunning(ctx context.Context, containerName string, image string) error {
	containerCreationLock.Lock()
	defer containerCreationLock.Unlock()

	isRunning, err := IsContainerRunning(ctx, containerName)
	if err != nil {
		return err
	}
	if !isRunning {
		if err := CreateContainer(
			ctx,
			containerName,
			image,
		); err != nil {
			return err
		}
	}
	return nil
}

func ExecuteCodeCommand(ctx context.Context, containerName string, command string, args []string, codePath string) *exec.Cmd {
	commandArguments := []string{
		"exec",
		"-i",
		containerName,
		command,
	}
	commandArguments = append(commandArguments, args...)
	commandArguments = append(commandArguments, codePath)
	execCmd := exec.CommandContext(ctx, "docker", commandArguments...)
	return execCmd
}

func RemoveContainers(ctx context.Context, containerNames []string) error {
	removeCmd := exec.CommandContext(ctx,
		"docker",
		"rm",
		"-f",
		strings.Join(containerNames, " "),
	)
	return removeCmd.Run()
}
