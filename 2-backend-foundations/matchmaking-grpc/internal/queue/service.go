package queue

import (
	"context"
	q "matchmaking/cmd/proto"
	"matchmaking/internal/auth"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type QueueService struct {
	q.UnimplementedMatchmakingServiceServer
	qm *QueueManager
}

func NewQueueService(qm *QueueManager) *QueueService {
	return &QueueService{qm: qm}
}

func (s *QueueService) Queue(req *q.QueueRequest, stream q.MatchmakingService_QueueServer) error {
	ctx := stream.Context()
	playerID := ctx.Value(auth.PlayerIDKey).(string)
	if playerID == "" {
		return status.Error(codes.Unauthenticated, "no player")
	}

	notify := make(chan MatchFound, 1)
	s.qm.Subscribe(playerID, notify)
	defer s.qm.Unsubscribe(playerID)

	if err := stream.Send(&q.QueueEvent{Type: q.QueueEvent_QUEUED}); err != nil {
		return err
	}

	s.qm.Enqueue(waiter{
		PlayerID: playerID,
		JoinedAt: time.Now(),
		Notify:   notify,
	})

	timeout := time.Duration(120)
	timer := time.NewTimer(timeout * time.Second)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		// best-effort cancel
		s.qm.Cancel(playerID)
		return status.Error(codes.Canceled, "client cancelled")

	case <-timer.C:
		s.qm.Cancel(playerID)
		return stream.Send(&q.QueueEvent{Type: q.QueueEvent_TIMEOUT})

	case mf := <-notify:
		return stream.Send(&q.QueueEvent{
			Type:       q.QueueEvent_MATCH_FOUND,
			MatchId:    mf.MatchID,
			OpponentId: mf.OpponentID,
		})
	}

}

func (s *QueueService) CancelQueue(ctx context.Context, _ *q.Empty) (*q.Empty, error) {
	playerID, _ := ctx.Value(auth.PlayerIDKey).(string)
	if playerID != "" {
		s.qm.Cancel(playerID)
	}
	return &q.Empty{}, nil
}
