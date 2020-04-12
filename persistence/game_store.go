package persistence

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"minesweeper-API/types"
	"os"
)

type GameStore struct {
	db *redis.Client
	logger *logrus.Logger
}

func NewGameStore(log *logrus.Logger) *GameStore {
	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	log.Printf("REDIS HOST: %s - REDIS PASSWORD: %s", redisHost, redisPassword)

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisHost,
		Password: redisPassword,
		DB: 0,
	})

	return &GameStore{db: redisClient, logger: log}
}

func (s *GameStore) Insert(game *types.Game) error {
	gameAsJson, err := json.Marshal(game)
	if err != nil {
		s.logger.Errorln(err)
	}

	err = s.db.Set(game.Name, gameAsJson, 0).Err()
	if err != nil {
		s.logger.Errorln(err)
	}

	s.logger.Infof("Game saved - Success: %s", game.Name)

	return nil
}

func (s *GameStore) Update(game *types.Game) error {
	val, err := s.db.Get(game.Name).Result()
	if val != "" && err != nil {
		s.logger.Errorf("game do not exist: %s", game.Name)
		return errors.New("game do not exist")
	}

	gameAsJson, err := json.Marshal(game)
	if err != nil {
		s.logger.Errorln(err)
	}

	err = s.db.Set(game.Name, gameAsJson, 0).Err()
	if err != nil {
		s.logger.Errorln(err)
	}

	s.logger.Infof("Game saved - Success: %s", game.Name)

	return nil
}

func (s *GameStore) GetByName(name string) (*types.Game, error) {

	val, err := s.db.Get(name).Result()
	if err != nil {
		s.logger.Errorln(err)
	}
	s.logger.Infoln("Game retrieved: %s", val)

	var gameStored types.Game
	parsingError := json.Unmarshal([]byte(val), &gameStored)

	if parsingError == nil {
		return &gameStored, nil
	}

	return nil, errors.New("game not found")
}
