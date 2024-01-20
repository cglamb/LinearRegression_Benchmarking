Author: Charles Lamb
Email: charlamb@gmail.com
Github address:  https://github.com/cglamb/LinearRegression_Benchmarking



***Introduction
This project tests linear regressiond performed in python, R, and Go for consistency and performance. 
We found all three programming languages produced consistent results.  We found that R had the shortest run-time while golang had the longest runtime
Results and a management recommendation are presented in the section immedately below. 
The balance of this README discusses experiement design and the files contained in this github.



***Results - Consistency
Consistiency was measured by ensuring each programming language produced the same linear fit on each of four data sets.
We present the intercept and slope for each dataset and programming language below.
The linear fit was found be identical across all three languages.  

			Data Set 1			Data Set 2			Data Set 3			Data Set 4
			int, slope			int, slope			int,slope			int, slope
Python			3.0001, 0.5001 			3.0009,  0.5000 		3.0025, 0.4997			3.0017, 0.4999 
R			3.0001, 0.5001			3.001,  0.500			3.0025, 0.4997			3.0017, 0.4999
Go			3.0001, 0.5001			3.0009, 0.5000			3.0025, 0.4997			3.0017, 0.4999



***Results - Benchmarking
To measure the speed at which the file executed, each file was ran 10 times.  The console was cleared between each execution to ensure no warmstarting.  Results per trial are shown below.  Times are shown in seconds.
We note that the statistics library used in the golang test (/montanaflynn/stats) did not (as far as this author can discern) have automatic reporting on line fitting.
The author built reporiting to generate the intercept and slope coefficients above and ran this reporting along with the regression in the benchmarked script.
We believe this additional reporting did not negatively impact runtime, as seperate testing that excluded the coeffienct calculations produced very similiar runtimes to those shown below.

Trial	.py	.r	.go
1	0.047	0.023	0.095
2	0.031	0.025	0.132
3	0.031	0.024	0.100
4	0.031	0.017	0.113
5	0.031	0.019	0.134
6	0.031	0.042	0.112
7	0.031	0.018	0.124
8	0.031	0.018	0.123
9	0.032	0.008	0.114
10	0.032	0.016	0.120
			
Avg	0.033	0.021	0.117

Note that go benchmarking using b.N interations produces a 0.128s runtime as opposed to the .117s above
	


***Management Recomendation
We recognize the benefit of the Company having a singular coding language.  Usage of a singular language cross-company should allow for greater collaberation between technical departments.  
However we do find some disadvantages from a data science team being moved to GoLang.  
The first disadvantage is longer runtimes.  In our limited study, we found a simple regression took signficantly longer to run in Golang than both Python and R (factor of 3-6 times slower).
A second disadvantage is modeling packages in golang appear less robust based on our limited study.  Example: to generate statistics as simple as the intercept and slope of a linear regression required building custom functions.
Finally while Golang is growing in popularity it is still less well know amongst the data science community than R or Python.  This may make talent acquisiton difficult and costly.
We recommend further testing around runtimes and robustness of modeling packages in Golang.  These challenges can be overcome but the Company should understand the investment it might require to overcome these hurdles.
Most importantly we recommend HR conduct market research around relative availability of data science professionals with prior experience working in golang.
If usage of golang would make talent acquistion difficult, we would recommend remaining with Python and Go



***Other Information*************************************************************************************************************************************************



***Datasets
Linear regressions were modeled on four different datasets from Anscombe (1973).  For convenience the four datasets are provided below:

{'x1' : [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
'x2' : [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
'x3' : [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
'x4' : [8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8],
'y1' : [8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26,10.84, 4.82, 5.68],
'y2' : [9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74],
'y3' : [7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73],
'y4' : [6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89]})

Anscombe, F. J. 1973, February. "Graphs in Statistical Analysis." The American Statistician 27(1): 17–21. Available online at https://www.sjsu.edu/faculty/gerstman/StatPrimer/anscombe1973.pdfLinks to an external site.



***Explanation of Files
linear_regresssion.py:  	Linear regression performed in python.  
linear_regression.r:		Linear regression performed in R
linear_regression.go:  		Linear regression performed in Go.  Uses "github.com/montanaflynn/stats" statistic package at https://pkg.go.dev/github.com/montanaflynn/stats
linear_regression_test.go:	Benchmarking for linear_regression.go
/Log/	
	PY LOG.txt: 	console output from running linear_regression.py
	R LOG.txt:	console output from running linear_regression.r	
	GO Log.txt:	console output for running linear_regression.go



***Hardware
All files were executed locally on the same hardware.  The hardware used a 13th Gen Intel(R) Core(TM) i5-13400F, 2500 Mhz, 10 Core(s), 16 Logical Processor(s) with 16.0GB of physical memory.



***Environements
The python code was executed in Spyder (Python 3.10.9 64-bit | Qt 5.15.2 | PyQt5 5.15.10 | Windows 10).  
The R script was executed in RStudio (Version 1.1.456).
The Go script was executed in Visual Code Studio (Version 1.85.1)