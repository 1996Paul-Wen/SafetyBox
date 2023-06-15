package model

func LoadModels() []interface{} {
	models := []interface{}{
		new(User),
		new(SafetyData),
	}
	return models
}
