package main

import "time"

type record struct {
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Data      string    `json:"data"`
	Timestamp time.Time `json:"time"`
}

type database interface {
	getAddr(name string) ([]*record, error)
	getName(addr string) ([]*record, error)
	put(name string, data string, recType string, ts time.Time) error
}

func newDatabase() database {
	return &memDatabase{
		addrRepo: map[string][]*record{},
		nameRepo: map[string][]*record{},
	}
}

type memDatabase struct {
	addrRepo map[string][]*record
	nameRepo map[string][]*record
}

func (x *memDatabase) getName(addr string) ([]*record, error) {
	// TODO: Implement read lock mechanism
	records, ok := x.addrRepo[addr]
	if !ok {
		return nil, nil
	}
	return records, nil
}

func (x *memDatabase) getAddr(name string) ([]*record, error) {
	// TODO: Implement read lock mechanism
	records, ok := x.nameRepo[name]
	if !ok {
		return nil, nil
	}
	return records, nil
}

func (x *memDatabase) put(name string, data string, recType string, ts time.Time) error {
	nameRecord := record{
		Name:      name,
		Data:      data,
		Type:      recType,
		Timestamp: ts,
	}
	addrRecord := nameRecord // clone

	logger.Debugw("Put record", "record", nameRecord)

	// TODO: Implement write lock mechanism
	x.nameRepo[name] = append(x.nameRepo[name], &nameRecord)
	x.addrRepo[data] = append(x.addrRepo[data], &addrRecord)

	return nil
}
