# Myth-promo 

## Architecture

The service is designed with an approach close to a DDD architecture. The application layer hasn't been created for simplicity.

*Data package* contains a simple wrapper for Mongodb, that could be updated to allow to choose other databases; a models subpackage with the structures of the documents in the database, and a repository with the queries used in the shop.

A mock repository has been included for testing purposes. The approach is quite simple: when a product or promos for a given product are requested, it will return a model matching the required data.

*Shop package* will include the business logic. It will limit the response to 5 products at most, as required. Nevertheless, it might be more efficient to pass this limitation to the data layer. We keep it in this package because the current specification is that the order doesn't matter, so we can pick any 5 products. But, in the future, this specification might change to require, let's say, the 5 products with the final lower price. Since we apply the discounts in this layer, we won't be able to query the database for the right products in advance.

Not included in the requirements, but the selection of the best promotion based in the highest percentage has been developed as a strategy, so it would be easy to have and apply different strategies such as cumulative discounts, etc.

*Api package* includes the request/response structs. The response doesn't differ from the struct returned by the shop, but it is included anyways and used to show the percentage as a text ("10%") instead of an integer.

To create the Http Api REST, the Gin web framework has been used.

## Compile and run

A Makefile has been included, some useful commands are:

`make build` compiles and creates a binary in the bin/ directory
`make tests` run the tests
`make run` will compile and execute the service

Please, note that the application will try to open a configuration file `config.toml` from the working directory (the one which the binary is called from). A `config.sample` is included.

To initialize the `products` and `promos` collections in the database, two JSON files has been included in `src/data/sources` directory.

## Improvements / TO-DO

-Capability to choose between different database drivers. Probably the Database interface should be updated, and the repositories are dependant on MongoDb. A query builder for each database would allow to implement the methods in the repository, but it exceeds the purpose of the project.

-Tests could be improved to cover more cases, and for a better code coverage. They are more a sample than a exhaustive test suite. Depending on the use case, some benchmarks might be helpful as well.

-Probably the packages internal structure could be better. 

-Better logs, error messages, error control...