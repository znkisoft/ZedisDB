# ZedisDB (WIP)

> a key-value memory database written in Go

## features

- [x] compatible with `RESP` (REdis Serialization Protocol), you can use `redis-cli` to connect the database.
- [ ] LRU Cache with time based eviction

## functions

### support data types

- Strings and Binary Data
- Numbers
- NULL
- Arrays (which may be nested)
- Dictionaries (which may be nested)
- Error messages

### support commands

[commands](config/commands.yml)

## reference

- [gomemcache](https://github.com/bradfitz/gomemcache/blob/master/memcache/memcache.go)
- [godis](https://github.com/HDT3213/godis)
- [redis](https://redis.io)