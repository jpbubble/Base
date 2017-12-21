/*
        bubble_getdata.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.21
*/
package bubble

import(
	"strings"
	"trickyunits/qll"
	)

func init(){
mkl.Version("Bubble Base - bubble_getdata.go","17.12.21")
mkl.Lic    ("Bubble Base - bubble_getdata.go","Mozilla Public License 2.0")
}


func GetData(a string) []byte {
	dpp:=strings.Index(a,":")
	from:="jcr"
	file:=a
	if dpp>0 {
		from=a[:dpp]
		file=a[dpp+1:]
	}
	switch from {
		case "real":
			return qff.GetData(file)
		case "jcr":
			if bubjcr==nil {
				buberror("No JCR file set to get data from")
				return []byte{}
			} 
			return JCR_B(bubjcr,a)
		}
	}
	buberror("bubble.GetData(\""+a+"\"): Complete malfunction!",true)
	return []byte{}
}
