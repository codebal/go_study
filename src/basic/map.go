package basic

import "fmt"

func studyMap() {
	{
		//make
		var idMap map[int]string
		idMap = make(map[int]string)
		_ = idMap

		map1 := make(map[int]string)
		fmt.Printf("%v\n", map1)
	}
	{
		//literal
		tickers := map[string]string{
			"GOOG": "Google Inc",
			"MSFT": "Microsoft",
			"FB":   "FaceBook",
		}
		fmt.Printf("%v , %s\n", tickers, tickers)
	}
	{
		//get, delete
		var m map[int]string

		m = make(map[int]string)
		m[901] = "Apple"
		m[134] = "Grape"
		m[777] = "Tomato"

		str := m[134]
		println("str=", str)

		noData := m[999]
		println("noData=", noData)

		delete(m, 777)
		fmt.Printf("%v\n", m)
	}
	{
		//exist
		tickers := map[string]string{
			"GOOG": "Google Inc",
			"MSFT": "Microsoft",
			"FB":   "FaceBook",
			"AMZN": "Amazon",
		}

		delete(tickers, "MSFT")

		val, exists := tickers["MSFT"]
		if !exists {
			println("No MSFT ticker")
		} else {
			println("MSFT is ", val)
		}
	}
	{
		//loop
		myMap := map[string]string{
			"A": "Apple",
			"B": "Banana",
			"C": "Charlie",
		}

		for key, val := range myMap {
			fmt.Println(key, val)
		}

	}
}
