# URL Shortening Service

## Overview

Welcome to the URL Shortening Service! This system is designed to generate short and unique aliases for long URLs, making them easier to share and manage. 

## Table of Contents



### Functional Requirements:

- Generate a shorter and unique alias for a given URL.
- Allow users to customize short links.
- Redirect users to the original link when accessing a short link.
- Set expiration times for short links.

### Non-Functional Requirements:

- High availability to ensure uninterrupted URL redirection.
- Real-time URL redirection with minimal latency.
- Shortened links should be unpredictable.
- Provide analytics on link redirections.
- Expose services through REST APIs.

### Extended Requirements:

- Analytics tracking for redirection statistics.

### Database Schema:

Two tables:
1. URL Mappings
2. User Data

### Database Choice: MongoDB


### Encoding Actual URL:

- Compute a unique hash (e.g., MD5 or SHA256) of the URL.
- Encode the hash using base64.
- Choose a fixed length for the short key (e.g., 6 characters).

### Generating Keys Offline:

- Use a Key Generation Service (KGS) to pre-generate random keys.
- Store generated keys in a key-DB.
- Ensure uniqueness and handle concurrency.

### Hash-Based Partitioning:

- Use a hash of the short link to determine the partition.
- Implement 'Consistent Hashing' to handle overloaded partitions.

### Cache

- Using a caching solution like Memcached.
- Cache frequently accessed URLs.
- Implement Least Recently Used (LRU) eviction policy.

### Load Balancer

- Implement load balancing at three levels.
- Initially use simple Round Robin.
- Consider server load for more intelligent load balancing.
