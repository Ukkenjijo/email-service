package service

import (
	"context"
	"log"
	"email-service/proto"
)

type EmailService struct {
	proto.UnimplementedEmailServiceServer
}

func (s *EmailService) SendEmail(ctx context.Context, req *proto.EmailRequest) (*proto.EmailResponse, error) {
	// Simulate email sending (Use SMTP in production)
	log.Printf("Sending email to %s: Subject: %s, Body: %s\n", req.To, req.Subject, req.Body)
	return &proto.EmailResponse{Message: "Email sent successfully"}, nil
}
