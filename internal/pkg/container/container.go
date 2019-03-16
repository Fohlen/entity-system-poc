package container

type Container struct {
	a bool
	b *int
	c *string
}

func (c Container) Switch(fa func(bool), fb func(int), fc func(string)) {
	switch {
	case c.c != nil:
		fc(*c.c)
	case c.b != nil:
		fb(*c.b)
	default:
		fa(c.a)
	}
}

func A(a bool) Container {
	return Container{a: a}
}

func B(b int)Container {
	return Container{b: &b}
}

func C(c string)Container  {
	return Container{c: &c}
}
