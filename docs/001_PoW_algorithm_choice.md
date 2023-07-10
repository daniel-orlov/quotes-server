# Introduction:

One of the requirements for the implementation of the server was to use Proof of Work (PoW) algorithms to protect
against DDoS attacks. The server should be able to generate a PoW challenge and validate the solution provided by the
client.

After initial research of exisitng PoW algorithms, the focus has been placed on two popular, battle-tested ones:
Hashcash and Scrypt. These algorithms have been chosen for their effectiveness in protecting against DDoS attacks and
their suitability for
the server's requirements. Let's delve into a comparison of these algorithms based on several key parameters.

## Comparison of Hashcash and Scrypt:

| Parameters                                | Hashcash                                                  | Scrypt                                                     |
|-------------------------------------------|-----------------------------------------------------------|------------------------------------------------------------|
| Difficulty of Implementation              | Relatively straightforward to implement                   | Moderate complexity in implementation                      |
| Memory Requirements                       | Minimal memory requirements                               | Requires a significant amount of memory                    |
| Computational Complexity to Solve         | High computational complexity                             | Moderate computational complexity                          |
| Resistance to ASIC-based Mining           | Less resistant to ASIC-based mining                       | Designed to be resistant to ASIC-based mining              |
| Security Against DDoS Attacks             | Effective protection against DDoS attacks                 | Offers good protection against DDoS attacks                |
| Availability of Libraries/Implementations | Abundance of libraries and implementations available      | Multiple libraries and implementations available           |
| Adoption and Community Support            | Widely adopted and well-supported                         | Relatively lower adoption compared to Hashcash             |
| Customization and Tuning Options          | Limited customization and tuning options                  | Offers more customization and tuning options               |
| Performance Considerations                | Faster computation, but less memory-hardness              | Slower computation, but more memory-hardness               |
| Energy Efficiency                         | Less energy-efficient due to higher computational demands | More energy-efficient due to memory-dependent computations |

## Conclusion and Final Choice:

Both Hashcash and Scrypt have their strengths and considerations to take into account. Hashcash is relatively easier to
implement and offers effective protection against DDoS attacks. On the other hand, Scrypt provides better resistance to
ASIC-based mining and offers more customization options. It is worth noting that Hashcash has wider adoption and better
community support, making it more readily accessible.

Considering the requirements of the "Word of Wisdom" server, I decided to choose Hashcash as the PoW algorithm
for implementation. Its straightforward implementation, effective DDoS protection, and availability of libraries and
community support make it a suitable choice. While Scrypt offers certain advantages, the specific needs and constraints
of this project lean towards Hashcash as the optimal solution.

## References:
- [Hashcash](http://hashcash.org/papers/hashcash.pdf)
- [Scrypt](https://www.tarsnap.com/scrypt/scrypt.pdf)