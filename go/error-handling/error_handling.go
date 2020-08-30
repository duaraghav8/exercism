package erratum

func Use(ro ResourceOpener, input string) (err error) {
	resource, err := ro()
	for {
		if _, ok := err.(TransientError); !ok {
			break
		}
		resource, err = ro()
	}
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			if fe, ok := r.(FrobError); ok {
				resource.Defrob(fe.defrobTag)
			}
			err = r.(error)
		}
		resource.Close()
	}()

	resource.Frob(input)
	return
}