DROP TABLE IF EXISTS vsp_tick, vsp, exchange_tick, exchange, pow_data;

CREATE TABLE IF NOT EXISTS exchange (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL
);


CREATE TABLE IF NOT EXISTS exchange_tick (
    id SERIAL PRIMARY KEY,
    exchange_id INT REFERENCES exchange(id) NOT NULL, 
	interval INT NOT NULL,
	high FLOAT NOT NULL,
	low FLOAT NOT NULL,
	open FLOAT NOT NULL,
	close FLOAT NOT NULL,
	volume FLOAT NOT NULL,
	currency_pair TEXT NOT NULL,
	time TIMESTAMPTZ NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS exchange_tick_idx ON exchange_tick (exchange_id, interval, currency_pair, time);

CREATE TABLE IF NOT EXISTS vsp (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	api_enabled BOOLEAN NOT NULL,
	api_versions_supported INT[] NOT NULL,
	network TEXT NOT NULL,
	url TEXT NOT NULL,
	launched TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS vsp_tick (
	id SERIAL PRIMARY KEY,
	vsp_id INT REFERENCES vsp(id) NOT NULL,
	immature INT NOT NULL,
	live INT NOT NULL,
	voted INT NOT NULL,
	missed INT NOT NULL,
	pool_fees FLOAT NOT NULL,
	proportion_live FLOAT NOT NULL,
	proportion_missed FLOAT NOT NULL,
	user_count INT NOT NULL,
	users_active INT NOT NULL,
	time TIMESTAMPTZ NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS vsp_tick_idx ON vsp_tick (vsp_id,immature,live,voted,missed,pool_fees,proportion_live,proportion_missed,user_count,users_active, time);

-- CREATE TABLE IF NOT EXISTS vsp_tick_time (
-- 	id SERIAL PRIMARY KEY,
-- 	vsp_tick_id INT REFERENCES vsp_tick(id) NOT NULL,
-- 	update_time TIMESTAMPTZ NOT NULL
-- );

-- CREATE UNIQUE INDEX IF NOT EXISTS vsp_tick_time_idx ON vsp_tick_time (vsp_tick_id, update_time);

CREATE TABLE IF NOT EXISTS pow_data (
	time INT,
	network_hashrate VARCHAR(25),
	pool_hashrate VARCHAR(25),
	workers INT,
	network_difficulty FLOAT8,
	coin_price VARCHAR(25),
	btc_price VARCHAR(25),
	source VARCHAR(25),
	PRIMARY KEY (time, source)
);
