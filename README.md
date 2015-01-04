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
<strong>OpenWonk &copy; 2015 </strong>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
</small>