Write SQL as you normally would, then [sqlc](https://github.com/kyleconroy/sqlc) generates type-safe Go code from it.

Then use [pgx](https://github.com/jackc/pgx) to handle the PostgreSQL database connection.

It seems to be widely recommended against using an ORM. The existing ORM solutions require a lot of effort to maintain and arent as performant as writing just plain SQL.

Using sqlc seems to hit the middle ground of writing plain SQL but still getting Go code generated for you so you dont have to write it all yourself.

Then pgx integrates with PostgreSQL much more than the stdlib sql package.
