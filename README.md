DBSample
========
DBSample creates database dumps which _make sense_.

**ALPHA RELEASE! USE WITH CAUTION!**

Dealing with large development/staging databases is unacceptable in an age of cloud services and containers, where personal dev platforms can be spun up in _seconds_ and then discarded. Containers have made it simple for each member of the dev team to work on their own copy of an application, but often the team is sharing a single dev/staging database at a remote location. Those databases often contain old data, and altering them is impossible without effecting everyone else on the team.

Random data generators are a common solution to the problem of creating small testable databases, but they generate data that is usually a poor representation of the real application data, and the generator itself is difficult to create and becomes another piece of software to be maintained. DBSample solves the problem by creating a snapshot of your real database with a small _sample_ of the _real_ data.

Currently supports MySQL 5+. Other drivers and versions may be supported in the future.

## Install
Linux and Windows builds are available on the [releases page](https://github.com/headzoo/dbsample/releases).

## Usage
```
usage: dbsample [<flags>] <database>

Flags:
      --help                 Show context-sensitive help (also try --help-long and --help-man).
      --version              Show application version.
  -h, --host="127.0.0.1"     The database host.
  -P, --port="3306"          The database port.
      --protocol="tcp"       The protocol to use for the connection (tcp, socket, pip, memory).
  -u, --user=USER            User for login if not current user.
  -p, --password=PASSWORD    Password to use when connecting to server. If password is not given it's asked from stderr.
      --routines             Dump procedures and functions.
      --triggers             Dump triggers.
  -l, --limit=100            Max number of rows from each table to dump.
  -n, --no-create-database   Disable adding CREATE DATABASE statement.
      --skip-lock-tables     Disable locking tables on read.
      --skip-add-drop-table  Disable adding DROP TABLE statements.
      --extended-insert      Use multiple-row INSERT syntax that include several VALUES lists.
      --rename-database=DUMP-NAME  
                             Use this database name in the dump.
  -c, --constraint=CONSTRAINT ...  
                             Assigns one or more foreign key constraints.
  -f, --filter=FILTER ...    Apply a filter to the output.

Args:
  <database>  Name of the database to dump.

Filters:
Filters alter column values in the dump. For example they can remove passwords or
other sensitive information. Each --filter flag should be passed the name of the
filter, e.g. "empty", the name of a table.column, e.g. "users.passwords", and one
or more arguments.

  --filter="empty table.column"
  --filter="repeat table.column <string>"

Examples:
dbsample --limit=100 blog > dump.sql
dbsample --limit=100 -h db1 -u admin -p blog > dump.sql
dbsample --limit=100 --rename-database=blog_dev blog > dump.sql
dbsample --limit=100 --filter="empty users.password" --filter="repeat users.email X" blog > dump.sql
dbsample --limit=100 -c "posts.user_id users.id" -c "posts.cat_id categories.id" blog > dump.sql
```
