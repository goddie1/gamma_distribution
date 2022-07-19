# This is a sample Python script.

# Press ⌃R to execute it or replace it with your code.
# Press Double ⇧ to search everywhere for classes, files, tool windows, actions, and settings.

import numpy as np
import matplotlib.pyplot as plt
from scipy import  stats
import pandas as pd
def print_hi(name):
    # Use a breakpoint in the code line below to debug your script.
    print(f'Hi, {name}')  # Press ⌘F8 to toggle the breakpoint.

def gamma_function(n):
    cal = 1
    for i in range(2,n):
        cal *=i
    return cal

def gamma(x,a,b):
    c = (b ** a) / gamma_function(a)
    y = c * (x ** (a - 1)) * np.exp(-b * x)
    return x,y,np.mean(y),np.std(y)

for ls in [(2,2)]:
    a ,b = 2,2
    x = np.arange(0,5,0.01,dtype=np.float64)
    x,y,u,s = gamma(x,a=a,b=b)
    plt.plot(x,y,label=r'$\mu=%.2f,\ \sigma=%.2f,'
             r'\ \alpha=%d,\ \beta=%d$' %(u,s,a,b))
    plt.legend()
    plt.title('gamma')
    plt.show()




# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    print_hi('PyCharm')

# See PyCharm help at https://www.jetbrains.com/help/pycharm/
