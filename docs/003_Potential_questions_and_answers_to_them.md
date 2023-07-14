# Potential question and answers to them

Since this project is a solution for a technical challenge that is supposed to be presented without any additional
explanations, I decided to write down some potential questions and answers to them.

The questions are listed in no particular order.

## Why did you choose this particular PoW algorithm?

Short answer: hashcash has a proven track record as an anti-spam mechanism, it is a good fit for the task, and it is
simple to implement.
Long answer: please see [docs/001_PoW_algorithm_choice.md](001_PoW_algorithm_choice.md).

## Can you explain how your solution is helping in preventing DDOS attacks?

Sure!
I based my solution on the combination of two mechanisms: PoW and rate limiting.
Both are common and proven mechanisms for preventing DDOS attacks, and they complement each other nicely.
Frankly, rate limiting might have been enough on its own, but since PoW is a requirement, it just made the server more
secure.

## Why have you used self-written mocks instead of generated ones?

While I have positive experience with generated mocks (like mockgen and mockery),
in recent years I have been transitioning towards self-written mocks for the following reasons:

- Generated mocks are not very readable and require additional time to understand them.
- Generated mocks have required me to pass controller, write all the expectations for them to know what to return,
  which polluted my tests with unnecessary code and made them less readable.
- Self-written mocks are more readable and require less code to write, even if we take into account the fact that
  they need to be maintained, and they need tests for themselves.

All things considered, I think that self-written mocks are more readable and maintainable, and I prefer them over the
generated ones.

## Why have you not used Redis/Memcached/Postgres for storing quotes/challenges?

Instead of using a database, I decided to use in-memory storage for quotes and challenges, but due to the fact that I
have used interfaces for storage, it is possible to easily switch to any other storage implementation, including Redis,
Memcached, Postgres, etc. This allowed me to focus on the core functionality of the project and not waste time on the
database, which wouldn't be a big challenge to add anyway.

## Do you really write that many comments in your code, or is it just for the sake of this project?

The latter.
I think that comments should be used sparingly and only when they add value,
but the requirements for this project state that "code should be fully covered in comments", so I did it.
I guess that the idea behind this requirement is to see how I write comments and how I structure my code,
as well as be able to see my reasoning and logic, because I will not be able to explain it in person in a demo.

## Why is this project written using Go 1.19, while the latest version is 1.20 at the moment?

Staying one version behind the latest Go release for a project offers benefits such as stability, compatibility with
third-party dependencies, ecosystem support, mature tooling, and reduced risk of encountering undiscovered bugs or
compatibility issues.
Even though, it may result in missing out on the latest language features and performance improvements,
I think that in most cases,
unless you really need a new feature or there's a critical bug fix or vulnerability,
staying one version behind is a good tactic.

## Why have you used X package/library?

Let's go over the list of dependencies and I will explain why I have chosen them.
github.com/JGLTechnologies/gin-rate-limit v1.5.4: Provides rate limiting middleware for Gin framework, allowing you to
control and limit the number of requests to your application.

- gin-goinc/gin: I like that it's fast and flexible, with a clean and intuitive API design.
  It also is a very popular framework, which means that it has a big community and a lot of third-party libraries.
- JGLTechnologies/gin-rate-limit: speaking of third-party libraries, I have chosen this one because it is the most
  popular
  rate limiting middleware for Gin framework, and it is actively maintained.
- kelseyhightower/envconfig:
  Simplifies the process
  of reading configuration data from environment variables by automatically mapping them to Go structs (and has
  defaults).
  Lovely, simple, and easy to use.
- oklog/ulid/v2: Implements Universally Unique Lexicographically Sortable Identifiers (ULID),
  which are highly efficient, and URL-safe alternatives to traditional UUIDs.
  Would be a good fit for most identifiers, but I have used it only for quote IDs.
- stretchr/testify: A popular and comprehensive testing toolkit for Go, providing powerful assertion functions and
  utilities to simplify writing tests and improve test coverage.
- ybbus/httpretry: Offers an easy way to perform retry logic for HTTP requests,
  allowing you to handle transient failures and improve the reliability of your applications.
  It is also very easy to use, and it complies with http.Client interface, so it can be used with almost any HTTP
  client.
- go.uber.org/zap: A fast, structured, and highly efficient logging library for Go, designed for high-performance
  applications with minimal memory allocations and low overhead. It is also very easy to use and has a lot of features.

## What would you do if you had more time?

My list of things to do would look like this (grouped by priority):

It might be that I have already added that by the time you are reading this:

+ Add more integration tests and client tests (I have already added a few, but I would like to add more)
+ Finish infrastructure setup (in another repo, since it is not a part of this
  project - [repo link](https://github.com/daniel-orlov/quotes-infra)

Will be adding in the future:

+ Improve error responses from server (currently they are not very informative and ad hoc defined, I would like client
  to be able to rely on the same error structure for all types of errors and for server to be able to manipulate and
  centrally control error type definition)
+ Add telemetry for more observability (currently there are logs, which are good, but not enough)
+ Add timeout middleware to server (to prevent hanging connections) — [this one](https://github.com/gin-contrib/timeout)
  seems nice
+ Make pipeline steps only run on relevant changes (currently all steps run on every change, which is not very
  efficient)
+ Improve CI/CD pipeline by adding continuous deployment to Cloud Run.

## My question is not listed here, what should I do?

Don't hesitate to ask me directly, I will be happy to answer any questions you might have.
The best way to reach me is via telegram: @danielorlov.
