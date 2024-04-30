# build
This directory contains all `CICD` related files.

## SQL
Using `PostgreSQL` as the primary database.

To support large scale data processing, we made some important decisions:
- Disable `soft delete` for all tables.
- Prefer to use `UUID` as primary key for all tables when the BIGINT type does not satisfy the requirements.
- Utilize reference tables to efficiently store extensive varchar data whenever feasible.
- Utilize sharding or partitioning to separate large datasets across tables, mitigating performance concerns.
- No foreign key constraints are used to avoid performance bottlenecks, manage data integrity at the application level.
