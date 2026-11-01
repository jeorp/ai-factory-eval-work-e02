package fixture
import ("errors"; "sync")
type Ledger struct { Balances map[string]int; Seen map[string]bool; mu sync.Mutex }
func (l *Ledger) Transfer(id, from, to string, amount int) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if id == "" || from == "" || to == "" || from == to || amount <= 0 {
		return errors.New("invalid transfer")
	}
	if l.Seen != nil && l.Seen[id] {
		return errors.New("duplicate id")
	}
	if l.Balances == nil || l.Balances[from] < amount {
		return errors.New("insufficient balance")
	}
	l.Balances[from] -= amount
	if l.Balances == nil {
		l.Balances = make(map[string]int)
	}
	l.Balances[to] += amount
	if l.Seen == nil {
		l.Seen = make(map[string]bool)
	}
	l.Seen[id] = true
	return nil
}
