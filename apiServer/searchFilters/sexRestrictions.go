package searchFilters

import (
	"MatchaServer/common"
)

func PrepareSexRestrictions(user common.User) string {
	// gay - только свой пол
	// натурал - только противоположный пол
	// би - без ограничений
	// нет инфы - без ограничений
	if user.Gender == "" {
		return ""
	}
	if user.Orientation == "bi" || user.Orientation == "" {
		if user.Gender == "male" {
			return "((gender='male' AND orientation='homo') OR (gender='female' AND orientation='hetero')" +
				" OR orientation='bi' OR gender='' OR orientation='')"
		} else {
			return "((gender='female' AND orientation='homo') OR (gender='male' AND orientation='hetero')" +
				" OR orientation='bi' OR gender='' OR orientation='')"
		}
	}
	if user.Orientation == "homo" {
		return "gender='" + user.Gender + "' AND (orientation='homo' OR orientation='bi' OR orientation='')"
	}
	// Остались только гетеросексуалы
	if user.Gender == "male" {
		return "gender='female' AND (orientation='hetero' OR orientation='bi' OR orientation='')"
	}
	if user.Gender == "female" {
		return "gender='male' AND (orientation='hetero' OR orientation='bi' OR orientation='')"
	}
	return ""
}
