# Deckard

[![Maintainability](https://api.codeclimate.com/v1/badges/3c85f35a6a7865558b27/maintainability)](https://codeclimate.com/github/bradcypert/deckard/maintainability)

### A framework agnostic tool for running database migrations.
###### Currently, Deckard only supports Postgres. More databases coming soon!

#### Usage
```bash
deckard create add_login_date_to_users
# modify the created files
deckard up --host=localhost --port=5432 --user=user --password=pass --database=app
deckard down --host=localhost --port=5432 --user=user --password=pass --database=app
```

#### TODO LIST:
- [x] Up Migrations for Postgres
- [x] Down Migrations for Postgres
- [x] Verify integrity for Postgres
- [x] Create new migrations from Deckard
- [x] Allow reading from Config file instead of cmd flags
- [ ] Support for MySQL

#### Managing your databases via YAML config.
Deckard also supports managing your databases via YAML.
Instead of writing
```bash
deckard up --host=localhost --port=5432 --user=user --password=pass --database=app
```

You can simply write
```bash
deckard up --key=prod
```
provided you have a a `.deckard.yml` in your home directory. In this instance, your YAML should look like this:
```yaml
prod:
    host: localhost
    port: 5432
    user: user
    password: pass
    database: app
```

Alternatively, you can provide deckard the path to the configuration value you want to use.
```bash
deckard up --config=/usr/app/deckard.yml --dbKey=prod
```

#### Verifying a migration was ran against the database.
Sometimes, you may find yourself curious if the migration was ran against the DB. You certainly can fire up your favorite database client and query for the metadata entry (or the schema change!), but Deckard also allows you to verify that a given migration has been ran against a given database. Simply use `deckard verify ~/path/to/my/migration.up.sql` and deckard will verify that that migration has been ran. A word of warning: We simply check to ensure the metadata table contains a matching entry for the migration. Basically, deckard is only verifying that the migration has been applied in the "UP" position. It won't verify that the schema is currently matching the changes that were introduced in that migration.

#### The Ups and Downs of Deckard
Before we ever run an `up` or a `down` migration, we have to verify that a metadata table exists and create it if it does not.

Deckard uses this metadata to keep track of what has and has not been ran via Deckard. Deckard works best when all schema changes are ran via migrations, and can not infer changes made outside of the Deckard toolchain.

When we run an up migration, we validate our database state against our metadata to figure out which migrations need to be ran. The same goes with down migrations as well. Finally, the up and down migrations additionally create or remove a row in the metadata table. If, for some reason, Deckard was to blow up after running a query but before successfully modifying the metadata table, you'll be able to get deckard back on the right path with a 1 row update. I say this as a caveat, but I've not yet ran into this issue.

#### Issues on Deckard
I use Github issues to track small improvements and work for Deckard. Additionally, issues are a great place for questions to be asked. Please check there if you have questions, concerns or are running into issues with the tool.
