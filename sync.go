package godict

type RangeIterator[K comparable, V any] func(key K, value V) bool

// Delete deletes the value for a key.
func (d *Dict[K, V]) Delete(key K) {
	d.m.Delete(key)
}

// Load returns the value stored in the map for a key, or default value of V if no
// value is present.
// The ok result indicates whether value was found in the map.
// In any case of failure, err contains the detailed error.
func (d *Dict[K, V]) Load(key K) (value V, ok bool, err error) {
	val, ok := d.m.Load(key)
	if !ok {
		return value, false, nil
	}

	value, ok = val.(V)
	if !ok {
		return value, true, ErrInvalidElement
	}

	return value, true, nil
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
// In any case of failure, err contains the detailed error.
func (d *Dict[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool, err error) {
	var (
		act any
		ok  bool
	)

	act, loaded = d.m.LoadOrStore(key, value)

	// stored value might change in runtime, so check for the result of casting to V.
	actual, ok = act.(V)
	if !ok {
		return actual, loaded, ErrInvalidElement
	}

	return
}

// Store sets the value for a key.
func (d *Dict[K, V]) Store(key K, value V) {
	d.m.Store(key, value)
}

// Range calls iter sequentially for each key and value present in the map.
// If iter returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently (including by f), Range may reflect any
// mapping for that key from any point during the Range call. Range does not
// block other methods on the receiver; even f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (d *Dict[K, V]) Range(iter RangeIterator[K, V]) {
	d.m.Range(func(key, value any) bool {
		return iter(key.(K), value.(V))
	})
}
