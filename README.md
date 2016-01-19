# Weatherman
Fetch and print forecast for Tromsø from yr.no. In Norwegian. 

# Usage 
First install [go](http://golang.org). Then: 

```
$ go get github.com/fjukstad/weatherman
$ weatherman
```

# Example output:
```
$ weatherman
Tirsdag 19.01.2016
	16:00-18:00 Delvis skyet og -14 grader. 0 mm nedbør
	18:00-00:00 Delvis skyet og -13 grader. 0 mm nedbør
Onsdag 20.01.2016
	00:00-06:00 Lettskyet og -11 grader. 0 mm nedbør
	06:00-12:00 Delvis skyet og -8 grader. 0 mm nedbør
	12:00-18:00 Lettskyet og -7 grader. 0 mm nedbør
	18:00-00:00 Lett snø og -7 grader. 0.7 mm nedbør
Torsdag 21.01.2016
	00:00-06:00 Delvis skyet og -5 grader. 0 mm nedbør
	06:00-12:00 Delvis skyet og -5 grader. 0 mm nedbør
	12:00-18:00 Skyet og -4 grader. 0 mm nedbør
	18:00-00:00 Skyet og -4 grader. 0 mm nedbør
Fredag 22.01.2016
	01:00-07:00 Skyet og -4 grader. 0 mm nedbør
	07:00-13:00 Skyet og -1 grader. 0 mm nedbør
	13:00-19:00 Delvis skyet og 0 grader. 0 mm nedbør
	19:00-01:00 Delvis skyet og 1 grader. 0 mm nedbør
...

```

# Todo
- [x] Make a simple prototype with weather for Tromsø. 
- [ ] Add support for emojis instead of text.
- [ ] Add support for text forecast, not just tabular. 
- [ ] Add command line options such as location and date

# Credit
Værvarsel fra yr.no, levert av NRK og Meteorologisk institutt. 
