# golang-sync-pool-example


"Pool

Pool is a scalable pool of temporary objects that is also concurrency safe. Any stored value in the pool can be deleted at any time without prior notification. In addition, the object pool can be dynamically expanded when the load is high, and when it is not in use or concurrency is not high, the object pool is shrunk.

The gist is to reuse objects to avoid the repeated creation and destruction that degrades performance.

But for what reason do we actually use it?
The intended purpose of the pool is to cache allocated but unused objects for later reuse, thus reducing the load on the garbage collector. This means it makes it easy to create efficient, thread-safe free lists. However, it is not suitable for all free lists.

The appropriate use of a pool is to manage a group of temporary items that are implicitly shared and possibly reused by concurrent independent clients of a package. Pool provides a way to spread the cost of allocation overhead across many clients.

It is important to note that Pool also introduces performance costs. Using sync.Pool is significantly slower than simply initialising it. In addition, a pool shall not be copied after the first use.

The sync.Pool, as you notice, is purely a temporary object pool, suitable for storing some temporary objects that are shared by different goroutines.
"

