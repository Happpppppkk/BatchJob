# BatchJob
A repo for Week 4 assignment, Go command line application that summarize input data

## Go, Python and R
The folder contains three source files that perform input csv data reading, summarizing data of mean, count, median, standard deviations and iterating 100 times print out the processing time for each application.
In batchJob.go, the application require user to input csv file name that already store in the path. This will allow user to choose which file they want to summarize. The application also ask user input for the iterations. In this assignement, the iteration throughout the three source files are 100.
Example of the input:
"Welcome to data summary application!"
"Now, please enter the path of your CSV file"
--housesInput.csv
"How many time you want this summary to run?"
--100

In this experiment, we implement the dataframe as the data wrangling method and using stat package from the assignment one. The reason of using dataframe is that easier in readability and transit from python using to Go. But there is also a trade-off in the performance. 
All functions are in main.go file, the usage of third party packages are gota that contains dataframe and series that help reading data into dataframe and define data type, and montanaflynn/stat package for the statistical functions. 

## Test result
The result coming from terminal are following:


The measurement are telling the best performance of python in this experiment and Go is slightly slower than Python. R is using the most processing time. Comparing to the assignment 1 where Go has the best performance, this experiment is not very well designed due to the usage of dataframe that is causing the slower process time in Go. Alternatively, to use efficient code, Ricardo Gerardi's colStats provides great outline for reading data from csv file.(https://pragprog.com/titles/rggo/powerful-command-line-applications-in-go/)

## More to concern
Python, with great performance in data wrangling and manipulation with outstanding performance in this test is hard to complete. Comparing to assignment one when purely testing in processing time of compiling code, Python has higher level library and better ecosystem in data manipulation. In this case, Go's advantage such as concurrency or type safety are not obvious.



