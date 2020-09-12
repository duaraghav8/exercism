package react

type cell struct {
	r         *reactor
	value     int
	onRefresh func()
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) setValue(v int) {
	c.value = v
	c.r.refresh()
}

type inputCell struct {
	*cell
}

func (ic *inputCell) SetValue(v int) {
	ic.setValue(v)
}

type computeCell struct {
	*cell
}

func (c *computeCell) AddCallback(cb func(int)) Canceler {
	return nil
}

type reactor struct {
	cells []*cell
}

func (r *reactor) CreateInput(i int) InputCell {
	c := &cell{
		r:         r,
		value:     i,
		onRefresh: func() {},
	}

	r.trackCell(c)
	return &inputCell{c}
}

func (r *reactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	cc := &cell{r: r}
	cc.onRefresh = func() {
		cc.value = f(c.Value())
	}
	cc.onRefresh()

	r.trackCell(cc)
	return &computeCell{cc}
}

func (r *reactor) CreateCompute2(c1 Cell, c2 Cell, f func(int, int) int) ComputeCell {
	cc := &cell{r: r}
	cc.onRefresh = func() {
		cc.value = f(c1.Value(), c2.Value())
	}
	cc.onRefresh()

	r.trackCell(cc)
	return &computeCell{cc}
}

func (r *reactor) trackCell(c *cell) {
	r.cells = append(r.cells, c)
}

// refresh forces all cells to update themselves if their dependency
// cells were updated.
func (r *reactor) refresh() {
	for _, c := range r.cells {
		c.onRefresh()
	}
}

func New() Reactor {
	return &reactor{cells: make([]*cell, 0, 0)}
}
