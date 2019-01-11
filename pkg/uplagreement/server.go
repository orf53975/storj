// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package uplagreement

import (
	"context"
	"crypto"
	"time"

	"go.uber.org/zap"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
)

var (
	mon = monkit.Package()
)

// DB stores bandwidth agreements.
type DB interface {
	// CreateAgreement adds a new bandwidth agreement.
	CreateAgreement(context.Context, string, Agreement) error
	// GetAgreements gets all bandwidth agreements.
	GetAgreements(context.Context) ([]Agreement, error)
	// GetSignature gets the public key of uplink corresponding to serial number
	GetSignature(ctx context.Context, serialnum string) (*Agreement, error)
	// GetAgreementsSince gets all bandwidth agreements since specific time.
	GetAgreementsSince(context.Context, time.Time) ([]Agreement, error)
}

// Server is an implementation of the pb.BandwidthServer interface
type Server struct {
	db     DB
	pkey   crypto.PublicKey
	logger *zap.Logger
}

// Agreement is a struct that contains a uplinks agreement info
type Agreement struct {
	Agreement []byte // uplink id
	Signature []byte // uplink public key
	CreatedAt time.Time
}

// NewServer creates instance of Server
func NewServer(db DB, logger *zap.Logger, pkey crypto.PublicKey) *Server {
	return &Server{
		db:     db,
		logger: logger,
		pkey:   pkey,
	}
}

// UplinkAgreements receives and stores bandwidth agreements from storage nodes
func (s *Server) UplinkAgreements(ctx context.Context, serialNum string, PubKey []byte, uplinkid storj.NodeID) (reply *pb.AgreementsSummary, err error) {
	defer mon.Task()(&ctx)(&err)

	s.logger.Debug("Received Agreement...")

	reply = &pb.AgreementsSummary{
		Status: pb.AgreementsSummary_FAIL,
	}

	err = s.db.CreateAgreement(ctx, serialNum, Agreement{
		Signature: PubKey,
		Agreement: uplinkid.Bytes(),
	})

	if err != nil {
		return reply, UplinkAgreementError.New("SerialNumber already exist in the PayerBandwidthAllocation")
	}

	reply.Status = pb.AgreementsSummary_OK

	s.logger.Debug("Stored Agreement...")

	return reply, nil
}
