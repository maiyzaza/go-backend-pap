package enum

type NumberMinMax int

const (
	RentDayMin        NumberMinMax = 1
	RentDayMax        NumberMinMax = 31
	RentSmallMonthMin NumberMinMax = 2
	RentSmallMonthMax NumberMinMax = 5
	RentBigMonthMin   NumberMinMax = 6
	RentBigMonthMax   NumberMinMax = 11
	RentYear          NumberMinMax = 1
	RentShowMin       NumberMinMax = 1
	RentShowMax       NumberMinMax = 1000000
)
