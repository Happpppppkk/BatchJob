import pandas as pd
import timeit
N = 100 
start_time = timeit.default_timer()
with open('housesOutputPy1.txt', 'wt') as outfile:
    for i in range(N):
        houses = pd.read_csv("housesInput.csv")
        outfile.write(houses.describe().to_string(header=True, index=True))
        outfile.write("\n")

end_time = timeit.default_timer()
elapsed_time = end_time - start_time

print("Python code Processing time: %f seconds" % elapsed_time)


