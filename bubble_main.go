/*
        bubble_main.go
	(c) 2017 Jeroen Petrus Broks.
	
	This Source Code Form is subject to the terms of the 
	Mozilla Public License, v. 2.0. If a copy of the MPL was not 
	distributed with this file, You can obtain one at 
	http://mozilla.org/MPL/2.0/.
        Version: 17.12.21
*/
package bubble

import(
	"trickyunits/mkl"
	"trickyunits/jcr6/jcr6main"
)

var bubjcr jcr6main.TJCR6Dir

func init(){
mkl.Version("Bubble Base - bubble_main.go","17.12.21")
mkl.Lic    ("Bubble Base - bubble_main.go","Mozilla Public License 2.0")
}

func SetJCR(jcrfile string) jcr6main.TJCR6Dir{
	bubjcr = jcr6main.Dir(jcrfile)
	return bubjcr
}
