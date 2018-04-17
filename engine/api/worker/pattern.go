package worker

import (
	"fmt"

	"github.com/go-gorp/gorp"

	"github.com/ovh/cds/sdk"
)

const patternColumns = `
	worker_model_pattern.id,
	worker_model_pattern.name,
	worker_model_pattern.type`

// InsertWorkerModelPattern insert a new worker model in database
func InsertWorkerModelPattern(db gorp.SqlExecutor, modelPattern *sdk.ModelPattern) error {
	dbmodelPattern := WorkerModelPattern(*modelPattern)
	if err := db.Insert(&dbmodelPattern); err != nil {
		return err
	}
	*modelPattern = sdk.ModelPattern(dbmodelPattern)
	return nil
}

// UpdateWorkerModelPattern insert a new worker model in database
func UpdateWorkerModelPattern(db gorp.SqlExecutor, modelPattern *sdk.ModelPattern) error {
	dbmodelPattern := WorkerModelPattern(*modelPattern)
	if _, err := db.Update(&dbmodelPattern); err != nil {
		return err
	}
	return nil
}

// DeleteWorkerModelPatter removes from database worker model pattern
func DeleteWorkerModelPattern(db gorp.SqlExecutor, ID int64) error {
	wmp := WorkerModelPattern(sdk.ModelPattern{ID: ID})
	count, err := db.Delete(&wmp)
	if err != nil {
		return err
	}
	if count == 0 {
		return sdk.ErrNotFound
	}
	return nil
}

// LoadWorkerModelPatterns retrieves model patterns from database
func LoadWorkerModelPatterns(db gorp.SqlExecutor) ([]sdk.ModelPattern, error) {
	var wmPatterns []WorkerModelPattern
	query := fmt.Sprintf(`SELECT %s from WORKER_MODEL_PATTERN ORDER BY name`, patternColumns)
	if _, err := db.Select(&wmPatterns, query); err != nil {
		return nil, sdk.WrapError(err, "LoadWorkerModelPatterns> ")
	}

	workerModelPatterns := make([]sdk.ModelPattern, len(wmPatterns))
	for i := range wmPatterns {
		if err := wmPatterns[i].PostSelect(db); err != nil {
			return nil, err
		}
		workerModelPatterns[i] = sdk.ModelPattern(wmPatterns[i])
	}
	return workerModelPatterns, nil
}

// LoadWorkerModelPatternByName retrieves model patterns from database given its name and type
func LoadWorkerModelPatternByName(db gorp.SqlExecutor, patternType, name string) (*sdk.ModelPattern, error) {
	var wmp WorkerModelPattern
	query := fmt.Sprintf(`SELECT %s FROM worker_model_pattern WHERE name = $1 AND type = $2`, patternColumns)
	if err := db.SelectOne(&wmp, query, name, patternType); err != nil {
		return nil, sdk.WrapError(err, "LoadWorkerModelPatternByName> ")
	}

	if err := wmp.PostSelect(db); err != nil {
		return nil, err
	}
	workerModelPattern := sdk.ModelPattern(wmp)

	return &workerModelPattern, nil
}
