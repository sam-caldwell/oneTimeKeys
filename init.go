package oneTimeKeys

import "time"

func init() {
	var consumptionRate float64
	var codesConsumed int
	startTime := time.Now().UnixNano()
	sleepInterval := 1

	go func() {
		for {
			inventoryLevel := otkQueue.Count()
			if inventoryLevel < minOneTimeKeys {
				codesConsumed++
				otkQueue.Push(generateNewCode())
			} else {
				if inventoryLevel >= maxOneTimeKeys {
					elapsedTime := time.Now().UnixNano() - startTime
					consumptionRate = float64(codesConsumed) / float64(elapsedTime)
					sleepInterval = int((1 - minimumInventoryMargin) / consumptionRate)
					time.Sleep(time.Duration(sleepInterval) * time.Nanosecond)
				}
			}
		}
	}()
}
