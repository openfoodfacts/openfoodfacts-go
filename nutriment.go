package openfoodfacts

type Nutriment struct {
	Sodium100G           string `json:"sodium_100g"`
	Sugars               string `json:"sugars"`
	Fat                  string `json:"fat"`
	Carbohydrates        string `json:"carbohydrates"`
	Fat100G              string `json:"fat_100g"`
	NutritionScoreFr100G int    `json:"nutrition-score-fr_100g"`
	Salt100G             string `json:"salt_100g"`
	Proteins100G         string `json:"proteins_100g"`
	CarbohydratesServing string `json:"carbohydrates_serving"`
	ProteinsUnit         string `json:"proteins_unit"`
	FiberServing         string `json:"fiber_serving"`
	SaturatedFat         string `json:"saturated-fat"`
	Proteins             string `json:"proteins"`
	NutritionScoreFr     int    `json:"nutrition-score-fr"`
	SaturatedFatUnit     string `json:"saturated-fat_unit"`
	Carbohydrates100G    string `json:"carbohydrates_100g"`
	FatUnit              string `json:"fat_unit"`
	SaturatedFatServing  string `json:"saturated-fat_serving"`
	CarbohydratesUnit    string `json:"carbohydrates_unit"`
	FiberUnit            string `json:"fiber_unit"`
	EnergyServing        string `json:"energy_serving"`
	Energy               string `json:"energy"`
	Sugars100G           string `json:"sugars_100g"`
	Salt                 string `json:"salt"`
	Sodium               string `json:"sodium"`
	Fiber                string `json:"fiber"`
	SodiumServing        string `json:"sodium_serving"`
	ProteinsServing      string `json:"proteins_serving"`
	Fiber100G            string `json:"fiber_100g"`
	SaturatedFat100G     string `json:"saturated-fat_100g"`
	Energy100G           string `json:"energy_100g"`
	SaltServing          string `json:"salt_serving"`
	SugarsUnit           string `json:"sugars_unit"`
	NutritionScoreUk     int    `json:"nutrition-score-uk"`
	SodiumUnit           string `json:"sodium_unit"`
	FatServing           string `json:"fat_serving"`
	SugarsServing        string `json:"sugars_serving"`
	NutritionScoreUk100G int    `json:"nutrition-score-uk_100g"`
	EnergyUnit           string `json:"energy_unit"`
}
