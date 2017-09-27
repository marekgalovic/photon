DROP KEYSPACE IF EXISTS photon_development;
CREATE KEYSPACE photon_development WITH replication={'class': 'SimpleStrategy', 'replication_factor': 1};
DROP KEYSPACE IF EXISTS photon_test;
CREATE KEYSPACE photon_test WITH replication={'class': 'SimpleStrategy', 'replication_factor': 1};
