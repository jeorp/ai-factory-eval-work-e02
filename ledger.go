package fixture
import ("errors"; "sync")
type Ledger struct { Balances map[string]int; Seen map[string]bool; mu sync.Mutex }
func (l *Ledger) Transfer(id, from, to string, amount int) error {
 l.mu.Lock(); defer l.mu.Unlock()
 if id=="" || from=="" || to=="" || from==to || amount<=0 { return errors.New("invalid transfer") }
 if l.Seen != nil && l.Seen[id] { return errors.New("duplicate transfer") }
 var fromBal int
 if l.Balances != nil { fromBal = l.Balances[from] }
 if fromBal < amount { return errors.New("insufficient balance") }
 if l.Balances == nil { l.Balances = map[string]int{} }
 l.Balances[from] -= amount
 l.Balances[to] += amount
 if l.Seen == nil { l.Seen = map[string]bool{} }
 l.Seen[id] = true
 return nil
}
