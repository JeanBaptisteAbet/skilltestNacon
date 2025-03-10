package grpcliveeventserver

import (
	"context"
	"time"

	pb "skilltestnacon/api/grpcserver/liveevents"
	"skilltestnacon/database"
)

type Server struct {
	DB database.DB
	pb.LiveEventsServiceServer
}

func (s *Server) CreateEvent(ctx context.Context, in *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	id, err := s.DB.CreateEvent(ctx, database.LiveEvent{
		Title:       in.Title,
		Description: in.Description,
		StartTime:   time.Unix(in.StartTime, 0),
		Rewards:     in.Rewards,
	})
	if err != nil {
		return &pb.CreateEventResponse{Id: 0}, err
	}

	return &pb.CreateEventResponse{Id: id}, nil
}

func (s *Server) UpdateEvent(ctx context.Context, in *pb.UpdateEventRequest) (*pb.UpdateEventResponse, error) {
	event := database.LiveEvent{
		ID:          int(in.Id),
		Title:       in.Title,
		Description: in.Description,
		Rewards:     in.Rewards,
	}

	if in.EndTime != 0 {
		endtime := time.Unix(in.EndTime, 0)
		event.EndTime = &endtime
	}

	err := s.DB.UpdateEvent(ctx, event)

	return &pb.UpdateEventResponse{}, err
}

func (s *Server) DeleteEvent(ctx context.Context, in *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	err := s.DB.DeleteEvent(ctx, int(in.Id))

	return &pb.DeleteEventResponse{}, err
}

func (s *Server) ListEvents(ctx context.Context, in *pb.ListEventsRequest) (*pb.ListEventsResponse, error) {
	events, err := s.DB.AllEvent(ctx)
	if err != nil {
		return &pb.ListEventsResponse{}, err
	}

	return &pb.ListEventsResponse{LiveEvents: dbEventsToPBEvents(events)}, nil
}

// transform a slice of event from type database.LiveEvent to grpc pb.LiveEvent
func dbEventsToPBEvents(dbEvents []database.LiveEvent) []*pb.LiveEvent {
	pbEvents := make([]*pb.LiveEvent, len(dbEvents))
	for i, dbEvent := range dbEvents {
		e := &pb.LiveEvent{
			Id:          int64(dbEvent.ID),
			Title:       dbEvent.Title,
			Description: dbEvent.Description,
			StartTime:   dbEvent.StartTime.Unix(),
			Rewards:     dbEvent.Rewards,
		}
		if dbEvent.EndTime != nil {
			e.EndTime = dbEvent.EndTime.Unix()
		}

		pbEvents[i] = e
	}
	return pbEvents
}
