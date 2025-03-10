package grpcliveeventserver

import (
	"context"
	"log"
	"net"
	"slices"
	"testing"
	"time"

	pb "skilltestnacon/api/grpcserver/liveevents"
	"skilltestnacon/database"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestCoreLogic(t *testing.T) {
	lis := bufconn.Listen(1024 * 1024)
	defer lis.Close()

	srv := grpc.NewServer()
	defer srv.Stop()

	db, err := database.InitDB(":memory:")
	if err != nil {
		t.Fatal(err)
	}

	pb.RegisterLiveEventsServiceServer(srv, &Server{DB: db})

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("srv.Serve %v", err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		conn.Close()
		t.Fatalf("grpc.DialContext %v", err)
	}
	defer conn.Close()

	client := pb.NewLiveEventsServiceClient(conn)

	event1 := pb.LiveEvent{
		Title:       "toto",
		Description: "description",
		StartTime:   time.Now().Unix(),
		Rewards:     "{rewards}",
	}

	//----------------------------------------- create one event ------------------------------- //
	resCreate, err := client.CreateEvent(context.Background(), &pb.CreateEventRequest{
		Title:       event1.Title,
		Description: event1.Description,
		StartTime:   event1.StartTime,
		Rewards:     event1.Rewards,
	})
	if err != nil {
		t.Errorf("client.CreateEvent %v", err)
	}

	if resCreate.Id != 1 {
		t.Errorf("Unexpected values %v", resCreate.Id)
	}

	//----------------------------------------- list the event inserted before ------------------------------- //
	resList, err := client.ListEvents(context.Background(), &pb.ListEventsRequest{})
	if err != nil {
		t.Fatalf("client.ListEvents %v", err)
	}

	if len(resList.LiveEvents) != 1 || !compare(resList.LiveEvents[0], &event1) {
		t.Errorf("client.ListEvents events doesn't match, got: %s,  want: %s", resList.LiveEvents[0].String(), event1.String())
	}

	//----------------------------------------- create a second event ------------------------------- //
	event2 := pb.LiveEvent{
		Title:       "toto2",
		Description: "description2",
		StartTime:   time.Now().Unix(),
		Rewards:     "{rewards}",
	}

	resCreate, err = client.CreateEvent(context.Background(), &pb.CreateEventRequest{
		Title:       event2.Title,
		Description: event2.Description,
		StartTime:   event2.StartTime,
		Rewards:     event2.Rewards,
	})
	if err != nil {
		t.Errorf("client.CreateEvent %v", err)
	}

	if resCreate.Id != 2 {
		t.Errorf("Unexpected values %v", resCreate.Id)
	}

	//----------------------------------------- update the first event ------------------------------- //

	endTime := time.Now().Unix()
	event1.EndTime = endTime

	_, err = client.UpdateEvent(context.Background(), &pb.UpdateEventRequest{
		Id:          1,
		Title:       event1.Title,
		Description: event1.Description,
		EndTime:     event1.EndTime,
		Rewards:     event1.Rewards,
	})
	if err != nil {
		t.Errorf("client.UpdateEvent %v", err)
	}

	//----------------------------------------- list all event ------------------------------- //
	resList, err = client.ListEvents(context.Background(), &pb.ListEventsRequest{})
	if err != nil {
		t.Errorf("client.ListEvents %v", err)
	}

	if !slices.EqualFunc(resList.LiveEvents, []*pb.LiveEvent{&event1, &event2}, compare) {
		t.Errorf("client.ListEvents events doesn't match")
	}

	//----------------------------------------- delete first event ------------------------------- //
	_, err = client.DeleteEvent(context.Background(), &pb.DeleteEventRequest{Id: 1})
	if err != nil {
		t.Errorf("client.ListEvents %v", err)
	}

	//----------------------------------------- list all event ------------------------------- //
	resList, err = client.ListEvents(context.Background(), &pb.ListEventsRequest{})
	if err != nil {
		t.Errorf("client.ListEvents %v", err)
	}

	if len(resList.LiveEvents) != 1 || !compare(resList.LiveEvents[0], &event2) {
		t.Errorf("client.ListEvents events doesn't match")
	}
}

func compare(a, b *pb.LiveEvent) bool {
	return a.Title == b.Title &&
		a.Description == b.Description &&
		a.StartTime == b.StartTime &&
		a.EndTime == b.EndTime &&
		a.Rewards == b.Rewards
}
