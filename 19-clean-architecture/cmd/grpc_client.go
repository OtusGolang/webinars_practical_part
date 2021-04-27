package cmd

import (
	"context"
	"log"
	"time"

	"github.com/OtusGolang/webinars_practical_part/19-clean-architecture/internal/adapters/grpc/api"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	server    string
	title     string
	text      string
	startTime string
	endTime   string
)

const tsLayout = "2006-01-02T15:04:05"

func parseTs(s string) (*timestamp.Timestamp, error) {
	t, err := time.Parse(tsLayout, s)
	if err != nil {
		return nil, err
	}
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

var GrpcClientCmd = &cobra.Command{
	Use:   "grpc_client",
	Short: "Run grpc client",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(server, grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		client := api.NewCalendarServiceClient(conn)
		st, err := parseTs(startTime)
		if err != nil {
			log.Fatal(err)
		}
		et, err := parseTs(endTime)
		if err != nil {
			log.Fatal(err)
		}
		req := &api.CreateEventRequest{
			Title:     title,
			Text:      text,
			StartTime: st,
			EndTime:   et,
		}
		resp, err := client.CreateEvent(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.GetError() != "" {
			log.Fatal(resp.GetError())
		} else {
			log.Println(resp.GetEvent().Id)
		}
	},
}

func init() {
	GrpcClientCmd.Flags().StringVar(&server, "server", "localhost:8080", "host:port to connect to")
	GrpcClientCmd.Flags().StringVar(&title, "title", "", "event title")
	GrpcClientCmd.Flags().StringVar(&text, "text", "", "event text")
	GrpcClientCmd.Flags().StringVar(&startTime, "start-time", "", "event start time, format: "+tsLayout)
	GrpcClientCmd.Flags().StringVar(&endTime, "end-time", "", "event end time, format: "+tsLayout)
}
