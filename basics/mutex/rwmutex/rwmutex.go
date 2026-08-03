package main

// RWMutex allows multiple readers but only one writer.
// Read karne wale bahut ho sakte hain. Lekin write ek hi karega.

/*
Scenario

Tumhara ek dictionary hai.

Har koi us dictionary ko read kar sakta hai. (Multiple Readers Allowed)

Magar jab koi update karega (Write), tab sabko rukna padega.

Visual:

Read → Reader 1

Read → Reader 2

Read → Reader 3

↓

(Sab saath mein read kar sakte hain)

Update → Writer 1

↓

(Sabko rukna padega)

Update → Writer 2

↓

(Sabko rukna padega)

Lock vs RWMutex
Mutex:

1 Reader → OK

100 Readers → Ek-ek karke

1 Writer → OK

RWMutex:

1 Reader → OK

100 Readers → Sab saath mein OK

1 Writer → OK

*/

// Methods

// mu.RLock()

// Read Lock

// mu.RUnlock()

// Kab Use Kare?

// Agar

// 95% Read

// 5% Write

// Ho.

// Database cache

// Configuration

// Product Catalog

// etc.