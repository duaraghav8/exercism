package react

type cell struct {
	r                 *reactor
	value             int
	onRefresh         func()
	onUpdateCallbacks []*callback
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) setValue(v int) {
	if c.value == v {
		return
	}
	c.value = v
	c.r.refresh()
}

type inputCell struct {
	*cell
}

func (ic *inputCell) SetValue(v int) {
	ic.setValue(v)
}

type callback struct {
	f       func(int)
	enabled bool
}

func (c *callback) Cancel() {
	c.enabled = false
}

func (c *callback) exec(i int) {
	if c.enabled {
		c.f(i)
	}
}

type computeCell struct {
	*cell
}

func (c *computeCell) AddCallback(f func(int)) Canceler {
	cb := &callback{f: f, enabled: true}
	c.onUpdateCallbacks = append(c.onUpdateCallbacks, cb)
	return cb
}

type reactor struct {
	cells []*cell
}

func (r *reactor) CreateInput(i int) InputCell {
	c := &cell{
		r:                 r,
		value:             i,
		onRefresh:         func() {},
		onUpdateCallbacks: make([]*callback, 0, 0),
	}

	r.trackCell(c)
	return &inputCell{c}
}

func (r *reactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	newCell := &cell{
		r:                 r,
		onUpdateCallbacks: make([]*callback, 0, 0),
	}
	newCell.onRefresh = func() {
		newCell.value = f(c.Value())

		for _, cb := range newCell.onUpdateCallbacks {
			cb.exec(newCell.value)
		}
	}
	newCell.onRefresh()

	r.trackCell(newCell)
	return &computeCell{newCell}
}

func (r *reactor) CreateCompute2(c1 Cell, c2 Cell, f func(int, int) int) ComputeCell {
	newCell := &cell{r: r}
	newCell.onRefresh = func() {
		newCell.value = f(c1.Value(), c2.Value())

		for _, cb := range newCell.onUpdateCallbacks {
			cb.exec(newCell.value)
		}
	}
	newCell.onRefresh()

	r.trackCell(newCell)
	return &computeCell{newCell}
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

// todo
// - onRefresh() functions currently compute the value everytime. Call them only if cell's value actually changes
// - refactor: move repeated code into a common place
