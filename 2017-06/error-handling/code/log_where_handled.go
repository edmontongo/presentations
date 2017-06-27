

logger.Debugf("entered function, about do DoSomething")

if err := t.DoSomething(); err != nil {
	logger.Errorf("failed to do something: %s", err) // redundant // HL
	return errors.Errorf("failed to do something: %s", err)
}
