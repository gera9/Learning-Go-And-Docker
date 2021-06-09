# Exercise 1

The `docker-compose.yml` file runs 4 containers:

* The first container runs a `MongoDB` server and it's named `mongo_compose`.

* The second container runs `Mongo Express` client and it's named `mexpress_compose`. This is connected to `MongoDB` by a network. This is also protected with user and password (only Mongo Express, not MongoDB) by using environment variables.

    * User: `DASistemas`
    * Password: `ex-especial567`

* The third container is a GO program that inserts the data from the `people.xml` file into the `MongoDB` database.

* The fourth container raise a server on the port `7777` by using the `GO Chi` library as router.

    * This has two endpoints:
        
        * The endpoint `/people` return all the records of the database as `JSON`.
        * The endpoint `/people/{id}` return one person (by id) as `JSON`.