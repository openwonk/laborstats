BLS API Client
=======
Client package that facilitates access to domestic labor data and statistics via the [US Bureau of Labor Statistics (BLS)](http://www.bls.gov/developers/).  For better access, be sure to obtain a free [registration key](http://data.bls.gov/registrationEngine/). 


### Usage
```go
package main

import (
	labor "github.com/openwonk/bls-api"
	"fmt"
)

func main() {
	s := labor.Series{
		"abcd1235e6789f0a987654bc3d21e0f1", // RegistrationKey
		"2010", // StartYear
		"2012", // EndYear
		false,  // Catalog
		true,   // Calculations
		true,   // AnnualAverage
		[]string{"LAUCN040010000000007", "LAUCN040010000000004"}, // Series
	}

	fmt.Println(s.Request())

}
```
<br>
<br>

<hr>
<small>
<strong>OpenWonk &copy; 2015 MIT License</strong>

</small>