// Queue Manager created with Chat GPT. I need to dig into the code to see what a queue manager does so I can code it solo.
package queue

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

type MatchFound struct {
	MatchID    string
	OpponentID string
}

type waiter struct {
	PlayerID string
	Mmr      int
	Region   string
	JoinedAt time.Time
	Notify   chan MatchFound // per-player signal channel
}

type enqueueReq struct{ W waiter }
type cancelReq struct{ PlayerID string }
type subscribeReq struct {
	PlayerID string
	Ch       chan MatchFound // used to attach the notify channel for a player (Queue call)
}
type unsubscribeReq struct{ PlayerID string }

type QueueManager struct {
	// API
	enqueueCh     chan enqueueReq
	cancelCh      chan cancelReq
	subscribeCh   chan subscribeReq
	unsubscribeCh chan unsubscribeReq

	// internal state (owned by run loop)
	waiting []waiter                   // FIFO per region (simple demo: one region queue)
	byID    map[string]waiter          // fast lookup to dedupe/cancel
	subs    map[string]chan MatchFound // playerID -> notify channel
	muStart sync.Once
}

func NewQueueManager() *QueueManager {
	return &QueueManager{
		enqueueCh:     make(chan enqueueReq),
		cancelCh:      make(chan cancelReq),
		subscribeCh:   make(chan subscribeReq),
		unsubscribeCh: make(chan unsubscribeReq),
		byID:          make(map[string]waiter),
		subs:          make(map[string]chan MatchFound),
	}
}

func (q *QueueManager) Start(ctx context.Context) {
	q.muStart.Do(func() {
		go q.loop(ctx)
	})
}

func (q *QueueManager) Enqueue(w waiter)       { q.enqueueCh <- enqueueReq{W: w} }
func (q *QueueManager) Cancel(playerID string) { q.cancelCh <- cancelReq{PlayerID: playerID} }
func (q *QueueManager) Subscribe(playerID string, ch chan MatchFound) {
	q.subscribeCh <- subscribeReq{PlayerID: playerID, Ch: ch}
}
func (q *QueueManager) Unsubscribe(playerID string) {
	q.unsubscribeCh <- unsubscribeReq{PlayerID: playerID}
}

func (q *QueueManager) loop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case s := <-q.subscribeCh:
			q.subs[s.PlayerID] = s.Ch
			log.Printf("%s has subscribed", s.PlayerID)

		case u := <-q.unsubscribeCh:
			delete(q.subs, u.PlayerID)
			log.Printf("%s has unsubscribed", u.PlayerID)

		case req := <-q.enqueueCh:
			// idempotency: if already queued, ignore
			if _, exists := q.byID[req.W.PlayerID]; exists {
				break
			}
			// If someone is waiting, pair immediately (1v1)
			if len(q.waiting) > 0 {
				opp := q.waiting[0]
				q.waiting = q.waiting[1:]
				delete(q.byID, opp.PlayerID)

				matchID := uuid.NewString()

				// notify both (only if still subscribed)
				if chOpp, ok := q.subs[opp.PlayerID]; ok {
					select {
					case chOpp <- MatchFound{MatchID: matchID, OpponentID: req.W.PlayerID}:
					default: /* drop if receiver not ready */
					}
				}
				if chReq, ok := q.subs[req.W.PlayerID]; ok {
					select {
					case chReq <- MatchFound{MatchID: matchID, OpponentID: opp.PlayerID}:
					default:
					}
				}
			} else {
				// else queue this player
				q.waiting = append(q.waiting, req.W)
				q.byID[req.W.PlayerID] = req.W
				log.Printf("%s has been queued", req.W.PlayerID)
			}

		case c := <-q.cancelCh:
			// remove from waiting if present
			if w, ok := q.byID[c.PlayerID]; ok {
				// compact waiting slice (small queue; fine for demo)
				n := make([]waiter, 0, len(q.waiting))
				for _, it := range q.waiting {
					if it.PlayerID != w.PlayerID {
						n = append(n, it)
					}
				}
				q.waiting = n
				delete(q.byID, w.PlayerID)
			}
		}
	}
}
