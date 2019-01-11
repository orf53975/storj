// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"
	"time"

	"storj.io/storj/pkg/uplagreement"
	dbx "storj.io/storj/satellite/satellitedb/dbx"
)

type uplinkagreement struct {
	db *dbx.DB
}

func (b *uplinkagreement) CreateAgreement(ctx context.Context, serialNum string, agreement uplagreement.Agreement) error {
	_, err := b.db.Create_Uplinkagreement(
		ctx,
		dbx.Uplinkagreement_Signature(agreement.Signature),
		dbx.Uplinkagreement_Serialnum(serialNum),
		dbx.Uplinkagreement_Data(agreement.Agreement),
	)
	return err
}

func (b *uplinkagreement) GetAgreements(ctx context.Context) ([]uplagreement.Agreement, error) {
	rows, err := b.db.All_Uplinkagreement(ctx)
	if err != nil {
		return nil, err
	}

	agreements := make([]uplagreement.Agreement, len(rows))
	for i, entry := range rows {
		agreement := &agreements[i]
		agreement.Signature = entry.Signature
		agreement.Agreement = entry.Data
		agreement.CreatedAt = entry.CreatedAt
	}
	return agreements, nil
}

func (b *uplinkagreement) GetSignature(ctx context.Context, serialnum string) (*uplagreement.Agreement, error) {
	dbxInfo, err := b.db.Get_Uplinkagreement_By_Serialnum(ctx, dbx.Uplinkagreement_Serialnum(serialnum))
	if err != nil {
		return &uplagreement.Agreement{}, err
	}

	return &uplagreement.Agreement{
		Agreement: dbxInfo.Data,      // Uplink ID
		Signature: dbxInfo.Signature, // Uplink Public Key
	}, nil
}

func (b *uplinkagreement) GetAgreementsSince(ctx context.Context, since time.Time) ([]uplagreement.Agreement, error) {
	rows, err := b.db.All_Uplinkagreement_By_CreatedAt_Greater(ctx, dbx.Uplinkagreement_CreatedAt(since))
	if err != nil {
		return nil, err
	}

	agreements := make([]uplagreement.Agreement, len(rows))
	for i, entry := range rows {
		agreement := &agreements[i]
		agreement.Signature = entry.Signature
		agreement.Agreement = entry.Data
		agreement.CreatedAt = entry.CreatedAt
	}
	return agreements, nil
}
