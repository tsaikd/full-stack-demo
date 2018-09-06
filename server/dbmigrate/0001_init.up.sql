CREATE EXTENSION "uuid-ossp";

CREATE TABLE coinHistory (
	coinHistoryID UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v1mc(),
	symbol TEXT NOT NULL,
	timestamp TIMESTAMP NOT NULL,
	price DOUBLE PRECISION NOT NULL
);

CREATE INDEX coinHistory_symbol ON coinHistory (symbol);
