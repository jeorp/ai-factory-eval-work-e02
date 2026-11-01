package fixture
import ("errors"; "sync")
type Ledger struct { Balances map[string]int; Seen map[string]bool; mu sync.Mutex }
func (l *Ledger) Transfer(id, from, to string, amount int) error {
		if id == "" || from == "" || to == "" || from == to || amount <= 0 {
			return errors.New("invalid transfer")
		}
		l.mu.Lock()
		defer l.mu.Unlock()
		if l.Seen != nil && l.Seen[id] {
			return errors.New("duplicate transfer")
		}
		if l.Balances == nil || l.Balances[from] < amount {
			return errors.New("insufficient balance")
		}
		// Ensure maps are initialized
		if l.Seen == nil {
			l.Seen = make(map[string]bool)
		}
		l.Seen[id] = true
		l.Balances[from] -= amount
		l.Balances[to] += amount
		return nil
	}
