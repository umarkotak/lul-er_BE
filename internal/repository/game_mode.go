package repository

import (
	"context"
	"errors"

	"github.com/umarkotak/lul-er_BE/internal/config"
	"github.com/umarkotak/lul-er_BE/internal/models"
)

func GetGameMode(mode string) (models.GameMode, error) {
	var gameMode models.GameMode

	if mode == "" {
		return gameMode, errors.New("mode can't be blank")
	}

	fbGameModesRef := config.GetConfig().FbGameModesRef
	fbGameModeRef := fbGameModesRef.Child(mode)
	fbGameModeRef.Get(context.Background(), &gameMode)

	return gameMode, nil
}
