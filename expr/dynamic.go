package expr

// Dynamic can be used to return an expression based on the address family
// the expression is used in.
type Dynamic struct {
	Expr func(fam uint8) Any
}

func (d *Dynamic) marshal(fam byte) ([]byte, error) {
	return d.Expr(fam).marshal(fam)
}

func (d *Dynamic) marshalData(fam byte) ([]byte, error) {
	return d.Expr(fam).marshalData(fam)
}

func (d *Dynamic) unmarshal(fam byte, data []byte) error {
	return d.Expr(fam).unmarshal(fam, data)
}
