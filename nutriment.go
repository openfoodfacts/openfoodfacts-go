// Copyright © 2016 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

import "encoding/json"

type Nutriment struct {
	Carbohydrates        json.Number `json:"carbohydrates"`
	Carbohydrates100G    json.Number `json:"carbohydrates_100g"`
	CarbohydratesServing json.Number `json:"carbohydrates_serving"`
	CarbohydratesUnit    string      `json:"carbohydrates_unit"`
	Energy               json.Number `json:"energy"`
	Energy100G           json.Number `json:"energy_100g"`
	EnergyServing        json.Number `json:"energy_serving"`
	EnergyUnit           string      `json:"energy_unit"`
	Fat                  json.Number `json:"fat"`
	Fat100G              json.Number `json:"fat_100g"`
	FatServing           json.Number `json:"fat_serving"`
	FatUnit              string      `json:"fat_unit"`
	Fiber                json.Number `json:"fiber"`
	Fiber100G            json.Number `json:"fiber_100g"`
	FiberServing         json.Number `json:"fiber_serving"`
	FiberUnit            string      `json:"fiber_unit"`
	Proteins             json.Number `json:"proteins"`
	Proteins100G         json.Number `json:"proteins_100g"`
	ProteinsServing      json.Number `json:"proteins_serving"`
	ProteinsUnit         string      `json:"proteins_unit"`
	Salt                 json.Number `json:"salt"`
	Salt100G             json.Number `json:"salt_100g"`
	SaltServing          json.Number `json:"salt_serving"`
	SaturatedFat         json.Number `json:"saturated-fat"`
	SaturatedFat100G     json.Number `json:"saturated-fat_100g"`
	SaturatedFatServing  json.Number `json:"saturated-fat_serving"`
	SaturatedFatUnit     string      `json:"saturated-fat_unit"`
	Sodium               json.Number `json:"sodium"`
	Sodium100G           json.Number `json:"sodium_100g"`
	SodiumServing        json.Number `json:"sodium_serving"`
	SodiumUnit           string      `json:"sodium_unit"`
	Sugars               json.Number `json:"sugars"`
	Sugars100G           json.Number `json:"sugars_100g"`
	SugarsServing        json.Number `json:"sugars_serving"`
	SugarsUnit           string      `json:"sugars_unit"`
	NutritionScoreFr     json.Number `json:"nutrition-score-fr"`
	NutritionScoreFr100G json.Number `json:"nutrition-score-fr_100g"`
	NutritionScoreUk     json.Number `json:"nutrition-score-uk"`
	NutritionScoreUk100G json.Number `json:"nutrition-score-uk_100g"`
}
