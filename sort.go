package dotenv_generator

func (x Environments) Len() int {
	return len(x)
}

func (x Environments) Less(i, j int) bool {
	return x[i].Name < x[j].Name
}

func (x Environments) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
