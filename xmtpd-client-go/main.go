package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	envUtils "github.com/xmtp/xmtpd/pkg/envelopes"
	"github.com/xmtp/xmtpd/pkg/proto/xmtpv4/envelopes"
	message_api "github.com/xmtp/xmtpd/pkg/proto/xmtpv4/message_api"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:5050", "xmtpd gRPC server address")
	flag.Parse()

	creds := credentials.NewClientTLSFromCert(nil, "")
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to connect to %s: %v", *addr, err)
	}
	defer conn.Close()

	client := message_api.NewReplicationApiClient(conn)

	req := &message_api.SubscribeEnvelopesRequest{
		Query: &message_api.EnvelopesQuery{
			OriginatorNodeIds: []uint32{100, 200, 0, 1},
			LastSeen: &envelopes.Cursor{
				NodeIdToSequenceId: map[uint32]uint64{
					100: 0,
					200: 0,
				},
			},
		},
	}

	stream, err := client.SubscribeEnvelopes(context.Background(), req)
	if err != nil {
		log.Fatalf("SubscribeEnvelopes error: %v", err)
	}

	fmt.Printf("✅ Subscribed, waiting for messages...\n")

	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("recv error: %v", err)
		}

		if len(resp.Envelopes) == 0 {
			continue
		}

		for _, envProto := range resp.Envelopes {
			env, err := envUtils.NewOriginatorEnvelope(envProto)
			if err != nil {
				log.Printf("Failed to unmarshal originator envelope")
			}

			fmt.Printf("%+v\n", env)
		}
	}
}
