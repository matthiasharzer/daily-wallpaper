package local

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const stateFileName = "state.json"

type State struct {
	CurrentWallpaperURL  string `json:"currentWallpaperURL"`
	CurrentWallpaperFile string `json:"currentWallpaperFile"`
}

func getStatePath() string {
	return filepath.Join(StorageDirectory, stateFileName)
}

type Storage struct{}

func (s *Storage) ReadState() (State, error) {
	statePath := getStatePath()

	_, err := os.Stat(statePath)
	if os.IsNotExist(err) {
		return State{}, nil
	} else if err != nil {
		return State{}, fmt.Errorf("failed to stat state file: %w", err)
	}

	data, err := os.ReadFile(statePath)
	if err != nil {
		return State{}, fmt.Errorf("failed to read state file: %w", err)
	}

	var state State
	err = json.Unmarshal(data, &state)
	if err != nil {
		return State{}, fmt.Errorf("failed to unmarshal state file: %w", err)
	}

	return state, nil
}

func (s *Storage) WriteState(state State) error {
	statePath := getStatePath()

	data, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("failed to marshal state: %w", err)
	}
	err = os.WriteFile(statePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write state file: %w", err)
	}
	return nil
}
