package migrate

import (
	"github.com/go-pg/pg/v10/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			CREATE TABLE users (
				id               SERIAL PRIMARY KEY,
				public_key       BYTEA,
				hide_everywhere  BOOL,
				hide_leaderboard BOOL,
				email            TEXT,
				email_verified   BOOL,
				phone_number     TEXT,
				phone_country    TEXT,
				phone_verified   BOOL,
				whitelist_posts  BOOL,
				verified         BOOL,
				graylisted       BOOL,
				blacklisted      BOOL
			)
		`)
		if err != nil {
			return err
		}

		_, err = db.Exec(`
			CREATE TABLE blocks (
				id                 SERIAL PRIMARY KEY,
                public_key         BYTEA,
				blocked_public_key BYTEA,
				blocked_at         BIGINT,

				PRIMARY KEY (public_key, blocked_public_key)
			)
		`)
		if err != nil {
			return err
		}

		_, err = db.Exec(`
			CREATE TABLE audit_logs (
				id     SERIAL PRIMARY KEY,
                actor  BYTEA,
				action TEXT,
				user   BYTEA,
				when   BIGINT
			)
		`)
		if err != nil {
			return err
		}

		_, err = db.Exec(`
			CREATE TABLE global_feed_posts (
				id        SERIAL PRIMARY KEY,
                post_hash BYTEA,
				pinned    BOOL,
				added_at  BIGINT,
				added_by  BYTEA
			)
		`)
		if err != nil {
			return err
		}

		_, err = db.Exec(`
			CREATE TABLE wyre_orders (
				wyre_order_id     TEXT PRIMARY KEY,
                public_key        BYTEA,
				last_payload      JSONB,
				last_wallet_order JSONB,
				bitclout_nanos    BIGINT,
				transfer_txn_hash BYTEA,
				processed         BOOL
			)
		`)
		if err != nil {
			return err
		}

		_, err = db.Exec(`
			CREATE TABLE node_params (
				id                         SERIAL PRIMARY KEY,
				bit_clout_reserve_price    BIGINT,
				bit_clout_fee_basis_points BIGINT,
				set_at                     BIGINT
			)
		`)
		if err != nil {
			return err
		}

		return nil
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			DROP TABLE users;
			DROP TABLE blocks;
			DROP TABLE audit_logs;
			DROP TABLE global_feed_posts;
			DROP TABLE wyre_orders;
			DROP TABLE node_params;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20210623152412_create_tables", up, down, opts)
}
