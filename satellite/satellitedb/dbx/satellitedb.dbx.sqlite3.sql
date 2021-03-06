-- AUTOGENERATED BY gopkg.in/spacemonkeygo/dbx.v1
-- DO NOT EDIT
CREATE TABLE accounting_raws (
	id INTEGER NOT NULL,
	node_id BLOB NOT NULL,
	interval_end_time TIMESTAMP NOT NULL,
	data_total REAL NOT NULL,
	data_type INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE accounting_rollups (
	id INTEGER NOT NULL,
	node_id BLOB NOT NULL,
	start_time TIMESTAMP NOT NULL,
	put_total INTEGER NOT NULL,
	get_total INTEGER NOT NULL,
	get_audit_total INTEGER NOT NULL,
	get_repair_total INTEGER NOT NULL,
	put_repair_total INTEGER NOT NULL,
	at_rest_total REAL NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE accounting_timestamps (
	name TEXT NOT NULL,
	value TIMESTAMP NOT NULL,
	PRIMARY KEY ( name )
);
CREATE TABLE bwagreements (
	signature BLOB NOT NULL,
	serialnum TEXT NOT NULL,
	data BLOB NOT NULL,
	created_at TIMESTAMP NOT NULL,
	expires_at TIMESTAMP NOT NULL,
	PRIMARY KEY ( signature ),
	UNIQUE ( serialnum )
);
CREATE TABLE injuredsegments (
	id INTEGER NOT NULL,
	info BLOB NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE irreparabledbs (
	segmentpath BLOB NOT NULL,
	segmentdetail BLOB NOT NULL,
	pieces_lost_count INTEGER NOT NULL,
	seg_damaged_unix_sec INTEGER NOT NULL,
	repair_attempt_count INTEGER NOT NULL,
	PRIMARY KEY ( segmentpath )
);
CREATE TABLE nodes (
	id BLOB NOT NULL,
	audit_success_count INTEGER NOT NULL,
	total_audit_count INTEGER NOT NULL,
	audit_success_ratio REAL NOT NULL,
	uptime_success_count INTEGER NOT NULL,
	total_uptime_count INTEGER NOT NULL,
	uptime_ratio REAL NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE overlay_cache_nodes (
	node_id BLOB NOT NULL,
	node_type INTEGER NOT NULL,
	address TEXT NOT NULL,
	protocol INTEGER NOT NULL,
	operator_email TEXT NOT NULL,
	operator_wallet TEXT NOT NULL,
	free_bandwidth INTEGER NOT NULL,
	free_disk INTEGER NOT NULL,
	latency_90 INTEGER NOT NULL,
	audit_success_ratio REAL NOT NULL,
	audit_uptime_ratio REAL NOT NULL,
	audit_count INTEGER NOT NULL,
	audit_success_count INTEGER NOT NULL,
	uptime_count INTEGER NOT NULL,
	uptime_success_count INTEGER NOT NULL,
	PRIMARY KEY ( node_id ),
	UNIQUE ( node_id )
);
CREATE TABLE projects (
	id BLOB NOT NULL,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	terms_accepted INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE users (
	id BLOB NOT NULL,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	email TEXT NOT NULL,
	password_hash BLOB NOT NULL,
	created_at TIMESTAMP NOT NULL,
	PRIMARY KEY ( id ),
	UNIQUE ( email )
);
CREATE TABLE api_keys (
	id BLOB NOT NULL,
	project_id BLOB NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
	key BLOB NOT NULL,
	name TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	PRIMARY KEY ( id ),
	UNIQUE ( key ),
	UNIQUE ( name, project_id )
);
CREATE TABLE bucket_infos (
	project_id BLOB NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
	name TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	PRIMARY KEY ( name )
);
CREATE TABLE project_members (
	member_id BLOB NOT NULL REFERENCES users( id ) ON DELETE CASCADE,
	project_id BLOB NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
	created_at TIMESTAMP NOT NULL,
	PRIMARY KEY ( member_id, project_id )
);
