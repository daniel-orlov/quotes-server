# Requirements

## Design and implement “Word of Wisdom” http server.

- [x] HTTP server should be protected from DDOS attacks with the Prof of
  Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.

- [x] The choice of the POW algorithm should be explained.

- [x] After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other
  collection of the quotes.

- [x] Docker file should be provided both for the server and for the client that solves the POW Challenge.

- [x] Code should be covered with unit tests.

- [x] Code should be fully covered in comments (methods, variables, tests, etc.).

- [x] Code should be published on GitHub.

- [x] Readme file should be provided with instructions on how to run the server and the client.

### Additional requirements (as derrived from the answers to the questions)

- [x] PoW algorithm should be balance between computational complexity and memory requirements.

- [x] No constraints on the choice of the storage for the quotes.

- [x] Additional layers of protection against DDoS attacks could be added, such as rate limiting or IP blocking.

- [x] Performance requirements for the server are not defined.

- [x] All documentation and comments are supposed to be in English.

- [x] Golang is the tool of choice.

