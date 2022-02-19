package main

import (
	"context"
	"log"
	"time"

	pb "mailinglist/proto"

	"github.com/alexflint/go-arg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var args struct {
	GrpcAddr string `arg:"env:MAILINGLIST_GRPC_ADDR"`
}

func createEmail(client pb.MailingListServiceClient, addr string) {
	log.Println("create email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreateEmail(ctx, &pb.CreateEmailRequest{EmailAddr: addr})
	if err != nil {
		log.Fatalf("  error: %v", err)
	}
	if res.EmailEntry == nil {
		log.Printf("  email not found")
	} else {
		log.Printf("  response: %v", res.EmailEntry)
	}
}

func getEmail(client pb.MailingListServiceClient) {
	log.Println("get email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetEmail(ctx, &pb.GetEmailRequest{EmailAddr: "example@example.com"})
	if err != nil {
		log.Fatalf("  error: %v", err)
	}
	if res.EmailEntry == nil {
		log.Printf("  email not found")
	} else {
		log.Printf("  response: %s", res.EmailEntry)
	}
}

func getEmailBatch(client pb.MailingListServiceClient) {
	log.Println("get email batch")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetEmailBatch(ctx, &pb.GetEmailBatchRequest{Count: 10, Page: 1})
	if err != nil {
		log.Fatalf("  error: %v", err)
	}
	log.Println("response:")
	for i := 0; i < len(res.EmailEntries); i++ {
		log.Printf("  item [%v of %v]: %s", i+1, len(res.EmailEntries), res.EmailEntries[i])
	}
}

func updateEmail(client pb.MailingListServiceClient) {
	log.Println("update email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	entry := pb.EmailEntry{
		Id:          1,
		Email:       "example@example.com",
		ConfirmedAt: 1,
		OptOut:      false,
	}

	res, err := client.UpdateEmail(ctx, &pb.UpdateEmailRequest{EmailEntry: &entry})
	if err != nil {
		log.Fatalf("  error: %v", err)
	}
	if res.EmailEntry == nil {
		log.Printf("  email not found")
	} else {
		log.Printf("  response: %v", res.EmailEntry)
	}
}

func deleteEmail(client pb.MailingListServiceClient) {
	log.Println("delete email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.DeleteEmail(ctx, &pb.DeleteEmailRequest{EmailAddr: "example@example.com"})
	if err != nil {
		log.Fatalf("  error: %v", err)
	}
	if res.EmailEntry == nil {
		log.Printf("  email not found")
	} else {
		log.Printf("  response: %v", res.EmailEntry)
	}
}

func main() {
	arg.MustParse(&args)

	if args.GrpcAddr == "" {
		args.GrpcAddr = ":8081"
	}

	conn, err := grpc.Dial(args.GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMailingListServiceClient(conn)

	createEmail(client, "sample1@example.com")
	getEmail(client)
	updateEmail(client)
	getEmailBatch(client)
	deleteEmail(client)
}
