# PAPCPT
📘 Course Syllabus: Programming Algorithms using Parallel and Concurrent Programming Techniques (Go Edition)

Course Description:

This crash course is an intensive, hands-on introduction to concurrent and parallel programming using the Go programming language. Students will explore key paradigms such as goroutines, channels, mutexes, condition variables, wait groups, and worker pools. The course culminates in advanced applications of these tools to solve complex optimization problems using metaheuristics.

Prerequisites:

- Basic knowledge of programming in Go
- Understanding of algorithms and data structures

Course Duration: 8 Weeks

Learning Outcomes:

- Understand Go's concurrency model and how it differs from traditional threads
- Build robust concurrent programs using goroutines and channels
- Develop parallel algorithms using worker pools, fan-in/fan-out patterns, and synchronization primitives
- Apply parallelism to real-world problems including simulation, sorting, graph algorithms, and metaheuristics

🗂 Week-by-Week Breakdown

Week 1: Introduction to Concurrency and Go’s Memory Model

Topics: Concurrency vs. Parallelism, goroutines, channels

Readings: Go Blog on concurrency patterns

Lab 1: Concurrent Web Crawler using Goroutines and Channels

Lab Goal: Demonstrate fan-out and fan-in concurrency patterns

Week 2: Synchronization Primitives and Race Conditions

Topics: Mutex, RWMutex, WaitGroup, race detector

Lab 2: Safe Bank Transactions using Mutexes and Condition Variables

Lab Goal: Prevent race conditions in a shared resource context

Week 3: Pipeline Patterns and Worker Pools

Topics: Channel pipelines, load balancing with worker pools

Lab 3: Concurrent Log Analyzer with a Scalable Worker Pool

Lab Goal: Analyze and process logs concurrently using load balancing

Week 4: Parallel Algorithms I – Sorting and Searching

Topics: Parallel Quicksort, concurrent binary search

Lab 4: Parallel Quicksort with Goroutines and Channels

Lab Goal: Break array into partitions and sort concurrently

Week 5: Parallel Algorithms II – Graph Traversal and Simulation

Topics: Parallel BFS/DFS, cellular automata

Lab 5: Concurrent Maze Solver with Shared State Synchronization

Lab Goal: Use goroutines to simulate agents traversing a maze

Week 6: Metaheuristics I – Tabu Search

Topics: Local search, neighborhood moves, tabu list, thread coordination

Lab 6: Multithreaded Tabu Search for the Traveling Salesman Problem (TSP)

Lab Goal: Launch concurrent neighborhood explorers with shared tabu memory

Week 7: Metaheuristics II – Genetic Algorithms

Topics: Population evolution, crossover, mutation, selection

Lab 7: Parallel Genetic Algorithm for Scheduling Problem

Lab Goal: Run fitness evaluation and selection in parallel across generations

Week 8: Capstone Project and Performance Tuning

Topics: Benchmarking, profiling (pprof), reducing contention

Final Project: Choose between optimizing a scientific simulation or solving an NP-hard problem using metaheuristics in parallel
