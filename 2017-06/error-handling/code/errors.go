// Errorf tracks cause of error and records stack frame information // HL
if err := os.RemoveAll("/var/syntropy/dvol/Win2008.PG"); err != nil {
	return errors.Errorf("failed to remove .dvol files: %s", err) // HL
}
