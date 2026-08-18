package store

import "sync"

type EventStore struct { mu sync.Mutex; events []string }
func NewEventStore() *EventStore { return &EventStore{} }
func (s *EventStore) Append(event string) { s.mu.Lock(); defer s.mu.Unlock(); s.events=append(s.events,event) }
func (s *EventStore) All() []string { s.mu.Lock(); defer s.mu.Unlock(); out:=make([]string,len(s.events));copy(out,s.events);return out }
func (s *EventStore) Count(event string) int { count:=0; for _,v:=range s.All(){if v==event{count++}};return count }
