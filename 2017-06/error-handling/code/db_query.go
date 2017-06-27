rows, err := db.Query(`SELECT Name, Status, Created, Updated, 
                       Notes FROM Job where JobID = ?`, jobId)
if err != nil {
	return errors.Errorf("failed to query Job table: %s", err)
}
defer rows.Close()

for rows.Next() {
	// e.g. rows.Scan(&name, &status, &created, &updated, &notes)
}

if err = rows.Err(); err != nil { // HL
	return errors.Errorf("job query error: %s", err)
} // HL
