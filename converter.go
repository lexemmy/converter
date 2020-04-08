package main
import(
	"fmt"
	"math"
) 

type converter struct{}

type feet float64
type centimeter float64
type seconds float64
type minutes float64
type celcius float64
type farenheit float64
type degree float64
type radian float64
type kilogram float64
type pound float64


func(cvr converter) centimetertofeet(c centimeter) feet {
	var result = feet(c / 30.48)
	return result
}

func(cvr converter) feettocentimeter(f feet) centimeter {
	var result = centimeter(f * 30.48)
	return result
}

func(cvr converter) minutestoseconds(m minutes) seconds {
	var result = seconds(m * 60)
	return result
}
func(cvr converter) secondstominutes(s seconds) minutes {
	var result = minutes(s / 60)
	return result
}
func(cvr converter) celciustofarenheit(c celcius) farenheit {
	var result = farenheit(c*9/5 + 32)
	return result
}
func(cvr converter) farenheittocelcius(f farenheit) celcius {
	var result = celcius((f - 32) * 5 / 9)
	return result
}
func(cvr converter) radiantodegree(r radian) degree {
	var result = degree(r * 180/ math.Pi)
	return result
}
func(cvr converter) degreetoradian(d degree) radian {
	var result = radian(d * math.Pi / 180)
	return result
}
func(cvr converter) kilogramtopound(kg kilogram) pound {
	var result = pound(kg * 2.205)
	return result
}
func(cvr converter) poundtokilogram(lb pound) kilogram {
	var result = kilogram(lb / 2.205)
	return result
}


func main() {

	cvr := converter{}
	fmt.Println(cvr.centimetertofeet(13.50))
	fmt.Println(cvr.feettocentimeter(0.44))
	fmt.Println(cvr.secondstominutes(60))
	fmt.Println(cvr.minutestoseconds(5))
	fmt.Println(cvr.farenheittocelcius(212))
	fmt.Println(cvr.celciustofarenheit(100))
	fmt.Println(cvr.radiantodegree(2))
	fmt.Println(cvr.degreetoradian(360))
	fmt.Println(cvr.kilogramtopound(5))
	fmt.Println(cvr.poundtokilogram(2))
}