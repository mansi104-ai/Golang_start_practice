package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	// panic("NeedsLicense not implemented")
    var result bool
    if kind == "car" || kind == "truck" {
        result = true   
    }else{
        result = false
    }
    return result
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	// panic("ChooseVehicle not implemented")
    const res = "is clearly the better choice."
    var result string
    if option1 < option2{
        result =  option1 +" "+ res
    }else{
        result =  option2 + " "+res
    }

    return result
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	// panic("CalculateResellPrice not implemented")
    var resultCost float64
    const age1 float64= 3
    const age2 float64 = 10
    
    if age < age1 {
        resultCost = originalPrice * 0.8
    } else if  age1 <= age && age< age2 {
        resultCost = originalPrice * 0.7
    } else {
        resultCost = originalPrice * 0.5
    }
    return resultCost
}