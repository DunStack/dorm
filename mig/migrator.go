package mig

import (
	"fmt"
	"os"
	"time"
)

type Migrator struct {
}

func (m *Migrator) GenerateSQLMigration(name string) error {
	mn := m.generateMigrationName(name)

	for _, t := range []string{"up", "down"} {
		f, err := os.Create(fmt.Sprintf("%s.%s.sql", mn, t))
		if err != nil {
			return err
		}
		defer f.Close()
	}

	return nil
}

func (m *Migrator) generateMigrationName(name string) string {
	return time.Now().Format("20060102150405") + "_" + name
}
