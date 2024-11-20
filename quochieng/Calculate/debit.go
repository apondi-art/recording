package Calculate

func (c *Status) Debit(amount int) {
	c.Balance += amount
}
