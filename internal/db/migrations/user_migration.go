package migrations

import (
	"fmt"
	"dot-portal-go/internal/db"
)

// MigrateUp runs all CREATE TABLE statements
func MigrateUp() error {
	stmts := []string{
		// Dependency Tables
		`CREATE TABLE IF NOT EXISTS subscriptions (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255),
			price FLOAT,
			duration_days INTEGER,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS payment_gateways (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			provider VARCHAR(255),
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS requests (
			id SERIAL PRIMARY KEY,
			user_id INTEGER,
			action VARCHAR(255),
			status VARCHAR(50),
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS ref_country_states (
			id SERIAL PRIMARY KEY,
			country_code VARCHAR(10),
			state_code VARCHAR(10),
			name VARCHAR(255),
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		// Main Tables
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			firstname VARCHAR(255),
			lastname VARCHAR(255),
			fullname VARCHAR(255),
			email VARCHAR(255) UNIQUE NOT NULL,
			phone VARCHAR(50),
			birthday DATE,
			email_verified_at TIMESTAMP,
			password VARCHAR(255) NOT NULL,
			is_active BOOLEAN NOT NULL DEFAULT TRUE,
			remember_token VARCHAR(100),
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS user_meta (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id),
			key VARCHAR(255) NOT NULL,
			value TEXT,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS password_reset_tokens (
			email VARCHAR(255) PRIMARY KEY,
			token VARCHAR(255) NOT NULL,
			created_at TIMESTAMP
		);`,

		`CREATE TABLE IF NOT EXISTS sessions (
			id VARCHAR(255) PRIMARY KEY,
			user_id INTEGER REFERENCES users(id),
			ip_address VARCHAR(45),
			user_agent TEXT,
			payload TEXT,
			last_activity INTEGER
		);`,

		`CREATE TABLE IF NOT EXISTS roles (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			slug VARCHAR(255) UNIQUE NOT NULL,
			description TEXT,
			is_default BOOLEAN NOT NULL DEFAULT FALSE,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS user_roles (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id),
			role_id INTEGER NOT NULL REFERENCES roles(id)
		);`,

		`CREATE TABLE IF NOT EXISTS user_company (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id),
			name VARCHAR(255),
			phone VARCHAR(50),
			dot_number VARCHAR(50),
			mc_number VARCHAR(50),
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS user_company_address (
			id SERIAL PRIMARY KEY,
			item_id INTEGER NOT NULL REFERENCES user_company(id),
			type VARCHAR(50),
			address1 VARCHAR(255),
			address2 VARCHAR(255),
			city VARCHAR(255),
			state_id INTEGER REFERENCES ref_country_states(id),
			zip VARCHAR(20),
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS user_address (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id),
			address1 VARCHAR(255),
			address2 VARCHAR(255),
			city VARCHAR(255),
			state_id INTEGER REFERENCES ref_country_states(id),
			zip VARCHAR(20),
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS user_payment_cards (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id),
			card_number VARCHAR(50),
			card_holder_name VARCHAR(255),
			expiry_date VARCHAR(10),
			payment_method_id INTEGER REFERENCES payment_gateways(id),
			"primary" BOOLEAN NOT NULL DEFAULT FALSE,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS user_subscription (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id),
			subscription_id INTEGER REFERENCES subscriptions(id),
			price FLOAT NOT NULL DEFAULT 0,
			discount FLOAT NOT NULL DEFAULT 0,
			payment_card_id INTEGER REFERENCES user_payment_cards(id),
			start_date TIMESTAMP,
			next_date TIMESTAMP,
			end_date TIMESTAMP,
			status VARCHAR(50),
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS user_subscription_meta (
			id SERIAL PRIMARY KEY,
			subscription_id INTEGER NOT NULL REFERENCES user_subscription(id),
			key VARCHAR(255) NOT NULL,
			value TEXT
		);`,

		`CREATE TABLE IF NOT EXISTS user_payment_card_meta (
			id SERIAL PRIMARY KEY,
			card_id INTEGER NOT NULL REFERENCES user_payment_cards(id),
			key VARCHAR(255) NOT NULL,
			value TEXT,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS user_payment_history (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id),
			payment_method_id INTEGER NOT NULL REFERENCES payment_gateways(id),
			subscription_id INTEGER REFERENCES user_subscription(id),
			type VARCHAR(50),
			amount FLOAT NOT NULL DEFAULT 0,
			payment_date TIMESTAMP,
			transaction_id VARCHAR(255),
			request_id INTEGER REFERENCES requests(id),
			status VARCHAR(50),
			notes TEXT,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS user_tasks (
			id SERIAL PRIMARY KEY,
			unique_code VARCHAR(255) UNIQUE NOT NULL,
			user_id INTEGER NOT NULL REFERENCES users(id),
			assigned_to INTEGER REFERENCES users(id),
			title VARCHAR(255),
			description TEXT,
			category VARCHAR(100),
			subcategory VARCHAR(100),
			status VARCHAR(50) NOT NULL DEFAULT 'pending',
			due_date TIMESTAMP,
			completed_at TIMESTAMP,
			priority VARCHAR(50) DEFAULT 'normal',
			link VARCHAR(255),
			entity VARCHAR(100),
			entity_id VARCHAR(100),
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);`,

		`CREATE TABLE IF NOT EXISTS user_task_meta (
			id SERIAL PRIMARY KEY,
			task_id INTEGER NOT NULL REFERENCES user_tasks(id),
			key VARCHAR(255) NOT NULL,
			value TEXT
		);`,
	}

	for _, s := range stmts {
		if _, err := db.CONN.Exec(s); err != nil {
			return fmt.Errorf("failed to exec statement: %w", err)
		}
	}
	return nil
}


// MigrateDown drops the tables in reverse dependency order
func MigrateDown() error {
	drops := []string{
		"DROP TABLE IF EXISTS user_task_meta;",
		"DROP TABLE IF EXISTS user_tasks;",
		"DROP TABLE IF EXISTS user_payment_history;",
		"DROP TABLE IF EXISTS user_payment_card_meta;",
		
		"DROP TABLE IF EXISTS user_subscription_meta;",
    "DROP TABLE IF EXISTS user_subscription;",
		"DROP TABLE IF EXISTS user_payment_cards;",
		"DROP TABLE IF EXISTS user_address;",
		"DROP TABLE IF EXISTS user_company_address;",
		"DROP TABLE IF EXISTS user_company;",
		"DROP TABLE IF EXISTS user_roles;",
		"DROP TABLE IF EXISTS roles;",
		"DROP TABLE IF EXISTS sessions;",
		"DROP TABLE IF EXISTS password_reset_tokens;",
		"DROP TABLE IF EXISTS user_meta;",
		"DROP TABLE IF EXISTS users;",
		"DROP TABLE IF EXISTS ref_country_states;",
		"DROP TABLE IF EXISTS requests;",
		"DROP TABLE IF EXISTS payment_gateways;",
		"DROP TABLE IF EXISTS subscriptions;",
	}

	for _, s := range drops {
		if _, err := db.CONN.Exec(s); err != nil {
			return fmt.Errorf("failed to exec drop: %w", err)
		}
	}
	return nil
}


