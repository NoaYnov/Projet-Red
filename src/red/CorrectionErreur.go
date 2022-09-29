package red

func (p *Personnage) IsAlpha(sentence string) bool { // fonction qui vérifie si la chaine de caractère est alphabétique
	result := false
	if len(sentence) == 0 {
		return false
	}
	for i := 0; i < len(sentence); i++ {
		if sentence[i] >= 65 && sentence[i] <= 90 {
			result = true
			continue
		}
		if sentence[i] >= 97 && sentence[i] <= 122 {
			result = true
			continue
		}
		return false
	}
	return result
}

func (p *Personnage) IsNumeric(sentence string) bool { // fonction qui vérifie si la chaine de caractère est numérique

	for i := 0; i < len(sentence); i++ {
		if sentence[i] > 47 && sentence[i] < 58 {
			return true
		}
	}
	return false
}
