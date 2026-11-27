package routes

func StatusC(ver string) int {
	switch ver {
	case "v1":
		return 200
	case "v2":
		return 501
	default:
		return 404
	}
}
func StatusC2(ver string) int {
	switch ver {
	case "v1":
		return 200
	case "v2":
		return 501
	default:
		return 404
	}
}
