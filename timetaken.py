import time
import json
import requests

avg=0.0
print("time to get countrynames")
for i in range(10):
    start=time.time()
    URL = "http://localhost:8000/countrynames"
    r = requests.get(url = URL) 
    data = r.json() 
    end=time.time()
    print("{:02d}: ".format(i+1), end-start)
    avg+=(end-start)
avg/=10
print("\nAverage: ",avg,"\n")



avg=0.0
print("time to get worldwide stats")
for i in range(10):
    start=time.time()
    URL = "http://localhost:8000/home"
    r = requests.get(url = URL) 
    data = r.json() 
    end=time.time()
    print("{:02d}: ".format(i+1), end-start)
    avg+=(end-start)
avg/=10
print("\nAverage: ",avg,"\n")



avg=0.0
print("time to get country stats")
URL = "http://localhost:8000/countrynames"
r = requests.get(url = URL) 
countrynames = r.json() 
for name in countrynames:
    if name=="Côte d'Ivoire":
        continue
    start=time.time()
    URL = "http://localhost:8000/country/"+name
    # try:
    r = requests.get(url = URL) 
    data = r.json() 
    end=time.time()
    print(name+": ", end-start)
    avg+=(end-start)
# except:
    #     continue
avg/=len(countrynames)
print("\nAverage: ",avg,"\n")



avg=0.0
print("time to get all data")
for i in range(10):
    start=time.time()
    URL = "http://localhost:8000/all"
    r = requests.get(url = URL) 
    data = r.json() 
    end=time.time()
    print("{:02d}: ".format(i+1), end-start)
    avg+=(end-start)
avg/=10
print("\nAverage: ",avg,"\n")
