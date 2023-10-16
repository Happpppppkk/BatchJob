N = 100 
start.time <- Sys.time() 
sink("housesOutputR2.txt")
for (i in 1:N) {
    houses = read.csv(file = "housesInput.csv", header = TRUE)
    print(summary(houses)) 
}
sink()
end.time <- Sys.time() 
elapsed.time <- end.time - start.time

cat("R code processing time: ", elapsed.time)