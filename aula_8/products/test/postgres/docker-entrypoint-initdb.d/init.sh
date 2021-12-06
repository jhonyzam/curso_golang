export PGPASSWORD=admin

# Database DDL
psql -U postgres -d products -a -f /docker-entrypoint-initdb.d/ddl/init.sql
psql -U postgres -d products -a -f /docker-entrypoint-initdb.d/ddl/products.sql

# Test Database DDL
psql -U postgres -d products_test -a -f /docker-entrypoint-initdb.d/ddl/products_test.sql
