package main

type MOError struct {
	Code    int
	FormatS string
	FormatV []string
}

type IMO interface {
	PreAdd(oldMO IMO, newMO IMO) (ret int, msg string)
	PreMod()
	PreRmv()
	PostAdd()
	PostMod()
	PostRmv(old IMO)
	GetRelated() []string
	Build() IMO
}

type BaseMO struct {
	IMO
	attr map[string]interface{}
}

func main() {
	return
}
