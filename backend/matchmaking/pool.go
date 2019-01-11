package matchmaking

import (
	"fmt"
	"sync"
	"sync/atomic"
	"github.com/sirupsen/logrus"

)

type Pool struct {
	Games map[int]*Game
	mu    sync.Mutex
	incr  uint32
}

func NewPool() *Pool {
	p := &Pool{
		Games: make(map[int]*Game),
		incr:  2,
	}

	return p
}

func (p *Pool) AddGame(game *Game) error {
	defer p.mu.Unlock()
	p.mu.Lock()

	id := p.nextID()
	game.ID = id

	// TODO: Use Lobby ID
	game.LobbyID = 1

	p.Games[id] = game
	logrus.Println("ADDED game w/ id: ", game)

	return nil
}

func (p *Pool) nextID() int {
	i := atomic.AddUint32(&p.incr, 1)
	p.incr = i
	return int(i)
}

func (p *Pool) GetGame(id int) (*Game, error) {
	game, ok := p.Games[id]
	if !ok {
		return nil, fmt.Errorf("Cannot find game with ID:%d via matchmaking", id)
	}
	return game, nil
}

func (p *Pool) FindAvailableGame() (*Game, error) {
	for _, game := range p.Games {
		return game, nil
	}

	return nil, fmt.Errorf("Cannot find any game")
}
