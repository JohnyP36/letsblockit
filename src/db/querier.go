// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type Querier interface {
	CountInstanceForUserAndFilter(ctx context.Context, arg CountInstanceForUserAndFilterParams) (int64, error)
	CreateInstanceForUserAndFilter(ctx context.Context, arg CreateInstanceForUserAndFilterParams) error
	CreateListForUser(ctx context.Context, userID uuid.UUID) (uuid.UUID, error)
	DeleteInstanceForUserAndFilter(ctx context.Context, arg DeleteInstanceForUserAndFilterParams) error
	GetActiveFiltersForUser(ctx context.Context, userID uuid.UUID) ([]string, error)
	GetInstanceForUserAndFilter(ctx context.Context, arg GetInstanceForUserAndFilterParams) (pgtype.JSONB, error)
	GetInstancesForList(ctx context.Context, filterListID int32) ([]GetInstancesForListRow, error)
	GetListForToken(ctx context.Context, token uuid.UUID) (GetListForTokenRow, error)
	GetListForUser(ctx context.Context, userID uuid.UUID) (GetListForUserRow, error)
	GetStats(ctx context.Context) (GetStatsRow, error)
	HasUserDownloadedList(ctx context.Context, userID uuid.UUID) (bool, error)
	MarkListDownloaded(ctx context.Context, id int32) error
	UpdateInstanceForUserAndFilter(ctx context.Context, arg UpdateInstanceForUserAndFilterParams) error
}

var _ Querier = (*Queries)(nil)
