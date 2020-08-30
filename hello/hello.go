package hello

// User user type
type User struct {
	ID   int64
	Name string
	Addr *Address
}

// Address address type
type Address struct {
	City   string
	ZIP    int
	LatLng [2]float64
}
var lat  [2]float64
var alex = User{2, "Pogo", &Address{ "Philadelphia", 19080, lat }}

// Bozo writes a welcome string
func Bozo() string {
	return "Yo... go-sample " + alex.Name + alex.Addr.City
}
