package constructor

func MakeLecture(g *Group, c *Cabinet, t *Teacher, s *Subject) *Lecture {
	l := &Lecture{
		Cabinet: c,
		Teacher: t,
		Group:   g,
		Subject: s,
	}
	return l
}
