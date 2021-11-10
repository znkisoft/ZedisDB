# ZedisDB

> a key-value memory database written in Go

## features

- [ ] compatible with `RESP` (REdis Serialization Protocol), you can use `redis-cli` to connect the database.

## functions

### support data types

- Strings and Binary Data
- Numbers
- NULL
- Arrays (which may be nested)
- Dictionaries (which may be nested)
- Error messages

### support commands

- [x] `PING`
- [x] `ECHO` <name>
- [ ] `GET` <key>
- [ ] `GET` <key>
- [ ] `SET` <key> <value>
- [ ] `DELETE` <key>
- [ ] `FLUSH`
- [ ] `MGET` <key1> ... <keyn>
- [ ] `MSET` <key1> <value1> ... <keyn> <valuen>

## available RESP commands

<!--#Anchor-->