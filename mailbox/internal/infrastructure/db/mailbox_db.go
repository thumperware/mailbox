package db

import (
	"context"

	"github.com/Hoomanfr/golib/database"
	"github.com/Hoomanfr/messaging/mailbox/internal/domain"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MailboxDB interface {
	Create(ctx context.Context, mailbox domain.Mailbox) error
	FindByUserId(ctx context.Context, userId string) ([]domain.Mailbox, error)
}

type MailboxDb struct {
	pgDb database.PgDB
}

func NewMailboxDb(pgDb database.PgDB) MailboxDB {
	return MailboxDb{
		pgDb: pgDb,
	}
}

func (db MailboxDb) Create(ctx context.Context, mailbox domain.Mailbox) error {
	err := db.pgDb.WithTransaction(ctx, func(tx pgx.Tx) error {
		sql, args, err := sb.Insert("mailboxes").
			Columns("id", "user_id", "email", "created_at").
			Values(mailbox.ID, mailbox.UserID, mailbox.Email, mailbox.CreatedAt).
			ToSql()
		if err != nil {
			return err
		}
		_, err = tx.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (db MailboxDb) FindByUserId(ctx context.Context, userId string) ([]domain.Mailbox, error) {
	var mailboxes []domain.Mailbox
	err := db.pgDb.WithConnection(ctx, func(c *pgxpool.Conn) error {
		sql, args, err := sb.Select("id", "user_id", "email", "created_at").
			From("mailboxes").
			Where(squirrel.Eq{"user_id": userId}).
			ToSql()
		if err != nil {
			return err
		}
		row, err := c.Query(ctx, sql, args...)
		if err != nil {
			return err
		}
		for row.Next() {
			var mailbox domain.Mailbox
			err = row.Scan(&mailbox.ID, &mailbox.UserID, &mailbox.Email, &mailbox.CreatedAt)
			if err != nil {
				return err
			}
			mailboxes = append(mailboxes, mailbox)
		}

		return nil
	})
	return mailboxes, err
}
