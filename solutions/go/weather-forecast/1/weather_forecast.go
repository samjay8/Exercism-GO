//Package weather is a package that provides tools to forecast weather accurately.
package weather

var (
    //CurrentCondition provides the current condition of the weather in the country.
	CurrentCondition string
    //CurrentLocation provides the current location at which the weather is being forecasted.
	CurrentLocation  string 
)
//Forecast function helps to forecast the weather condition of a particular city in the country.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
